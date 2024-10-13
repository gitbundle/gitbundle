// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/auth"
	"github.com/gitbundle/server/pkg/auth/manager/source/oauth2"
	"github.com/gitbundle/server/pkg/auth/manager/source/smtp"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"

	_ "github.com/gitbundle/server/pkg/auth/manager/source/db"   // register the sources (and below)
	_ "github.com/gitbundle/server/pkg/auth/manager/source/ldap" // register the ldap source
	_ "github.com/gitbundle/server/pkg/auth/manager/source/pam"  // register the pam source
	_ "github.com/gitbundle/server/pkg/auth/manager/source/sspi" // register the sspi source
)

// UserSignIn validates user name and password.
func UserSignIn(username, password string) (*user_model.User, *auth.Source, error) {
	var user *user_model.User
	isEmail := false
	if strings.Contains(username, "@") {
		isEmail = true
		emailAddress := user_model.EmailAddress{LowerEmail: strings.ToLower(strings.TrimSpace(username))}
		// check same email
		has, err := db.GetEngine(db.DefaultContext).Get(&emailAddress)
		if err != nil {
			return nil, nil, err
		}
		if has {
			if !emailAddress.IsActivated {
				return nil, nil, user_model.ErrEmailAddressNotExist{
					Email: username,
				}
			}
			user = &user_model.User{ID: emailAddress.UID}
		}
	} else {
		trimmedUsername := strings.TrimSpace(username)
		if len(trimmedUsername) == 0 {
			return nil, nil, user_model.ErrUserNotExist{Name: username}
		}

		user = &user_model.User{LowerName: strings.ToLower(trimmedUsername)}
	}

	if user != nil {
		hasUser, err := user_model.GetUser(user)
		if err != nil {
			return nil, nil, err
		}

		if hasUser {
			source, err := auth.GetSourceByID(user.LoginSource)
			if err != nil {
				return nil, nil, err
			}

			if !source.IsActive {
				return nil, nil, oauth2.ErrAuthSourceNotActived
			}

			authenticator, ok := source.Cfg.(PasswordAuthenticator)
			if !ok {
				return nil, nil, smtp.ErrUnsupportedLoginType
			}

			user, err := authenticator.Authenticate(user, user.LoginName, password)
			if err != nil {
				return nil, nil, err
			}

			// WARN: DON'T check user.IsActive, that will be checked on reqSign so that
			// user could be hint to resend confirm email.
			if user.ProhibitLogin {
				return nil, nil, user_model.ErrUserProhibitLogin{UID: user.ID, Name: user.Name}
			}

			return user, source, nil
		}
	}

	sources, err := auth.AllActiveSources()
	if err != nil {
		return nil, nil, err
	}

	for _, source := range sources {
		if !source.IsActive {
			// don't try to authenticate non-active sources
			continue
		}

		authenticator, ok := source.Cfg.(PasswordAuthenticator)
		if !ok {
			continue
		}

		authUser, err := authenticator.Authenticate(nil, username, password)

		if err == nil {
			if !authUser.ProhibitLogin {
				return authUser, source, nil
			}
			err = user_model.ErrUserProhibitLogin{UID: authUser.ID, Name: authUser.Name}
		}

		if user_model.IsErrUserNotExist(err) {
			log.Debug("Failed to login '%s' via '%s': %v", username, source.Name, err)
		} else {
			log.Warn("Failed to login '%s' via '%s': %v", username, source.Name, err)
		}
	}

	if isEmail {
		return nil, nil, user_model.ErrEmailAddressNotExist{Email: username}
	}

	return nil, nil, user_model.ErrUserNotExist{Name: username}
}
