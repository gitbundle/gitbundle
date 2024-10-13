// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package web

import (
	"net/http"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/context"
)

// tplSwaggerV1Json swagger v1 json template
const tplSwaggerV1Json base.TplName = "swagger/v1_json"

// SwaggerV1Json render swagger v1 json
func SwaggerV1Json(ctx *context.Context) {
	t := ctx.Render.TemplateLookup(string(tplSwaggerV1Json))
	ctx.Resp.Header().Set("Content-Type", "application/json")
	if err := t.Execute(ctx.Resp, ctx.Data); err != nil {
		log.Error("%v", err)
		ctx.Error(http.StatusInternalServerError)
	}
}
