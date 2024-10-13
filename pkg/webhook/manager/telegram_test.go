// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package webhook_test

import (
	"testing"

	api "github.com/gitbundle/api/pkg/structs"
	webhook_model "github.com/gitbundle/server/pkg/webhook"
	webhook "github.com/gitbundle/server/pkg/webhook/manager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTelegramPayload(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		p := createTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Create(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `[<a href="http://localhost:3000/test/repo">test/repo</a>] branch <a href="http://localhost:3000/test/repo/src/test">test</a> created`, pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Delete", func(t *testing.T) {
		p := deleteTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Delete(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `[<a href="http://localhost:3000/test/repo">test/repo</a>] branch <a href="http://localhost:3000/test/repo/src/test">test</a> deleted`, pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Fork", func(t *testing.T) {
		p := forkTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Fork(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `test/repo2 is forked to <a href="http://localhost:3000/test/repo">test/repo</a>`, pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Push", func(t *testing.T) {
		p := pushTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Push(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[<a href=\"http://localhost:3000/test/repo\">test/repo</a>:<a href=\"http://localhost:3000/test/repo/src/test\">test</a>] 2 new commits\n[<a href=\"http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778\">2020558</a>] commit message - user1\n[<a href=\"http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778\">2020558</a>] commit message - user1", pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Issue", func(t *testing.T) {
		p := issueTestPayload()

		d := new(webhook.TelegramPayload)
		p.Action = api.HookIssueOpened
		pl, err := d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[<a href=\"http://localhost:3000/test/repo\">test/repo</a>] Issue opened: <a href=\"http://localhost:3000/test/repo/issues/2\">#2 crash</a> by <a href=\"https://try.gitea.io/user1\">user1</a>\n\nissue body", pl.(*webhook.TelegramPayload).Message)

		p.Action = api.HookIssueClosed
		pl, err = d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `[<a href="http://localhost:3000/test/repo">test/repo</a>] Issue closed: <a href="http://localhost:3000/test/repo/issues/2">#2 crash</a> by <a href="https://try.gitea.io/user1">user1</a>`, pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("IssueComment", func(t *testing.T) {
		p := issueCommentTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[<a href=\"http://localhost:3000/test/repo\">test/repo</a>] New comment on issue <a href=\"http://localhost:3000/test/repo/issues/2\">#2 crash</a> by <a href=\"https://try.gitea.io/user1\">user1</a>\nmore info needed", pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("PullRequest", func(t *testing.T) {
		p := pullRequestTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.PullRequest(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[<a href=\"http://localhost:3000/test/repo\">test/repo</a>] Pull request opened: <a href=\"http://localhost:3000/test/repo/pulls/12\">#12 Fix bug</a> by <a href=\"https://try.gitea.io/user1\">user1</a>\nfixes bug #2", pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("PullRequestComment", func(t *testing.T) {
		p := pullRequestCommentTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[<a href=\"http://localhost:3000/test/repo\">test/repo</a>] New comment on pull request <a href=\"http://localhost:3000/test/repo/pulls/12\">#12 Fix bug</a> by <a href=\"https://try.gitea.io/user1\">user1</a>\nchanges requested", pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Review", func(t *testing.T) {
		p := pullRequestTestPayload()
		p.Action = api.HookIssueReviewed

		d := new(webhook.TelegramPayload)
		pl, err := d.Review(p, webhook_model.HookEventPullRequestReviewApproved)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, "[test/repo] Pull request review approved: #12 Fix bug\ngood job", pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Repository", func(t *testing.T) {
		p := repositoryTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Repository(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `[<a href="http://localhost:3000/test/repo">test/repo</a>] Repository created`, pl.(*webhook.TelegramPayload).Message)
	})

	t.Run("Release", func(t *testing.T) {
		p := pullReleaseTestPayload()

		d := new(webhook.TelegramPayload)
		pl, err := d.Release(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.TelegramPayload{}, pl)

		assert.Equal(t, `[<a href="http://localhost:3000/test/repo">test/repo</a>] Release created: <a href="http://localhost:3000/test/repo/releases/tag/v1.0">v1.0</a> by <a href="https://try.gitea.io/user1">user1</a>`, pl.(*webhook.TelegramPayload).Message)
	})
}

func TestTelegramJSONPayload(t *testing.T) {
	p := pushTestPayload()

	pl, err := new(webhook.TelegramPayload).Push(p)
	require.NoError(t, err)
	require.NotNil(t, pl)
	require.IsType(t, &webhook.TelegramPayload{}, pl)

	json, err := pl.JSONPayload()
	require.NoError(t, err)
	assert.NotEmpty(t, json)
}
