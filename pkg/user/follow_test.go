// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package user_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestIsFollowing(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	assert.True(t, user_model.IsFollowing(4, 2))
	assert.False(t, user_model.IsFollowing(2, 4))
	assert.False(t, user_model.IsFollowing(5, unittest.NonexistentID))
	assert.False(t, user_model.IsFollowing(unittest.NonexistentID, 5))
	assert.False(t, user_model.IsFollowing(unittest.NonexistentID, unittest.NonexistentID))
}
