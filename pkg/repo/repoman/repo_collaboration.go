// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repoman

import (
	"context"
	"fmt"

	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/perm"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"

	"xorm.io/builder"
)

func addCollaborator(ctx context.Context, repo *repo_model.Repository, u *user_model.User) error {
	collaboration := &repo_model.Collaboration{
		RepoID: repo.ID,
		UserID: u.ID,
	}

	has, err := db.GetByBean(ctx, collaboration)
	if err != nil {
		return err
	} else if has {
		return nil
	}
	collaboration.Mode = perm.AccessModeWrite

	if err = db.Insert(ctx, collaboration); err != nil {
		return err
	}

	return access_model.RecalculateUserAccess(ctx, repo, u.ID)
}

// AddCollaborator adds new collaboration to a repository with default access mode.
func AddCollaborator(repo *repo_model.Repository, u *user_model.User) error {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if err := addCollaborator(ctx, repo, u); err != nil {
		return err
	}

	return committer.Commit()
}

// DeleteCollaboration removes collaboration relation between the user and repository.
func DeleteCollaboration(repo *repo_model.Repository, uid int64) (err error) {
	collaboration := &repo_model.Collaboration{
		RepoID: repo.ID,
		UserID: uid,
	}

	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if has, err := db.GetEngine(ctx).Delete(collaboration); err != nil || has == 0 {
		return err
	} else if err = access_model.RecalculateAccesses(ctx, repo); err != nil {
		return err
	}

	if err = repo_model.WatchRepo(ctx, uid, repo.ID, false); err != nil {
		return err
	}

	if err = reconsiderWatches(ctx, repo, uid); err != nil {
		return err
	}

	// Unassign a user from any issue (s)he has been assigned to in the repository
	if err := reconsiderRepoIssuesAssignee(ctx, repo, uid); err != nil {
		return err
	}

	return committer.Commit()
}

func reconsiderRepoIssuesAssignee(ctx context.Context, repo *repo_model.Repository, uid int64) error {
	user, err := user_model.GetUserByIDCtx(ctx, uid)
	if err != nil {
		return err
	}

	if canAssigned, err := access_model.CanBeAssigned(ctx, user, repo, true); err != nil || canAssigned {
		return err
	}

	if _, err := db.GetEngine(ctx).Where(builder.Eq{"assignee_id": uid}).
		In("issue_id", builder.Select("id").From("issue").Where(builder.Eq{"repo_id": repo.ID})).
		Delete(&issues_model.IssueAssignees{}); err != nil {
		return fmt.Errorf("Could not delete assignee[%d] %v", uid, err)
	}
	return nil
}

func reconsiderWatches(ctx context.Context, repo *repo_model.Repository, uid int64) error {
	if has, err := access_model.HasAccess(ctx, uid, repo); err != nil || has {
		return err
	}
	if err := repo_model.WatchRepo(ctx, uid, repo.ID, false); err != nil {
		return err
	}

	// Remove all IssueWatches a user has subscribed to in the repository
	return issues_model.RemoveIssueWatchersByRepoID(ctx, uid, repo.ID)
}
