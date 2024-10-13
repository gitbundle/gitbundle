// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"os"
	"testing"

	"github.com/gitbundle/modules/setting"
	migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestMigrateLocalPath(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	adminUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{Name: "user1"}).(*user_model.User)

	old := setting.ImportLocalPaths
	setting.ImportLocalPaths = true

	lowercasePath, err := os.MkdirTemp("", "lowercase") // may not be lowercase because MkdirTemp creates a random directory name which may be mixedcase
	assert.NoError(t, err)
	defer os.RemoveAll(lowercasePath)

	err = migrations.IsMigrateURLAllowed(lowercasePath, adminUser)
	assert.NoError(t, err, "case lowercase path")

	mixedcasePath, err := os.MkdirTemp("", "mIxeDCaSe")
	assert.NoError(t, err)
	defer os.RemoveAll(mixedcasePath)

	err = migrations.IsMigrateURLAllowed(mixedcasePath, adminUser)
	assert.NoError(t, err, "case mixedcase path")

	setting.ImportLocalPaths = old
}
