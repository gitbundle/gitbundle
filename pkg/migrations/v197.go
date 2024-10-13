// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addRenamedBranchTable(x *xorm.Engine) error {
	type RenamedBranch struct {
		ID          int64 `xorm:"pk autoincr"`
		RepoID      int64 `xorm:"INDEX NOT NULL"`
		From        string
		To          string
		CreatedUnix int64 `xorm:"created"`
	}
	return x.Sync2(new(RenamedBranch))
}
