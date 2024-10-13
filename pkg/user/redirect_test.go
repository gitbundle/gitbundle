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

func TestLookupUserRedirect(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	userID, err := user_model.LookupUserRedirect("olduser1")
	assert.NoError(t, err)
	assert.EqualValues(t, 1, userID)

	_, err = user_model.LookupUserRedirect("doesnotexist")
	assert.True(t, user_model.IsErrUserRedirectNotExist(err))
}
