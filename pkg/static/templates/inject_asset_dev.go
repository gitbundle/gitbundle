// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build fast_dev

package templates

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
)

var directory string

const (
	emptyAsset = template.HTML("")
	assetsHost = "http://localhost:4001"
)

// LinkTag make css link tag from url
func cssAsset(url string) string {
	return `<link rel="stylesheet" href="` + assetsHost + setting.StaticURLPrefix + `/css/web_src/` + Escape(url) + `?v=` + setting.AssetVersion + `" />`
}

// ScriptTag make js script tag from url
func jsAsset(url string) string {
	return `<script type="module" src="` + assetsHost + setting.StaticURLPrefix + `/js/web_src/` + Escape(url) + `?v=` + setting.AssetVersion + `" onerror="alert('Failed to load asset files from ' + this.src + ', please make sure the asset files can be accessed and the ROOT_URL setting in app.ini is correct.')"></script>`
}

// AssetTag make js or css tag from url
func UrlAssets(url string) template.HTML {
	asset := ""
	ext := filepath.Ext(url)
	switch ext {
	case ".css":
		resp, err := http.Head(assetsHost + setting.StaticURLPrefix + `/css/web_src/` + url)
		if err != nil || resp.StatusCode != http.StatusOK {
			return emptyAsset
		}

		asset = cssAsset(url)
	case ".js":
		resp, err := http.Head(assetsHost + setting.StaticURLPrefix + `/js/web_src/` + url)
		if err != nil || resp.StatusCode != http.StatusOK {
			return emptyAsset
		}

		asset = jsAsset(url)
	default:
		log.Warn("unsupported asset kind: " + ext)
	}

	return template.HTML(asset)
}

func LiveReload() template.HTML {
	return template.HTML(`<script src="http://localhost:35729/livereload.js?host=localhost"></script>`)
}

const (
	showPricing = true
	showApi     = true
)
