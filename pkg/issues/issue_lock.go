// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issues

import (
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

// IssueLockOptions defines options for locking and/or unlocking an issue/PR
type IssueLockOptions struct {
	Doer   *user_model.User
	Issue  *Issue
	Reason string
}

// LockIssue locks an issue. This would limit commenting abilities to
// users with write access to the repo
func LockIssue(opts *IssueLockOptions) error {
	return updateIssueLock(opts, true)
}

// UnlockIssue unlocks a previously locked issue.
func UnlockIssue(opts *IssueLockOptions) error {
	return updateIssueLock(opts, false)
}

func updateIssueLock(opts *IssueLockOptions, lock bool) error {
	if opts.Issue.IsLocked == lock {
		return nil
	}

	opts.Issue.IsLocked = lock
	var commentType CommentType
	if opts.Issue.IsLocked {
		commentType = CommentTypeLock
	} else {
		commentType = CommentTypeUnlock
	}

	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if err := UpdateIssueCols(ctx, opts.Issue, "is_locked"); err != nil {
		return err
	}

	opt := &CreateCommentOptions{
		Doer:    opts.Doer,
		Issue:   opts.Issue,
		Repo:    opts.Issue.Repo,
		Type:    commentType,
		Content: opts.Reason,
	}
	if _, err := CreateCommentCtx(ctx, opt); err != nil {
		return err
	}

	return committer.Commit()
}
