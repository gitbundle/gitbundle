// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package feed

import (
	"net/http"
	"time"

	"github.com/gitbundle/server/pkg/action"
	"github.com/gitbundle/server/pkg/context"

	"github.com/gorilla/feeds"
)

// ShowUserFeedRSS show user activity as RSS feed
func ShowUserFeedRSS(ctx *context.Context) {
	showUserFeed(ctx, "rss")
}

// ShowUserFeedAtom show user activity as Atom feed
func ShowUserFeedAtom(ctx *context.Context) {
	showUserFeed(ctx, "atom")
}

// showUserFeed show user activity as RSS / Atom feed
func showUserFeed(ctx *context.Context, formatType string) {
	actions, err := action.GetFeeds(ctx, action.GetFeedsOptions{
		RequestedUser:   ctx.ContextUser,
		Actor:           ctx.Doer,
		IncludePrivate:  false,
		OnlyPerformedBy: !ctx.ContextUser.IsOrganization(),
		IncludeDeleted:  false,
		Date:            ctx.FormString("date"),
	})
	if err != nil {
		ctx.ServerError("GetFeeds", err)
		return
	}

	feed := &feeds.Feed{
		Title:       ctx.Tr("home.feed_of", ctx.ContextUser.DisplayName()),
		Link:        &feeds.Link{Href: ctx.ContextUser.HTMLURL()},
		Description: ctx.ContextUser.Description,
		Created:     time.Now(),
	}

	feed.Items, err = feedActionsToFeedItems(ctx, actions)
	if err != nil {
		ctx.ServerError("convert feed", err)
		return
	}

	writeFeed(ctx, feed, formatType)
}

// writeFeed write a feeds.Feed as atom or rss to ctx.Resp
func writeFeed(ctx *context.Context, feed *feeds.Feed, formatType string) {
	ctx.Resp.WriteHeader(http.StatusOK)
	if formatType == "atom" {
		ctx.Resp.Header().Set("Content-Type", "application/atom+xml;charset=utf-8")
		if err := feed.WriteAtom(ctx.Resp); err != nil {
			ctx.ServerError("Render Atom failed", err)
		}
	} else {
		ctx.Resp.Header().Set("Content-Type", "application/rss+xml;charset=utf-8")
		if err := feed.WriteRss(ctx.Resp); err != nil {
			ctx.ServerError("Render RSS failed", err)
		}
	}
}
