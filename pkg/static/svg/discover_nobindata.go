// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !bindata

package svg

import (
	"os"
	"path/filepath"

	"github.com/gitbundle/modules/setting"
)

// Discover returns a map of discovered SVG icons in the file system
func Discover() map[string]string {
	svgs := make(map[string]string)

	files, _ := filepath.Glob(filepath.Join(setting.StaticRootPath, "public", "img", "svg", "*.svg"))
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err == nil {
			filename := filepath.Base(file)
			svgs[filename[:len(filename)-4]] = string(content)
		}
	}

	return svgs
}
