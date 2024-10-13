// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package gitdiff_test

import (
	"path/filepath"
	"testing"
	_ "unsafe"

	"github.com/gitbundle/server/pkg/git/gitdiff"
	"github.com/gitbundle/server/pkg/unittest"

	_ "github.com/gitbundle/server/pkg"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

//go:linkname diffToHTML github.com/gitbundle/server/pkg/services/gitdiff.diffToHTML
func diffToHTML(fileName string, diffs []diffmatchpatch.Diff, lineType gitdiff.DiffLineType) gitdiff.DiffInline
