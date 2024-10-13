// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issue

import (
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/notification"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	user_model "github.com/gitbundle/server/pkg/user"
)

// ClearLabels clears all of an issue's labels
func ClearLabels(issue *issues_model.Issue, doer *user_model.User) (err error) {
	if err = issues_model.ClearIssueLabels(issue, doer); err != nil {
		return
	}

	notification.NotifyIssueClearLabels(doer, issue)

	return nil
}

// AddLabel adds a new label to the issue.
func AddLabel(issue *issues_model.Issue, doer *user_model.User, label *issues_model.Label) error {
	if err := issues_model.NewIssueLabel(issue, label, doer); err != nil {
		return err
	}

	notification.NotifyIssueChangeLabels(doer, issue, []*issues_model.Label{label}, nil)
	return nil
}

// AddLabels adds a list of new labels to the issue.
func AddLabels(issue *issues_model.Issue, doer *user_model.User, labels []*issues_model.Label) error {
	if err := issues_model.NewIssueLabels(issue, labels, doer); err != nil {
		return err
	}

	notification.NotifyIssueChangeLabels(doer, issue, labels, nil)
	return nil
}

// RemoveLabel removes a label from issue by given ID.
func RemoveLabel(issue *issues_model.Issue, doer *user_model.User, label *issues_model.Label) error {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if err := issue.LoadRepo(ctx); err != nil {
		return err
	}

	perm, err := access_model.GetUserRepoPermission(ctx, issue.Repo, doer)
	if err != nil {
		return err
	}
	if !perm.CanWriteIssuesOrPulls(issue.IsPull) {
		if label.OrgID > 0 {
			return issues_model.ErrOrgLabelNotExist{}
		}
		return issues_model.ErrRepoLabelNotExist{}
	}

	if err := issues_model.DeleteIssueLabel(ctx, issue, label, doer); err != nil {
		return err
	}

	if err := committer.Commit(); err != nil {
		return err
	}

	notification.NotifyIssueChangeLabels(doer, issue, nil, []*issues_model.Label{label})
	return nil
}

// ReplaceLabels removes all current labels and add new labels to the issue.
func ReplaceLabels(issue *issues_model.Issue, doer *user_model.User, labels []*issues_model.Label) error {
	old, err := issues_model.GetLabelsByIssueID(db.DefaultContext, issue.ID)
	if err != nil {
		return err
	}

	if err := issues_model.ReplaceIssueLabels(issue, labels, doer); err != nil {
		return err
	}

	notification.NotifyIssueChangeLabels(doer, issue, labels, old)
	return nil
}
