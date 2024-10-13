// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build fast_dev

package public

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/x"
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

	assetsDevServerAddr, _ := url.Parse("http://localhost:4001")
	rs := NewAssetsReverseProxy(assetsDevServerAddr, strings.TrimSuffix(opts.Prefix, "/"))

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

		rs.ServeHTTP(resp, req)
	}
}

type reverseProxyWrapper struct {
	target *url.URL
	prefix string
	inner  *httputil.ReverseProxy
}

func NewAssetsReverseProxy(target *url.URL, trimPrefix string) reverseProxyWrapper {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = x.JoinURLPath(target, req.URL)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, trimPrefix)
		req.URL.RawPath = strings.TrimPrefix(req.URL.RawPath, trimPrefix)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return reverseProxyWrapper{
		target: target,
		prefix: trimPrefix,
		inner:  &httputil.ReverseProxy{Director: director},
	}
}

func (h reverseProxyWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(error); ok {
				if errors.Is(err, http.ErrAbortHandler) {
					return
				}
			}
			panic(p)
		}
	}()

	h.inner.ServeHTTP(w, r)
}
