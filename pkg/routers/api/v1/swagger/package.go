// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package swagger

import (
	api "github.com/gitbundle/api/pkg/structs"
)

// Package
// swagger:response Package
type swaggerResponsePackage struct {
	// in:body
	Body api.Package `json:"body"`
}

// PackageList
// swagger:response PackageList
type swaggerResponsePackageList struct {
	// in:body
	Body []api.Package `json:"body"`
}

// PackageFileList
// swagger:response PackageFileList
type swaggerResponsePackageFileList struct {
	// in:body
	Body []api.PackageFile `json:"body"`
}
