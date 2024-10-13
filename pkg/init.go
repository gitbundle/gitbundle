// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package pkg

import (
	_ "github.com/gitbundle/server/pkg/access_token"
	_ "github.com/gitbundle/server/pkg/action"
	_ "github.com/gitbundle/server/pkg/admin"
	_ "github.com/gitbundle/server/pkg/appstate"
	_ "github.com/gitbundle/server/pkg/asymkey"
	_ "github.com/gitbundle/server/pkg/auth"
	_ "github.com/gitbundle/server/pkg/avatars"
	_ "github.com/gitbundle/server/pkg/db"
	_ "github.com/gitbundle/server/pkg/eventsource"
	_ "github.com/gitbundle/server/pkg/foreignreference"
	_ "github.com/gitbundle/server/pkg/git"
	_ "github.com/gitbundle/server/pkg/issues"
	_ "github.com/gitbundle/server/pkg/organization"
	_ "github.com/gitbundle/server/pkg/packages"
	_ "github.com/gitbundle/server/pkg/perm/access"
	_ "github.com/gitbundle/server/pkg/project"
	_ "github.com/gitbundle/server/pkg/pull"
	_ "github.com/gitbundle/server/pkg/repo"
	_ "github.com/gitbundle/server/pkg/repo/release"
	_ "github.com/gitbundle/server/pkg/repo/repoman"
	_ "github.com/gitbundle/server/pkg/session"
	_ "github.com/gitbundle/server/pkg/task"
	_ "github.com/gitbundle/server/pkg/unittest"
	_ "github.com/gitbundle/server/pkg/user"
	_ "github.com/gitbundle/server/pkg/webhook"
)
