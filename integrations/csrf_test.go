// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestCsrfProtection(t *testing.T) {
	defer prepareTestEnv(t)()

	// test web form csrf via form
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	session := loginUser(t, user.Name)
	req := NewRequestWithValues(t, "POST", "/user/settings", map[string]string{
		"_csrf": "fake_csrf",
	})
	session.MakeRequest(t, req, http.StatusSeeOther)

	resp := session.MakeRequest(t, req, http.StatusSeeOther)
	loc := resp.Header().Get("Location")
	assert.Equal(t, setting.AppSubURL+"/", loc)
	resp = session.MakeRequest(t, NewRequest(t, "GET", loc), http.StatusOK)
	htmlDoc := NewHTMLParser(t, resp.Body)
	assert.Equal(t, "Bad Request: invalid CSRF token",
		strings.TrimSpace(htmlDoc.doc.Find(".ui.message").Text()),
	)

	// test web form csrf via header. TODO: should use an UI api to test
	req = NewRequest(t, "POST", "/user/settings")
	req.Header.Add("X-Csrf-Token", "fake_csrf")
	session.MakeRequest(t, req, http.StatusSeeOther)

	resp = session.MakeRequest(t, req, http.StatusSeeOther)
	loc = resp.Header().Get("Location")
	assert.Equal(t, setting.AppSubURL+"/", loc)
	resp = session.MakeRequest(t, NewRequest(t, "GET", loc), http.StatusOK)
	htmlDoc = NewHTMLParser(t, resp.Body)
	assert.Equal(t, "Bad Request: invalid CSRF token",
		strings.TrimSpace(htmlDoc.doc.Find(".ui.message").Text()),
	)
}
