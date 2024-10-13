// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addUserRepoMissingColumns(x *xorm.Engine) error {
	type VisibleType int
	type User struct {
		PasswdHashAlgo string      `xorm:"NOT NULL DEFAULT 'pbkdf2'"`
		Visibility     VisibleType `xorm:"NOT NULL DEFAULT 0"`
	}

	type Repository struct {
		IsArchived bool     `xorm:"INDEX"`
		Topics     []string `xorm:"TEXT JSON"`
	}

	return x.Sync2(new(User), new(Repository))
}
