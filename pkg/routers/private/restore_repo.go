// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package private

import (
	"io"
	"net/http"

	"github.com/gitbundle/modules/json"
	myCtx "github.com/gitbundle/server/pkg/context"
	migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/repo/private"
)

// RestoreRepo restore a repository from data
func RestoreRepo(ctx *myCtx.PrivateContext) {
	bs, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
		return
	}
	params := struct {
		RepoDir    string
		OwnerName  string
		RepoName   string
		Units      []string
		Validation bool
	}{}
	if err = json.Unmarshal(bs, &params); err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
		return
	}

	if err := migrations.RestoreRepository(
		ctx,
		params.RepoDir,
		params.OwnerName,
		params.RepoName,
		params.Units,
		params.Validation,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}
