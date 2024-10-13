// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issue_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	issue_service "github.com/gitbundle/server/pkg/issues/manager"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNotPassedAssignee(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	// Fake issue with assignees
	issue, err := issues_model.GetIssueWithAttrsByID(1)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(issue.Assignees))

	user1, err := user_model.GetUserByID(1) // This user is already assigned (see the definition in fixtures), so running  UpdateAssignee should unassign him
	assert.NoError(t, err)

	// Check if he got removed
	isAssigned, err := issues_model.IsUserAssignedToIssue(db.DefaultContext, issue, user1)
	assert.NoError(t, err)
	assert.True(t, isAssigned)

	// Clean everyone
	err = issue_service.DeleteNotPassedAssignee(issue, user1, []*user_model.User{})
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(issue.Assignees))

	// Check they're gone
	assert.NoError(t, issue.LoadAssignees(db.DefaultContext))
	assert.EqualValues(t, 0, len(issue.Assignees))
	assert.Empty(t, issue.Assignee)
}
