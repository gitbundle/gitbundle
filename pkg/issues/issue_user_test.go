// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issues_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func Test_NewIssueUsers(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	newIssue := &issues_model.Issue{
		RepoID:   repo.ID,
		PosterID: 4,
		Index:    6,
		Title:    "newTestIssueTitle",
		Content:  "newTestIssueContent",
	}

	// artificially insert new issue
	unittest.AssertSuccessfulInsert(t, newIssue)

	assert.NoError(t, issues_model.NewIssueUsers(db.DefaultContext, repo, newIssue))

	// issue_user table should now have entries for new issue
	unittest.AssertExistsAndLoadBean(t, &issues_model.IssueUser{IssueID: newIssue.ID, UID: newIssue.PosterID})
	unittest.AssertExistsAndLoadBean(t, &issues_model.IssueUser{IssueID: newIssue.ID, UID: repo.OwnerID})
}

func TestUpdateIssueUserByRead(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 1}).(*issues_model.Issue)

	assert.NoError(t, issues_model.UpdateIssueUserByRead(4, issue.ID))
	unittest.AssertExistsAndLoadBean(t, &issues_model.IssueUser{IssueID: issue.ID, UID: 4}, "is_read=1")

	assert.NoError(t, issues_model.UpdateIssueUserByRead(4, issue.ID))
	unittest.AssertExistsAndLoadBean(t, &issues_model.IssueUser{IssueID: issue.ID, UID: 4}, "is_read=1")

	assert.NoError(t, issues_model.UpdateIssueUserByRead(unittest.NonexistentID, unittest.NonexistentID))
}

func TestUpdateIssueUsersByMentions(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 1}).(*issues_model.Issue)

	uids := []int64{2, 5}
	assert.NoError(t, issues_model.UpdateIssueUsersByMentions(db.DefaultContext, issue.ID, uids))
	for _, uid := range uids {
		unittest.AssertExistsAndLoadBean(t, &issues_model.IssueUser{IssueID: issue.ID, UID: uid}, "is_mentioned=1")
	}
}
