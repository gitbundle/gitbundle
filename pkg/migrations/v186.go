// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"github.com/gitbundle/modules/timeutil"

	"xorm.io/xorm"
)

func createProtectedTagTable(x *xorm.Engine) error {
	type ProtectedTag struct {
		ID               int64 `xorm:"pk autoincr"`
		RepoID           int64
		NamePattern      string
		AllowlistUserIDs []int64 `xorm:"JSON TEXT"`
		AllowlistTeamIDs []int64 `xorm:"JSON TEXT"`

		CreatedUnix timeutil.TimeStamp `xorm:"created"`
		UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
	}

	return x.Sync2(new(ProtectedTag))
}
