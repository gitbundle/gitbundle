// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/url"
	"os"
	"testing"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/util"
	git_model "github.com/gitbundle/server/pkg/git"
	repo_model "github.com/gitbundle/server/pkg/repo"
	release_model "github.com/gitbundle/server/pkg/repo/release"
	release "github.com/gitbundle/server/pkg/repo/release_manager"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTagProtected(t *testing.T) {
	defer prepareTestEnv(t)()

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	owner := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: repo.OwnerID}).(*user_model.User)

	t.Run("API", func(t *testing.T) {
		defer PrintCurrentTest(t)()

		err := release.CreateNewTag(git.DefaultContext, owner, repo, "master", "v-1", "first tag")
		assert.NoError(t, err)

		err = git_model.InsertProtectedTag(&git_model.ProtectedTag{
			RepoID:      repo.ID,
			NamePattern: "v-*",
		})
		assert.NoError(t, err)
		err = git_model.InsertProtectedTag(&git_model.ProtectedTag{
			RepoID:           repo.ID,
			NamePattern:      "v-1.1",
			AllowlistUserIDs: []int64{repo.OwnerID},
		})
		assert.NoError(t, err)

		err = release.CreateNewTag(git.DefaultContext, owner, repo, "master", "v-2", "second tag")
		assert.Error(t, err)
		assert.True(t, release_model.IsErrProtectedTagName(err))

		err = release.CreateNewTag(git.DefaultContext, owner, repo, "master", "v-1.1", "third tag")
		assert.NoError(t, err)
	})

	t.Run("Git", func(t *testing.T) {
		onGiteaRun(t, func(t *testing.T, u *url.URL) {
			username := "user2"
			httpContext := NewAPITestContext(t, username, "repo1")

			dstPath, err := os.MkdirTemp("", httpContext.Reponame)
			assert.NoError(t, err)
			defer util.RemoveAll(dstPath)

			u.Path = httpContext.GitPath()
			u.User = url.UserPassword(username, userPassword)

			doGitClone(dstPath, u)(t)

			_, _, err = git.NewCommand(git.DefaultContext, "tag", "v-2").RunStdString(&git.RunOpts{Dir: dstPath})
			assert.NoError(t, err)

			_, _, err = git.NewCommand(git.DefaultContext, "push", "--tags").RunStdString(&git.RunOpts{Dir: dstPath})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "Tag v-2 is protected")
		})
	})

	// Cleanup
	releases, err := release_model.GetReleasesByRepoID(repo.ID, release_model.FindReleasesOptions{
		IncludeTags: true,
		TagNames:    []string{"v-1", "v-1.1"},
	})
	assert.NoError(t, err)

	for _, release := range releases {
		err = release_model.DeleteReleaseByID(release.ID)
		assert.NoError(t, err)
	}

	protectedTags, err := git_model.GetProtectedTags(repo.ID)
	assert.NoError(t, err)

	for _, protectedTag := range protectedTags {
		err = git_model.DeleteProtectedTag(protectedTag)
		assert.NoError(t, err)
	}
}
