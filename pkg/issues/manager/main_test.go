// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issue_test

import (
	"path/filepath"
	"testing"
	_ "unsafe"

	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/unittest"

	_ "github.com/gitbundle/server/pkg/access_token"
	_ "github.com/gitbundle/server/pkg/action"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

//go:linkname deleteIssue github.com/gitbundle/server/pkg/issues/manager.deleteIssue
func deleteIssue(issue *issues_model.Issue) error
