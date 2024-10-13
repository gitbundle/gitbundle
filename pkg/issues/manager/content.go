// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issue

import (
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/notification"
	user_model "github.com/gitbundle/server/pkg/user"
)

// ChangeContent changes issue content, as the given user.
func ChangeContent(issue *issues_model.Issue, doer *user_model.User, content string) (err error) {
	oldContent := issue.Content

	if err := issues_model.ChangeIssueContent(issue, doer, content); err != nil {
		return err
	}

	notification.NotifyIssueChangeContent(doer, issue, oldContent)

	return nil
}
