// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/log"
	git_model "github.com/gitbundle/server/pkg/git"
	"github.com/gitbundle/server/pkg/notification"
	pull_service "github.com/gitbundle/server/pkg/pull/manager"
	repo_model "github.com/gitbundle/server/pkg/repo"
	repo_module "github.com/gitbundle/server/pkg/repo/repository"
	user_model "github.com/gitbundle/server/pkg/user"
)

// CreateNewBranch creates a new repository branch
func CreateNewBranch(ctx context.Context, doer *user_model.User, repo *repo_model.Repository, oldBranchName, branchName string) (err error) {
	// Check if branch name can be used
	if err := checkBranchName(ctx, repo, branchName); err != nil {
		return err
	}

	if !git.IsBranchExist(ctx, repo.RepoPath(), oldBranchName) {
		return repo_model.ErrBranchDoesNotExist{
			BranchName: oldBranchName,
		}
	}

	if err := git.Push(ctx, repo.RepoPath(), git.PushOptions{
		Remote: repo.RepoPath(),
		Branch: fmt.Sprintf("%s%s:%s%s", git.BranchPrefix, oldBranchName, git.BranchPrefix, branchName),
		Env:    repo_module.PushingEnvironment(doer, repo),
	}); err != nil {
		if git.IsErrPushOutOfDate(err) || git.IsErrPushRejected(err) {
			return err
		}
		return fmt.Errorf("Push: %v", err)
	}

	return nil
}

// GetBranches returns branches from the repository, skipping skip initial branches and
// returning at most limit branches, or all branches if limit is 0.
func GetBranches(ctx context.Context, repo *repo_model.Repository, skip, limit int) ([]*git.Branch, int, error) {
	return git.GetBranchesByPath(ctx, repo.RepoPath(), skip, limit)
}

// checkBranchName validates branch name with existing repository branches
func checkBranchName(ctx context.Context, repo *repo_model.Repository, name string) error {
	_, err := git.WalkReferences(ctx, repo.RepoPath(), func(_, refName string) error {
		branchRefName := strings.TrimPrefix(refName, git.BranchPrefix)
		switch {
		case branchRefName == name:
			return repo_model.ErrBranchAlreadyExists{
				BranchName: name,
			}
		// If branchRefName like a/b but we want to create a branch named a then we have a conflict
		case strings.HasPrefix(branchRefName, name+"/"):
			return repo_model.ErrBranchNameConflict{
				BranchName: branchRefName,
			}
			// Conversely if branchRefName like a but we want to create a branch named a/b then we also have a conflict
		case strings.HasPrefix(name, branchRefName+"/"):
			return repo_model.ErrBranchNameConflict{
				BranchName: branchRefName,
			}
		case refName == git.TagPrefix+name:
			return repo_model.ErrTagAlreadyExists{
				TagName: name,
			}
		}
		return nil
	})

	return err
}

// CreateNewBranchFromCommit creates a new repository branch
func CreateNewBranchFromCommit(ctx context.Context, doer *user_model.User, repo *repo_model.Repository, commit, branchName string) (err error) {
	// Check if branch name can be used
	if err := checkBranchName(ctx, repo, branchName); err != nil {
		return err
	}

	if err := git.Push(ctx, repo.RepoPath(), git.PushOptions{
		Remote: repo.RepoPath(),
		Branch: fmt.Sprintf("%s:%s%s", commit, git.BranchPrefix, branchName),
		Env:    repo_module.PushingEnvironment(doer, repo),
	}); err != nil {
		if git.IsErrPushOutOfDate(err) || git.IsErrPushRejected(err) {
			return err
		}
		return fmt.Errorf("Push: %v", err)
	}

	return nil
}

// RenameBranch rename a branch
func RenameBranch(repo *repo_model.Repository, doer *user_model.User, gitRepo *git.Repository, from, to string) (string, error) {
	if from == to {
		return "target_exist", nil
	}

	if gitRepo.IsBranchExist(to) {
		return "target_exist", nil
	}

	if !gitRepo.IsBranchExist(from) {
		return "from_not_exist", nil
	}

	if err := git_model.RenameBranch(repo, from, to, func(isDefault bool) error {
		err2 := gitRepo.RenameBranch(from, to)
		if err2 != nil {
			return err2
		}

		if isDefault {
			err2 = gitRepo.SetDefaultBranch(to)
			if err2 != nil {
				return err2
			}
		}

		return nil
	}); err != nil {
		return "", err
	}
	refID, err := gitRepo.GetRefCommitID(git.BranchPrefix + to)
	if err != nil {
		return "", err
	}

	notification.NotifyDeleteRef(doer, repo, "branch", git.BranchPrefix+from)
	notification.NotifyCreateRef(doer, repo, "branch", git.BranchPrefix+to, refID)

	return "", nil
}

// enmuerates all branch related errors
var (
	ErrBranchIsDefault   = errors.New("branch is default")
	ErrBranchIsProtected = errors.New("branch is protected")
)

// DeleteBranch delete branch
func DeleteBranch(doer *user_model.User, repo *repo_model.Repository, gitRepo *git.Repository, branchName string) error {
	if branchName == repo.DefaultBranch {
		return ErrBranchIsDefault
	}

	isProtected, err := git_model.IsProtectedBranch(repo.ID, branchName)
	if err != nil {
		return err
	}

	if isProtected {
		return ErrBranchIsProtected
	}

	commit, err := gitRepo.GetBranchCommit(branchName)
	if err != nil {
		return err
	}

	if err := gitRepo.DeleteBranch(branchName, git.DeleteBranchOptions{
		Force: true,
	}); err != nil {
		return err
	}

	if err := pull_service.CloseBranchPulls(doer, repo.ID, branchName); err != nil {
		return err
	}

	// Don't return error below this
	if err := PushUpdate(
		&repo_module.PushUpdateOptions{
			RefFullName:  git.BranchPrefix + branchName,
			OldCommitID:  commit.ID.String(),
			NewCommitID:  git.EmptySHA,
			PusherID:     doer.ID,
			PusherName:   doer.Name,
			RepoUserName: repo.OwnerName,
			RepoName:     repo.Name,
		}); err != nil {
		log.Error("Update: %v", err)
	}

	if err := git_model.AddDeletedBranch(repo.ID, branchName, commit.ID.String(), doer.ID); err != nil {
		log.Warn("AddDeletedBranch: %v", err)
	}

	return nil
}
