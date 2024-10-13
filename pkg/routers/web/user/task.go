// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package user

import (
	"net/http"
	"strconv"

	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/server/pkg/context"
	task_model "github.com/gitbundle/server/pkg/task"
)

// TaskStatus returns task's status
func TaskStatus(ctx *context.Context) {
	task, opts, err := task_model.GetMigratingTaskByID(ctx.ParamsInt64("task"), ctx.Doer.ID)
	if err != nil {
		if task_model.IsErrTaskDoesNotExist(err) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "task `" + strconv.FormatInt(ctx.ParamsInt64("task"), 10) + "` does not exist",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err,
		})
		return
	}

	message := task.Message

	if task.Message != "" && task.Message[0] == '{' {
		// assume message is actually a translatable string
		var translatableMessage task_model.TranslatableMessage
		if err := json.Unmarshal([]byte(message), &translatableMessage); err != nil {
			translatableMessage = task_model.TranslatableMessage{
				Format: "migrate.migrating_failed.error",
				Args:   []interface{}{task.Message},
			}
		}
		message = ctx.Tr(translatableMessage.Format, translatableMessage.Args...)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":    task.Status,
		"message":   message,
		"repo-id":   task.RepoID,
		"repo-name": opts.RepoName,
		"start":     task.StartTime,
		"end":       task.EndTime,
	})
}
