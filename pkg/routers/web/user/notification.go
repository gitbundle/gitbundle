// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package user

import (
	goctx "context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/repo/repoman"
)

const (
	tplNotification    base.TplName = "user/notification/notification"
	tplNotificationDiv base.TplName = "user/notification/notification_div"
)

// GetNotificationCount is the middleware that sets the notification count in the context
func GetNotificationCount(c *context.Context) {
	if strings.HasPrefix(c.Req.URL.Path, "/api") {
		return
	}

	if !c.IsSigned {
		return
	}

	c.Data["NotificationUnreadCount"] = func() int64 {
		count, err := repoman.GetNotificationCount(c, c.Doer, repoman.NotificationStatusUnread)
		if err != nil {
			if err != goctx.Canceled {
				log.Error("Unable to GetNotificationCount for user:%-v: %v", c.Doer, err)
			}
			return -1
		}

		return count
	}
}

// Notifications is the notifications page
func Notifications(c *context.Context) {
	getNotifications(c)
	if c.Written() {
		return
	}
	if c.FormBool("div-only") {
		c.Data["SequenceNumber"] = c.FormString("sequence-number")
		c.HTML(http.StatusOK, tplNotificationDiv)
		return
	}
	c.HTML(http.StatusOK, tplNotification)
}

func getNotifications(c *context.Context) {
	var (
		keyword = c.FormTrim("q")
		status  repoman.NotificationStatus
		page    = c.FormInt("page")
		perPage = c.FormInt("perPage")
	)
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 20
	}

	switch keyword {
	case "read":
		status = repoman.NotificationStatusRead
	default:
		status = repoman.NotificationStatusUnread
	}

	total, err := repoman.GetNotificationCount(c, c.Doer, status)
	if err != nil {
		c.ServerError("ErrGetNotificationCount", err)
		return
	}

	// redirect to last page if request page is more than total pages
	pager := context.NewPagination(int(total), perPage, page, 5)
	if pager.Paginater.Current() < page {
		c.Redirect(fmt.Sprintf("%s/notifications?q=%s&page=%d", setting.AppSubURL, url.QueryEscape(c.FormString("q")), pager.Paginater.Current()))
		return
	}

	statuses := []repoman.NotificationStatus{status, repoman.NotificationStatusPinned}
	notifications, err := repoman.NotificationsForUser(c, c.Doer, statuses, page, perPage)
	if err != nil {
		c.ServerError("ErrNotificationsForUser", err)
		return
	}

	failCount := 0

	repos, failures, err := notifications.LoadRepos()
	if err != nil {
		c.ServerError("LoadRepos", err)
		return
	}
	notifications = notifications.Without(failures)
	if err := repos.LoadAttributes(); err != nil {
		c.ServerError("LoadAttributes", err)
		return
	}
	failCount += len(failures)

	failures, err = notifications.LoadIssues()
	if err != nil {
		c.ServerError("LoadIssues", err)
		return
	}
	notifications = notifications.Without(failures)
	failCount += len(failures)

	failures, err = notifications.LoadComments()
	if err != nil {
		c.ServerError("LoadComments", err)
		return
	}
	notifications = notifications.Without(failures)
	failCount += len(failures)

	if failCount > 0 {
		c.Flash.Error(fmt.Sprintf("ERROR: %d notifications were removed due to missing parts - check the logs", failCount))
	}

	c.Data["Title"] = c.Tr("notifications")
	c.Data["Keyword"] = keyword
	c.Data["Status"] = status
	c.Data["Notifications"] = notifications

	pager.SetDefaultParams(c)
	c.Data["Page"] = pager
}

// NotificationStatusPost is a route for changing the status of a notification
func NotificationStatusPost(c *context.Context) {
	var (
		notificationID = c.FormInt64("notification_id")
		statusStr      = c.FormString("status")
		status         repoman.NotificationStatus
	)

	switch statusStr {
	case "read":
		status = repoman.NotificationStatusRead
	case "unread":
		status = repoman.NotificationStatusUnread
	case "pinned":
		status = repoman.NotificationStatusPinned
	default:
		c.ServerError("InvalidNotificationStatus", errors.New("Invalid notification status"))
		return
	}

	if _, err := repoman.SetNotificationStatus(notificationID, c.Doer, status); err != nil {
		c.ServerError("SetNotificationStatus", err)
		return
	}

	if !c.FormBool("noredirect") {
		url := fmt.Sprintf("%s/notifications?page=%s", setting.AppSubURL, url.QueryEscape(c.FormString("page")))
		c.Redirect(url, http.StatusSeeOther)
	}

	getNotifications(c)
	if c.Written() {
		return
	}
	c.Data["Link"] = setting.AppURL + "notifications"
	c.Data["SequenceNumber"] = c.Req.PostFormValue("sequence-number")

	c.HTML(http.StatusOK, tplNotificationDiv)
}

// NotificationPurgePost is a route for 'purging' the list of notifications - marking all unread as read
func NotificationPurgePost(c *context.Context) {
	err := repoman.UpdateNotificationStatuses(c.Doer, repoman.NotificationStatusUnread, repoman.NotificationStatusRead)
	if err != nil {
		c.ServerError("ErrUpdateNotificationStatuses", err)
		return
	}

	c.Redirect(setting.AppSubURL+"/notifications", http.StatusSeeOther)
}

// NewAvailable returns the notification counts
func NewAvailable(ctx *context.Context) {
	ctx.JSON(http.StatusOK, structs.NotificationCount{New: repoman.CountUnread(ctx, ctx.Doer.ID)})
}
