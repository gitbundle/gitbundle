// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/organization"
	"github.com/gitbundle/server/pkg/perm"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

// GetReviewerTeams get all teams can be requested to review
func GetReviewerTeams(repo *repo_model.Repository) ([]*organization.Team, error) {
	if err := repo.GetOwner(db.DefaultContext); err != nil {
		return nil, err
	}
	if !repo.Owner.IsOrganization() {
		return nil, nil
	}

	return organization.GetTeamsWithAccessToRepo(db.DefaultContext, repo.OwnerID, repo.ID, perm.AccessModeRead)
}
