// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package routers

import (
	"context"
	"reflect"
	"runtime"

	"github.com/gitbundle/modules/cache"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/highlight"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/markup"
	"github.com/gitbundle/modules/markup/external"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/storage"
	"github.com/gitbundle/modules/translation"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/appstate"
	asymkey_model "github.com/gitbundle/server/pkg/asymkey"
	auth "github.com/gitbundle/server/pkg/auth/manager"
	"github.com/gitbundle/server/pkg/auth/manager/source/oauth2"
	"github.com/gitbundle/server/pkg/cron"
	"github.com/gitbundle/server/pkg/eventsource"
	code_indexer "github.com/gitbundle/server/pkg/indexer/code"
	issue_indexer "github.com/gitbundle/server/pkg/indexer/issues"
	stats_indexer "github.com/gitbundle/server/pkg/indexer/stats"
	"github.com/gitbundle/server/pkg/mailer"
	repo_migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/notification"
	"github.com/gitbundle/server/pkg/pull/automerge"
	pull_service "github.com/gitbundle/server/pkg/pull/manager"
	mirror_service "github.com/gitbundle/server/pkg/repo/mirror"
	"github.com/gitbundle/server/pkg/repo/repoman"
	repo_service "github.com/gitbundle/server/pkg/repo/repository_manager"
	"github.com/gitbundle/server/pkg/repo/repository_manager/archiver"
	packages_router "github.com/gitbundle/server/pkg/routers/api/packages"
	apiv1 "github.com/gitbundle/server/pkg/routers/api/v1"
	"github.com/gitbundle/server/pkg/routers/common"
	"github.com/gitbundle/server/pkg/routers/private"
	web_routers "github.com/gitbundle/server/pkg/routers/web"
	"github.com/gitbundle/server/pkg/ssh"
	"github.com/gitbundle/server/pkg/static/options"
	"github.com/gitbundle/server/pkg/static/svg"
	task "github.com/gitbundle/server/pkg/task/manager"
	"github.com/gitbundle/server/pkg/web"
	webhook "github.com/gitbundle/server/pkg/webhook/manager"
	"github.com/gitbundle/server/pkg/x"

	_ "github.com/gitbundle/server/pkg" // register all models

	"github.com/go-chi/chi/v5/middleware"
)

func mustInit(fn func() error) {
	err := fn()
	if err != nil {
		ptr := reflect.ValueOf(fn).Pointer()
		fi := runtime.FuncForPC(ptr)
		log.Fatal("%s failed: %v", fi.Name(), err)
	}
}

func mustInitCtx(ctx context.Context, fn func(ctx context.Context) error) {
	err := fn(ctx)
	if err != nil {
		ptr := reflect.ValueOf(fn).Pointer()
		fi := runtime.FuncForPC(ptr)
		log.Fatal("%s(ctx) failed: %v", fi.Name(), err)
	}
}

// InitGitServices init new services for git, this is also called in `contrib/pr/checkout.go`
func InitGitServices() {
	setting.NewServices()
	mustInit(storage.Init)
	mustInit(repo_service.Init)
}

func syncAppPathForGit(ctx context.Context) error {
	runtimeState := new(appstate.RuntimeState)
	if err := appstate.AppState2.Get(runtimeState); err != nil {
		return err
	}
	if runtimeState.LastAppPath != setting.AppPath {
		log.Info("AppPath changed from '%s' to '%s'", runtimeState.LastAppPath, setting.AppPath)

		log.Info("re-sync repository hooks ...")
		mustInitCtx(ctx, repo_service.SyncRepositoryHooks)

		log.Info("re-write ssh public keys ...")
		mustInit(asymkey_model.RewriteAllPublicKeys)

		runtimeState.LastAppPath = setting.AppPath
		return appstate.AppState2.Set(runtimeState)
	}
	return nil
}

// GlobalInitInstalled is for global installed configuration.
func GlobalInitInstalled(ctx context.Context) {
	if !setting.InstallLock {
		log.Fatal("GitBundle is not installed")
	}

	mustInitCtx(ctx, git.InitOnceWithSync)
	log.Info("Git Version: %s (home: %s)", git.VersionInfo(), git.HomeDir())

	git.CheckLFSVersion()
	log.Info("AppPath: %s", setting.AppPath)
	log.Info("AppWorkPath: %s", setting.AppWorkPath)
	log.Info("Custom path: %s", setting.CustomPath)
	log.Info("Log path: %s", setting.LogRootPath)
	log.Info("Configuration file: %s", setting.CustomConf)
	log.Info("Run Mode: %s", util.ToTitleCase(setting.RunMode))

	// Setup i18n
	translation.InitLocales(options.Dir, options.Locale)

	setting.NewServices()
	mustInit(storage.Init)

	mailer.NewContext()
	mustInit(cache.NewContext)
	notification.NewContext()
	mustInit(archiver.Init)

	highlight.NewContext()
	external.RegisterRenderers()
	markup.Init()

	if setting.EnableSQLite3 {
		log.Info("SQLite3 support is enabled")
	} else if setting.Database.UseSQLite3 {
		log.Fatal("SQLite3 support is disabled, but it is used for database setting. Please get or build a GitBundle release with SQLite3 support.")
	}

	mustInitCtx(ctx, common.InitDBEngine)
	log.Info("ORM engine initialization successful!")
	mustInit(appstate.Init)
	mustInit(oauth2.Init)

	// mustInit(x.InitCIRpcClient)
	mustInit(x.InitRedisClient)

	repoman.NewRepoContext()
	mustInit(repo_service.Init)

	// Booting long running goroutines.
	issue_indexer.InitIssueIndexer(false)
	code_indexer.Init()
	mustInit(stats_indexer.Init)

	mirror_service.InitSyncMirrors()
	mustInit(webhook.Init)
	mustInit(pull_service.Init)
	mustInit(automerge.Init)
	mustInit(task.Init)
	mustInit(repo_migrations.Init)
	eventsource.GetManager().Init()

	mustInitCtx(ctx, syncAppPathForGit)

	mustInit(ssh.Init)

	auth.Init()
	svg.Init()

	// NOTE:
	// At last, register the bundled service, like plugin and k8s deployment helper service
	// need redis and nsq configured first
	mustInit(x.Init)

	// Finally start up the cron
	cron.NewContext(ctx)
}

// NormalRoutes represents non install routes
func NormalRoutes() *web.Route {
	r := web.NewRoute()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(middleware.StripSlashes)
	//r.Use(middleware.Logger)

	for _, middle := range common.Middlewares() {
		r.Use(middle)
	}

	r.Mount("/", web_routers.Routes())
	r.Mount(common.ApiV1Prefix, apiv1.Routes())
	r.Mount("/api/internal", private.Routes())

	// TODO: When we have checked, we remove this line
	// if setting.Packages.Enabled {
	r.Mount("/api/packages", packages_router.Routes())
	r.Mount("/v2", packages_router.ContainerRoutes())
	// }
	return r
}
