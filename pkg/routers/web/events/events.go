// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package events

import (
	"net/http"
	"time"

	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/eventsource"
	"github.com/gitbundle/server/pkg/routers/web/auth"
)

// Events listens for events
func Events(ctx *context.Context) {
	// FIXME: Need to check if resp is actually a http.Flusher! - how though?

	// Set the headers related to event streaming.
	ctx.Resp.Header().Set("Content-Type", "text/event-stream")
	ctx.Resp.Header().Set("Cache-Control", "no-cache")
	ctx.Resp.Header().Set("Connection", "keep-alive")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")
	ctx.Resp.WriteHeader(http.StatusOK)

	if !ctx.IsSigned {
		// Return unauthorized status event
		event := &eventsource.Event{
			Name: "close",
			Data: "unauthorized",
		}
		_, _ = event.WriteTo(ctx)
		ctx.Resp.Flush()
		return
	}

	// Listen to connection close and un-register messageChan
	notify := ctx.Done()
	ctx.Resp.Flush()

	shutdownCtx := graceful.GetManager().ShutdownContext()

	uid := ctx.Doer.ID

	messageChan := eventsource.GetManager().Register(uid)

	unregister := func() {
		eventsource.GetManager().Unregister(uid, messageChan)
		// ensure the messageChan is closed
		for {
			_, ok := <-messageChan
			if !ok {
				break
			}
		}
	}

	if _, err := ctx.Resp.Write([]byte("\n")); err != nil {
		log.Error("Unable to write to EventStream: %v", err)
		unregister()
		return
	}

	timer := time.NewTicker(30 * time.Second)

loop:
	for {
		select {
		case <-timer.C:
			event := &eventsource.Event{
				Name: "ping",
			}
			_, err := event.WriteTo(ctx.Resp)
			if err != nil {
				log.Error("Unable to write to EventStream for user %s: %v", ctx.Doer.Name, err)
				go unregister()
				break loop
			}
			ctx.Resp.Flush()
		case <-notify:
			go unregister()
			break loop
		case <-shutdownCtx.Done():
			go unregister()
			break loop
		case event, ok := <-messageChan:
			if !ok {
				break loop
			}

			// Handle logout
			if event.Name == "logout" {
				if ctx.Session.ID() == event.Data {
					_, _ = (&eventsource.Event{
						Name: "logout",
						Data: "here",
					}).WriteTo(ctx.Resp)
					ctx.Resp.Flush()
					go unregister()
					auth.HandleSignOut(ctx)
					break loop
				}
				// Replace the event - we don't want to expose the session ID to the user
				event = &eventsource.Event{
					Name: "logout",
					Data: "elsewhere",
				}
			}

			_, err := event.WriteTo(ctx.Resp)
			if err != nil {
				log.Error("Unable to write to EventStream for user %s: %v", ctx.Doer.Name, err)
				go unregister()
				break loop
			}
			ctx.Resp.Flush()
		}
	}
	timer.Stop()
}
