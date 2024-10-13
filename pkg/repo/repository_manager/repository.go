// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"fmt"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	admin_model "github.com/gitbundle/server/pkg/admin"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/notification"
	"github.com/gitbundle/server/pkg/organization"
	packages_model "github.com/gitbundle/server/pkg/packages"
	pull_service "github.com/gitbundle/server/pkg/pull/manager"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/release"
	"github.com/gitbundle/server/pkg/repo/repoman"
	repo_module "github.com/gitbundle/server/pkg/repo/repository"
	"github.com/gitbundle/server/pkg/unit"
	user_model "github.com/gitbundle/server/pkg/user"
)

// CreateRepository creates a repository for the user/organization.
func CreateRepository(doer, owner *user_model.User, opts repoman.CreateRepoOptions) (*repo_model.Repository, error) {
	repo, err := repo_module.CreateRepository(doer, owner, opts)
	if err != nil {
		// No need to rollback here we should do this in CreateRepository...
		return nil, err
	}

	notification.NotifyCreateRepository(doer, owner, repo)

	return repo, nil
}

// DeleteRepository deletes a repository for a user or organization.
func DeleteRepository(ctx context.Context, doer *user_model.User, repo *repo_model.Repository, notify bool) error {
	if err := pull_service.CloseRepoBranchesPulls(ctx, doer, repo); err != nil {
		log.Error("CloseRepoBranchesPulls failed: %v", err)
	}

	if notify {
		// If the repo itself has webhooks, we need to trigger them before deleting it...
		notification.NotifyDeleteRepository(doer, repo)
	}

	if err := repoman.DeleteRepository(doer, repo.OwnerID, repo.ID); err != nil {
		return err
	}

	return packages_model.UnlinkRepositoryFromAllPackages(ctx, repo.ID)
}

// PushCreateRepo creates a repository when a new repository is pushed to an appropriate namespace
func PushCreateRepo(authUser, owner *user_model.User, repoName string) (*repo_model.Repository, error) {
	if !authUser.IsAdmin {
		if owner.IsOrganization() {
			if ok, err := organization.CanCreateOrgRepo(owner.ID, authUser.ID); err != nil {
				return nil, err
			} else if !ok {
				return nil, fmt.Errorf("cannot push-create repository for org")
			}
		} else if authUser.ID != owner.ID {
			return nil, fmt.Errorf("cannot push-create repository for another user")
		}
	}

	repo, err := CreateRepository(authUser, owner, repoman.CreateRepoOptions{
		Name:      repoName,
		IsPrivate: setting.Repository.DefaultPushCreatePrivate,
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// Init start repository service
func Init() error {
	repo_module.LoadRepoConfig()
	admin_model.RemoveAllWithNotice(db.DefaultContext, "Clean up temporary repository uploads", setting.Repository.Upload.TempPath)
	admin_model.RemoveAllWithNotice(db.DefaultContext, "Clean up temporary repositories", repo_module.LocalCopyPath())
	return initPushQueue()
}

// UpdateRepository updates a repository
func UpdateRepository(repo *repo_model.Repository, visibilityChanged bool) (err error) {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if err = repo_module.UpdateRepository(ctx, repo, visibilityChanged); err != nil {
		return fmt.Errorf("updateRepository: %v", err)
	}

	return committer.Commit()
}

// LinkedRepository returns the linked repo if any
func LinkedRepository(a *repo_model.Attachment) (*repo_model.Repository, unit.Type, error) {
	if a.IssueID != 0 {
		iss, err := issues_model.GetIssueByID(db.DefaultContext, a.IssueID)
		if err != nil {
			return nil, unit.TypeIssues, err
		}
		repo, err := repo_model.GetRepositoryByID(iss.RepoID)
		unitType := unit.TypeIssues
		if iss.IsPull {
			unitType = unit.TypePullRequests
		}
		return repo, unitType, err
	} else if a.ReleaseID != 0 {
		rel, err := release.GetReleaseByID(db.DefaultContext, a.ReleaseID)
		if err != nil {
			return nil, unit.TypeReleases, err
		}
		repo, err := repo_model.GetRepositoryByID(rel.RepoID)
		return repo, unit.TypeReleases, err
	}
	return nil, -1, nil
}
