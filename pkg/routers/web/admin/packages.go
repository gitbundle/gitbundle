// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package admin

import (
	"net/http"
	"net/url"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/db"
	packages_model "github.com/gitbundle/server/pkg/packages"
	packages_service "github.com/gitbundle/server/pkg/packages/manager"
)

const (
	tplPackagesList base.TplName = "admin/packages/list"
)

// Packages shows all packages
func Packages(ctx *context.Context) {
	page := ctx.FormInt("page")
	if page <= 1 {
		page = 1
	}
	query := ctx.FormTrim("q")
	packageType := ctx.FormTrim("type")
	sort := ctx.FormTrim("sort")

	pvs, total, err := packages_model.SearchVersions(ctx, &packages_model.PackageSearchOptions{
		Type: packages_model.Type(packageType),
		Name: packages_model.SearchValue{Value: query},
		Sort: sort,
		Paginator: &db.ListOptions{
			PageSize: setting.UI.PackagesPagingNum,
			Page:     page,
		},
	})
	if err != nil {
		ctx.ServerError("SearchVersions", err)
		return
	}

	pds, err := packages_model.GetPackageDescriptors(ctx, pvs)
	if err != nil {
		ctx.ServerError("GetPackageDescriptors", err)
		return
	}

	totalBlobSize, err := packages_model.GetTotalBlobSize()
	if err != nil {
		ctx.ServerError("GetTotalBlobSize", err)
		return
	}

	ctx.Data["Title"] = ctx.Tr("packages.title")
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminPackages"] = true
	ctx.Data["Query"] = query
	ctx.Data["PackageType"] = packageType
	ctx.Data["SortType"] = sort
	ctx.Data["PackageDescriptors"] = pds
	ctx.Data["Total"] = total
	ctx.Data["TotalBlobSize"] = totalBlobSize

	pager := context.NewPagination(int(total), setting.UI.PackagesPagingNum, page, 5)
	pager.AddParamString("q", query)
	pager.AddParamString("type", packageType)
	pager.AddParamString("sort", sort)
	ctx.Data["Page"] = pager

	ctx.HTML(http.StatusOK, tplPackagesList)
}

// DeletePackageVersion deletes a package version
func DeletePackageVersion(ctx *context.Context) {
	pv, err := packages_model.GetVersionByID(db.DefaultContext, ctx.FormInt64("id"))
	if err != nil {
		ctx.ServerError("GetRepositoryByID", err)
		return
	}

	if err := packages_service.RemovePackageVersion(ctx.Doer, pv); err != nil {
		ctx.ServerError("RemovePackageVersion", err)
		return
	}

	ctx.Flash.Success(ctx.Tr("packages.settings.delete.success"))
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"redirect": setting.AppSubURL + "/admin/packages?page=" + url.QueryEscape(ctx.FormString("page")) + "&q=" + url.QueryEscape(ctx.FormString("q")) + "&type=" + url.QueryEscape(ctx.FormString("type")),
	})
}
