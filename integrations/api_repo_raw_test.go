// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"

	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"
)

func TestAPIReposRaw(t *testing.T) {
	defer prepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session)

	for _, ref := range [...]string{
		"master", // Branch
		"v1.1",   // Tag
		"65f1bf27bc3bf70f64657658635e66094edbcb4d", // Commit
	} {
		req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/raw/%s/README.md?token="+token, user.Name, ref)
		session.MakeRequest(t, req, http.StatusOK)
	}
	// Test default branch
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/raw/README.md?token="+token, user.Name)
	session.MakeRequest(t, req, http.StatusOK)
}
