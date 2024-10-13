// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import "xorm.io/xorm"

func addIndexOnRepositoryAndComment(x *xorm.Engine) error {
	type Repository struct {
		ID      int64 `xorm:"pk autoincr"`
		OwnerID int64 `xorm:"index"`
	}

	if err := x.Sync2(new(Repository)); err != nil {
		return err
	}

	type Comment struct {
		ID       int64 `xorm:"pk autoincr"`
		Type     int   `xorm:"index"`
		ReviewID int64 `xorm:"index"`
	}

	return x.Sync2(new(Comment))
}
