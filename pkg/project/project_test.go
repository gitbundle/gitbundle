// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package project

import (
	"testing"

	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestIsProjectTypeValid(t *testing.T) {
	const UnknownType Type = 15

	cases := []struct {
		typ   Type
		valid bool
	}{
		{TypeIndividual, false},
		{TypeRepository, true},
		{TypeOrganization, false},
		{UnknownType, false},
	}

	for _, v := range cases {
		assert.Equal(t, v.valid, IsTypeValid(v.typ))
	}
}

func TestGetProjects(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	projects, _, err := GetProjects(db.DefaultContext, SearchOptions{RepoID: 1})
	assert.NoError(t, err)

	// 1 value for this repo exists in the fixtures
	assert.Len(t, projects, 1)

	projects, _, err = GetProjects(db.DefaultContext, SearchOptions{RepoID: 3})
	assert.NoError(t, err)

	// 1 value for this repo exists in the fixtures
	assert.Len(t, projects, 1)
}

func TestProject(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	project := &Project{
		Type:        TypeRepository,
		BoardType:   BoardTypeBasicKanban,
		Title:       "New Project",
		RepoID:      1,
		CreatedUnix: timeutil.TimeStampNow(),
		CreatorID:   2,
	}

	assert.NoError(t, NewProject(project))

	_, err := GetProjectByID(db.DefaultContext, project.ID)
	assert.NoError(t, err)

	// Update project
	project.Title = "Updated title"
	assert.NoError(t, UpdateProject(db.DefaultContext, project))

	projectFromDB, err := GetProjectByID(db.DefaultContext, project.ID)
	assert.NoError(t, err)

	assert.Equal(t, project.Title, projectFromDB.Title)

	assert.NoError(t, ChangeProjectStatus(project, true))

	// Retrieve from DB afresh to check if it is truly closed
	projectFromDB, err = GetProjectByID(db.DefaultContext, project.ID)
	assert.NoError(t, err)

	assert.True(t, projectFromDB.IsClosed)
}
