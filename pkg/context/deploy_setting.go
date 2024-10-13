// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"github.com/gitbundle/modules/log"
)

func ForkedRepoNotAllowed(ctx *Context) {
	if ctx.Repo.Repository.IsFork {
		log.Warn("services/context: unexpected request")
		ctx.ServerError(ctx.Tr("repo.forked.not_allowed"), nil)
		return
	}
}
