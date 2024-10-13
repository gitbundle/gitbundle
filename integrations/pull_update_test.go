// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	pull_service "github.com/gitbundle/server/pkg/pull/manager"
	"github.com/gitbundle/server/pkg/repo/repoman"
	repo_service "github.com/gitbundle/server/pkg/repo/repository_manager"
	files_service "github.com/gitbundle/server/pkg/repo/repository_manager/files"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestAPIPullUpdate(t *testing.T) {
	onGiteaRun(t, func(t *testing.T, giteaURL *url.URL) {
		// Create PR to test
		user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
		org26 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 26}).(*user_model.User)
		pr := createOutdatedPR(t, user, org26)

		// Test GetDiverging
		diffCount, err := pull_service.GetDiverging(git.DefaultContext, pr)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, diffCount.Behind)
		assert.EqualValues(t, 1, diffCount.Ahead)
		assert.NoError(t, pr.LoadBaseRepo())
		assert.NoError(t, pr.LoadIssue())

		session := loginUser(t, "user2")
		token := getTokenForLoggedInUser(t, session)
		req := NewRequestf(t, "POST", "/api/v1/repos/%s/%s/pulls/%d/update?token="+token, pr.BaseRepo.OwnerName, pr.BaseRepo.Name, pr.Issue.Index)
		session.MakeRequest(t, req, http.StatusOK)

		// Test GetDiverging after update
		diffCount, err = pull_service.GetDiverging(git.DefaultContext, pr)
		assert.NoError(t, err)
		assert.EqualValues(t, 0, diffCount.Behind)
		assert.EqualValues(t, 2, diffCount.Ahead)
	})
}

func TestAPIPullUpdateByRebase(t *testing.T) {
	onGiteaRun(t, func(t *testing.T, giteaURL *url.URL) {
		// Create PR to test
		user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
		org26 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 26}).(*user_model.User)
		pr := createOutdatedPR(t, user, org26)

		// Test GetDiverging
		diffCount, err := pull_service.GetDiverging(git.DefaultContext, pr)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, diffCount.Behind)
		assert.EqualValues(t, 1, diffCount.Ahead)
		assert.NoError(t, pr.LoadBaseRepo())
		assert.NoError(t, pr.LoadIssue())

		session := loginUser(t, "user2")
		token := getTokenForLoggedInUser(t, session)
		req := NewRequestf(t, "POST", "/api/v1/repos/%s/%s/pulls/%d/update?style=rebase&token="+token, pr.BaseRepo.OwnerName, pr.BaseRepo.Name, pr.Issue.Index)
		session.MakeRequest(t, req, http.StatusOK)

		// Test GetDiverging after update
		diffCount, err = pull_service.GetDiverging(git.DefaultContext, pr)
		assert.NoError(t, err)
		assert.EqualValues(t, 0, diffCount.Behind)
		assert.EqualValues(t, 1, diffCount.Ahead)
	})
}

func createOutdatedPR(t *testing.T, actor, forkOrg *user_model.User) *issues_model.PullRequest {
	baseRepo, err := repo_service.CreateRepository(actor, actor, repoman.CreateRepoOptions{
		Name:        "repo-pr-update",
		Description: "repo-tmp-pr-update description",
		AutoInit:    true,
		Gitignores:  "C,C++",
		License:     "MIT",
		Readme:      "Default",
		IsPrivate:   false,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, baseRepo)

	headRepo, err := repo_service.ForkRepository(git.DefaultContext, actor, forkOrg, repo_service.ForkRepoOptions{
		BaseRepo:    baseRepo,
		Name:        "repo-pr-update",
		Description: "desc",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, headRepo)

	// create a commit on base Repo
	_, err = files_service.CreateOrUpdateRepoFile(git.DefaultContext, baseRepo, actor, &files_service.UpdateRepoFileOptions{
		TreePath:  "File_A",
		Message:   "Add File A",
		Content:   "File A",
		IsNewFile: true,
		OldBranch: "master",
		NewBranch: "master",
		Author: &files_service.IdentityOptions{
			Name:  actor.Name,
			Email: actor.Email,
		},
		Committer: &files_service.IdentityOptions{
			Name:  actor.Name,
			Email: actor.Email,
		},
		Dates: &files_service.CommitDateOptions{
			Author:    time.Now(),
			Committer: time.Now(),
		},
	})
	assert.NoError(t, err)

	// create a commit on head Repo
	_, err = files_service.CreateOrUpdateRepoFile(git.DefaultContext, headRepo, actor, &files_service.UpdateRepoFileOptions{
		TreePath:  "File_B",
		Message:   "Add File on PR branch",
		Content:   "File B",
		IsNewFile: true,
		OldBranch: "master",
		NewBranch: "newBranch",
		Author: &files_service.IdentityOptions{
			Name:  actor.Name,
			Email: actor.Email,
		},
		Committer: &files_service.IdentityOptions{
			Name:  actor.Name,
			Email: actor.Email,
		},
		Dates: &files_service.CommitDateOptions{
			Author:    time.Now(),
			Committer: time.Now(),
		},
	})
	assert.NoError(t, err)

	// create Pull
	pullIssue := &issues_model.Issue{
		RepoID:   baseRepo.ID,
		Title:    "Test Pull -to-update-",
		PosterID: actor.ID,
		Poster:   actor,
		IsPull:   true,
	}
	pullRequest := &issues_model.PullRequest{
		HeadRepoID: headRepo.ID,
		BaseRepoID: baseRepo.ID,
		HeadBranch: "newBranch",
		BaseBranch: "master",
		HeadRepo:   headRepo,
		BaseRepo:   baseRepo,
		Type:       issues_model.PullRequestGitea,
	}
	err = pull_service.NewPullRequest(git.DefaultContext, baseRepo, pullIssue, nil, nil, pullRequest, nil)
	assert.NoError(t, err)

	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{Title: "Test Pull -to-update-"}).(*issues_model.Issue)
	pr, err := issues_model.GetPullRequestByIssueID(db.DefaultContext, issue.ID)
	assert.NoError(t, err)

	return pr
}
