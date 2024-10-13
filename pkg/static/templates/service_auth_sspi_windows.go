// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package templates

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/auth"
	authm "github.com/gitbundle/server/pkg/auth/manager"
	"github.com/gitbundle/server/pkg/auth/manager/source/sspi"
	"github.com/gitbundle/server/pkg/avatars"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/mailer"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/web/middleware"

	gouuid "github.com/google/uuid"
	"github.com/quasoft/websspi"
	"github.com/unrolled/render"
)

const (
	tplSignIn base.TplName = "user/auth/signin"
)

var (
	// sspiAuth is a global instance of the websspi authentication package,
	// which is used to avoid acquiring the server credential handle on
	// every request
	sspiAuth *websspi.Authenticator

	// Ensure the struct implements the interface.
	_ authm.Method        = &SSPI{}
	_ authm.Named         = &SSPI{}
	_ authm.Initializable = &SSPI{}
	_ authm.Freeable      = &SSPI{}
)

// SSPI implements the SingleSignOn interface and authenticates requests
// via the built-in SSPI module in Windows for SPNEGO authentication.
// On successful authentication returns a valid user object.
// Returns nil if authentication fails.
type SSPI struct {
	rnd *render.Render
}

// Init creates a new global websspi.Authenticator object
func (s *SSPI) Init() error {
	config := websspi.NewConfig()
	var err error
	sspiAuth, err = websspi.New(config)
	if err != nil {
		return err
	}
	s.rnd = render.New(render.Options{
		Extensions:    []string{".html"},
		Directory:     "templates",
		Funcs:         NewFuncMap(),
		Asset:         GetAsset,
		AssetNames:    GetAssetNames,
		IsDevelopment: !setting.IsProd,
	})
	return nil
}

// Name represents the name of auth method
func (s *SSPI) Name() string {
	return "sspi"
}

// Free releases resources used by the global websspi.Authenticator object
func (s *SSPI) Free() error {
	return sspiAuth.Free()
}

// Verify uses SSPI (Windows implementation of SPNEGO) to authenticate the request.
// If authentication is successful, returns the corresponding user object.
// If negotiation should continue or authentication fails, immediately returns a 401 HTTP
// response code, as required by the SPNEGO protocol.
func (s *SSPI) Verify(req *http.Request, w http.ResponseWriter, store authm.DataStore, sess authm.SessionStore) *user_model.User {
	if !s.shouldAuthenticate(req) {
		return nil
	}

	cfg, err := s.getConfig()
	if err != nil {
		log.Error("could not get SSPI config: %v", err)
		return nil
	}

	log.Trace("SSPI Authorization: Attempting to authenticate")
	userInfo, outToken, err := sspiAuth.Authenticate(req, w)
	if err != nil {
		log.Warn("Authentication failed with error: %v\n", err)
		sspiAuth.AppendAuthenticateHeader(w, outToken)

		// Include the user login page in the 401 response to allow the user
		// to login with another authentication method if SSPI authentication
		// fails
		store.GetData()["Flash"] = map[string]string{
			"ErrorMsg": err.Error(),
		}
		store.GetData()["EnableOpenIDSignIn"] = setting.Service.EnableOpenIDSignIn
		store.GetData()["EnableSSPI"] = true

		err := s.rnd.HTML(w, http.StatusUnauthorized, string(tplSignIn), context.BaseVars().Merge(store.GetData()))
		if err != nil {
			log.Error("%v", err)
		}

		return nil
	}
	if outToken != "" {
		sspiAuth.AppendAuthenticateHeader(w, outToken)
	}

	username := sanitizeUsername(userInfo.Username, cfg)
	if len(username) == 0 {
		return nil
	}
	log.Info("Authenticated as %s\n", username)

	user, err := user_model.GetUserByName(req.Context(), username)
	if err != nil {
		if !user_model.IsErrUserNotExist(err) {
			log.Error("GetUserByName: %v", err)
			return nil
		}
		if !cfg.AutoCreateUsers {
			log.Error("User '%s' not found", username)
			return nil
		}
		user, err = s.newUser(username, cfg)
		if err != nil {
			log.Error("CreateUser: %v", err)
			return nil
		}
	}

	// Make sure requests to API paths and PWA resources do not create a new session
	if !middleware.IsAPIPath(req) && !authm.IsAttachmentDownload(req) {
		authm.HandleSignIn(w, req, sess, user)
	}

	log.Trace("SSPI Authorization: Logged in user %-v", user)
	return user
}

