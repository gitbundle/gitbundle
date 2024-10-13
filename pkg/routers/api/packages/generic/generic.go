// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package generic

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gitbundle/modules/log"
	packages_module "github.com/gitbundle/modules/packages"
	"github.com/gitbundle/server/pkg/context"
	packages_model "github.com/gitbundle/server/pkg/packages"
	packages_service "github.com/gitbundle/server/pkg/packages/manager"
	"github.com/gitbundle/server/pkg/routers/api/packages/helper"
)

var (
	packageNameRegex = regexp.MustCompile(`\A[A-Za-z0-9\.\_\-\+]+\z`)
	filenameRegex    = packageNameRegex
)

func apiError(ctx *context.Context, status int, obj interface{}) {
	helper.LogAndProcessError(ctx, status, obj, func(message string) {
		ctx.PlainText(status, message)
	})
}

// DownloadPackageFile serves the specific generic package.
func DownloadPackageFile(ctx *context.Context) {
	packageName, packageVersion, filename, err := sanitizeParameters(ctx)
	if err != nil {
		apiError(ctx, http.StatusBadRequest, err)
		return
	}

	s, pf, err := packages_service.GetFileStreamByPackageNameAndVersion(
		ctx,
		&packages_service.PackageInfo{
			Owner:       ctx.Package.Owner,
			PackageType: packages_model.TypeGeneric,
			Name:        packageName,
			Version:     packageVersion,
		},
		&packages_service.PackageFileInfo{
			Filename: filename,
		},
	)
	if err != nil {
		if err == packages_model.ErrPackageNotExist || err == packages_model.ErrPackageFileNotExist {
			apiError(ctx, http.StatusNotFound, err)
			return
		}
		apiError(ctx, http.StatusInternalServerError, err)
		return
	}
	defer s.Close()

	ctx.ServeStream(s, pf.Name)
}

// UploadPackage uploads the specific generic package.
// Duplicated packages get rejected.
func UploadPackage(ctx *context.Context) {
	packageName, packageVersion, filename, err := sanitizeParameters(ctx)
	if err != nil {
		apiError(ctx, http.StatusBadRequest, err)
		return
	}

	upload, close, err := ctx.UploadStream()
	if err != nil {
		apiError(ctx, http.StatusInternalServerError, err)
		return
	}
	if close {
		defer upload.Close()
	}

	buf, err := packages_module.CreateHashedBufferFromReader(upload, 32*1024*1024)
	if err != nil {
		log.Error("Error creating hashed buffer: %v", err)
		apiError(ctx, http.StatusInternalServerError, err)
		return
	}
	defer buf.Close()

	_, _, err = packages_service.CreatePackageAndAddFile(
		&packages_service.PackageCreationInfo{
			PackageInfo: packages_service.PackageInfo{
				Owner:       ctx.Package.Owner,
				PackageType: packages_model.TypeGeneric,
				Name:        packageName,
				Version:     packageVersion,
			},
			Creator: ctx.Doer,
		},
		&packages_service.PackageFileCreationInfo{
			PackageFileInfo: packages_service.PackageFileInfo{
				Filename: filename,
			},
			Data:   buf,
			IsLead: true,
		},
	)
	if err != nil {
		if err == packages_model.ErrDuplicatePackageVersion {
			apiError(ctx, http.StatusBadRequest, err)
			return
		}
		apiError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

// DeletePackage deletes the specific generic package.
func DeletePackage(ctx *context.Context) {
	packageName, packageVersion, _, err := sanitizeParameters(ctx)
	if err != nil {
		apiError(ctx, http.StatusBadRequest, err)
		return
	}

	err = packages_service.RemovePackageVersionByNameAndVersion(
		ctx.Doer,
		&packages_service.PackageInfo{
			Owner:       ctx.Package.Owner,
			PackageType: packages_model.TypeGeneric,
			Name:        packageName,
			Version:     packageVersion,
		},
	)
	if err != nil {
		if err == packages_model.ErrPackageNotExist {
			apiError(ctx, http.StatusNotFound, err)
			return
		}
		apiError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func sanitizeParameters(ctx *context.Context) (string, string, string, error) {
	packageName := ctx.Params("packagename")
	filename := ctx.Params("filename")

	if !packageNameRegex.MatchString(packageName) || !filenameRegex.MatchString(filename) {
		return "", "", "", errors.New("Invalid package name or filename")
	}

	packageVersion := strings.TrimSpace(ctx.Params("packageversion"))
	if packageVersion == "" {
		return "", "", "", errors.New("Invalid package version")
	}

	return packageName, packageVersion, filename, nil
}
