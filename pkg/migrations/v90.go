// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import "xorm.io/xorm"

func changeSomeColumnsLengthOfRepo(x *xorm.Engine) error {
	type Repository struct {
		ID          int64  `xorm:"pk autoincr"`
		Description string `xorm:"TEXT"`
		Website     string `xorm:"VARCHAR(2048)"`
		OriginalURL string `xorm:"VARCHAR(2048)"`
	}

	return x.Sync2(new(Repository))
}
