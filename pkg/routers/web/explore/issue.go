// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package explore

import (
	"net/http"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/routers/web/user"
	"github.com/gitbundle/server/pkg/unit"
)

const (
	tplIssues base.TplName = "explore/issues"
)

func Issues(ctx *context.Context) {
	if unit.TypeIssues.UnitGlobalDisabled() {
		log.Debug("Issues overview page not available as it is globally disabled.")
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Data["Title"] = ctx.Tr("issues")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreIssues"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	user.BuildIssueOverview(ctx, unit.TypeIssues)
	ctx.HTML(http.StatusOK, tplIssues)
}

func Pulls(ctx *context.Context) {
	if unit.TypePullRequests.UnitGlobalDisabled() {
		log.Debug("Pull request overview page not available as it is globally disabled.")
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Data["Title"] = ctx.Tr("pull_requests")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExplorePulls"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	user.BuildIssueOverview(ctx, unit.TypePullRequests)
	ctx.HTML(http.StatusOK, tplIssues)
}
