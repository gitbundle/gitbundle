// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/db"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestStarRepo(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	const userID = 2
	const repoID = 1
	unittest.AssertNotExistsBean(t, &repo_model.Star{UID: userID, RepoID: repoID})
	assert.NoError(t, repo_model.StarRepo(userID, repoID, true))
	unittest.AssertExistsAndLoadBean(t, &repo_model.Star{UID: userID, RepoID: repoID})
	assert.NoError(t, repo_model.StarRepo(userID, repoID, true))
	unittest.AssertExistsAndLoadBean(t, &repo_model.Star{UID: userID, RepoID: repoID})
	assert.NoError(t, repo_model.StarRepo(userID, repoID, false))
	unittest.AssertNotExistsBean(t, &repo_model.Star{UID: userID, RepoID: repoID})
}

func TestIsStaring(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	assert.True(t, repo_model.IsStaring(db.DefaultContext, 2, 4))
	assert.False(t, repo_model.IsStaring(db.DefaultContext, 3, 4))
}

func TestRepository_GetStargazers(t *testing.T) {
	// repo with stargazers
	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 4}).(*repo_model.Repository)
	gazers, err := repo_model.GetStargazers(repo, db.ListOptions{Page: 0})
	assert.NoError(t, err)
	if assert.Len(t, gazers, 1) {
		assert.Equal(t, int64(2), gazers[0].ID)
	}
}

func TestRepository_GetStargazers2(t *testing.T) {
	// repo with stargazers
	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 3}).(*repo_model.Repository)
	gazers, err := repo_model.GetStargazers(repo, db.ListOptions{Page: 0})
	assert.NoError(t, err)
	assert.Len(t, gazers, 0)
}
