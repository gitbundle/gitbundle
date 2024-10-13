// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func extendTrackedTimes(x *xorm.Engine) error {
	type TrackedTime struct {
		Time    int64 `xorm:"NOT NULL"`
		Deleted bool  `xorm:"NOT NULL DEFAULT false"`
	}

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if _, err := sess.Exec("DELETE FROM tracked_time WHERE time IS NULL"); err != nil {
		return err
	}

	if err := sess.Sync2(new(TrackedTime)); err != nil {
		return err
	}

	return sess.Commit()
}
