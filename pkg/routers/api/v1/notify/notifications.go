// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package notify

import (
	"net/http"
	"strings"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/routers/api/v1/utils"
)

// NewAvailable check if unread notifications exist
func NewAvailable(ctx *context.APIContext) {
	// swagger:operation GET /notifications/new notification notifyNewAvailable
	// ---
	// summary: Check if unread notifications exist
	// responses:
	//   "200":
	//     "$ref": "#/responses/NotificationCount"
	ctx.JSON(http.StatusOK, api.NotificationCount{New: repoman.CountUnread(ctx, ctx.Doer.ID)})
}

func getFindNotificationOptions(ctx *context.APIContext) *repoman.FindNotificationOptions {
	before, since, err := context.GetQueryBeforeSince(ctx.Context)
	if err != nil {
		ctx.Error(http.StatusUnprocessableEntity, "GetQueryBeforeSince", err)
		return nil
	}
	opts := &repoman.FindNotificationOptions{
		ListOptions:       utils.GetListOptions(ctx),
		UserID:            ctx.Doer.ID,
		UpdatedBeforeUnix: before,
		UpdatedAfterUnix:  since,
	}
	if !ctx.FormBool("all") {
		statuses := ctx.FormStrings("status-types")
		opts.Status = statusStringsToNotificationStatuses(statuses, []string{"unread", "pinned"})
	}

	subjectTypes := ctx.FormStrings("subject-type")
	if len(subjectTypes) != 0 {
		opts.Source = subjectToSource(subjectTypes)
	}

	return opts
}

func subjectToSource(value []string) (result []repoman.NotificationSource) {
	for _, v := range value {
		switch strings.ToLower(v) {
		case "issue":
			result = append(result, repoman.NotificationSourceIssue)
		case "pull":
			result = append(result, repoman.NotificationSourcePullRequest)
		case "commit":
			result = append(result, repoman.NotificationSourceCommit)
		case "repository":
			result = append(result, repoman.NotificationSourceRepository)
		}
	}
	return
}
