// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package forms

import (
	"net/http"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/web/middleware"

	"gitea.com/go-chi/binding"
)

// AdminCreateUserForm form for admin to create user
type AdminCreateUserForm struct {
	LoginType          string `binding:"Required"`
	LoginName          string
	UserName           string `binding:"Required;AlphaDashDot;MaxSize(40)"`
	Email              string `binding:"Required;Email;MaxSize(254)"`
	Password           string `binding:"MaxSize(255)"`
	SendNotify         bool
	MustChangePassword bool
	Visibility         structs.VisibleType
}

// Validate validates form fields
func (f *AdminCreateUserForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// AdminEditUserForm form for admin to create user
type AdminEditUserForm struct {
	LoginType               string `binding:"Required"`
	UserName                string `binding:"AlphaDashDot;MaxSize(40)"`
	LoginName               string
	FullName                string `binding:"MaxSize(100)"`
	Email                   string `binding:"Required;Email;MaxSize(254)"`
	Password                string `binding:"MaxSize(255)"`
	Website                 string `binding:"ValidUrl;MaxSize(255)"`
	Location                string `binding:"MaxSize(50)"`
	MaxRepoCreation         int
	Active                  bool
	Admin                   bool
	Restricted              bool
	AllowGitHook            bool
	AllowImportLocal        bool
	AllowCreateOrganization bool
	ProhibitLogin           bool
	Reset2FA                bool `                                      form:"reset_2fa"`
	Visibility              structs.VisibleType
}

// Validate validates form fields
func (f *AdminEditUserForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// AdminDashboardForm form for admin dashboard operations
type AdminDashboardForm struct {
	Op   string `binding:"required"`
	From string
}

// Validate validates form fields
func (f *AdminDashboardForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

type PluginConfig struct {
	ServerAddress string `binding:"Required" form:"serverAddress" json:"serverAddress"`
	PluginIcon    string `form:"pluginIcon" json:"pluginIcon"`
	Arch          string `form:"arch" json:"arch"`
	Platform      string `form:"platform" json:"platform"`
	RouteIndex    int32  `form:"routeIndex" json:"routeIndex"`
}

func (f *PluginConfig) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

/*
func (f *AddKubeconfig) patchKubeClusters(form url.Values) (errors binding.Errors) {
	for key, values := range form {
		re := regexp.MustCompile("(.*)\\[([0-9]+)\\]\\[(.*)\\]")
		matches := re.FindStringSubmatch(key)

		var bindErr binding.Error
		if len(matches) == 4 {
			if matches[1] != "kube_clusters" {
				bindErr.Message = fmt.Sprintf("Unknown field (%v)", matches[1])
				continue
			}

			index, err := strconv.Atoi(matches[2])
			if err != nil {
				bindErr.Message = fmt.Sprintf("strconv.Atoi (%v), err: %v", matches[2], err)
				continue
			}

			for index >= len(f.KubeClusters) {
				f.KubeClusters = append(f.KubeClusters, deployment.KubeClusterMetainfo{})
			}

			switch matches[3] {
			case "name":
				f.KubeClusters[index].Name = values[0]
			case "priority":
				priority, err := strconv.Atoi(values[0])
				if err != nil {
					bindErr.Message = fmt.Sprintf("strconv.Atoi priority (%v), err: %v", values[0], err)
				} else {
					f.KubeClusters[index].Priority = priority
				}
			case "mode":
				mode := perm.ParseAccessMode(values[0])
				if mode == perm.AccessModeNone {
					bindErr.Message = fmt.Sprintf("Unknown field (%s), mode (%v)", matches[3], values[0])
				} else {
					f.KubeClusters[index].Mode = mode
				}
			default:
				bindErr.Message = fmt.Sprintf("Unknown field (%s), value (%#v)", matches[3], values)
			}
		} else {
			bindErr.Message = fmt.Sprintf("Unmatched (%s): (%v)", key, values)
		}

		if bindErr.Message != "" {
			errors = append(errors, bindErr)
		}
	}

	return errors
}

func (f *AddKubeconfig) Bind() http.HandlerFunc {
	return web.Wrap(func(ctx *context.Context) {
		binding.Bind(ctx.Req, f)
		f.patchKubeClusters(ctx.Req.Form)
		web.SetForm(ctx, f)
		middleware.AssignForm(f, ctx.Data)
	})
}
*/
