// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"crypto/sha1"
	"fmt"

	"xorm.io/xorm"
)

func hashContext(context string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(context)))
}

func addCommitStatusContext(x *xorm.Engine) error {
	type CommitStatus struct {
		ID          int64  `xorm:"pk autoincr"`
		ContextHash string `xorm:"char(40) index"`
		Context     string `xorm:"TEXT"`
	}

	if err := x.Sync2(new(CommitStatus)); err != nil {
		return err
	}

	sess := x.NewSession()
	defer sess.Close()

	start := 0
	for {
		statuses := make([]*CommitStatus, 0, 100)
		err := sess.OrderBy("id").Limit(100, start).Find(&statuses)
		if err != nil {
			return err
		}
		if len(statuses) == 0 {
			break
		}

		if err = sess.Begin(); err != nil {
			return err
		}

		for _, status := range statuses {
			status.ContextHash = hashContext(status.Context)
			if _, err := sess.ID(status.ID).Cols("context_hash").Update(status); err != nil {
				return err
			}
		}

		if err := sess.Commit(); err != nil {
			return err
		}

		if len(statuses) < 100 {
			break
		}

		start += len(statuses)
	}

	return nil
}
