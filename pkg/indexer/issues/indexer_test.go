// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issues_test

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"sync"
	"testing"
	"time"
	_ "unsafe"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/indexer/issues"
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

func TestBleveSearchIssues(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	setting.Cfg = ini.Empty()

	tmpIndexerDir, err := os.MkdirTemp("", "issues-indexer")
	if err != nil {
		assert.Fail(t, "Unable to create temporary directory: %v", err)
		return
	}

	setting.Cfg.Section("queue.issue_indexer").Key("DATADIR").MustString(path.Join(tmpIndexerDir, "issues.queue"))

	oldIssuePath := setting.Indexer.IssuePath
	setting.Indexer.IssuePath = path.Join(tmpIndexerDir, "issues.queue")
	defer func() {
		setting.Indexer.IssuePath = oldIssuePath
		util.RemoveAll(tmpIndexerDir)
	}()

	setting.Indexer.IssueType = "bleve"
	setting.NewQueueService()
	issues.InitIssueIndexer(true)
	defer func() {
		indexer := holder.get()
		if bleveIndexer, ok := indexer.(*issues.BleveIndexer); ok {
			bleveIndexer.Close()
		}
	}()

	time.Sleep(5 * time.Second)

	ids, err := issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "issue2")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{2}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "first")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{1}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "for")
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int64{1, 2, 3, 5, 11}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "good")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{1}, ids)
}

func TestDBSearchIssues(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	setting.Indexer.IssueType = "db"
	issues.InitIssueIndexer(true)

	ids, err := issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "issue2")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{2}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "first")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{1}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "for")
	assert.NoError(t, err)
	assert.ElementsMatch(t, []int64{1, 2, 3, 5, 11}, ids)

	ids, err = issues.SearchIssuesByKeyword(context.TODO(), []int64{1}, "good")
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{1}, ids)
}

type indexerHolder struct {
	indexer   issues.Indexer
	mutex     sync.RWMutex
	cond      *sync.Cond
	cancelled bool
}

//go:linkname (*indexerHolder).get github.com/gitbundle/server/pkg/indexer/issues.(*indexerHolder).get
func (h *indexerHolder) get() issues.Indexer

//go:linkname holder github.com/gitbundle/server/pkg/indexer/issues.holder
var holder indexerHolder
