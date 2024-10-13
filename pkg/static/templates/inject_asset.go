// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !fast_dev

package templates

import (
	"html/template"

	"github.com/gitbundle/modules/setting"
)

var directory string

const emptyAsset = template.HTML("")

// LinkTag make css link tag from url
func cssAsset(url string) string {
	return `<link rel="stylesheet" href="` + setting.StaticURLPrefix + `/assets/` + Escape(url) + `?v=` + setting.AssetVersion + `" />`
}

// ScriptTag make js script tag from url
func jsAsset(url string) string {
	return `<script type="module" src="` + setting.StaticURLPrefix + `/assets/` + Escape(url) + `?v=` + setting.AssetVersion + `" onerror="alert('Failed to load asset files from ' + this.src + ', please make sure the asset files can be accessed and the ROOT_URL setting in app.ini is correct.')"></script>`
}

// AssetTag make js or css tag from url
func UrlAssets(url string) template.HTML {
	return emptyAsset
}

func LiveReload() template.HTML {
	return emptyAsset
}

const (
	showPricing = true
	showApi     = false
)
