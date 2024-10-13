// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package install

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gitbundle/modules/httpcache"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/forms"
	"github.com/gitbundle/server/pkg/routers/common"
	"github.com/gitbundle/server/pkg/routers/web/healthcheck"
	"github.com/gitbundle/server/pkg/static/public"
	"github.com/gitbundle/server/pkg/static/templates"
	"github.com/gitbundle/server/pkg/web"
	"github.com/gitbundle/server/pkg/web/middleware"

	"gitea.com/go-chi/session"
)

type dataStore map[string]interface{}

func (d *dataStore) GetData() map[string]interface{} {
	return *d
}

func installRecovery() func(next http.Handler) http.Handler {
	rnd := templates.HTMLRenderer()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				// Why we need this? The first recover will try to render a beautiful
				// error page for user, but the process can still panic again, then
				// we have to just recover twice and send a simple error page that
				// should not panic any more.
				defer func() {
					if err := recover(); err != nil {
						combinedErr := fmt.Sprintf("PANIC: %v\n%s", err, log.Stack(2))
						log.Error("%s", combinedErr)
						if setting.IsProd {
							http.Error(
								w,
								http.StatusText(http.StatusInternalServerError),
								http.StatusInternalServerError,
							)
						} else {
							http.Error(w, combinedErr, http.StatusInternalServerError)
						}
					}
				}()

				if err := recover(); err != nil {
					combinedErr := fmt.Sprintf("PANIC: %v\n%s", err, log.Stack(2))
					log.Error("%s", combinedErr)

					lc := middleware.Locale(w, req)
					store := dataStore{
						"Language":       lc.Language(),
						"CurrentURL":     setting.AppSubURL + req.URL.RequestURI(),
						"i18n":           lc,
						"SignedUserID":   int64(0),
						"SignedUserName": "",
					}

					httpcache.AddCacheControlToHeader(w.Header(), 0, "no-transform")
					w.Header().Set(`X-Frame-Options`, setting.CORSConfig.XFrameOptions)

					if !setting.IsProd {
						store["ErrorMsg"] = combinedErr
					}
					err = rnd.HTML(
						w,
						http.StatusInternalServerError,
						"status/500",
						context.BaseVars().Merge(store),
					)
					if err != nil {
						log.Error("%v", err)
					}
				}
			}()

			next.ServeHTTP(w, req)
		})
	}
}

// Routes registers the install routes
func Routes() *web.Route {
	r := web.NewRoute()
	for _, middle := range common.Middlewares() {
		r.Use(middle)
	}

	r.Use(web.WrapWithPrefix(public.AssetsURLPathPrefix, public.AssetsHandlerFunc(&public.Options{
		Directory: path.Join(setting.StaticRootPath, "public"),
		Prefix:    public.AssetsURLPathPrefix,
	}), "InstallAssetsHandler"))

	r.Use(session.Sessioner(session.Options{
		Provider:       setting.SessionConfig.Provider,
		ProviderConfig: setting.SessionConfig.ProviderConfig,
		CookieName:     setting.SessionConfig.CookieName,
		CookiePath:     setting.SessionConfig.CookiePath,
		Gclifetime:     setting.SessionConfig.Gclifetime,
		Maxlifetime:    setting.SessionConfig.Maxlifetime,
		Secure:         setting.SessionConfig.Secure,
		SameSite:       setting.SessionConfig.SameSite,
		Domain:         setting.SessionConfig.Domain,
	}))

	r.Use(installRecovery())
	r.Use(Init)
	r.Get("/", Install)
	r.Post("/", web.Bind(forms.InstallForm{}), SubmitInstall)
	r.Get("/api/healthz", healthcheck.Check)

	r.NotFound(web.Wrap(installNotFound))
	return r
}

func installNotFound(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, setting.AppURL, http.StatusFound)
}
