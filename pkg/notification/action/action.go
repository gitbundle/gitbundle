// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
	"path"
	"strings"

	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/process"
	"github.com/gitbundle/modules/util"
	action_model "github.com/gitbundle/server/pkg/action"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/notification/base"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/release"
	"github.com/gitbundle/server/pkg/repo/repository"
	user_model "github.com/gitbundle/server/pkg/user"
)

type actionNotifier struct {
	base.NullNotifier
}

var _ base.Notifier = &actionNotifier{}

// NewNotifier create a new actionNotifier notifier
func NewNotifier() base.Notifier {
	return &actionNotifier{}
}

func (a *actionNotifier) NotifyNewIssue(issue *issues_model.Issue, mentions []*user_model.User) {
	if err := issue.LoadPoster(); err != nil {
		log.Error("issue.LoadPoster: %v", err)
		return
	}
	if err := issue.LoadRepo(db.DefaultContext); err != nil {
		log.Error("issue.LoadRepo: %v", err)
		return
	}
	repo := issue.Repo

	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: issue.Poster.ID,
		ActUser:   issue.Poster,
		OpType:    action_model.ActionCreateIssue,
		Content:   fmt.Sprintf("%d|%s", issue.Index, issue.Title),
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
	}); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

// NotifyIssueChangeStatus notifies close or reopen issue to notifiers
func (a *actionNotifier) NotifyIssueChangeStatus(doer *user_model.User, issue *issues_model.Issue, actionComment *issues_model.Comment, closeOrReopen bool) {
	// Compose comment action, could be plain comment, close or reopen issue/pull request.
	// This object will be used to notify watchers in the end of function.
	act := &action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		Content:   fmt.Sprintf("%d|%s", issue.Index, ""),
		RepoID:    issue.Repo.ID,
		Repo:      issue.Repo,
		Comment:   actionComment,
		CommentID: actionComment.ID,
		IsPrivate: issue.Repo.IsPrivate,
	}
	// Check comment type.
	if closeOrReopen {
		act.OpType = action_model.ActionCloseIssue
		if issue.IsPull {
			act.OpType = action_model.ActionClosePullRequest
		}
	} else {
		act.OpType = action_model.ActionReopenIssue
		if issue.IsPull {
			act.OpType = action_model.ActionReopenPullRequest
		}
	}

	// Notify watchers for whatever action comes in, ignore if no action type.
	if err := action_model.NotifyWatchers(act); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

// NotifyCreateIssueComment notifies comment on an issue to notifiers
func (a *actionNotifier) NotifyCreateIssueComment(doer *user_model.User, repo *repo_model.Repository,
	issue *issues_model.Issue, comment *issues_model.Comment, mentions []*user_model.User,
) {
	act := &action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		RepoID:    issue.Repo.ID,
		Repo:      issue.Repo,
		Comment:   comment,
		CommentID: comment.ID,
		IsPrivate: issue.Repo.IsPrivate,
	}

	truncatedContent, truncatedRight := util.SplitStringAtByteN(comment.Content, 200)
	if truncatedRight != "" {
		// in case the content is in a Latin family language, we remove the last broken word.
		lastSpaceIdx := strings.LastIndex(truncatedContent, " ")
		if lastSpaceIdx != -1 && (len(truncatedContent)-lastSpaceIdx < 15) {
			truncatedContent = truncatedContent[:lastSpaceIdx] + "…"
		}
	}
	act.Content = fmt.Sprintf("%d|%s", issue.Index, truncatedContent)

	if issue.IsPull {
		act.OpType = action_model.ActionCommentPull
	} else {
		act.OpType = action_model.ActionCommentIssue
	}

	// Notify watchers for whatever action comes in, ignore if no action type.
	if err := action_model.NotifyWatchers(act); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyNewPullRequest(pull *issues_model.PullRequest, mentions []*user_model.User) {
	if err := pull.LoadIssue(); err != nil {
		log.Error("pull.LoadIssue: %v", err)
		return
	}
	if err := pull.Issue.LoadRepo(db.DefaultContext); err != nil {
		log.Error("pull.Issue.LoadRepo: %v", err)
		return
	}
	if err := pull.Issue.LoadPoster(); err != nil {
		log.Error("pull.Issue.LoadPoster: %v", err)
		return
	}

	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: pull.Issue.Poster.ID,
		ActUser:   pull.Issue.Poster,
		OpType:    action_model.ActionCreatePullRequest,
		Content:   fmt.Sprintf("%d|%s", pull.Issue.Index, pull.Issue.Title),
		RepoID:    pull.Issue.Repo.ID,
		Repo:      pull.Issue.Repo,
		IsPrivate: pull.Issue.Repo.IsPrivate,
	}); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyRenameRepository(doer *user_model.User, repo *repo_model.Repository, oldRepoName string) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionRenameRepo,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		Content:   oldRepoName,
	}); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyTransferRepository(doer *user_model.User, repo *repo_model.Repository, oldOwnerName string) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionTransferRepo,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		Content:   path.Join(oldOwnerName, repo.Name),
	}); err != nil {
		log.Error("NotifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyCreateRepository(doer, u *user_model.User, repo *repo_model.Repository) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionCreateRepo,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
	}); err != nil {
		log.Error("notify watchers '%d/%d': %v", doer.ID, repo.ID, err)
	}
}

