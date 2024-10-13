// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package user

import (
	"strings"
	"time"

	"github.com/gitbundle/modules/httpcache"
	"github.com/gitbundle/server/pkg/avatars"
	"github.com/gitbundle/server/pkg/context"
	user_model "github.com/gitbundle/server/pkg/user"
)

func cacheableRedirect(ctx *context.Context, location string) {
	// here we should not use `setting.StaticCacheTime`, it is pretty long (default: 6 hours)
	// we must make sure the redirection cache time is short enough, otherwise a user won't see the updated avatar in 6 hours
	// it's OK to make the cache time short, it is only a redirection, and doesn't cost much to make a new request
	httpcache.AddCacheControlToHeader(ctx.Resp.Header(), 5*time.Minute)
	ctx.Redirect(location)
}

// AvatarByUserName redirect browser to user avatar of requested size
func AvatarByUserName(ctx *context.Context) {
	userName := ctx.Params(":username")
	size := int(ctx.ParamsInt64(":size"))

	var user *user_model.User
	if strings.ToLower(userName) != "ghost" {
		var err error
		if user, err = user_model.GetUserByName(ctx, userName); err != nil {
			ctx.ServerError("Invalid user: "+userName, err)
			return
		}
	} else {
		user = user_model.NewGhostUser()
	}

	cacheableRedirect(ctx, user.AvatarLinkWithSize(size))
}

// AvatarByEmailHash redirects the browser to the email avatar link
func AvatarByEmailHash(ctx *context.Context) {
	hash := ctx.Params(":hash")
	email, err := avatars.GetEmailForHash(hash)
	if err != nil {
		ctx.ServerError("invalid avatar hash: "+hash, err)
		return
	}
	size := ctx.FormInt("size")
	cacheableRedirect(ctx, avatars.GenerateEmailAvatarFinalLink(email, size))
}
