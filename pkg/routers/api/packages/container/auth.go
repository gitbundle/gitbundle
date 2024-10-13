// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package container

import (
	"net/http"

	"github.com/gitbundle/modules/log"
	auth "github.com/gitbundle/server/pkg/auth/manager"
	packages "github.com/gitbundle/server/pkg/packages/manager"
	user_model "github.com/gitbundle/server/pkg/user"
)

type Auth struct{}

func (a *Auth) Name() string {
	return "container"
}

// Verify extracts the user from the Bearer token
// If it's an anonymous session a ghost user is returned
func (a *Auth) Verify(req *http.Request, w http.ResponseWriter, store auth.DataStore, sess auth.SessionStore) *user_model.User {
	uid, err := packages.ParseAuthorizationToken(req)
	if err != nil {
		log.Trace("ParseAuthorizationToken: %v", err)
		return nil
	}

	if uid == 0 {
		return nil
	}
	if uid == -1 {
		return user_model.NewGhostUser()
	}

	u, err := user_model.GetUserByID(uid)
	if err != nil {
		log.Error("GetUserByID:  %v", err)
		return nil
	}

	return u
}
