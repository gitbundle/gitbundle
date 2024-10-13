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

func TestFeishuPayload(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		p := createTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Create(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, `[test/repo] branch test created`, pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Delete", func(t *testing.T) {
		p := deleteTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Delete(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, `[test/repo] branch test deleted`, pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Fork", func(t *testing.T) {
		p := forkTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Fork(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, `test/repo2 is forked to test/repo`, pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Push", func(t *testing.T) {
		p := pushTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Push(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "[test/repo:test] \r\n[2020558](http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778) commit message - user1\r\n[2020558](http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778) commit message - user1", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Issue", func(t *testing.T) {
		p := issueTestPayload()

		d := new(webhook.FeishuPayload)
		p.Action = api.HookIssueOpened
		pl, err := d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "#2 crash\r\n[test/repo] Issue opened: #2 crash by user1\r\n\r\nissue body", pl.(*webhook.FeishuPayload).Content.Text)

		p.Action = api.HookIssueClosed
		pl, err = d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "#2 crash\r\n[test/repo] Issue closed: #2 crash by user1", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("IssueComment", func(t *testing.T) {
		p := issueCommentTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "#2 crash\r\n[test/repo] New comment on issue #2 crash by user1\r\n\r\nmore info needed", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("PullRequest", func(t *testing.T) {
		p := pullRequestTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.PullRequest(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "#12 Fix bug\r\n[test/repo] Pull request opened: #12 Fix bug by user1\r\n\r\nfixes bug #2", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("PullRequestComment", func(t *testing.T) {
		p := pullRequestCommentTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "#12 Fix bug\r\n[test/repo] New comment on pull request #12 Fix bug by user1\r\n\r\nchanges requested", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Review", func(t *testing.T) {
		p := pullRequestTestPayload()
		p.Action = api.HookIssueReviewed

		d := new(webhook.FeishuPayload)
		pl, err := d.Review(p, webhook_model.HookEventPullRequestReviewApproved)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "[test/repo] Pull request review approved : #12 Fix bug\r\n\r\ngood job", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Repository", func(t *testing.T) {
		p := repositoryTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Repository(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "[test/repo] Repository created", pl.(*webhook.FeishuPayload).Content.Text)
	})

	t.Run("Release", func(t *testing.T) {
		p := pullReleaseTestPayload()

		d := new(webhook.FeishuPayload)
		pl, err := d.Release(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.FeishuPayload{}, pl)

		assert.Equal(t, "[test/repo] Release created: v1.0 by user1", pl.(*webhook.FeishuPayload).Content.Text)
	})
}

func TestFeishuJSONPayload(t *testing.T) {
	p := pushTestPayload()

	pl, err := new(webhook.FeishuPayload).Push(p)
	require.NoError(t, err)
	require.NotNil(t, pl)
	require.IsType(t, &webhook.FeishuPayload{}, pl)

	json, err := pl.JSONPayload()
	require.NoError(t, err)
	assert.NotEmpty(t, json)
}
