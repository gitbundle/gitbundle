// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repoman_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/organization"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestUser_RemoveMember(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	org := unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 3}).(*organization.Organization)

	// remove a user that is a member
	unittest.AssertExistsAndLoadBean(t, &organization.OrgUser{UID: 4, OrgID: 3})
	prevNumMembers := org.NumMembers
	assert.NoError(t, repoman.RemoveOrgUser(org.ID, 4))
	unittest.AssertNotExistsBean(t, &organization.OrgUser{UID: 4, OrgID: 3})
	org = unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 3}).(*organization.Organization)
	assert.Equal(t, prevNumMembers-1, org.NumMembers)

	// remove a user that is not a member
	unittest.AssertNotExistsBean(t, &organization.OrgUser{UID: 5, OrgID: 3})
	prevNumMembers = org.NumMembers
	assert.NoError(t, repoman.RemoveOrgUser(org.ID, 5))
	unittest.AssertNotExistsBean(t, &organization.OrgUser{UID: 5, OrgID: 3})
	org = unittest.AssertExistsAndLoadBean(t, &organization.Organization{ID: 3}).(*organization.Organization)
	assert.Equal(t, prevNumMembers, org.NumMembers)

	unittest.CheckConsistencyFor(t, &user_model.User{}, &organization.Team{})
}

func TestRemoveOrgUser(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	testSuccess := func(orgID, userID int64) {
		org := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: orgID}).(*user_model.User)
		expectedNumMembers := org.NumMembers
		if unittest.BeanExists(t, &organization.OrgUser{OrgID: orgID, UID: userID}) {
			expectedNumMembers--
		}
		assert.NoError(t, repoman.RemoveOrgUser(orgID, userID))
		unittest.AssertNotExistsBean(t, &organization.OrgUser{OrgID: orgID, UID: userID})
		org = unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: orgID}).(*user_model.User)
		assert.EqualValues(t, expectedNumMembers, org.NumMembers)
	}
	testSuccess(3, 4)
	testSuccess(3, 4)

	err := repoman.RemoveOrgUser(7, 5)
	assert.Error(t, err)
	assert.True(t, organization.IsErrLastOrgOwner(err))
	unittest.AssertExistsAndLoadBean(t, &organization.OrgUser{OrgID: 7, UID: 5})
	unittest.CheckConsistencyFor(t, &user_model.User{}, &organization.Team{})
}
