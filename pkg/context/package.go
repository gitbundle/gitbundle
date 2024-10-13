// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/organization"
	packages_model "github.com/gitbundle/server/pkg/packages"
	"github.com/gitbundle/server/pkg/perm"
	"github.com/gitbundle/server/pkg/unit"
	user_model "github.com/gitbundle/server/pkg/user"
)

var errUserPackageNotEnabled = errors.New("User or organization package registry feature is disabled now, you can't access")

// Package contains owner, access mode and optional the package descriptor
type Package struct {
	Owner      *user_model.User
	AccessMode perm.AccessMode
	Descriptor *packages_model.PackageDescriptor
}

// PackageAssignment returns a middleware to handle Context.Package assignment
func PackageAssignment() func(ctx *Context) {
	return func(ctx *Context) {
		packageAssignment(ctx, func(status int, title string, obj interface{}) {
			err, ok := obj.(error)
			if !ok {
				err = fmt.Errorf("%s", obj)
			}
			if status == http.StatusNotFound {
				ctx.NotFound(title, err)
			} else {
				ctx.ServerError(title, err)
			}
		})
	}
}

// PackageAssignmentAPI returns a middleware to handle Context.Package assignment
func PackageAssignmentAPI() func(ctx *APIContext) {
	return func(ctx *APIContext) {
		packageAssignment(ctx.Context, ctx.Error)
	}
}

func packageAssignment(ctx *Context, errCb func(int, string, interface{})) {
	if !ctx.ContextUser.EnabledPackages && !setting.Packages.EnableForAllOrgs {
		errCb(http.StatusForbidden, "", errUserPackageNotEnabled)
		return
	}

	ctx.Package = &Package{
		Owner: ctx.ContextUser,
	}

	if ctx.Package.Owner.IsOrganization() {
		org := organization.OrgFromUser(ctx.Package.Owner)

		// 1. Get user max authorize level for the org (may be none, if user is not member of the org)
		if ctx.Doer != nil {
			var err error
			ctx.Package.AccessMode, err = org.GetOrgUserMaxAuthorizeLevel(ctx.Doer.ID)
			if err != nil {
				errCb(http.StatusInternalServerError, "GetOrgUserMaxAuthorizeLevel", err)
				return
			}
			// If access mode is less than write check every team for more permissions
			if ctx.Package.AccessMode < perm.AccessModeWrite {
				teams, err := organization.GetUserOrgTeams(ctx, org.ID, ctx.Doer.ID)
				if err != nil {
					errCb(http.StatusInternalServerError, "GetUserOrgTeams", err)
					return
				}
				for _, t := range teams {
					perm := t.UnitAccessModeCtx(ctx, unit.TypePackages)
					if ctx.Package.AccessMode < perm {
						ctx.Package.AccessMode = perm
					}
				}
			}
		}
		// 2. If authorize level is none, check if org is visible to user
		if ctx.Package.AccessMode == perm.AccessModeNone && organization.HasOrgOrUserVisible(ctx, ctx.Package.Owner, ctx.Doer) {
			ctx.Package.AccessMode = perm.AccessModeRead
		}
	} else {
		if ctx.Doer != nil && !ctx.Doer.IsGhost() {
			// 1. Check if user is package owner
			if ctx.Doer.ID == ctx.Package.Owner.ID {
				ctx.Package.AccessMode = perm.AccessModeOwner
			} else if ctx.Package.Owner.Visibility == structs.VisibleTypePublic || ctx.Package.Owner.Visibility == structs.VisibleTypeLimited { // 2. Check if package owner is public or limited
				ctx.Package.AccessMode = perm.AccessModeRead
			}
		} else if ctx.Package.Owner.Visibility == structs.VisibleTypePublic { // 3. Check if package owner is public
			ctx.Package.AccessMode = perm.AccessModeRead
		}
	}

	packageType := ctx.Params("type")
	name := ctx.Params("name")
	version := ctx.Params("version")
	if packageType != "" && name != "" && version != "" {
		pv, err := packages_model.GetVersionByNameAndVersion(ctx, ctx.Package.Owner.ID, packages_model.Type(packageType), name, version)
		if err != nil {
			if err == packages_model.ErrPackageNotExist {
				errCb(http.StatusNotFound, "GetVersionByNameAndVersion", err)
			} else {
				errCb(http.StatusInternalServerError, "GetVersionByNameAndVersion", err)
			}
			return
		}

		ctx.Package.Descriptor, err = packages_model.GetPackageDescriptor(ctx, pv)
		if err != nil {
			errCb(http.StatusInternalServerError, "GetPackageDescriptor", err)
			return
		}
	}
}
