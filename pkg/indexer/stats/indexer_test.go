// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package stats_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/gitbundle/modules/queue"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/indexer/stats"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	_ "github.com/gitbundle/server/pkg"

	"github.com/stretchr/testify/assert"
	"gopkg.in/ini.v1"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", "..", ".."),
	})
}

func TestRepoStatsIndex(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	setting.Cfg = ini.Empty()

	setting.NewQueueService()

	err := stats.Init()
	assert.NoError(t, err)

	repo, err := repo_model.GetRepositoryByID(1)
	assert.NoError(t, err)

	err = stats.UpdateRepoIndexer(repo)
	assert.NoError(t, err)

	queue.GetManager().FlushAll(context.Background(), 5*time.Second)

	status, err := repo_model.GetIndexerStatus(db.DefaultContext, repo, repo_model.RepoIndexerTypeStats)
	assert.NoError(t, err)
	assert.Equal(t, "65f1bf27bc3bf70f64657658635e66094edbcb4d", status.CommitSha)
	langs, err := repo_model.GetTopLanguageStats(repo, 5)
	assert.NoError(t, err)
	assert.Empty(t, langs)
}
