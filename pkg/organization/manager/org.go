// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package org

import (
	"fmt"

	"github.com/gitbundle/modules/storage"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/organization"
	packages_model "github.com/gitbundle/server/pkg/packages"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"
)

// DeleteOrganization completely and permanently deletes everything of organization.
func DeleteOrganization(org *organization.Organization) error {
	ctx, commiter, err := db.TxContext()
	if err != nil {
		return err
	}
	defer commiter.Close()

	// Check ownership of repository.
	count, err := repo_model.CountRepositories(ctx, repo_model.CountRepositoryOptions{OwnerID: org.ID})
	if err != nil {
		return fmt.Errorf("GetRepositoryCount: %v", err)
	} else if count > 0 {
		return user_model.ErrUserOwnRepos{UID: org.ID}
	}

	// Check ownership of packages.
	if ownsPackages, err := packages_model.HasOwnerPackages(ctx, org.ID); err != nil {
		return fmt.Errorf("HasOwnerPackages: %v", err)
	} else if ownsPackages {
		return user_model.ErrUserOwnPackages{UID: org.ID}
	}

	if err := organization.DeleteOrganization(ctx, org); err != nil {
		return fmt.Errorf("DeleteOrganization: %v", err)
	}

	if err := commiter.Commit(); err != nil {
		return err
	}

	// FIXME: system notice
	// Note: There are something just cannot be roll back,
	//	so just keep error logs of those operations.
	path := user_model.UserPath(org.Name)

	if err := util.RemoveAll(path); err != nil {
		return fmt.Errorf("Failed to RemoveAll %s: %v", path, err)
	}

	if len(org.Avatar) > 0 {
		avatarPath := org.CustomAvatarRelativePath()
		if err := storage.Avatars.Delete(avatarPath); err != nil {
			return fmt.Errorf("Failed to remove %s: %v", avatarPath, err)
		}
	}

	return nil
}
