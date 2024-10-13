// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"net/http"

	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/convert"
	release_model "github.com/gitbundle/server/pkg/repo/release"
	releaseservice "github.com/gitbundle/server/pkg/repo/release_manager"
)

// GetReleaseByTag get a single release of a repository by tag name
func GetReleaseByTag(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/releases/tags/{tag} repository repoGetReleaseByTag
	// ---
	// summary: Get a release by tag name
	// produces:
	// - application/json
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
	// - name: tag
	//   in: path
	//   description: tag name of the release to get
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/Release"
	//   "404":
	//     "$ref": "#/responses/notFound"

	tag := ctx.Params(":tag")

	release, err := release_model.GetRelease(ctx.Repo.Repository.ID, tag)
	if err != nil {
		if release_model.IsErrReleaseNotExist(err) {
			ctx.NotFound()
			return
		}
		ctx.Error(http.StatusInternalServerError, "GetRelease", err)
		return
	}

	if release.IsTag {
		ctx.NotFound()
		return
	}

	if err = release.LoadAttributes(); err != nil {
		ctx.Error(http.StatusInternalServerError, "LoadAttributes", err)
		return
	}
	ctx.JSON(http.StatusOK, convert.ToRelease(release))
}

// DeleteReleaseByTag delete a release from a repository by tag name
func DeleteReleaseByTag(ctx *context.APIContext) {
	// swagger:operation DELETE /repos/{owner}/{repo}/releases/tags/{tag} repository repoDeleteReleaseByTag
	// ---
	// summary: Delete a release by tag name
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
	// - name: tag
	//   in: path
	//   description: tag name of the release to delete
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//     "$ref": "#/responses/empty"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "405":
	//     "$ref": "#/responses/empty"

	tag := ctx.Params(":tag")

	release, err := release_model.GetRelease(ctx.Repo.Repository.ID, tag)
	if err != nil {
		if release_model.IsErrReleaseNotExist(err) {
			ctx.NotFound()
			return
		}
		ctx.Error(http.StatusInternalServerError, "GetRelease", err)
		return
	}

	if release.IsTag {
		ctx.NotFound()
		return
	}

	if err = releaseservice.DeleteReleaseByID(ctx, release.ID, ctx.Doer, false); err != nil {
		if release_model.IsErrProtectedTagName(err) {
			ctx.Error(http.StatusMethodNotAllowed, "delTag", "user not allowed to delete protected tag")
			return
		}
		ctx.Error(http.StatusInternalServerError, "DeleteReleaseByID", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
