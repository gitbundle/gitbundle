// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/organization"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"
)

// CanUserDelete returns true if user could delete the repository
func CanUserDelete(repo *repo_model.Repository, user *user_model.User) (bool, error) {
	if user.IsAdmin || user.ID == repo.OwnerID {
		return true, nil
	}

	if err := repo.GetOwner(db.DefaultContext); err != nil {
		return false, err
	}

	if repo.Owner.IsOrganization() {
		isOwner, err := organization.OrgFromUser(repo.Owner).IsOwnedBy(user.ID)
		if err != nil {
			return false, err
		}
		return isOwner, nil
	}

	return false, nil
}
