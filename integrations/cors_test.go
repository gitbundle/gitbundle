// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSNotSet(t *testing.T) {
	defer prepareTestEnv(t)()
	req := NewRequestf(t, "GET", "/api/v1/version")
	session := loginUser(t, "user2")
	resp := session.MakeRequest(t, req, http.StatusOK)
	assert.Equal(t, resp.Code, http.StatusOK)
	corsHeader := resp.Header().Get("Access-Control-Allow-Origin")
	assert.Equal(t, corsHeader, "", "Access-Control-Allow-Origin: generated header should match") // header not set
}