func (a *actionNotifier) NotifyForkRepository(doer *user_model.User, oldRepo, repo *repo_model.Repository) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionCreateRepo,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
	}); err != nil {
		log.Error("notify watchers '%d/%d': %v", doer.ID, repo.ID, err)
	}
}

func (a *actionNotifier) NotifyPullRequestReview(pr *issues_model.PullRequest, review *issues_model.Review, comment *issues_model.Comment, mentions []*user_model.User) {
	ctx, _, finished := process.GetManager().AddContext(graceful.GetManager().HammerContext(), fmt.Sprintf("actionNotifier.NotifyPullRequestReview Pull[%d] #%d in [%d]", pr.ID, pr.Index, pr.BaseRepoID))
	defer finished()

	if err := review.LoadReviewer(); err != nil {
		log.Error("LoadReviewer '%d/%d': %v", review.ID, review.ReviewerID, err)
		return
	}
	if err := review.LoadCodeComments(ctx); err != nil {
		log.Error("LoadCodeComments '%d/%d': %v", review.Reviewer.ID, review.ID, err)
		return
	}

	actions := make([]*action_model.Action, 0, 10)
	for _, lines := range review.CodeComments {
		for _, comments := range lines {
			for _, comm := range comments {
				actions = append(actions, &action_model.Action{
					ActUserID: review.Reviewer.ID,
					ActUser:   review.Reviewer,
					Content:   fmt.Sprintf("%d|%s", review.Issue.Index, strings.Split(comm.Content, "\n")[0]),
					OpType:    action_model.ActionCommentPull,
					RepoID:    review.Issue.RepoID,
					Repo:      review.Issue.Repo,
					IsPrivate: review.Issue.Repo.IsPrivate,
					Comment:   comm,
					CommentID: comm.ID,
				})
			}
		}
	}

	if review.Type != issues_model.ReviewTypeComment || strings.TrimSpace(comment.Content) != "" {
		action := &action_model.Action{
			ActUserID: review.Reviewer.ID,
			ActUser:   review.Reviewer,
			Content:   fmt.Sprintf("%d|%s", review.Issue.Index, strings.Split(comment.Content, "\n")[0]),
			RepoID:    review.Issue.RepoID,
			Repo:      review.Issue.Repo,
			IsPrivate: review.Issue.Repo.IsPrivate,
			Comment:   comment,
			CommentID: comment.ID,
		}

		switch review.Type {
		case issues_model.ReviewTypeApprove:
			action.OpType = action_model.ActionApprovePullRequest
		case issues_model.ReviewTypeReject:
			action.OpType = action_model.ActionRejectPullRequest
		default:
			action.OpType = action_model.ActionCommentPull
		}

		actions = append(actions, action)
	}

	if err := action_model.NotifyWatchersActions(actions); err != nil {
		log.Error("notify watchers '%d/%d': %v", review.Reviewer.ID, review.Issue.RepoID, err)
	}
}

func (*actionNotifier) NotifyMergePullRequest(pr *issues_model.PullRequest, doer *user_model.User) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionMergePullRequest,
		Content:   fmt.Sprintf("%d|%s", pr.Issue.Index, pr.Issue.Title),
		RepoID:    pr.Issue.Repo.ID,
		Repo:      pr.Issue.Repo,
		IsPrivate: pr.Issue.Repo.IsPrivate,
	}); err != nil {
		log.Error("NotifyWatchers [%d]: %v", pr.ID, err)
	}
}

