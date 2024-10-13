// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package swagger

import (
	api "github.com/gitbundle/api/pkg/structs"
)

// NotificationThread
// swagger:response NotificationThread
type swaggerNotificationThread struct {
	// in:body
	Body api.NotificationThread `json:"body"`
}

// NotificationThreadList
// swagger:response NotificationThreadList
type swaggerNotificationThreadList struct {
	// in:body
	Body []api.NotificationThread `json:"body"`
}

// Number of unread notifications
// swagger:response NotificationCount
type swaggerNotificationCount struct {
	// in:body
	Body api.NotificationCount `json:"body"`
}
