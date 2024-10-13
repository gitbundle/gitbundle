// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package templates

import (
	"html/template"
	"os"
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/util"

	"github.com/unrolled/render"
)

func getDirAssetNames(dir string) []string {
	var tmpls []string
	f, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return tmpls
		}
		log.Warn("Unable to check if templates dir %s is a directory. Error: %v", dir, err)
		return tmpls
	}
	if !f.IsDir() {
		log.Warn("Templates dir %s is a not directory.", dir)
		return tmpls
	}

	files, err := util.StatDir(dir)
	if err != nil {
		log.Warn("Failed to read %s templates dir. %v", dir, err)
		return tmpls
	}
	for _, filePath := range files {
		if strings.HasPrefix(filePath, "mail/") {
			continue
		}

		if !strings.HasSuffix(filePath, ".html") {
			continue
		}

		tmpls = append(tmpls, "templates/"+filePath)
	}
	return tmpls
}

// HTMLRenderer returns a render.
func HTMLRenderer() *render.Render {
	return render.New(render.Options{
		Extensions:                []string{".html"},
		Directory:                 "templates",
		Funcs:                     NewFuncMap(),
		Asset:                     GetAsset,
		AssetNames:                GetAssetNames,
		IsDevelopment:             !setting.IsProd,
		DisableHTTPErrorRendering: true,
	})
}

func IndexCss() template.HTML { return template.HTML(cssAsset("css/index.css")) }

func IndexJs() template.HTML { return template.HTML(jsAsset("js/index.js")) }
