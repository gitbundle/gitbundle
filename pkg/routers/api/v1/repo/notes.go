// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"fmt"
	"net/http"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/validation"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/convert"
)

// GetNote Get a note corresponding to a single commit from a repository
func GetNote(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/git/notes/{sha} repository repoGetNote
	// ---
	// summary: Get a note corresponding to a single commit from a repository
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
	// - name: sha
	//   in: path
	//   description: a git ref or commit sha
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/Note"
	//   "422":
	//     "$ref": "#/responses/validationError"
	//   "404":
	//     "$ref": "#/responses/notFound"

	sha := ctx.Params(":sha")
	if (validation.GitRefNamePatternInvalid.MatchString(sha) || !validation.CheckGitRefAdditionalRulesValid(sha)) && !git.SHAPattern.MatchString(sha) {
		ctx.Error(http.StatusUnprocessableEntity, "no valid ref or sha", fmt.Sprintf("no valid ref or sha: %s", sha))
		return
	}
	getNote(ctx, sha)
}

func getNote(ctx *context.APIContext, identifier string) {
	if ctx.Repo.GitRepo == nil {
		ctx.InternalServerError(fmt.Errorf("no open git repo"))
		return
	}

	var note git.Note
	if err := git.GetNote(ctx, ctx.Repo.GitRepo, identifier, &note); err != nil {
		if git.IsErrNotExist(err) {
			ctx.NotFound(identifier)
			return
		}
		ctx.Error(http.StatusInternalServerError, "GetNote", err)
		return
	}

	cmt, err := convert.ToCommit(ctx.Repo.Repository, ctx.Repo.GitRepo, note.Commit, nil)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "ToCommit", err)
		return
	}
	apiNote := api.Note{Message: string(note.Message), Commit: cmt}
	ctx.JSON(http.StatusOK, apiNote)
}
