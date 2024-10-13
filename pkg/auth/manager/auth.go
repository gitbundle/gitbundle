// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/auth/webauthn"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/session"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/web/middleware"
)

// Init should be called exactly once when the application starts to allow plugins
// to allocate necessary resources
func Init() {
	webauthn.Init()
}

// IsAttachmentDownload check if request is a file download (GET) with URL to an attachment
func IsAttachmentDownload(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/attachments/") && req.Method == "GET"
}

// isContainerPath checks if the request targets the container endpoint
func isContainerPath(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/v2/")
}

var (
	gitRawReleasePathRe = regexp.MustCompile(`^/[a-zA-Z0-9_.-]+/[a-zA-Z0-9_.-]+/(?:(?:git-(?:(?:upload)|(?:receive))-pack$)|(?:info/refs$)|(?:HEAD$)|(?:objects/)|(?:raw/)|(?:releases/download/))`)
	lfsPathRe           = regexp.MustCompile(`^/[a-zA-Z0-9_.-]+/[a-zA-Z0-9_.-]+/info/lfs/`)
)

func isGitRawReleaseOrLFSPath(req *http.Request) bool {
	if gitRawReleasePathRe.MatchString(req.URL.Path) {
		return true
	}
	if setting.LFS.StartServer {
		return lfsPathRe.MatchString(req.URL.Path)
	}
	return false
}

// HandleSignIn clears existing session variables and stores new ones for the specified user object
func HandleSignIn(resp http.ResponseWriter, req *http.Request, sess SessionStore, user *user_model.User) {
	// We need to regenerate the session...
	newSess, err := session.RegenerateSession(resp, req)
	if err != nil {
		log.Error(fmt.Sprintf("Error regenerating session: %v", err))
	} else {
		sess = newSess
	}

	_ = sess.Delete("openid_verified_uri")
	_ = sess.Delete("openid_signin_remember")
	_ = sess.Delete("openid_determined_email")
	_ = sess.Delete("openid_determined_username")
	_ = sess.Delete("twofaUid")
	_ = sess.Delete("twofaRemember")
	_ = sess.Delete("webauthnAssertion")
	_ = sess.Delete("linkAccount")
	err = sess.Set("uid", user.ID)
	if err != nil {
		log.Error(fmt.Sprintf("Error setting session: %v", err))
	}
	err = sess.Set("uname", user.Name)
	if err != nil {
		log.Error(fmt.Sprintf("Error setting session: %v", err))
	}

	// Language setting of the user overwrites the one previously set
	// If the user does not have a locale set, we save the current one.
	if len(user.Language) == 0 {
		lc := middleware.Locale(resp, req)
		user.Language = lc.Language()
		if err := user_model.UpdateUserCols(db.DefaultContext, user, "language"); err != nil {
			log.Error(fmt.Sprintf("Error updating user language [user: %d, locale: %s]", user.ID, user.Language))
			return
		}
	}

	middleware.SetLocaleCookie(resp, user.Language, 0)

	// Clear whatever CSRF has right now, force to generate a new one
	middleware.DeleteCSRFCookie(resp)
}
