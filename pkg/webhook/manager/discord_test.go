// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package webhook_test

import (
	"testing"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/setting"
	webhook_model "github.com/gitbundle/server/pkg/webhook"
	webhook "github.com/gitbundle/server/pkg/webhook/manager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiscordPayload(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		p := createTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Create(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] branch test created", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Empty(t, pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/src/test", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Delete", func(t *testing.T) {
		p := deleteTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Delete(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] branch test deleted", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Empty(t, pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/src/test", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Fork", func(t *testing.T) {
		p := forkTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Fork(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "test/repo2 is forked to test/repo", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Empty(t, pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Push", func(t *testing.T) {
		p := pushTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Push(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo:test] 2 new commits", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "[2020558](http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778) commit message - user1\n[2020558](http://localhost:3000/test/repo/commit/2020558fe2e34debb818a514715839cabd25e778) commit message - user1", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/src/test", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Issue", func(t *testing.T) {
		p := issueTestPayload()

		d := new(webhook.DiscordPayload)
		p.Action = api.HookIssueOpened
		pl, err := d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Issue opened: #2 crash", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "issue body", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/issues/2", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)

		p.Action = api.HookIssueClosed
		pl, err = d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Issue closed: #2 crash", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Empty(t, pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/issues/2", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("IssueComment", func(t *testing.T) {
		p := issueCommentTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] New comment on issue #2 crash", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "more info needed", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/issues/2#issuecomment-4", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("PullRequest", func(t *testing.T) {
		p := pullRequestTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.PullRequest(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Pull request opened: #12 Fix bug", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "fixes bug #2", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/pulls/12", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("PullRequestComment", func(t *testing.T) {
		p := pullRequestCommentTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] New comment on pull request #12 Fix bug", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "changes requested", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/pulls/12#issuecomment-4", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Review", func(t *testing.T) {
		p := pullRequestTestPayload()
		p.Action = api.HookIssueReviewed

		d := new(webhook.DiscordPayload)
		pl, err := d.Review(p, webhook_model.HookEventPullRequestReviewApproved)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Pull request review approved: #12 Fix bug", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "good job", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo/pulls/12", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Repository", func(t *testing.T) {
		p := repositoryTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Repository(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Repository created", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Empty(t, pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/test/repo", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})

	t.Run("Release", func(t *testing.T) {
		p := pullReleaseTestPayload()

		d := new(webhook.DiscordPayload)
		pl, err := d.Release(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &webhook.DiscordPayload{}, pl)

		assert.Len(t, pl.(*webhook.DiscordPayload).Embeds, 1)
		assert.Equal(t, "[test/repo] Release created: v1.0", pl.(*webhook.DiscordPayload).Embeds[0].Title)
		assert.Equal(t, "Note of first stable release", pl.(*webhook.DiscordPayload).Embeds[0].Description)
		assert.Equal(t, "http://localhost:3000/api/v1/repos/test/repo/releases/2", pl.(*webhook.DiscordPayload).Embeds[0].URL)
		assert.Equal(t, p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.Name)
		assert.Equal(t, setting.AppURL+p.Sender.UserName, pl.(*webhook.DiscordPayload).Embeds[0].Author.URL)
		assert.Equal(t, p.Sender.AvatarURL, pl.(*webhook.DiscordPayload).Embeds[0].Author.IconURL)
	})
}

func TestDiscordJSONPayload(t *testing.T) {
	p := pushTestPayload()

	pl, err := new(webhook.DiscordPayload).Push(p)
	require.NoError(t, err)
	require.NotNil(t, pl)
	require.IsType(t, &webhook.DiscordPayload{}, pl)

	json, err := pl.JSONPayload()
	require.NoError(t, err)
	assert.NotEmpty(t, json)
}
