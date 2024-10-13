// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addProjectIssueSorting(x *xorm.Engine) error {
	// ProjectIssue saves relation from issue to a project
	type ProjectIssue struct {
		Sorting int64 `xorm:"NOT NULL DEFAULT 0"`
	}

	return x.Sync2(new(ProjectIssue))
}
