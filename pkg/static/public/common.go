// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package public

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gitbundle/modules/httpcache"
	"github.com/gitbundle/modules/log"
)

// Options represents the available options to configure the handler.
type Options struct {
	Directory   string
	Prefix      string
	CorsHandler func(http.Handler) http.Handler
}

// AssetsURLPathPrefix is the path prefix for static asset files
const AssetsURLPathPrefix = "/assets/"

const PluginAssetsURLPathPrefix = "/assets/-"

// parseAcceptEncoding parse Accept-Encoding: deflate, gzip;q=1.0, *;q=0.5 as compress methods
func parseAcceptEncoding(val string) map[string]bool {
	parts := strings.Split(val, ";")
	types := make(map[string]bool)
	for _, v := range strings.Split(parts[0], ",") {
		types[strings.TrimSpace(v)] = true
	}
	return types
}

// setWellKnownContentType will set the Content-Type if the file is a well-known type.
// See the comments of detectWellKnownMimeType
func setWellKnownContentType(w http.ResponseWriter, file string) {
	mimeType := detectWellKnownMimeType(filepath.Ext(file))
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
}

func (opts *Options) handle(w http.ResponseWriter, req *http.Request, fs http.FileSystem, file string) bool {
	// use clean to keep the file is a valid path with no . or ..
	f, err := fs.Open(path.Clean(file))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("[Static] Open %q failed: %v", file, err)
		return true
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("[Static] %q exists, but fails to open: %v", file, err)
		return true
	}

	// Try to serve index file
	if fi.IsDir() {
		w.WriteHeader(http.StatusNotFound)
		return true
	}

	if httpcache.HandleFileETagCache(req, w, fi) {
		return true
	}

	setWellKnownContentType(w, file)

	serveContent(w, req, fi, fi.ModTime(), f)
	return true
}
