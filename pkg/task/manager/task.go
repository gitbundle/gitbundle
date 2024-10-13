// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package task

import (
	"fmt"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/queue"
	"github.com/gitbundle/modules/secret"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/repoman"
	repo_module "github.com/gitbundle/server/pkg/repo/repository"
	task_model "github.com/gitbundle/server/pkg/task"
	user_model "github.com/gitbundle/server/pkg/user"
)

// taskQueue is a global queue of tasks
var taskQueue queue.Queue

// Run a task
func Run(t *task_model.Task) error {
	switch t.Type {
	case structs.TaskTypeMigrateRepo:
		return runMigrateTask(t)
	default:
		return fmt.Errorf("Unknown task type: %d", t.Type)
	}
}

// Init will start the service to get all unfinished tasks and run them
func Init() error {
	taskQueue = queue.CreateQueue("task", handle, &task_model.Task{})

	if taskQueue == nil {
		return fmt.Errorf("Unable to create Task Queue")
	}

	go graceful.GetManager().RunWithShutdownFns(taskQueue.Run)

	return nil
}

func handle(data ...queue.Data) []queue.Data {
	for _, datum := range data {
		task := datum.(*task_model.Task)
		if err := Run(task); err != nil {
			log.Error("Run task failed: %v", err)
		}
	}
	return nil
}

// MigrateRepository add migration repository to task
func MigrateRepository(doer, u *user_model.User, opts schemas.MigrateOptions) error {
	task, err := CreateMigrateTask(doer, u, opts)
	if err != nil {
		return err
	}

	return taskQueue.Push(task)
}

// CreateMigrateTask creates a migrate task
func CreateMigrateTask(doer, u *user_model.User, opts schemas.MigrateOptions) (*task_model.Task, error) {
	// encrypt credentials for persistence
	var err error
	opts.CloneAddrEncrypted, err = secret.EncryptSecret(setting.SecretKey, opts.CloneAddr)
	if err != nil {
		return nil, err
	}
	opts.CloneAddr = util.SanitizeCredentialURLs(opts.CloneAddr)
	opts.AuthPasswordEncrypted, err = secret.EncryptSecret(setting.SecretKey, opts.AuthPassword)
	if err != nil {
		return nil, err
	}
	opts.AuthPassword = ""
	opts.AuthTokenEncrypted, err = secret.EncryptSecret(setting.SecretKey, opts.AuthToken)
	if err != nil {
		return nil, err
	}
	opts.AuthToken = ""
	bs, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	task := &task_model.Task{
		DoerID:         doer.ID,
		OwnerID:        u.ID,
		Type:           structs.TaskTypeMigrateRepo,
		Status:         structs.TaskStatusQueue,
		PayloadContent: string(bs),
	}

	if err := task_model.CreateTask(task); err != nil {
		return nil, err
	}

	repo, err := repo_module.CreateRepository(doer, u, repoman.CreateRepoOptions{
		Name:           opts.RepoName,
		Description:    opts.Description,
		OriginalURL:    opts.OriginalURL,
		GitServiceType: opts.GitServiceType,
		IsPrivate:      opts.Private,
		IsMirror:       opts.Mirror,
		Status:         repo_model.RepositoryBeingMigrated,
	})
	if err != nil {
		task.EndTime = timeutil.TimeStampNow()
		task.Status = structs.TaskStatusFailed
		err2 := task.UpdateCols("end_time", "status")
		if err2 != nil {
			log.Error("UpdateCols Failed: %v", err2.Error())
		}
		return nil, err
	}

	task.RepoID = repo.ID
	if err = task.UpdateCols("repo_id"); err != nil {
		return nil, err
	}

	return task, nil
}
