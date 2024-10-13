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

func addSessionTable(x *xorm.Engine) error {
	type Session struct {
		Key    string `xorm:"pk CHAR(16)"`
		Data   []byte `xorm:"BLOB"`
		Expiry timeutil.TimeStamp
	}
	return x.Sync2(new(Session))
}
