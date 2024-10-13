// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package db

import (
	"github.com/gitbundle/server/pkg/auth"
	user_model "github.com/gitbundle/server/pkg/user"
)

// Source is a password authentication service
type Source struct{}

// FromDB fills up an OAuth2Config from serialized format.
func (source *Source) FromDB(bs []byte) error {
	return nil
}

// ToDB exports an SMTPConfig to a serialized format.
func (source *Source) ToDB() ([]byte, error) {
	return nil, nil
}

// Authenticate queries if login/password is valid against the PAM,
// and create a local user if success when enabled.
func (source *Source) Authenticate(user *user_model.User, login, password string) (*user_model.User, error) {
	return Authenticate(user, login, password)
}

func init() {
	auth.RegisterTypeConfig(auth.NoType, &Source{})
	auth.RegisterTypeConfig(auth.Plain, &Source{})
}
