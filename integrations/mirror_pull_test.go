// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"context"
	"testing"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
	repo_model "github.com/gitbundle/server/pkg/repo"
	mirror_service "github.com/gitbundle/server/pkg/repo/mirror"
	release_model "github.com/gitbundle/server/pkg/repo/release"
	release_service "github.com/gitbundle/server/pkg/repo/release_manager"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/repo/repository"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestMirrorPull(t *testing.T) {
	defer prepareTestEnv(t)()

	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	repoPath := repo_model.RepoPath(user.Name, repo.Name)

	opts := schemas.MigrateOptions{
		RepoName:    "test_mirror",
		Description: "Test mirror",
		Private:     false,
		Mirror:      true,
		CloneAddr:   repoPath,
		Wiki:        true,
		Releases:    false,
	}

	mirrorRepo, err := repository.CreateRepository(user, user, repoman.CreateRepoOptions{
		Name:        opts.RepoName,
		Description: opts.Description,
		IsPrivate:   opts.Private,
		IsMirror:    opts.Mirror,
		Status:      repo_model.RepositoryBeingMigrated,
	})
	assert.NoError(t, err)
	assert.True(t, mirrorRepo.IsMirror, "expected pull-mirror repo to be marked as a mirror immediately after its creation")

	ctx := context.Background()

	mirror, err := repository.MigrateRepositoryGitData(ctx, user, mirrorRepo, opts, nil)
	assert.NoError(t, err)

	gitRepo, err := git.OpenRepository(git.DefaultContext, repoPath)
	assert.NoError(t, err)
	defer gitRepo.Close()

	findOptions := release_model.FindReleasesOptions{IncludeDrafts: true, IncludeTags: true}
	initCount, err := release_model.GetReleaseCountByRepoID(mirror.ID, findOptions)
	assert.NoError(t, err)

	assert.NoError(t, release_service.CreateRelease(gitRepo, &release_model.Release{
		RepoID:       repo.ID,
		Repo:         repo,
		PublisherID:  user.ID,
		Publisher:    user,
		TagName:      "v0.2",
		Target:       "master",
		Title:        "v0.2 is released",
		Note:         "v0.2 is released",
		IsDraft:      false,
		IsPrerelease: false,
		IsTag:        true,
	}, nil, ""))

	_, err = repo_model.GetMirrorByRepoID(ctx, mirror.ID)
	assert.NoError(t, err)

	ok := mirror_service.SyncPullMirror(ctx, mirror.ID)
	assert.True(t, ok)

	count, err := release_model.GetReleaseCountByRepoID(mirror.ID, findOptions)
	assert.NoError(t, err)
	assert.EqualValues(t, initCount+1, count)

	release, err := release_model.GetRelease(repo.ID, "v0.2")
	assert.NoError(t, err)
	assert.NoError(t, release_service.DeleteReleaseByID(ctx, release.ID, user, true))

	ok = mirror_service.SyncPullMirror(ctx, mirror.ID)
	assert.True(t, ok)

	count, err = release_model.GetReleaseCountByRepoID(mirror.ID, findOptions)
	assert.NoError(t, err)
	assert.EqualValues(t, initCount, count)
}
