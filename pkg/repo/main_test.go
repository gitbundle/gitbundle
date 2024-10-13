// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo_test

import (
	"path/filepath"
	"testing"

	"github.com/gitbundle/server/pkg/unittest"

	_ "github.com/gitbundle/server/pkg"             // register table model
	_ "github.com/gitbundle/server/pkg/perm/access" // register table model
	_ "github.com/gitbundle/server/pkg/repo"        // register table model
	_ "github.com/gitbundle/server/pkg/user"        // register table model
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}
