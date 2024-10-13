// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build bindata

package schemas

import (
	"io"
	"path"
)

func openSchema(filename string) (io.ReadCloser, error) {
	return Assets.Open(path.Base(filename))
}
