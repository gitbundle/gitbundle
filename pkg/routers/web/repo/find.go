// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"net/http"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/server/pkg/context"
)

const (
	tplFindFiles base.TplName = "repo/find/files"
)

// FindFiles render the page to find repository files
func FindFiles(ctx *context.Context) {
	path := ctx.Params("*")
	ctx.Data["TreeLink"] = ctx.Repo.RepoLink + "/src/" + path
	ctx.Data["DataLink"] = ctx.Repo.RepoLink + "/tree-list/" + path
	ctx.HTML(http.StatusOK, tplFindFiles)
}
