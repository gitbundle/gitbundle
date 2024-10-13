// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package x

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
)

type reverseProxyWrapper struct {
	target *url.URL
	inner  *httputil.ReverseProxy
}

func NewReverseProxy(target *url.URL, routePrefix string) reverseProxyWrapper {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = JoinURLPath(target, req.URL)
		if routePrefix != "" {
			req.URL.Path = strings.TrimPrefix(req.URL.Path, routePrefix)
			req.URL.RawPath = strings.TrimPrefix(req.URL.RawPath, routePrefix)
		}
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

func SingleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func JoinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return SingleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

var cookieHeaders = map[string]string{
	setting.CSRFCookieName: "X-CSRF-Token",
	"lang":                 "X-Lang",
}

func CopyCookies(r *http.Request) error {
	for k, v := range cookieHeaders {
		c, err := r.Cookie(k)
		if err != nil {
			log.Error("routers/common: error (%v)", err)
			return err
		}
		r.Header.Set(v, c.Value)
	}
	r.Header.Del("Cookie")
	return nil
}
