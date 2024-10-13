// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"
)

func TestExploreRepos(t *testing.T) {
	defer prepareTestEnv(t)()

	req := NewRequest(t, "GET", "/explore/repos")
	MakeRequest(t, req, http.StatusOK)
}
