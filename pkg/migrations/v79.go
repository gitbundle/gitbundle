// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"github.com/gitbundle/modules/setting"

	"xorm.io/xorm"
)

func addCanCloseIssuesViaCommitInAnyBranch(x *xorm.Engine) error {
	type Repository struct {
		ID                              int64 `xorm:"pk autoincr"`
		CloseIssuesViaCommitInAnyBranch bool  `xorm:"NOT NULL DEFAULT false"`
	}

	if err := x.Sync2(new(Repository)); err != nil {
		return err
	}

	_, err := x.Exec("UPDATE repository SET close_issues_via_commit_in_any_branch = ?",
		setting.Repository.DefaultCloseIssuesViaCommitsInAnyBranch)
	return err
}
