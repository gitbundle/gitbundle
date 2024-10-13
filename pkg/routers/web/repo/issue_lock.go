// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/forms"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/web"
)

// LockIssue locks an issue. This would limit commenting abilities to
// users with write access to the repo.
func LockIssue(ctx *context.Context) {
	form := web.GetForm(ctx).(*forms.IssueLockForm)
	issue := GetActionIssue(ctx)
	if ctx.Written() {
		return
	}

	if issue.IsLocked {
		ctx.Flash.Error(ctx.Tr("repo.issues.lock_duplicate"))
		ctx.Redirect(issue.HTMLURL())
		return
	}

	if !form.HasValidReason() {
		ctx.Flash.Error(ctx.Tr("repo.issues.lock.unknown_reason"))
		ctx.Redirect(issue.HTMLURL())
		return
	}

	if err := issues_model.LockIssue(&issues_model.IssueLockOptions{
		Doer:   ctx.Doer,
		Issue:  issue,
		Reason: form.Reason,
	}); err != nil {
		ctx.ServerError("LockIssue", err)
		return
	}

	ctx.Redirect(issue.HTMLURL())
}

// UnlockIssue unlocks a previously locked issue.
func UnlockIssue(ctx *context.Context) {
	issue := GetActionIssue(ctx)
	if ctx.Written() {
		return
	}

	if !issue.IsLocked {
		ctx.Flash.Error(ctx.Tr("repo.issues.unlock_error"))
		ctx.Redirect(issue.HTMLURL())
		return
	}

	if err := issues_model.UnlockIssue(&issues_model.IssueLockOptions{
		Doer:  ctx.Doer,
		Issue: issue,
	}); err != nil {
		ctx.ServerError("UnlockIssue", err)
		return
	}

	ctx.Redirect(issue.HTMLURL())
}
