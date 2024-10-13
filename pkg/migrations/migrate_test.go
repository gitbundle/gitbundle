// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"strconv"
	"testing"

	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/foreignreference"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/migrations"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/release"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestMigrate_InsertMilestones(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	reponame := "repo1"
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{Name: reponame}).(*repo_model.Repository)
	name := "milestonetest1"
	ms := &issues_model.Milestone{
		RepoID: repo.ID,
		Name:   name,
	}
	err := migrations.InsertMilestones(ms)
	assert.NoError(t, err)
	unittest.AssertExistsAndLoadBean(t, ms)
	repoModified := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: repo.ID}).(*repo_model.Repository)
	assert.EqualValues(t, repo.NumMilestones+1, repoModified.NumMilestones)

	unittest.CheckConsistencyFor(t, &issues_model.Milestone{})
}

func assertCreateIssues(t *testing.T, isPull bool) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	reponame := "repo1"
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{Name: reponame}).(*repo_model.Repository)
	owner := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: repo.OwnerID}).(*user_model.User)
	label := unittest.AssertExistsAndLoadBean(t, &issues_model.Label{ID: 1}).(*issues_model.Label)
	milestone := unittest.AssertExistsAndLoadBean(t, &issues_model.Milestone{ID: 1}).(*issues_model.Milestone)
	assert.EqualValues(t, milestone.ID, 1)
	reaction := &issues_model.Reaction{
		Type:   "heart",
		UserID: owner.ID,
	}

	foreignIndex := int64(12345)
	title := "issuetitle1"
	is := &issues_model.Issue{
		RepoID:      repo.ID,
		MilestoneID: milestone.ID,
		Repo:        repo,
		Title:       title,
		Content:     "issuecontent1",
		IsPull:      isPull,
		PosterID:    owner.ID,
		Poster:      owner,
		IsClosed:    true,
		Labels:      []*issues_model.Label{label},
		Reactions:   []*issues_model.Reaction{reaction},
		ForeignReference: &foreignreference.ForeignReference{
			ForeignIndex: strconv.FormatInt(foreignIndex, 10),
			RepoID:       repo.ID,
			Type:         foreignreference.TypeIssue,
		},
	}
	err := migrations.InsertIssues(is)
	assert.NoError(t, err)

	i := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{Title: title}).(*issues_model.Issue)
	assert.Nil(t, i.ForeignReference)
	err = i.LoadAttributes(db.DefaultContext)
	assert.NoError(t, err)
	assert.EqualValues(t, strconv.FormatInt(foreignIndex, 10), i.ForeignReference.ForeignIndex)
	unittest.AssertExistsAndLoadBean(t, &issues_model.Reaction{Type: "heart", UserID: owner.ID, IssueID: i.ID})
}

func TestMigrate_CreateIssuesIsPullFalse(t *testing.T) {
	assertCreateIssues(t, false)
}

func TestMigrate_CreateIssuesIsPullTrue(t *testing.T) {
	assertCreateIssues(t, true)
}

func TestMigrate_InsertIssueComments(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 1}).(*issues_model.Issue)
	_ = issue.LoadRepo(db.DefaultContext)
	owner := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: issue.Repo.OwnerID}).(*user_model.User)
	reaction := &issues_model.Reaction{
		Type:   "heart",
		UserID: owner.ID,
	}

	comment := &issues_model.Comment{
		PosterID:  owner.ID,
		Poster:    owner,
		IssueID:   issue.ID,
		Issue:     issue,
		Reactions: []*issues_model.Reaction{reaction},
	}

	err := migrations.InsertIssueComments([]*issues_model.Comment{comment})
	assert.NoError(t, err)

	issueModified := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 1}).(*issues_model.Issue)
	assert.EqualValues(t, issue.NumComments+1, issueModified.NumComments)

	unittest.CheckConsistencyFor(t, &issues_model.Issue{})
}

func TestMigrate_InsertPullRequests(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	reponame := "repo1"
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{Name: reponame}).(*repo_model.Repository)
	owner := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: repo.OwnerID}).(*user_model.User)

	i := &issues_model.Issue{
		RepoID:   repo.ID,
		Repo:     repo,
		Title:    "title1",
		Content:  "issuecontent1",
		IsPull:   true,
		PosterID: owner.ID,
		Poster:   owner,
	}

	p := &issues_model.PullRequest{
		Issue: i,
	}

	err := migrations.InsertPullRequests(p)
	assert.NoError(t, err)

	_ = unittest.AssertExistsAndLoadBean(t, &issues_model.PullRequest{IssueID: i.ID}).(*issues_model.PullRequest)

	unittest.CheckConsistencyFor(t, &issues_model.Issue{}, &issues_model.PullRequest{})
}

func TestMigrate_InsertReleases(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a := &repo_model.Attachment{
		UUID: "a0eebc91-9c0c-4ef7-bb6e-6bb9bd380a12",
	}
	r := &release.Release{
		Attachments: []*repo_model.Attachment{a},
	}

	err := migrations.InsertReleases(r)
	assert.NoError(t, err)
}