// getConfig retrieves the SSPI configuration from login sources
func (s *SSPI) getConfig() (*sspi.Source, error) {
	sources, err := auth.ActiveSources(auth.SSPI)
	if err != nil {
		return nil, err
	}
	if len(sources) == 0 {
		return nil, errors.New("no active login sources of type SSPI found")
	}
	if len(sources) > 1 {
		return nil, errors.New("more than one active login source of type SSPI found")
	}
	return sources[0].Cfg.(*sspi.Source), nil
}

func (s *SSPI) shouldAuthenticate(req *http.Request) (shouldAuth bool) {
	shouldAuth = false
	path := strings.TrimSuffix(req.URL.Path, "/")
	if path == "/user/login" {
		if req.FormValue("user_name") != "" && req.FormValue("password") != "" {
			shouldAuth = false
		} else if req.FormValue("auth_with_sspi") == "1" {
			shouldAuth = true
		}
	} else if middleware.IsAPIPath(req) || authm.IsAttachmentDownload(req) {
		shouldAuth = true
	}
	return
}

// newUser creates a new user object for the purpose of automatic registration
// and populates its name and email with the information present in request headers.
func (s *SSPI) newUser(username string, cfg *sspi.Source) (*user_model.User, error) {
	email := gouuid.New().String() + "@localhost.localdomain"
	user := &user_model.User{
		Name:            username,
		Email:           email,
		Passwd:          gouuid.New().String(),
		Language:        cfg.DefaultLanguage,
		UseCustomAvatar: true,
		Avatar:          avatars.DefaultAvatarLink(),
	}
	emailNotificationPreference := user_model.EmailNotificationsDisabled
	overwriteDefault := &user_model.CreateUserOverwriteOptions{
		IsActive:                     util.OptionalBoolOf(cfg.AutoActivateUsers),
		KeepEmailPrivate:             util.OptionalBoolTrue,
		EmailNotificationsPreference: &emailNotificationPreference,
	}
	if err := user_model.CreateUser(user, overwriteDefault); err != nil {
		return nil, err
	}

	mailer.SendRegisterNotifyMail(user)

	return user, nil
}

// stripDomainNames removes NETBIOS domain name and separator from down-level logon names
// (eg. "DOMAIN\user" becomes "user"), and removes the UPN suffix (domain name) and separator
// from UPNs (eg. "user@domain.local" becomes "user")
func stripDomainNames(username string) string {
	if strings.Contains(username, "\\") {
		parts := strings.SplitN(username, "\\", 2)
		if len(parts) > 1 {
			username = parts[1]
		}
	} else if strings.Contains(username, "@") {
		parts := strings.Split(username, "@")
		if len(parts) > 1 {
			username = parts[0]
		}
	}
	return username
}

func replaceSeparators(username string, cfg *sspi.Source) string {
	newSep := cfg.SeparatorReplacement
	username = strings.ReplaceAll(username, "\\", newSep)
	username = strings.ReplaceAll(username, "/", newSep)
	username = strings.ReplaceAll(username, "@", newSep)
	return username
}

func sanitizeUsername(username string, cfg *sspi.Source) string {
	if len(username) == 0 {
		return ""
	}
	if cfg.StripDomainNames {
		username = stripDomainNames(username)
	}
	// Replace separators even if we have already stripped the domain name part,
	// as the username can contain several separators: eg. "MICROSOFT\useremail@live.com"
	username = replaceSeparators(username, cfg)
	return username
}
