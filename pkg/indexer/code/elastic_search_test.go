// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package code

import (
	"os"
	"testing"

	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestESIndexAndSearch(t *testing.T) {
	unittest.PrepareTestEnv(t)

	u := os.Getenv("TEST_INDEXER_CODE_ES_URL")
	if u == "" {
		t.SkipNow()
		return
	}

	indexer, _, err := NewElasticSearchIndexer(u, "gitea_codes")
	if err != nil {
		assert.Fail(t, "Unable to create ES indexer Error: %v", err)
		if indexer != nil {
			indexer.Close()
		}
		return
	}
	defer indexer.Close()

	testIndexer("elastic_search", t, indexer)
}

func TestIndexPos(t *testing.T) {
	startIdx, endIdx := indexPos("test index start and end", "start", "end")
	assert.EqualValues(t, 11, startIdx)
	assert.EqualValues(t, 24, endIdx)
}