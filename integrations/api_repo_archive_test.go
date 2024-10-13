// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestAPIDownloadArchive(t *testing.T) {
	defer prepareTestEnv(t)()

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	user2 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	session := loginUser(t, user2.LowerName)
	token := getTokenForLoggedInUser(t, session)

	link, _ := url.Parse(fmt.Sprintf("/api/v1/repos/%s/%s/archive/master.zip", user2.Name, repo.Name))
	link.RawQuery = url.Values{"token": {token}}.Encode()
	resp := MakeRequest(t, NewRequest(t, "GET", link.String()), http.StatusOK)
	bs, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, 320, len(bs))

	link, _ = url.Parse(fmt.Sprintf("/api/v1/repos/%s/%s/archive/master.tar.gz", user2.Name, repo.Name))
	link.RawQuery = url.Values{"token": {token}}.Encode()
	resp = MakeRequest(t, NewRequest(t, "GET", link.String()), http.StatusOK)
	bs, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, 266, len(bs))

	link, _ = url.Parse(fmt.Sprintf("/api/v1/repos/%s/%s/archive/master.bundle", user2.Name, repo.Name))
	link.RawQuery = url.Values{"token": {token}}.Encode()
	resp = MakeRequest(t, NewRequest(t, "GET", link.String()), http.StatusOK)
	bs, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, 382, len(bs))

	link, _ = url.Parse(fmt.Sprintf("/api/v1/repos/%s/%s/archive/master", user2.Name, repo.Name))
	link.RawQuery = url.Values{"token": {token}}.Encode()
	MakeRequest(t, NewRequest(t, "GET", link.String()), http.StatusBadRequest)
}
