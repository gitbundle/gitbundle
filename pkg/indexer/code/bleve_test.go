// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package code

import (
	"os"
	"testing"

	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestBleveIndexAndSearch(t *testing.T) {
	unittest.PrepareTestEnv(t)

	dir, err := os.MkdirTemp("", "bleve.index")
	assert.NoError(t, err)
	if err != nil {
		assert.Fail(t, "Unable to create temporary directory")
		return
	}
	defer util.RemoveAll(dir)

	idx, _, err := NewBleveIndexer(dir)
	if err != nil {
		assert.Fail(t, "Unable to create bleve indexer Error: %v", err)
		if idx != nil {
			idx.Close()
		}
		return
	}
	defer idx.Close()

	testIndexer("beleve", t, idx)
}
