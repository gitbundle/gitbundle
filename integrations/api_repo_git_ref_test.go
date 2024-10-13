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

func TestAPIReposGitRefs(t *testing.T) {
	defer prepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session)

	for _, ref := range [...]string{
		"refs/heads/master", // Branch
		"refs/tags/v1.1",    // Tag
	} {
		req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/%s?token="+token, user.Name, ref)
		session.MakeRequest(t, req, http.StatusOK)
	}
	// Test getting all refs
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/refs?token="+token, user.Name)
	session.MakeRequest(t, req, http.StatusOK)
	// Test getting non-existent refs
	req = NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/refs/heads/unknown?token="+token, user.Name)
	session.MakeRequest(t, req, http.StatusNotFound)
}
