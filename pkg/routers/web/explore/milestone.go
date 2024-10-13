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
	tplMilestones base.TplName = "explore/milestones"
)

func Milestones(ctx *context.Context) {
	if unit.TypeIssues.UnitGlobalDisabled() && unit.TypePullRequests.UnitGlobalDisabled() {
		log.Debug(
			"Milestones overview page not available as both issues and pull requests are globally disabled",
		)
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Data["Title"] = ctx.Tr("milestones")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreMilestones"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	user.BuildMilestonesOverview(ctx)
	ctx.HTML(http.StatusOK, tplMilestones)
}
