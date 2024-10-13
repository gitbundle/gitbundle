// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package admin

import (
	"net/http"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/cron"
	"github.com/gitbundle/server/pkg/routers/api/v1/utils"
)

// ListCronTasks api for getting cron tasks
func ListCronTasks(ctx *context.APIContext) {
	// swagger:operation GET /admin/cron admin adminCronList
	// ---
	// summary: List cron tasks
	// produces:
	// - application/json
	// parameters:
	// - name: page
	//   in: query
	//   description: page number of results to return (1-based)
	//   type: integer
	// - name: limit
	//   in: query
	//   description: page size of results
	//   type: integer
	// responses:
	//   "200":
	//     "$ref": "#/responses/CronList"
	//   "403":
	//     "$ref": "#/responses/forbidden"
	tasks := cron.ListTasks()
	count := len(tasks)

	listOpts := utils.GetListOptions(ctx)
	tasks = util.PaginateSlice(tasks, listOpts.Page, listOpts.PageSize).(cron.TaskTable)

	res := make([]structs.Cron, len(tasks))
	for i, task := range tasks {
		res[i] = structs.Cron{
			Name:      task.Name,
			Schedule:  task.Spec,
			Next:      task.Next,
			Prev:      task.Prev,
			ExecTimes: task.ExecTimes,
		}
	}

	ctx.SetTotalCountHeader(int64(count))
	ctx.JSON(http.StatusOK, res)
}

// PostCronTask api for getting cron tasks
func PostCronTask(ctx *context.APIContext) {
	// swagger:operation POST /admin/cron/{task} admin adminCronRun
	// ---
	// summary: Run cron task
	// produces:
	// - application/json
	// parameters:
	// - name: task
	//   in: path
	//   description: task to run
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//     "$ref": "#/responses/empty"
	//   "404":
	//     "$ref": "#/responses/notFound"
	task := cron.GetTask(ctx.Params(":task"))
	if task == nil {
		ctx.NotFound()
		return
	}
	task.Run()
	log.Trace("Cron Task %s started by admin(%s)", task.Name, ctx.Doer.Name)

	ctx.Status(http.StatusNoContent)
}
