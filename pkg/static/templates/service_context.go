// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package templates

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	mc "github.com/gitbundle/modules/cache"
	"github.com/gitbundle/modules/httpcache"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/translation"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/static/public"
	"github.com/gitbundle/server/pkg/unit"
	"github.com/gitbundle/server/pkg/web/middleware"

	"gitea.com/go-chi/session"
)

var reverseProxyPathPrefixes = []string{
	public.PluginAssetsURLPathPrefix,
	"/hpi/-",
	"/ppi/-",
	"/marketplace/api",
	"/admin/plugins/assets",
	"/admin/plugins/api",
}

// Contexter initializes a classic context for a request.
func Contexter() func(next http.Handler) http.Handler {
	rnd := HTMLRenderer()
	csrfOpts := getCsrfOpts()
	if !setting.IsProd {
		context.CsrfTokenRegenerationInterval = 5 * time.Second // in dev, re-generate the tokens more aggressively for debug purpose
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			locale := middleware.Locale(resp, req)
			startTime := time.Now()
			link := setting.AppSubURL + strings.TrimSuffix(req.URL.EscapedPath(), "/")

			ctx := context.Context{
				Resp:    context.NewResponse(resp),
				Cache:   mc.GetCache(),
				Locale:  locale,
				Link:    link,
				Render:  rnd,
				Session: session.GetSession(req),
				Repo: &context.Repository{
					PullRequest: &context.PullRequest{},
				},
				Org: &context.Organization{},
				Data: map[string]interface{}{
					"CurrentURL":    setting.AppSubURL + req.URL.RequestURI(),
					"PageStartTime": startTime,
					"Link":          link,
					"RunModeIsProd": setting.IsProd,
				},
			}
			defer ctx.Close()

			// PageData is passed by reference, and it will be rendered to `window.config.pageData` in `head.tmpl` for JavaScript modules
			ctx.PageData = map[string]interface{}{}
			ctx.Data["PageData"] = ctx.PageData
			ctx.Data["Context"] = &ctx

			ctx.Req = context.WithContext(req, &ctx)
			var found bool
			for _, prefix := range reverseProxyPathPrefixes {
				if strings.HasPrefix(req.URL.Path, prefix) || strings.HasPrefix(req.URL.RawPath, prefix) {
					found = true
					break
				}
			}
			if found {
				// For reverse proxy handler
				ctx.ReqRaw = context.WithContext(req.Clone(req.Context()), &ctx)
			}
			ctx.Csrf = context.PrepareCSRFProtector(csrfOpts, &ctx)

			// Get flash.
			flashCookie := ctx.GetCookie("macaron_flash")
			vals, _ := url.ParseQuery(flashCookie)
			if len(vals) > 0 {
				f := &middleware.Flash{
					DataStore:  &ctx,
					Values:     vals,
					ErrorMsg:   vals.Get("error"),
					SuccessMsg: vals.Get("success"),
					InfoMsg:    vals.Get("info"),
					WarningMsg: vals.Get("warning"),
				}
				ctx.Data["Flash"] = f
			}

			f := &middleware.Flash{
				DataStore:  &ctx,
				Values:     url.Values{},
				ErrorMsg:   "",
				WarningMsg: "",
				InfoMsg:    "",
				SuccessMsg: "",
			}
			ctx.Resp.Before(func(resp context.ResponseWriter) {
				if flash := f.Encode(); len(flash) > 0 {
					middleware.SetCookie(resp, "macaron_flash", flash, 0,
						setting.SessionConfig.CookiePath,
						middleware.Domain(setting.SessionConfig.Domain),
						middleware.HTTPOnly(true),
						middleware.Secure(setting.SessionConfig.Secure),
						middleware.SameSite(setting.SessionConfig.SameSite),
					)
					return
				}

				middleware.SetCookie(ctx.Resp, "macaron_flash", "", -1,
					setting.SessionConfig.CookiePath,
					middleware.Domain(setting.SessionConfig.Domain),
					middleware.HTTPOnly(true),
					middleware.Secure(setting.SessionConfig.Secure),
					middleware.SameSite(setting.SessionConfig.SameSite),
				)
			})

			ctx.Flash = f

			// If request sends files, parse them here otherwise the Query() can't be parsed and the CsrfToken will be invalid.
			if ctx.Req.Method == "POST" && strings.Contains(ctx.Req.Header.Get("Content-Type"), "multipart/form-data") {
				if ctx.ReqRaw != nil {
					body, err := io.ReadAll(ctx.Req.Body)
					if err != nil {
						ctx.ServerError("io.ReadAll", err)
						return
					}
					ctx.Req.Body = io.NopCloser(bytes.NewBuffer(body))
					if err := ctx.Req.ParseMultipartForm(setting.Attachment.MaxSize << 20); err != nil && !strings.Contains(err.Error(), "EOF") { // 32MB max size
						ctx.ServerError("ParseMultipartForm", err)
						return
					}
					ctx.ReqRaw.Body = io.NopCloser(bytes.NewBuffer(body))
				} else {
					if err := ctx.Req.ParseMultipartForm(setting.Attachment.MaxSize << 20); err != nil && !strings.Contains(err.Error(), "EOF") { // 32MB max size
						ctx.ServerError("ParseMultipartForm", err)
						return
					}
				}
			}

			httpcache.AddCacheControlToHeader(ctx.Resp.Header(), 0, "no-transform")
			ctx.Resp.Header().Set(`X-Frame-Options`, setting.CORSConfig.XFrameOptions)

			ctx.Data["CsrfToken"] = ctx.Csrf.GetToken()
			ctx.Data["CsrfTokenHtml"] = template.HTML(`<input type="hidden" name="_csrf" value="` + ctx.Data["CsrfToken"].(string) + `">`)

			// FIXME: do we really always need these setting? There should be someway to have to avoid having to always set these
			ctx.Data["IsLandingPageHome"] = setting.LandingPageURL == setting.LandingPageHome
			ctx.Data["IsLandingPageExplore"] = setting.LandingPageURL == setting.LandingPageExplore
			ctx.Data["IsLandingPageOrganizations"] = setting.LandingPageURL == setting.LandingPageOrganizations

			ctx.Data["ShowRegistrationButton"] = setting.Service.ShowRegistrationButton
			ctx.Data["ShowMilestonesDashboardPage"] = setting.Service.ShowMilestonesDashboardPage
			ctx.Data["ShowFooterBranding"] = setting.ShowFooterBranding
			ctx.Data["ShowFooterVersion"] = setting.ShowFooterVersion

			ctx.Data["EnableSwagger"] = setting.API.EnableSwagger
			ctx.Data["EnableOpenIDSignIn"] = setting.Service.EnableOpenIDSignIn
			ctx.Data["DisableMigrations"] = setting.Repository.DisableMigrations
			ctx.Data["DisableStars"] = setting.Repository.DisableStars
			ctx.Data["ManifestData"] = setting.ManifestData

			ctx.Data["UnitWikiGlobalDisabled"] = unit.TypeWiki.UnitGlobalDisabled()
			ctx.Data["UnitIssuesGlobalDisabled"] = unit.TypeIssues.UnitGlobalDisabled()
			ctx.Data["UnitPullsGlobalDisabled"] = unit.TypePullRequests.UnitGlobalDisabled()
			ctx.Data["UnitProjectsGlobalDisabled"] = unit.TypeProjects.UnitGlobalDisabled()

			ctx.Data["i18n"] = locale
			ctx.Data["AllLangs"] = translation.AllLangs()

			next.ServeHTTP(ctx.Resp, ctx.Req)

			// Handle adding signedUserName to the context for the AccessLogger
			usernameInterface := ctx.Data["SignedUserName"]
			identityPtrInterface := ctx.Req.Context().Value(context.SignedUserNameStringPointerKey)
			if usernameInterface != nil && identityPtrInterface != nil {
				username := usernameInterface.(string)
				identityPtr := identityPtrInterface.(*string)
				if identityPtr != nil && username != "" {
					*identityPtr = username
				}
			}
		})
	}
}

func getCsrfOpts() context.CsrfOptions {
	return context.CsrfOptions{
		Secret:         setting.SecretKey,
		Cookie:         setting.CSRFCookieName,
		SetCookie:      true,
		Secure:         setting.SessionConfig.Secure,
		CookieHTTPOnly: setting.CSRFCookieHTTPOnly,
		Header:         "X-Csrf-Token",
		CookieDomain:   setting.SessionConfig.Domain,
		CookiePath:     setting.SessionConfig.CookiePath,
		SameSite:       setting.SessionConfig.SameSite,
	}
}

// PackageContexter initializes a package context for a request.
func PackageContexter() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			ctx := context.Context{
				Resp:   context.NewResponse(resp),
				Data:   map[string]interface{}{},
				Render: HTMLRenderer(),
			}
			defer ctx.Close()

			ctx.Req = context.WithContext(req, &ctx)

			next.ServeHTTP(ctx.Resp, ctx.Req)
		})
	}
}
