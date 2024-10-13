// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build bindata

package svg

import (
	"path/filepath"

	"github.com/gitbundle/server/pkg/static/public"
)

// Discover returns a map of discovered SVG icons in bindata
func Discover() map[string]string {
	svgs := make(map[string]string)

	for _, file := range public.AssetNames() {
		matched, _ := filepath.Match("img/svg/*.svg", file)
		if matched {
			content, err := public.Asset(file)
			if err == nil {
				filename := filepath.Base(file)
				svgs[filename[:len(filename)-4]] = string(content)
			}
		}
	}

	return svgs
}
