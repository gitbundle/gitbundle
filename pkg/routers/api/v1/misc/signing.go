// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package misc

import (
	"fmt"
	"net/http"

	asymkey_service "github.com/gitbundle/server/pkg/asymkey/manager"
	"github.com/gitbundle/server/pkg/context"
)

// SigningKey returns the public key of the default signing key if it exists
func SigningKey(ctx *context.APIContext) {
	// swagger:operation GET /signing-key.gpg miscellaneous getSigningKey
	// ---
	// summary: Get default signing-key.gpg
	// produces:
	//     - text/plain
	// responses:
	//   "200":
	//     description: "GPG armored public key"
	//     schema:
	//       type: string

	// swagger:operation GET /repos/{owner}/{repo}/signing-key.gpg repository repoSigningKey
	// ---
	// summary: Get signing-key.gpg for given repository
	// produces:
	//     - text/plain
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     description: "GPG armored public key"
	//     schema:
	//       type: string

	path := ""
	if ctx.Repo != nil && ctx.Repo.Repository != nil {
		path = ctx.Repo.Repository.RepoPath()
	}

	content, err := asymkey_service.PublicSigningKey(ctx, path)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "gpg export", err)
		return
	}
	_, err = ctx.Write([]byte(content))
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "gpg export", fmt.Errorf("Error writing key content %v", err))
	}
}
