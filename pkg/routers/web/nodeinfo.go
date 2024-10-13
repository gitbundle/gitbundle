// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"net/http"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/context"
)

type nodeInfoLinks struct {
	Links []nodeInfoLink `json:"links"`
}

type nodeInfoLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// NodeInfoLinks returns links to the node info endpoint
func NodeInfoLinks(ctx *context.Context) {
	nodeinfolinks := &nodeInfoLinks{
		Links: []nodeInfoLink{{
			fmt.Sprintf("%sapi/v1/nodeinfo", setting.AppURL),
			"http://nodeinfo.diaspora.software/ns/schema/2.1",
		}},
	}
	ctx.JSON(http.StatusOK, nodeinfolinks)
}
