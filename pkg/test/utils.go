// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package test

import (
	"net/http"
)

// RedirectURL returns the redirect URL of a http response.
func RedirectURL(resp http.ResponseWriter) string {
	return resp.Header().Get("Location")
}
