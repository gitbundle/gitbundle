// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	repo_model "github.com/gitbundle/server/pkg/repo"
	files_service "github.com/gitbundle/server/pkg/repo/repository_manager/files"
	user_model "github.com/gitbundle/server/pkg/user"
)

func createFileInBranch(user *user_model.User, repo *repo_model.Repository, treePath, branchName, content string) (*api.FileResponse, error) {
	opts := &files_service.UpdateRepoFileOptions{
		OldBranch: branchName,
		TreePath:  treePath,
		Content:   content,
		IsNewFile: true,
		Author:    nil,
		Committer: nil,
	}
	return files_service.CreateOrUpdateRepoFile(git.DefaultContext, repo, user, opts)
}

func createFile(user *user_model.User, repo *repo_model.Repository, treePath string) (*api.FileResponse, error) {
	return createFileInBranch(user, repo, treePath, repo.DefaultBranch, "This is a NEW file")
}
