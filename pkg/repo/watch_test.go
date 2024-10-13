// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo_test

import (
	"testing"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/db"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestIsWatching(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	assert.True(t, repo_model.IsWatching(1, 1))
	assert.True(t, repo_model.IsWatching(4, 1))
	assert.True(t, repo_model.IsWatching(11, 1))

	assert.False(t, repo_model.IsWatching(1, 5))
	assert.False(t, repo_model.IsWatching(8, 1))
	assert.False(t, repo_model.IsWatching(unittest.NonexistentID, unittest.NonexistentID))
}

func TestGetWatchers(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	watches, err := repo_model.GetWatchers(db.DefaultContext, repo.ID)
	assert.NoError(t, err)
	// One watchers are inactive, thus minus 1
	assert.Len(t, watches, repo.NumWatches-1)
	for _, watch := range watches {
		assert.EqualValues(t, repo.ID, watch.RepoID)
	}

	watches, err = repo_model.GetWatchers(db.DefaultContext, unittest.NonexistentID)
	assert.NoError(t, err)
	assert.Len(t, watches, 0)
}

func TestRepository_GetWatchers(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	watchers, err := repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, repo.NumWatches)
	for _, watcher := range watchers {
		unittest.AssertExistsAndLoadBean(t, &repo_model.Watch{UserID: watcher.ID, RepoID: repo.ID})
	}

	repo = unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 9}).(*repo_model.Repository)
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, 0)
}

func TestWatchIfAuto(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	watchers, err := repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, repo.NumWatches)

	setting.Service.AutoWatchOnChanges = false

	prevCount := repo.NumWatches

	// Must not add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 8, 1, true))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)

	// Should not add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 10, 1, true))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)

	setting.Service.AutoWatchOnChanges = true

	// Must not add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 8, 1, true))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)

	// Should not add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 12, 1, false))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)

	// Should add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 12, 1, true))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount+1)

	// Should remove watch, inhibit from adding auto
	assert.NoError(t, repo_model.WatchRepo(db.DefaultContext, 12, 1, false))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)

	// Must not add watch
	assert.NoError(t, repo_model.WatchIfAuto(db.DefaultContext, 12, 1, true))
	watchers, err = repo_model.GetRepoWatchers(repo.ID, db.ListOptions{Page: 1})
	assert.NoError(t, err)
	assert.Len(t, watchers, prevCount)
}

func TestWatchRepoMode(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1}, 0)

	assert.NoError(t, repo_model.WatchRepoMode(12, 1, repo_model.WatchModeAuto))
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1}, 1)
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1, Mode: repo_model.WatchModeAuto}, 1)

	assert.NoError(t, repo_model.WatchRepoMode(12, 1, repo_model.WatchModeNormal))
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1}, 1)
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1, Mode: repo_model.WatchModeNormal}, 1)

	assert.NoError(t, repo_model.WatchRepoMode(12, 1, repo_model.WatchModeDont))
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1}, 1)
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1, Mode: repo_model.WatchModeDont}, 1)

	assert.NoError(t, repo_model.WatchRepoMode(12, 1, repo_model.WatchModeNone))
	unittest.AssertCount(t, &repo_model.Watch{UserID: 12, RepoID: 1}, 0)
}
