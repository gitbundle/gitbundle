// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/view_history"
)

// UserAssignmentWeb returns a middleware to handle context-user assignment for web routes
func UserAssignmentWeb() func(ctx *Context) {
	return func(ctx *Context) {
		userAssignment(ctx, func(status int, title string, obj interface{}) {
			err, ok := obj.(error)
			if !ok {
				err = fmt.Errorf("%s", obj)
			}
			if status == http.StatusNotFound {
				ctx.NotFound(title, err)
			} else {
				ctx.ServerError(title, err)
			}
		})
	}
}

// UserAssignmentAPI returns a middleware to handle context-user assignment for api routes
func UserAssignmentAPI() func(ctx *APIContext) {
	return func(ctx *APIContext) {
		userAssignment(ctx.Context, ctx.Error)
	}
}

func userAssignment(ctx *Context, errCb func(int, string, interface{})) {
	username := ctx.Params(":username")
	var err error

	if ctx.IsSigned && ctx.Doer.LowerName == strings.ToLower(username) {
		ctx.ContextUser = ctx.Doer
	} else {
		ctx.ContextUser, err = user_model.GetUserByName(ctx, username)
		if err != nil {
			if user_model.IsErrUserNotExist(err) {
				if redirectUserID, err := user_model.LookupUserRedirect(username); err == nil {
					RedirectToUser(ctx, username, redirectUserID)
				} else if user_model.IsErrUserRedirectNotExist(err) {
					errCb(http.StatusNotFound, "GetUserByName", err)
				} else {
					errCb(http.StatusInternalServerError, "LookupUserRedirect", err)
				}
			} else {
				errCb(http.StatusInternalServerError, "GetUserByName", err)
			}
		}
	}

	if err == nil && ctx.IsSigned && ctx.Doer.LowerName != strings.ToLower(username) && username != "" {
		content, _ := json.Marshal(map[string]interface{}{
			"name":   username,
			"avatar": ctx.ContextUser.AvatarLink(),
			"email":  ctx.ContextUser.GetEmail(),
		})
		vh := &view_history.ViewHistory{
			DoerID:    ctx.Doer.ID,
			UserID:    ctx.ContextUser.ID,
			VisitType: view_history.UserType,
			Content:   content,
		}
		err := vh.CreateOrUpdate(ctx)
		if err != nil {
			errCb(http.StatusInternalServerError, "CreateOrUpdateUserViewHistory", err)
		}
	}
}
