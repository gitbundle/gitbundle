// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addOwnerNameOnRepository(x *xorm.Engine) error {
	type Repository struct {
		OwnerName string
	}
	if err := x.Sync2(new(Repository)); err != nil {
		return err
	}
	_, err := x.Exec("UPDATE repository SET owner_name = (SELECT name FROM `user` WHERE `user`.id = repository.owner_id)")
	return err
}
