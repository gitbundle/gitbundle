// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addTeamReviewRequestSupport(x *xorm.Engine) error {
	type Review struct {
		ReviewerTeamID int64 `xorm:"NOT NULL DEFAULT 0"`
	}

	type Comment struct {
		AssigneeTeamID int64 `xorm:"NOT NULL DEFAULT 0"`
	}

	if err := x.Sync2(new(Review)); err != nil {
		return err
	}

	return x.Sync2(new(Comment))
}
