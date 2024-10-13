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
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/convert"
	"github.com/gitbundle/server/pkg/organization"
	"github.com/gitbundle/server/pkg/perm"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	repo_model "github.com/gitbundle/server/pkg/repo"
	repo_service "github.com/gitbundle/server/pkg/repo/repository_manager"
	"github.com/gitbundle/server/pkg/routers/api/v1/utils"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/web"
)

// ListForks list a repository's forks
func ListForks(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/forks repository listForks
	// ---
	// summary: List a repository's forks
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
	// - name: page
	//   in: query
	//   description: page number of results to return (1-based)
	//   type: integer
	// - name: limit
	//   in: query
	//   description: page size of results
	//   type: integer
	// responses:
	//   "200":
	//     "$ref": "#/responses/RepositoryList"

	forks, err := repo_model.GetForks(ctx.Repo.Repository, utils.GetListOptions(ctx))
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetForks", err)
		return
	}
	apiForks := make([]*api.Repository, len(forks))
	for i, fork := range forks {
		access, err := access_model.AccessLevel(ctx.Doer, fork)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "AccessLevel", err)
			return
		}
		apiForks[i] = convert.ToRepo(fork, access)
	}

	ctx.SetTotalCountHeader(int64(ctx.Repo.Repository.NumForks))
	ctx.JSON(http.StatusOK, apiForks)
}

// CreateFork create a fork of a repo
func CreateFork(ctx *context.APIContext) {
	// swagger:operation POST /repos/{owner}/{repo}/forks repository createFork
	// ---
	// summary: Fork a repository
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo to fork
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo to fork
	//   type: string
	//   required: true
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/CreateForkOption"
	// responses:
	//   "202":
	//     "$ref": "#/responses/Repository"
	//   "403":
	//     "$ref": "#/responses/forbidden"
	//   "409":
	//     description: The repository with the same name already exists.
	//   "422":
	//     "$ref": "#/responses/validationError"

	form := web.GetForm(ctx).(*api.CreateForkOption)
	repo := ctx.Repo.Repository
	var forker *user_model.User // user/org that will own the fork
	if form.Organization == nil {
		forker = ctx.Doer
	} else {
		org, err := organization.GetOrgByName(*form.Organization)
		if err != nil {
			if organization.IsErrOrgNotExist(err) {
				ctx.Error(http.StatusUnprocessableEntity, "", err)
			} else {
				ctx.Error(http.StatusInternalServerError, "GetOrgByName", err)
			}
			return
		}
		isMember, err := org.IsOrgMember(ctx.Doer.ID)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "IsOrgMember", err)
			return
		} else if !isMember {
			ctx.Error(http.StatusForbidden, "isMemberNot", fmt.Sprintf("User is no Member of Organisation '%s'", org.Name))
			return
		}
		forker = org.AsUser()
	}

	var name string
	if form.Name == nil {
		name = repo.Name
	} else {
		name = *form.Name
	}

	fork, err := repo_service.ForkRepository(ctx, ctx.Doer, forker, repo_service.ForkRepoOptions{
		BaseRepo:    repo,
		Name:        name,
		Description: repo.Description,
	})
	if err != nil {
		if repo_model.IsErrRepoAlreadyExist(err) {
			ctx.Error(http.StatusConflict, "ForkRepository", err)
		} else {
			ctx.Error(http.StatusInternalServerError, "ForkRepository", err)
		}
		return
	}

	// TODO change back to 201
	ctx.JSON(http.StatusAccepted, convert.ToRepo(fork, perm.AccessModeOwner))
}
