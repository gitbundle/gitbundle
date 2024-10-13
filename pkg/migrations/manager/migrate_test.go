// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"path/filepath"
	"testing"

	"github.com/gitbundle/modules/setting"
	migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestMigrateWhiteBlocklist(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	adminUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{Name: "user1"}).(*user_model.User)
	nonAdminUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{Name: "user2"}).(*user_model.User)

	setting.Migrations.AllowedDomains = "github.com"
	setting.Migrations.AllowLocalNetworks = false
	assert.NoError(t, migrations.Init())

	err := migrations.IsMigrateURLAllowed("https://gitlab.com/gitlab/gitlab.git", nonAdminUser)
	assert.Error(t, err)

	err = migrations.IsMigrateURLAllowed("https://github.com/go-gitea/gitea.git", nonAdminUser)
	assert.NoError(t, err)

	err = migrations.IsMigrateURLAllowed("https://gITHUb.com/go-gitea/gitea.git", nonAdminUser)
	assert.NoError(t, err)

	setting.Migrations.AllowedDomains = ""
	setting.Migrations.BlockedDomains = "github.com"
	assert.NoError(t, migrations.Init())

	err = migrations.IsMigrateURLAllowed("https://gitlab.com/gitlab/gitlab.git", nonAdminUser)
	assert.NoError(t, err)

	err = migrations.IsMigrateURLAllowed("https://github.com/go-gitea/gitea.git", nonAdminUser)
	assert.Error(t, err)

	err = migrations.IsMigrateURLAllowed("https://10.0.0.1/go-gitea/gitea.git", nonAdminUser)
	assert.Error(t, err)

	setting.Migrations.AllowLocalNetworks = true
	assert.NoError(t, migrations.Init())
	err = migrations.IsMigrateURLAllowed("https://10.0.0.1/go-gitea/gitea.git", nonAdminUser)
	assert.NoError(t, err)

	old := setting.ImportLocalPaths
	setting.ImportLocalPaths = false

	err = migrations.IsMigrateURLAllowed("/home/foo/bar/goo", adminUser)
	assert.Error(t, err)

	setting.ImportLocalPaths = true
	abs, err := filepath.Abs(".")
	assert.NoError(t, err)

	err = migrations.IsMigrateURLAllowed(abs, adminUser)
	assert.NoError(t, err)

	err = migrations.IsMigrateURLAllowed(abs, nonAdminUser)
	assert.Error(t, err)

	nonAdminUser.AllowImportLocal = true
	err = migrations.IsMigrateURLAllowed(abs, nonAdminUser)
	assert.NoError(t, err)

	setting.ImportLocalPaths = old
}
