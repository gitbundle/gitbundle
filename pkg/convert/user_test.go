// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert_test

import (
	"testing"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestUser_ToUser(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	user1 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1, IsAdmin: true}).(*user_model.User)

	apiUser := toUser(user1, true, true)
	assert.True(t, apiUser.IsAdmin)
	assert.Contains(t, apiUser.AvatarURL, "://")

	user2 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2, IsAdmin: false}).(*user_model.User)

	apiUser = toUser(user2, true, true)
	assert.False(t, apiUser.IsAdmin)

	apiUser = toUser(user1, false, false)
	assert.False(t, apiUser.IsAdmin)
	assert.EqualValues(t, api.VisibleTypePublic.String(), apiUser.Visibility)

	user31 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 31, IsAdmin: false, Visibility: api.VisibleTypePrivate}).(*user_model.User)

	apiUser = toUser(user31, true, true)
	assert.False(t, apiUser.IsAdmin)
	assert.EqualValues(t, api.VisibleTypePrivate.String(), apiUser.Visibility)
}
