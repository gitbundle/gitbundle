// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package asymkey

import (
	"path/filepath"
	"testing"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/unittest"
)

func init() {
	setting.SetCustomPathAndConf("", "", "")
	setting.LoadForTest()
}

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
		FixtureFiles: []string{
			"gpg_key.yml",
			"public_key.yml",
			"deploy_key.yml",
			"gpg_key_import.yml",
			"user.yml",
			"email_address.yml",
		},
	})
}
