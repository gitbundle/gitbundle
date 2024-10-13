// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"github.com/gitbundle/modules/log"

	"xorm.io/builder"
	"xorm.io/xorm"
)

func setIsArchivedToFalse(x *xorm.Engine) error {
	type Repository struct {
		IsArchived bool `xorm:"INDEX"`
	}
	count, err := x.Where(builder.IsNull{"is_archived"}).Cols("is_archived").Update(&Repository{
		IsArchived: false,
	})
	if err == nil {
		log.Debug("Updated %d repositories with is_archived IS NULL", count)
	}
	return err
}
