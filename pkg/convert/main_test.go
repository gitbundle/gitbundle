// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert_test

import (
	"path/filepath"
	"testing"
	_ "unsafe"

	api "github.com/gitbundle/api/pkg/structs"
	_ "github.com/gitbundle/server/pkg/access_token"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

//go:linkname toUser github.com/gitbundle/server/pkg/convert.toUser
func toUser(user *user_model.User, signed, authed bool) *api.User
