// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package swagger

import (
	api "github.com/gitbundle/api/pkg/structs"
)

// OAuth2Application
// swagger:response OAuth2Application
type swaggerResponseOAuth2Application struct {
	// in:body
	Body api.OAuth2Application `json:"body"`
}

// AccessToken represents an API access token.
// swagger:response AccessToken
type swaggerResponseAccessToken struct {
	// in:body
	Body api.AccessToken `json:"body"`
}
