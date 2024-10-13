// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"github.com/gitbundle/modules/log"

	"xorm.io/builder"
	"xorm.io/xorm"
)

func updateMatrixWebhookHTTPMethod(x *xorm.Engine) error {
	matrixHookTaskType := 9 // value comes from the models package
	type Webhook struct {
		HTTPMethod string
	}

	cond := builder.Eq{"hook_task_type": matrixHookTaskType}.And(builder.Neq{"http_method": "PUT"})
	count, err := x.Where(cond).Cols("http_method").Update(&Webhook{HTTPMethod: "PUT"})
	if err == nil {
		log.Debug("Updated %d Matrix webhooks with http_method 'PUT'", count)
	}
	return err
}
