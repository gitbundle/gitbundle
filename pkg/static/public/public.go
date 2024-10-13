// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !fast_dev

package public

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gitbundle/modules/setting"
)

// AssetsHandlerFunc implements the static handler for serving custom or original assets.
func AssetsHandlerFunc(opts *Options) http.HandlerFunc {
	custPath := filepath.Join(setting.CustomPath, "public")
	if !filepath.IsAbs(custPath) {
		custPath = filepath.Join(setting.AppWorkPath, custPath)
	}
	if !filepath.IsAbs(opts.Directory) {
		opts.Directory = filepath.Join(setting.AppWorkPath, opts.Directory)
	}
	if !strings.HasSuffix(opts.Prefix, "/") {
		opts.Prefix += "/"
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" && req.Method != "HEAD" {
			resp.WriteHeader(http.StatusNotFound)
			return
		}

		file := req.URL.Path
		file = file[len(opts.Prefix):]
		if len(file) == 0 {
			resp.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.Contains(file, "\\") {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		file = "/" + file

		var written bool
		if opts.CorsHandler != nil {
			written = true
			opts.CorsHandler(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				written = false
			})).ServeHTTP(resp, req)
		}
		if written {
			return
		}

		// custom files
		if opts.handle(resp, req, http.Dir(custPath), file) {
			return
		}

		// internal files
		if opts.handle(resp, req, fileSystem(opts.Directory), file) {
			return
		}

		resp.WriteHeader(http.StatusNotFound)
	}
}
