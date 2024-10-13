// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build vendor

package main

// Libraries that are included to vendor utilities used during build.
// These libraries will not be included in a normal compilation.

import (
	// for embed
	_ "github.com/shurcooL/vfsgen"

	// for cover merge
	_ "golang.org/x/tools/cover"

	// for vet
	_ "code.gitea.io/gitea-vet"

	// for swagger
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)
