// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"strings"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/timeutil"
	access_token_model "github.com/gitbundle/server/pkg/access_token"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/web/middleware"
)

// Ensure the struct implements the interface.
var (
	_ Method = &Basic{}
	_ Named  = &Basic{}
)

// BasicMethodName is the constant name of the basic authentication method
const BasicMethodName = "basic"

// Basic implements the Auth interface and authenticates requests (API requests
// only) by looking for Basic authentication data or "x-oauth-basic" token in the "Authorization"
// header.
type Basic struct{}

// Name represents the name of auth method
func (b *Basic) Name() string {
	return BasicMethodName
}

// Verify extracts and validates Basic data (username and password/token) from the
// "Authorization" header of the request and returns the corresponding user object for that
// name/token on successful validation.
// Returns nil if header is empty or validation fails.
func (b *Basic) Verify(req *http.Request, w http.ResponseWriter, store DataStore, sess SessionStore) *user_model.User {
	// Basic authentication should only fire on API, Download or on Git or LFSPaths
	if !middleware.IsAPIPath(req) && !isContainerPath(req) && !IsAttachmentDownload(req) && !isGitRawReleaseOrLFSPath(req) {
		return nil
	}

	baHead := req.Header.Get("Authorization")
	if len(baHead) == 0 {
		return nil
	}

	auths := strings.SplitN(baHead, " ", 2)
	if len(auths) != 2 || (strings.ToLower(auths[0]) != "basic") {
		return nil
	}

	uname, passwd, _ := base.BasicAuthDecode(auths[1])
	return b.getUserByBasicAuth(uname, passwd, store)
}

func (b *Basic) getUserByBasicAuth(uname, passwd string, store DataStore) *user_model.User {
	// Check if username or password is a token
	isUsernameToken := len(passwd) == 0 || passwd == "x-oauth-basic"
	// Assume username is token
	authToken := uname
	if !isUsernameToken {
		log.Trace("Basic Authorization: Attempting login for: %s", uname)
		// Assume password is token
		authToken = passwd
	} else {
		log.Trace("Basic Authorization: Attempting login with username as token")
	}

	uid := CheckOAuthAccessToken(authToken)
	if uid != 0 {
		log.Trace("Basic Authorization: Valid OAuthAccessToken for user[%d]", uid)

		u, err := user_model.GetUserByID(uid)
		if err != nil {
			log.Error("GetUserByID:  %v", err)
			return nil
		}

		store.GetData()["IsApiToken"] = true
		return u
	}

	token, err := access_token_model.GetAccessTokenBySHA(authToken)
	if err == nil {
		log.Trace("Basic Authorization: Valid AccessToken for user[%d]", uid)
		u, err := user_model.GetUserByID(token.UID)
		if err != nil {
			log.Error("GetUserByID:  %v", err)
			return nil
		}

		token.UpdatedUnix = timeutil.TimeStampNow()
		if err = access_token_model.UpdateAccessToken(token); err != nil {
			log.Error("UpdateAccessToken:  %v", err)
		}

		store.GetData()["IsApiToken"] = true
		return u
	} else if !access_token_model.IsErrAccessTokenNotExist(err) && !access_token_model.IsErrAccessTokenEmpty(err) {
		log.Error("GetAccessTokenBySha: %v", err)
	}

	if !setting.Service.EnableBasicAuth {
		return nil
	}

	u, source, err := UserSignIn(uname, passwd)
	if err != nil {
		if !user_model.IsErrUserNotExist(err) {
			log.Error("UserSignIn: %v", err)
		}
		return nil
	}

	if skipper, ok := source.Cfg.(LocalTwoFASkipper); ok && skipper.IsSkipLocalTwoFA() {
		store.GetData()["SkipLocalTwoFA"] = true
	}

	log.Trace("Basic Authorization: Signin user[%-v]", u.Name)
	return u
}
