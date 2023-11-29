// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package organization

import (
	"context"

	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/perm"
	"github.com/gitbundle/server/pkg/unit"
)

// TeamUnit describes all units of a repository
type TeamUnit struct {
	ID         int64     `xorm:"pk autoincr"`
	OrgID      int64     `xorm:"INDEX"`
	TeamID     int64     `xorm:"UNIQUE(s)"`
	Type       unit.Type `xorm:"UNIQUE(s)"`
	AccessMode perm.AccessMode
}

// Unit returns Unit
func (t *TeamUnit) Unit() unit.Unit {
	return unit.Units[t.Type]
}

func getUnitsByTeamID(ctx context.Context, teamID int64) (units []*TeamUnit, err error) {
	return units, db.GetEngine(ctx).Where("team_id = ?", teamID).Find(&units)
}

// UpdateTeamUnits updates a teams's units
func UpdateTeamUnits(team *Team, units []TeamUnit) (err error) {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if _, err = db.GetEngine(ctx).Where("team_id = ?", team.ID).Delete(new(TeamUnit)); err != nil {
		return err
	}

	if len(units) > 0 {
		if err = db.Insert(ctx, units); err != nil {
			return err
		}
	}

	return committer.Commit()
}