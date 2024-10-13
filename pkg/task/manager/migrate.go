// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package task

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/process"
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/db"
	migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
	"github.com/gitbundle/server/pkg/notification"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/repoman"
	task_model "github.com/gitbundle/server/pkg/task"
	user_model "github.com/gitbundle/server/pkg/user"
)

func handleCreateError(owner *user_model.User, err error) error {
	switch {
	case repo_model.IsErrReachLimitOfRepo(err):
		return fmt.Errorf("You have already reached your limit of %d repositories", owner.MaxCreationLimit())
	case repo_model.IsErrRepoAlreadyExist(err):
		return errors.New("The repository name is already used")
	case db.IsErrNameReserved(err):
		return fmt.Errorf("The repository name '%s' is reserved", err.(db.ErrNameReserved).Name)
	case db.IsErrNamePatternNotAllowed(err):
		return fmt.Errorf("The pattern '%s' is not allowed in a repository name", err.(db.ErrNamePatternNotAllowed).Pattern)
	default:
		return err
	}
}

func runMigrateTask(t *task_model.Task) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PANIC whilst trying to do migrate task: %v", e)
			log.Critical("PANIC during runMigrateTask[%d] by DoerID[%d] to RepoID[%d] for OwnerID[%d]: %v\nStacktrace: %v", t.ID, t.DoerID, t.RepoID, t.OwnerID, e, log.Stack(2))
		}

		if err == nil {
			err = task_model.FinishMigrateTask(t)
			if err == nil {
				notification.NotifyMigrateRepository(t.Doer, t.Owner, t.Repo)
				return
			}

			log.Error("FinishMigrateTask[%d] by DoerID[%d] to RepoID[%d] for OwnerID[%d] failed: %v", t.ID, t.DoerID, t.RepoID, t.OwnerID, err)
		}

		t.EndTime = timeutil.TimeStampNow()
		t.Status = structs.TaskStatusFailed
		t.Message = err.Error()
		// Ensure that the repo loaded before we zero out the repo ID from the task - thus ensuring that we can delete it
		_ = t.LoadRepo()

		t.RepoID = 0
		if err := t.UpdateCols("status", "errors", "repo_id", "end_time"); err != nil {
			log.Error("Task UpdateCols failed: %v", err)
		}

		if t.Repo != nil {
			if errDelete := repoman.DeleteRepository(t.Doer, t.OwnerID, t.Repo.ID); errDelete != nil {
				log.Error("DeleteRepository: %v", errDelete)
			}
		}
	}()

	if err = t.LoadRepo(); err != nil {
		return
	}

	// if repository is ready, then just finish the task
	if t.Repo.Status == repo_model.RepositoryReady {
		return nil
	}

	if err = t.LoadDoer(); err != nil {
		return
	}
	if err = t.LoadOwner(); err != nil {
		return
	}

	var opts *schemas.MigrateOptions
	opts, err = t.MigrateConfig()
	if err != nil {
		return
	}

	opts.MigrateToRepoID = t.RepoID

	pm := process.GetManager()
	ctx, _, finished := pm.AddContext(graceful.GetManager().ShutdownContext(), fmt.Sprintf("MigrateTask: %s/%s", t.Owner.Name, opts.RepoName))
	defer finished()

	t.StartTime = timeutil.TimeStampNow()
	t.Status = structs.TaskStatusRunning
	if err = t.UpdateCols("start_time", "status"); err != nil {
		return
	}

	t.Repo, err = migrations.MigrateRepository(ctx, t.Doer, t.Owner.Name, *opts, func(format string, args ...interface{}) {
		message := task_model.TranslatableMessage{
			Format: format,
			Args:   args,
		}
		bs, _ := json.Marshal(message)
		t.Message = string(bs)
		_ = t.UpdateCols("message")
	})
	if err == nil {
		log.Trace("Repository migrated [%d]: %s/%s", t.Repo.ID, t.Owner.Name, t.Repo.Name)
		return
	}

	if repo_model.IsErrRepoAlreadyExist(err) {
		err = errors.New("The repository name is already used")
		return
	}

	// remoteAddr may contain credentials, so we sanitize it
	err = util.SanitizeErrorCredentialURLs(err)
	if strings.Contains(err.Error(), "Authentication failed") ||
		strings.Contains(err.Error(), "could not read Username") {
		return fmt.Errorf("Authentication failed: %v", err.Error())
	} else if strings.Contains(err.Error(), "fatal:") {
		return fmt.Errorf("Migration failed: %v", err.Error())
	}

	// do not be tempted to coalesce this line with the return
	err = handleCreateError(t.Owner, err)
	return
}
