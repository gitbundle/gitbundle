// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package oauth2

import (
	"github.com/gitbundle/server/pkg/auth/manager/source/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

// Authenticate falls back to the db authenticator
func (source *Source) Authenticate(user *user_model.User, login, password string) (*user_model.User, error) {
	return db.Authenticate(user, login, password)
}

// NB: Oauth2 does not implement LocalTwoFASkipper for password authentication
// as its password authentication drops to db authentication
