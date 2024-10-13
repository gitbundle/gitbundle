// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert

import (
	"context"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/packages"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	user_model "github.com/gitbundle/server/pkg/user"
)

// ToPackage convert a packages.PackageDescriptor to api.Package
func ToPackage(ctx context.Context, pd *packages.PackageDescriptor, doer *user_model.User) (*api.Package, error) {
	var repo *api.Repository
	if pd.Repository != nil {
		permission, err := access_model.GetUserRepoPermission(ctx, pd.Repository, doer)
		if err != nil {
			return nil, err
		}

		if permission.HasAccess() {
			repo = ToRepo(pd.Repository, permission.AccessMode)
		}
	}

	return &api.Package{
		ID:         pd.Version.ID,
		Owner:      ToUser(pd.Owner, doer),
		Repository: repo,
		Creator:    ToUser(pd.Creator, doer),
		Type:       string(pd.Package.Type),
		Name:       pd.Package.Name,
		Version:    pd.Version.Version,
		CreatedAt:  pd.Version.CreatedUnix.AsTime(),
	}, nil
}

// ToPackageFile converts packages.PackageFileDescriptor to api.PackageFile
func ToPackageFile(pfd *packages.PackageFileDescriptor) *api.PackageFile {
	return &api.PackageFile{
		ID:         pfd.File.ID,
		Size:       pfd.Blob.Size,
		Name:       pfd.File.Name,
		HashMD5:    pfd.Blob.HashMD5,
		HashSHA1:   pfd.Blob.HashSHA1,
		HashSHA256: pfd.Blob.HashSHA256,
		HashSHA512: pfd.Blob.HashSHA512,
	}
}