func (*actionNotifier) NotifyPullRevieweDismiss(doer *user_model.User, review *issues_model.Review, comment *issues_model.Comment) {
	reviewerName := review.Reviewer.Name
	if len(review.OriginalAuthor) > 0 {
		reviewerName = review.OriginalAuthor
	}
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    action_model.ActionPullReviewDismissed,
		Content:   fmt.Sprintf("%d|%s|%s", review.Issue.Index, reviewerName, comment.Content),
		RepoID:    review.Issue.Repo.ID,
		Repo:      review.Issue.Repo,
		IsPrivate: review.Issue.Repo.IsPrivate,
		CommentID: comment.ID,
		Comment:   comment,
	}); err != nil {
		log.Error("NotifyWatchers [%d]: %v", review.Issue.ID, err)
	}
}

func (a *actionNotifier) NotifyPushCommits(pusher *user_model.User, repo *repo_model.Repository, opts *repository.PushUpdateOptions, commits *repository.PushCommits) {
	data, err := json.Marshal(commits)
	if err != nil {
		log.Error("Marshal: %v", err)
		return
	}

	opType := action_model.ActionCommitRepo

	// Check it's tag push or branch.
	if opts.IsTag() {
		opType = action_model.ActionPushTag
		if opts.IsDelRef() {
			opType = action_model.ActionDeleteTag
		}
	} else if opts.IsDelRef() {
		opType = action_model.ActionDeleteBranch
	}

	if err = action_model.NotifyWatchers(&action_model.Action{
		ActUserID: pusher.ID,
		ActUser:   pusher,
		OpType:    opType,
		Content:   string(data),
		RepoID:    repo.ID,
		Repo:      repo,
		RefName:   opts.RefFullName,
		IsPrivate: repo.IsPrivate,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyCreateRef(doer *user_model.User, repo *repo_model.Repository, refType, refFullName, refID string) {
	opType := action_model.ActionCommitRepo
	if refType == "tag" {
		// has sent same action in `NotifyPushCommits`, so skip it.
		return
	}
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    opType,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		RefName:   refFullName,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyDeleteRef(doer *user_model.User, repo *repo_model.Repository, refType, refFullName string) {
	opType := action_model.ActionDeleteBranch
	if refType == "tag" {
		// has sent same action in `NotifyPushCommits`, so skip it.
		return
	}
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: doer.ID,
		ActUser:   doer,
		OpType:    opType,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		RefName:   refFullName,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifySyncPushCommits(pusher *user_model.User, repo *repo_model.Repository, opts *repository.PushUpdateOptions, commits *repository.PushCommits) {
	data, err := json.Marshal(commits)
	if err != nil {
		log.Error("json.Marshal: %v", err)
		return
	}

	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: repo.OwnerID,
		ActUser:   repo.MustOwner(),
		OpType:    action_model.ActionMirrorSyncPush,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		RefName:   opts.RefFullName,
		Content:   string(data),
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifySyncCreateRef(doer *user_model.User, repo *repo_model.Repository, refType, refFullName, refID string) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: repo.OwnerID,
		ActUser:   repo.MustOwner(),
		OpType:    action_model.ActionMirrorSyncCreate,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		RefName:   refFullName,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifySyncDeleteRef(doer *user_model.User, repo *repo_model.Repository, refType, refFullName string) {
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: repo.OwnerID,
		ActUser:   repo.MustOwner(),
		OpType:    action_model.ActionMirrorSyncDelete,
		RepoID:    repo.ID,
		Repo:      repo,
		IsPrivate: repo.IsPrivate,
		RefName:   refFullName,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}

func (a *actionNotifier) NotifyNewRelease(rel *release.Release) {
	if err := rel.LoadAttributes(); err != nil {
		log.Error("NotifyNewRelease: %v", err)
		return
	}
	if err := action_model.NotifyWatchers(&action_model.Action{
		ActUserID: rel.PublisherID,
		ActUser:   rel.Publisher,
		OpType:    action_model.ActionPublishRelease,
		RepoID:    rel.RepoID,
		Repo:      rel.Repo,
		IsPrivate: rel.Repo.IsPrivate,
		Content:   rel.Title,
		RefName:   rel.TagName,
	}); err != nil {
		log.Error("notifyWatchers: %v", err)
	}
}
