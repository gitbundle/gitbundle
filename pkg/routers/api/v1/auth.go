// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !windows

package v1

import auth_service "github.com/gitbundle/server/pkg/auth/manager"

func specialAdd(group *auth_service.Group) {}
