// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"xorm.io/xorm"
)

func addAuthorizeColForTeamUnit(x *xorm.Engine) error {
	type TeamUnit struct {
		ID         int64 `xorm:"pk autoincr"`
		OrgID      int64 `xorm:"INDEX"`
		TeamID     int64 `xorm:"UNIQUE(s)"`
		Type       int   `xorm:"UNIQUE(s)"`
		AccessMode int
	}

	if err := x.Sync2(new(TeamUnit)); err != nil {
		return fmt.Errorf("sync2: %v", err)
	}

	// migrate old permission
	_, err := x.Exec("UPDATE team_unit SET access_mode = (SELECT authorize FROM team WHERE team.id = team_unit.team_id)")
	return err
}
