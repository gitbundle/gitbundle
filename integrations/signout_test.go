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

func TestSignOut(t *testing.T) {
	defer prepareTestEnv(t)()

	session := loginUser(t, "user2")

	req := NewRequest(t, "POST", "/user/logout")
	session.MakeRequest(t, req, http.StatusSeeOther)

	// try to view a private repo, should fail
	req = NewRequest(t, "GET", "/user2/repo2")
	session.MakeRequest(t, req, http.StatusNotFound)

	// invalidate cached cookies for user2, for subsequent tests
	delete(loginSessionCache, "user2")
}
