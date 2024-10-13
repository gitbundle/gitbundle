// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func removeLabelUneededCols(x *xorm.Engine) error {
	// Make sure the columns exist before dropping them
	type Label struct {
		QueryString string
		IsSelected  bool
	}
	if err := x.Sync2(new(Label)); err != nil {
		return err
	}

	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	if err := dropTableColumns(sess, "label", "query_string"); err != nil {
		return err
	}
	if err := dropTableColumns(sess, "label", "is_selected"); err != nil {
		return err
	}
	return sess.Commit()
}
