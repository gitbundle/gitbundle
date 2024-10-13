// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"net/http"
	"time"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/server/pkg/context"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/repository_manager/files"
	"github.com/gitbundle/server/pkg/web"
)

// ApplyDiffPatch handles API call for applying a patch
func ApplyDiffPatch(ctx *context.APIContext) {
	// swagger:operation POST /repos/{owner}/{repo}/diffpatch repository repoApplyDiffPatch
	// ---
	// summary: Apply diff patch to repository
	// consumes:
	// - application/json
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
	// - name: body
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/UpdateFileOptions"
	// responses:
	//   "200":
	//     "$ref": "#/responses/FileResponse"
	apiOpts := web.GetForm(ctx).(*api.ApplyDiffPatchFileOptions)

	opts := &files.ApplyDiffPatchOptions{
		Content:   apiOpts.Content,
		SHA:       apiOpts.SHA,
		Message:   apiOpts.Message,
		OldBranch: apiOpts.BranchName,
		NewBranch: apiOpts.NewBranchName,
		Committer: &files.IdentityOptions{
			Name:  apiOpts.Committer.Name,
			Email: apiOpts.Committer.Email,
		},
		Author: &files.IdentityOptions{
			Name:  apiOpts.Author.Name,
			Email: apiOpts.Author.Email,
		},
		Dates: &files.CommitDateOptions{
			Author:    apiOpts.Dates.Author,
			Committer: apiOpts.Dates.Committer,
		},
		Signoff: apiOpts.Signoff,
	}
	if opts.Dates.Author.IsZero() {
		opts.Dates.Author = time.Now()
	}
	if opts.Dates.Committer.IsZero() {
		opts.Dates.Committer = time.Now()
	}

	if opts.Message == "" {
		opts.Message = "apply-patch"
	}

	if !canWriteFiles(ctx, apiOpts.BranchName) {
		ctx.Error(http.StatusInternalServerError, "ApplyPatch", repo_model.ErrUserDoesNotHaveAccessToRepo{
			UserID:   ctx.Doer.ID,
			RepoName: ctx.Repo.Repository.LowerName,
		})
		return
	}

	fileResponse, err := files.ApplyDiffPatch(ctx, ctx.Repo.Repository, ctx.Doer, opts)
	if err != nil {
		if repo_model.IsErrUserCannotCommit(err) || repo_model.IsErrFilePathProtected(err) {
			ctx.Error(http.StatusForbidden, "Access", err)
			return
		}
		if repo_model.IsErrBranchAlreadyExists(err) || repo_model.IsErrFilenameInvalid(err) || repo_model.IsErrSHADoesNotMatch(err) ||
			repo_model.IsErrFilePathInvalid(err) || repo_model.IsErrRepoFileAlreadyExists(err) {
			ctx.Error(http.StatusUnprocessableEntity, "Invalid", err)
			return
		}
		if repo_model.IsErrBranchDoesNotExist(err) || git.IsErrBranchNotExist(err) {
			ctx.Error(http.StatusNotFound, "BranchDoesNotExist", err)
			return
		}
		ctx.Error(http.StatusInternalServerError, "ApplyPatch", err)
	} else {
		ctx.JSON(http.StatusCreated, fileResponse)
	}
}
