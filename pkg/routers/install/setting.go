// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package install

import (
	"context"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/translation"
	"github.com/gitbundle/server/pkg/routers/common"
	"github.com/gitbundle/server/pkg/static/options"
	"github.com/gitbundle/server/pkg/static/svg"
)

// PreloadSettings preloads the configuration to check if we need to run install
func PreloadSettings(ctx context.Context) bool {
	setting.LoadAllowEmpty()
	if !setting.InstallLock {
		log.Info("AppPath: %s", setting.AppPath)
		log.Info("AppWorkPath: %s", setting.AppWorkPath)
		log.Info("Custom path: %s", setting.CustomPath)
		log.Info("Log path: %s", setting.LogRootPath)
		log.Info("Configuration file: %s", setting.CustomConf)
		log.Info("Prepare to run install page")
		translation.InitLocales(options.Dir, options.Locale)
		if setting.EnableSQLite3 {
			log.Info("SQLite3 is supported")
		}
		setting.InitDBConfig()
		setting.NewServicesForInstall()
		svg.Init()
	}

	return !setting.InstallLock
}

// reloadSettings reloads the existing settings and starts up the database
func reloadSettings(ctx context.Context) {
	setting.LoadFromExisting()
	setting.InitDBConfig()
	if setting.InstallLock {
		if err := common.InitDBEngine(ctx); err == nil {
			log.Info("ORM engine initialization successful!")
		} else {
			log.Fatal("ORM engine initialization failed: %v", err)
		}
		svg.Init()
	}
}
