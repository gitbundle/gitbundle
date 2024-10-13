// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"strings"

	"github.com/gitbundle/modules/cache"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/setting"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

func getRefName(fullRefName string) string {
	if strings.HasPrefix(fullRefName, git.TagPrefix) {
		return fullRefName[len(git.TagPrefix):]
	} else if strings.HasPrefix(fullRefName, git.BranchPrefix) {
		return fullRefName[len(git.BranchPrefix):]
	}
	return ""
}

// CacheRef cachhe last commit information of the branch or the tag
func CacheRef(ctx context.Context, repo *repo_model.Repository, gitRepo *git.Repository, fullRefName string) error {
	if !setting.CacheService.LastCommit.Enabled {
		return nil
	}

	commit, err := gitRepo.GetCommit(fullRefName)
	if err != nil {
		return err
	}

	commitsCount, err := cache.GetInt64(repo.GetCommitsCountCacheKey(getRefName(fullRefName), true), commit.CommitsCount)
	if err != nil {
		return err
	}
	if commitsCount < setting.CacheService.LastCommit.CommitsCount {
		return nil
	}

	commitCache := git.NewLastCommitCache(repo.FullName(), gitRepo, setting.LastCommitCacheTTLSeconds, cache.GetCache())

	return commitCache.CacheCommit(ctx, commit)
}
