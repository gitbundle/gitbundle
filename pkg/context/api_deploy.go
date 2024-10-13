// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"github.com/gitbundle/server/pkg/system"
)

var ApiDeploy = apiDeploy{}

type apiDeploy struct{}

func (apiDeploy) SystemSettingsMap(keys ...string) func(ctx *APIContext) {
	return func(ctx *APIContext) {
		res, err := system.GetSystemSettingsByKeys(keys...)
		if err != nil {
			ctx.InternalServerError(err)
			return
		}

		ctx.SystemSettingMap = res
	}
}

func (apiDeploy) ForkedRepoNotAllowed(ctx *APIContext) {
	if ctx.Repo.Repository.IsFork {
		ctx.NotFound()
		return
	}
}
