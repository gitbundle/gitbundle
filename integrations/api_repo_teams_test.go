// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"fmt"
	"net/http"
	"testing"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/util"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unit"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestAPIRepoTeams(t *testing.T) {
	defer prepareTestEnv(t)()

	// publicOrgRepo = user3/repo21
	publicOrgRepo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 32}).(*repo_model.Repository)
	// user4
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 4}).(*user_model.User)
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session)

	// ListTeams
	url := fmt.Sprintf("/api/v1/repos/%s/teams?token=%s", publicOrgRepo.FullName(), token)
	req := NewRequest(t, "GET", url)
	res := session.MakeRequest(t, req, http.StatusOK)
	var teams []*api.Team
	DecodeJSON(t, res, &teams)
	if assert.Len(t, teams, 2) {
		assert.EqualValues(t, "Owners", teams[0].Name)
		assert.True(t, teams[0].CanCreateOrgRepo)
		assert.True(t, util.IsEqualSlice(unit.AllUnitKeyNames(), teams[0].Units), fmt.Sprintf("%v == %v", unit.AllUnitKeyNames(), teams[0].Units))
		assert.EqualValues(t, "owner", teams[0].Permission)

		assert.EqualValues(t, "test_team", teams[1].Name)
		assert.False(t, teams[1].CanCreateOrgRepo)
		assert.EqualValues(t, []string{"repo.issues"}, teams[1].Units)
		assert.EqualValues(t, "write", teams[1].Permission)
	}

	// IsTeam
	url = fmt.Sprintf("/api/v1/repos/%s/teams/%s?token=%s", publicOrgRepo.FullName(), "Test_Team", token)
	req = NewRequest(t, "GET", url)
	res = session.MakeRequest(t, req, http.StatusOK)
	var team *api.Team
	DecodeJSON(t, res, &team)
	assert.EqualValues(t, teams[1], team)

	url = fmt.Sprintf("/api/v1/repos/%s/teams/%s?token=%s", publicOrgRepo.FullName(), "NonExistingTeam", token)
	req = NewRequest(t, "GET", url)
	session.MakeRequest(t, req, http.StatusNotFound)

	// AddTeam with user4
	url = fmt.Sprintf("/api/v1/repos/%s/teams/%s?token=%s", publicOrgRepo.FullName(), "team1", token)
	req = NewRequest(t, "PUT", url)
	session.MakeRequest(t, req, http.StatusForbidden)

	// AddTeam with user2
	user = unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	session = loginUser(t, user.Name)
	token = getTokenForLoggedInUser(t, session)
	url = fmt.Sprintf("/api/v1/repos/%s/teams/%s?token=%s", publicOrgRepo.FullName(), "team1", token)
	req = NewRequest(t, "PUT", url)
	session.MakeRequest(t, req, http.StatusNoContent)
	session.MakeRequest(t, req, http.StatusUnprocessableEntity) // test duplicate request

	// DeleteTeam
	url = fmt.Sprintf("/api/v1/repos/%s/teams/%s?token=%s", publicOrgRepo.FullName(), "team1", token)
	req = NewRequest(t, "DELETE", url)
	session.MakeRequest(t, req, http.StatusNoContent)
	session.MakeRequest(t, req, http.StatusUnprocessableEntity) // test duplicate request
}
