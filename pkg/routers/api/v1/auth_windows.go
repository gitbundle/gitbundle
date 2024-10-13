// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gitbundle/server/pkg/auth"
	authm "github.com/gitbundle/server/pkg/auth/manager"
	"github.com/gitbundle/server/pkg/static/templates"
)

// specialAdd registers the SSPI auth method as the last method in the list.
// The SSPI plugin is expected to be executed last, as it returns 401 status code if negotiation
// fails (or if negotiation should continue), which would prevent other authentication methods
// to execute at all.
func specialAdd(group *authm.Group) {
	if auth.IsSSPIEnabled() {
		group.Add(&templates.SSPI{})
	}
}
