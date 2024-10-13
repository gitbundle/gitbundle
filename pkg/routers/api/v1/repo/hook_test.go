// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"net/http"
	"testing"

	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/test"
	"github.com/gitbundle/server/pkg/unittest"
	"github.com/gitbundle/server/pkg/webhook"

	"github.com/stretchr/testify/assert"
)

func TestTestHook(t *testing.T) {
	unittest.PrepareTestEnv(t)

	ctx := test.MockContext(t, "user2/repo1/wiki/_pages")
	ctx.SetParams(":id", "1")
	test.LoadRepo(t, ctx, 1)
	test.LoadRepoCommit(t, ctx)
	test.LoadUser(t, ctx, 2)
	TestHook(&context.APIContext{Context: ctx, Org: nil})
	assert.EqualValues(t, http.StatusNoContent, ctx.Resp.Status())

	unittest.AssertExistsAndLoadBean(t, &webhook.HookTask{
		RepoID: 1,
		HookID: 1,
	}, unittest.Cond("is_delivered=?", false))
}
