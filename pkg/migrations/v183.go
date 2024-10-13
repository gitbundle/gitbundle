// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"
	"time"

	"github.com/gitbundle/modules/timeutil"

	"xorm.io/xorm"
)

func createPushMirrorTable(x *xorm.Engine) error {
	type PushMirror struct {
		ID         int64 `xorm:"pk autoincr"`
		RepoID     int64 `xorm:"INDEX"`
		RemoteName string

		Interval       time.Duration
		CreatedUnix    timeutil.TimeStamp `xorm:"created"`
		LastUpdateUnix timeutil.TimeStamp `xorm:"INDEX last_update"`
		LastError      string             `xorm:"text"`
	}

	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	if err := sess.Sync2(new(PushMirror)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}

	return sess.Commit()
}
