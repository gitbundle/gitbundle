// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"net/url"
	"testing"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/routers"

	"github.com/stretchr/testify/assert"
)

func TestNodeinfo(t *testing.T) {
	setting.Federation.Enabled = true
	c = routers.NormalRoutes()
	defer func() {
		setting.Federation.Enabled = false
		c = routers.NormalRoutes()
	}()

	onGiteaRun(t, func(*testing.T, *url.URL) {
		req := NewRequestf(t, "GET", "/api/v1/nodeinfo")
		resp := MakeRequest(t, req, http.StatusOK)
		var nodeinfo api.NodeInfo
		DecodeJSON(t, resp, &nodeinfo)
		assert.True(t, nodeinfo.OpenRegistrations)
		assert.Equal(t, "gitea", nodeinfo.Software.Name)
		assert.Equal(t, 23, nodeinfo.Usage.Users.Total)
		assert.Equal(t, 17, nodeinfo.Usage.LocalPosts)
		assert.Equal(t, 2, nodeinfo.Usage.LocalComments)
	})
}
