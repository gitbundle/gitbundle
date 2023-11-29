// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package mailer

import (
	"context"
	"fmt"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/action"
	issues_model "github.com/gitbundle/server/pkg/issues"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/unit"
	user_model "github.com/gitbundle/server/pkg/user"
)

func fallbackMailSubject(issue *issues_model.Issue) string {
	return fmt.Sprintf("[%s] %s (#%d)", issue.Repo.FullName(), issue.Title, issue.Index)
}

type mailCommentContext struct {
	context.Context
	Issue      *issues_model.Issue
	Doer       *user_model.User
	ActionType action.ActionType
	Content    string
	Comment    *issues_model.Comment
}

const (
	// MailBatchSize set the batch size used in mailIssueCommentBatch
	MailBatchSize = 100
)

// mailIssueCommentToParticipants can be used for both new issue creation and comment.
// This function sends two list of emails:
// 1. Repository watchers (except for WIP pull requests) and users who are participated in comments.
// 2. Users who are not in 1. but get mentioned in current issue/comment.
func mailIssueCommentToParticipants(ctx *mailCommentContext, mentions []*user_model.User) error {
	// Required by the mail composer; make sure to load these before calling the async function
	if err := ctx.Issue.LoadRepo(ctx); err != nil {
		return fmt.Errorf("LoadRepo(): %v", err)
	}
	if err := ctx.Issue.LoadPoster(); err != nil {
		return fmt.Errorf("LoadPoster(): %v", err)
	}
	if err := ctx.Issue.LoadPullRequest(); err != nil {
		return fmt.Errorf("LoadPullRequest(): %v", err)
	}

	// Enough room to avoid reallocations
	unfiltered := make([]int64, 1, 64)

	// =========== Original poster ===========
	unfiltered[0] = ctx.Issue.PosterID

	// =========== Assignees ===========
	ids, err := issues_model.GetAssigneeIDsByIssue(ctx.Issue.ID)
	if err != nil {
		return fmt.Errorf("GetAssigneeIDsByIssue(%d): %v", ctx.Issue.ID, err)
	}
	unfiltered = append(unfiltered, ids...)

	// =========== Participants (i.e. commenters, reviewers) ===========
	ids, err = issues_model.GetParticipantsIDsByIssueID(ctx.Issue.ID)
	if err != nil {
		return fmt.Errorf("GetParticipantsIDsByIssueID(%d): %v", ctx.Issue.ID, err)
	}
	unfiltered = append(unfiltered, ids...)

	// =========== Issue watchers ===========
	ids, err = issues_model.GetIssueWatchersIDs(ctx, ctx.Issue.ID, true)
	if err != nil {
		return fmt.Errorf("GetIssueWatchersIDs(%d): %v", ctx.Issue.ID, err)
	}
	unfiltered = append(unfiltered, ids...)

	// =========== Repo watchers ===========
	// Make repo watchers last, since it's likely the list with the most users
	if !(ctx.Issue.IsPull && ctx.Issue.PullRequest.IsWorkInProgress() && ctx.ActionType != action.ActionCreatePullRequest) {
		ids, err = repo_model.GetRepoWatchersIDs(ctx, ctx.Issue.RepoID)
		if err != nil {
			return fmt.Errorf("GetRepoWatchersIDs(%d): %v", ctx.Issue.RepoID, err)
		}
		unfiltered = append(ids, unfiltered...)
	}

	visited := make(map[int64]bool, len(unfiltered)+len(mentions)+1)

	// Avoid mailing the doer
	visited[ctx.Doer.ID] = true

	// =========== Mentions ===========
	if err = mailIssueCommentBatch(ctx, mentions, visited, true); err != nil {
		return fmt.Errorf("mailIssueCommentBatch() mentions: %v", err)
	}

	// Avoid mailing explicit unwatched
	ids, err = issues_model.GetIssueWatchersIDs(ctx, ctx.Issue.ID, false)
	if err != nil {
		return fmt.Errorf("GetIssueWatchersIDs(%d): %v", ctx.Issue.ID, err)
	}
	for _, i := range ids {
		visited[i] = true
	}

	unfilteredUsers, err := user_model.GetMaileableUsersByIDs(unfiltered, false)
	if err != nil {
		return err
	}
	if err = mailIssueCommentBatch(ctx, unfilteredUsers, visited, false); err != nil {
		return fmt.Errorf("mailIssueCommentBatch(): %v", err)
	}

	return nil
}

func mailIssueCommentBatch(ctx *mailCommentContext, users []*user_model.User, visited map[int64]bool, fromMention bool) error {
	checkUnit := unit.TypeIssues
	if ctx.Issue.IsPull {
		checkUnit = unit.TypePullRequests
	}

	langMap := make(map[string][]*user_model.User)
	for _, user := range users {
		if !user.IsActive {
			// Exclude deactivated users
			continue
		}
		// At this point we exclude:
		// user that don't have all mails enabled or users only get mail on mention and this is one ...
		if !(user.EmailNotificationsPreference == user_model.EmailNotificationsEnabled ||
			fromMention && user.EmailNotificationsPreference == user_model.EmailNotificationsOnMention) {
			continue
		}

		// if we have already visited this user we exclude them
		if _, ok := visited[user.ID]; ok {
			continue
		}

		// now mark them as visited
		visited[user.ID] = true

		// test if this user is allowed to see the issue/pull
		if !repoman.CheckRepoUnitUser(ctx, ctx.Issue.Repo, user, checkUnit) {
			continue
		}

		langMap[user.Language] = append(langMap[user.Language], user)
	}

	for lang, receivers := range langMap {
		// because we know that the len(receivers) > 0 and we don't care about the order particularly
		// working backwards from the last (possibly) incomplete batch. If len(receivers) can be 0 this
		// starting condition will need to be changed slightly
		for i := ((len(receivers) - 1) / MailBatchSize) * MailBatchSize; i >= 0; i -= MailBatchSize {
			msgs, err := composeIssueCommentMessages(ctx, lang, receivers[i:], fromMention, "issue comments")
			if err != nil {
				return err
			}
			SendAsyncs(msgs)
			receivers = receivers[:i]
		}
	}

	return nil
}

// MailParticipants sends new issue thread created emails to repository watchers
// and mentioned people.
func MailParticipants(issue *issues_model.Issue, doer *user_model.User, opType action.ActionType, mentions []*user_model.User) error {
	if setting.MailService == nil {
		// No mail service configured
		return nil
	}

	content := issue.Content
	if opType == action.ActionCloseIssue || opType == action.ActionClosePullRequest ||
		opType == action.ActionReopenIssue || opType == action.ActionReopenPullRequest ||
		opType == action.ActionMergePullRequest {
		content = ""
	}
	if err := mailIssueCommentToParticipants(
		&mailCommentContext{
			Context:    context.TODO(), // TODO: use a correct context
			Issue:      issue,
			Doer:       doer,
			ActionType: opType,
			Content:    content,
			Comment:    nil,
		}, mentions); err != nil {
		log.Error("mailIssueCommentToParticipants: %v", err)
	}
	return nil
}