// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert_test

import (
	"fmt"
	"testing"
	"time"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/convert"
	issues_model "github.com/gitbundle/server/pkg/issues"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestLabel_ToLabel(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	label := unittest.AssertExistsAndLoadBean(t, &issues_model.Label{ID: 1}).(*issues_model.Label)
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: label.RepoID}).(*repo_model.Repository)
	assert.Equal(t, &api.Label{
		ID:    label.ID,
		Name:  label.Name,
		Color: "abcdef",
		URL:   fmt.Sprintf("%sapi/v1/repos/user2/repo1/labels/%d", setting.AppURL, label.ID),
	}, convert.ToLabel(label, repo, nil))
}

func TestMilestone_APIFormat(t *testing.T) {
	milestone := &issues_model.Milestone{
		ID:              3,
		RepoID:          4,
		Name:            "milestoneName",
		Content:         "milestoneContent",
		IsClosed:        false,
		NumOpenIssues:   5,
		NumClosedIssues: 6,
		CreatedUnix:     timeutil.TimeStamp(time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		UpdatedUnix:     timeutil.TimeStamp(time.Date(1999, time.March, 1, 0, 0, 0, 0, time.UTC).Unix()),
		DeadlineUnix:    timeutil.TimeStamp(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
	}
	assert.Equal(t, api.Milestone{
		ID:           milestone.ID,
		State:        api.StateOpen,
		Title:        milestone.Name,
		Description:  milestone.Content,
		OpenIssues:   milestone.NumOpenIssues,
		ClosedIssues: milestone.NumClosedIssues,
		Created:      milestone.CreatedUnix.AsTime(),
		Updated:      milestone.UpdatedUnix.AsTimePtr(),
		Deadline:     milestone.DeadlineUnix.AsTimePtr(),
	}, *convert.ToAPIMilestone(milestone))
}
