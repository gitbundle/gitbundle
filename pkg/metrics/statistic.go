// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package metrics

import (
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/action"
	asymkey_model "github.com/gitbundle/server/pkg/asymkey"
	"github.com/gitbundle/server/pkg/auth"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/organization"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	project_model "github.com/gitbundle/server/pkg/project"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/release"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/webhook"
)

// Statistic contains the database statistics
type Statistic struct {
	Counter struct {
		User, Org, PublicKey,
		Repo, Watch, Star, Action, Access,
		Issue, IssueClosed, IssueOpen,
		Comment, Oauth, Follow,
		Mirror, Release, AuthSource, Webhook,
		Milestone, Label, HookTask,
		Team, UpdateTask, Project,
		ProjectBoard, Attachment int64
		IssueByLabel      []IssueByLabelCount
		IssueByRepository []IssueByRepositoryCount
	}
}

// IssueByLabelCount contains the number of issue group by label
type IssueByLabelCount struct {
	Count int64
	Label string
}

// IssueByRepositoryCount contains the number of issue group by repository
type IssueByRepositoryCount struct {
	Count      int64
	OwnerName  string
	Repository string
}

// GetStatistic returns the database statistics
func GetStatistic() (stats Statistic) {
	e := db.GetEngine(db.DefaultContext)
	stats.Counter.User = user_model.CountUsers(nil)
	stats.Counter.Org, _ = organization.CountOrgs(organization.FindOrgOptions{IncludePrivate: true})
	stats.Counter.PublicKey, _ = e.Count(new(asymkey_model.PublicKey))
	stats.Counter.Repo, _ = repo_model.CountRepositories(db.DefaultContext, repo_model.CountRepositoryOptions{})
	stats.Counter.Watch, _ = e.Count(new(repo_model.Watch))
	stats.Counter.Star, _ = e.Count(new(repo_model.Star))
	stats.Counter.Action, _ = db.EstimateCount(db.DefaultContext, new(action.Action))
	stats.Counter.Access, _ = e.Count(new(access_model.Access))

	type IssueCount struct {
		Count    int64
		IsClosed bool
	}

	if setting.Metrics.EnabledIssueByLabel {
		stats.Counter.IssueByLabel = []IssueByLabelCount{}

		_ = e.Select("COUNT(*) AS count, l.name AS label").
			Join("LEFT", "label l", "l.id=il.label_id").
			Table("issue_label il").
			GroupBy("l.name").
			Find(&stats.Counter.IssueByLabel)
	}

	if setting.Metrics.EnabledIssueByRepository {
		stats.Counter.IssueByRepository = []IssueByRepositoryCount{}

		_ = e.Select("COUNT(*) AS count, r.owner_name, r.name AS repository").
			Join("LEFT", "repository r", "r.id=i.repo_id").
			Table("issue i").
			GroupBy("r.owner_name, r.name").
			Find(&stats.Counter.IssueByRepository)
	}

	issueCounts := []IssueCount{}

	_ = e.Select("COUNT(*) AS count, is_closed").Table("issue").GroupBy("is_closed").Find(&issueCounts)
	for _, c := range issueCounts {
		if c.IsClosed {
			stats.Counter.IssueClosed = c.Count
		} else {
			stats.Counter.IssueOpen = c.Count
		}
	}

	stats.Counter.Issue = stats.Counter.IssueClosed + stats.Counter.IssueOpen

	stats.Counter.Comment, _ = e.Count(new(issues_model.Comment))
	stats.Counter.Oauth = 0
	stats.Counter.Follow, _ = e.Count(new(user_model.Follow))
	stats.Counter.Mirror, _ = e.Count(new(repo_model.Mirror))
	stats.Counter.Release, _ = e.Count(new(release.Release))
	stats.Counter.AuthSource = auth.CountSources()
	stats.Counter.Webhook, _ = e.Count(new(webhook.Webhook))
	stats.Counter.Milestone, _ = e.Count(new(issues_model.Milestone))
	stats.Counter.Label, _ = e.Count(new(issues_model.Label))
	stats.Counter.HookTask, _ = e.Count(new(webhook.HookTask))
	stats.Counter.Team, _ = e.Count(new(organization.Team))
	stats.Counter.Attachment, _ = e.Count(new(repo_model.Attachment))
	stats.Counter.Project, _ = e.Count(new(project_model.Project))
	stats.Counter.ProjectBoard, _ = e.Count(new(project_model.Board))
	return
}
