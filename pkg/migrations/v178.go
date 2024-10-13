// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addLFSMirrorColumns(x *xorm.Engine) error {
	type Mirror struct {
		LFS         bool   `xorm:"lfs_enabled NOT NULL DEFAULT false"`
		LFSEndpoint string `xorm:"lfs_endpoint TEXT"`
	}

	return x.Sync2(new(Mirror))
}
