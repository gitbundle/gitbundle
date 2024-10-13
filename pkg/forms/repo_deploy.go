// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package forms

import (
	"net/http"

	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/web/middleware"

	"gitea.com/go-chi/binding"
)

//type DeploySettingForm struct {
//	RegistryHost   string
//	DockerfilePath string
//	HasDockerfile  bool
//	BranchChange   string
//	DeployBranch   string
//	DeployPort     int64
//	DeployProtocol deployment.Protocol
//}
//
//func (f *DeploySettingForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
//	ctx := context.GetContext(req)
//	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
//}

//type DeploySettingMetaInfos struct {
//	Metainfo structs.DeployMetainfo `json:"metainfo"`
//}
//
//func (f *DeploySettingMetaInfos) Validate(req *http.Request, errs binding.Errors) binding.Errors {
//	ctx := context.GetContext(req)
//	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
//}

//type DeployLockForm struct {
//	Reason string
//}
//
//func (f *DeployLockForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
//	ctx := context.GetContext(req)
//	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
//}

type SimplePullRequestForm struct {
	Index int64  `json:"index"`
	Sha   string `json:"sha"`
}

func (f *SimplePullRequestForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// type MetricsQueryOpts struct {
// 	Category     string `json:"category"`
// 	Nodes        string `json:"nodes"`
// 	Pods         string `json:"pods"`
// 	Pvc          string `json:"pvc"`
// 	Selector     string `json:"selector"`
// 	Ingress      string `json:"ingress"`
// 	QueryName    string `json:"query_name"`
// 	RateAccuracy string `json:"rate_accuracy"`
// 	Start        int64  `json:"start"`
// 	End          int64  `json:"end"`
//
// 	// Category, Nodes, Pods, Pvc, Selector, Ingress, QueryName, RateAccuracy string
// 	// Start, End                                                             int64
// }
//
// func (f *MetricsQueryOpts) Validate(req *http.Request, errs binding.Errors) binding.Errors {
// 	ctx := context.GetContext(req)
// 	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
// }
