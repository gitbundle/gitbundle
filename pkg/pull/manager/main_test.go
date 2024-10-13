// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package pull_test

import (
	"path/filepath"
	"regexp"
	"testing"
	_ "unsafe"

	"github.com/gitbundle/modules/queue"
	"github.com/gitbundle/server/pkg/unittest"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

//go:linkname commitMessageTrailersPattern github.com/gitbundle/server/pkg/pull/manager.commitMessageTrailersPattern
var commitMessageTrailersPattern *regexp.Regexp

//go:linkname prPatchCheckerQueue github.com/gitbundle/server/pkg/pull/manager.prPatchCheckerQueue
var prPatchCheckerQueue queue.UniqueQueue
