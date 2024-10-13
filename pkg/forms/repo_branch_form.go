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

// NewBranchForm form for creating a new branch
type NewBranchForm struct {
	NewBranchName string `binding:"Required;MaxSize(100);GitRefName"`
	CreateTag     bool
}

// Validate validates the fields
func (f *NewBranchForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}

// RenameBranchForm form for rename a branch
type RenameBranchForm struct {
	From string `binding:"Required;MaxSize(100);GitRefName"`
	To   string `binding:"Required;MaxSize(100);GitRefName"`
}

// Validate validates the fields
func (f *RenameBranchForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}
