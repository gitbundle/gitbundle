// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addTeamIncludesAllRepositories(x *xorm.Engine) error {
	type Team struct {
		ID                      int64 `xorm:"pk autoincr"`
		IncludesAllRepositories bool  `xorm:"NOT NULL DEFAULT false"`
	}

	if err := x.Sync2(new(Team)); err != nil {
		return err
	}

	_, err := x.Exec("UPDATE `team` SET `includes_all_repositories` = ? WHERE `name`=?",
		true, "Owners")
	return err
}
