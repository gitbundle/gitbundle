// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package db

import (
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

// Authenticate authenticates the provided user against the DB
func Authenticate(user *user_model.User, login, password string) (*user_model.User, error) {
	if user == nil {
		return nil, user_model.ErrUserNotExist{Name: login}
	}

	if !user.IsPasswordSet() || !user.ValidatePassword(password) {
		return nil, user_model.ErrUserNotExist{UID: user.ID, Name: user.Name}
	}

	// Update password hash if server password hash algorithm have changed
	// Or update the password when the salt length doesn't match the current
	// recommended salt length, this in order to migrate user's salts to a more secure salt.
	if user.PasswdHashAlgo != setting.PasswordHashAlgo || len(user.Salt) != user_model.SaltByteLength*2 {
		if err := user.SetPassword(password); err != nil {
			return nil, err
		}
		if err := user_model.UpdateUserCols(db.DefaultContext, user, "passwd", "passwd_hash_algo", "salt"); err != nil {
			return nil, err
		}
	}

	// WARN: DON'T check user.IsActive, that will be checked on reqSign so that
	// user could be hint to resend confirm email.
	if user.ProhibitLogin {
		return nil, user_model.ErrUserProhibitLogin{
			UID:  user.ID,
			Name: user.Name,
		}
	}

	return user, nil
}
