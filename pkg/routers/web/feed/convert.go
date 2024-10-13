// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package feed

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gitbundle/modules/markup"
	"github.com/gitbundle/modules/markup/markdown"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/action"
	"github.com/gitbundle/server/pkg/context"
	"github.com/gitbundle/server/pkg/static/templates"

	"github.com/gorilla/feeds"
)

func toBranchLink(act *action.Action) string {
	return act.GetRepoLink() + "/src/branch/" + util.PathEscapeSegments(act.GetBranch())
}

func toTagLink(act *action.Action) string {
	return act.GetRepoLink() + "/src/tag/" + util.PathEscapeSegments(act.GetTag())
}

func toIssueLink(act *action.Action) string {
	return act.GetRepoLink() + "/issues/" + url.PathEscape(act.GetIssueInfos()[0])
}

func toPullLink(act *action.Action) string {
	return act.GetRepoLink() + "/pulls/" + url.PathEscape(act.GetIssueInfos()[0])
}

func toSrcLink(act *action.Action) string {
	return act.GetRepoLink() + "/src/" + util.PathEscapeSegments(act.GetBranch())
}

func toReleaseLink(act *action.Action) string {
	return act.GetRepoLink() + "/releases/tag/" + util.PathEscapeSegments(act.GetBranch())
}

// renderMarkdown creates a minimal markdown render context from an action.
// If rendering fails, the original markdown text is returned
func renderMarkdown(ctx *context.Context, act *action.Action, content string) string {
	markdownCtx := &markup.RenderContext{
		Ctx:       ctx,
		URLPrefix: act.GetRepoLink(),
		Type:      markdown.MarkupName,
		Metas: map[string]string{
			"user": act.GetRepoUserName(),
			"repo": act.GetRepoName(),
		},
	}
	markdown, err := markdown.RenderString(markdownCtx, content)
	if err != nil {
		return content
	}
	return markdown
}

// feedActionsToFeedItems convert gitea's Action feed to feeds Item
func feedActionsToFeedItems(ctx *context.Context, actions action.ActionList) (items []*feeds.Item, err error) {
	for _, act := range actions {
		act.LoadActUser()

		content, desc, title := "", "", ""

		link := &feeds.Link{Href: act.GetCommentLink()}

		// title
		title = act.ActUser.DisplayName() + " "
		switch act.OpType {
		case action.ActionCreateRepo:
			title += ctx.TrHTMLEscapeArgs("action.create_repo", act.GetRepoLink(), act.ShortRepoPath())
			link.Href = act.GetRepoLink()
		case action.ActionRenameRepo:
			title += ctx.TrHTMLEscapeArgs("action.rename_repo", act.GetContent(), act.GetRepoLink(), act.ShortRepoPath())
			link.Href = act.GetRepoLink()
		case action.ActionCommitRepo:
			link.Href = toBranchLink(act)
			if len(act.Content) != 0 {
				title += ctx.TrHTMLEscapeArgs("action.commit_repo", act.GetRepoLink(), link.Href, act.GetBranch(), act.ShortRepoPath())
			} else {
				title += ctx.TrHTMLEscapeArgs("action.create_branch", act.GetRepoLink(), link.Href, act.GetBranch(), act.ShortRepoPath())
			}
		case action.ActionCreateIssue:
			link.Href = toIssueLink(act)
			title += ctx.TrHTMLEscapeArgs("action.create_issue", link.Href, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionCreatePullRequest:
			link.Href = toPullLink(act)
			title += ctx.TrHTMLEscapeArgs("action.create_pull_request", link.Href, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionTransferRepo:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.transfer_repo", act.GetContent(), act.GetRepoLink(), act.ShortRepoPath())
		case action.ActionPushTag:
			link.Href = toTagLink(act)
			title += ctx.TrHTMLEscapeArgs("action.push_tag", act.GetRepoLink(), link.Href, act.GetTag(), act.ShortRepoPath())
		case action.ActionCommentIssue:
			issueLink := toIssueLink(act)
			if link.Href == "#" {
				link.Href = issueLink
			}
			title += ctx.TrHTMLEscapeArgs("action.comment_issue", issueLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionMergePullRequest:
			pullLink := toPullLink(act)
			if link.Href == "#" {
				link.Href = pullLink
			}
			title += ctx.TrHTMLEscapeArgs("action.merge_pull_request", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionCloseIssue:
			issueLink := toIssueLink(act)
			if link.Href == "#" {
				link.Href = issueLink
			}
			title += ctx.TrHTMLEscapeArgs("action.close_issue", issueLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionReopenIssue:
			issueLink := toIssueLink(act)
			if link.Href == "#" {
				link.Href = issueLink
			}
			title += ctx.TrHTMLEscapeArgs("action.reopen_issue", issueLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionClosePullRequest:
			pullLink := toPullLink(act)
			if link.Href == "#" {
				link.Href = pullLink
			}
			title += ctx.TrHTMLEscapeArgs("action.close_pull_request", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionReopenPullRequest:
			pullLink := toPullLink(act)
			if link.Href == "#" {
				link.Href = pullLink
			}
			title += ctx.TrHTMLEscapeArgs("action.reopen_pull_request", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionDeleteTag:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.delete_tag", act.GetRepoLink(), act.GetTag(), act.ShortRepoPath())
		case action.ActionDeleteBranch:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.delete_branch", act.GetRepoLink(), html.EscapeString(act.GetBranch()), act.ShortRepoPath())
		case action.ActionMirrorSyncPush:
			srcLink := toSrcLink(act)
			if link.Href == "#" {
				link.Href = srcLink
			}
			title += ctx.TrHTMLEscapeArgs("action.mirror_sync_push", act.GetRepoLink(), srcLink, act.GetBranch(), act.ShortRepoPath())
		case action.ActionMirrorSyncCreate:
			srcLink := toSrcLink(act)
			if link.Href == "#" {
				link.Href = srcLink
			}
			title += ctx.TrHTMLEscapeArgs("action.mirror_sync_create", act.GetRepoLink(), srcLink, act.GetBranch(), act.ShortRepoPath())
		case action.ActionMirrorSyncDelete:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.mirror_sync_delete", act.GetRepoLink(), act.GetBranch(), act.ShortRepoPath())
		case action.ActionApprovePullRequest:
			pullLink := toPullLink(act)
			title += ctx.TrHTMLEscapeArgs("action.approve_pull_request", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionRejectPullRequest:
			pullLink := toPullLink(act)
			title += ctx.TrHTMLEscapeArgs("action.reject_pull_request", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionCommentPull:
			pullLink := toPullLink(act)
			title += ctx.TrHTMLEscapeArgs("action.comment_pull", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath())
		case action.ActionPublishRelease:
			releaseLink := toReleaseLink(act)
			if link.Href == "#" {
				link.Href = releaseLink
			}
			title += ctx.TrHTMLEscapeArgs("action.publish_release", act.GetRepoLink(), releaseLink, act.ShortRepoPath(), act.Content)
		case action.ActionPullReviewDismissed:
			pullLink := toPullLink(act)
			title += ctx.TrHTMLEscapeArgs("action.review_dismissed", pullLink, act.GetIssueInfos()[0], act.ShortRepoPath(), act.GetIssueInfos()[1])
		case action.ActionStarRepo:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.starred_repo", act.GetRepoLink(), act.GetRepoPath())
		case action.ActionWatchRepo:
			link.Href = act.GetRepoLink()
			title += ctx.TrHTMLEscapeArgs("action.watched_repo", act.GetRepoLink(), act.GetRepoPath())
		default:
			return nil, fmt.Errorf("unknown action type: %v", act.OpType)
		}

		// description & content
		{
			switch act.OpType {
			case action.ActionCommitRepo, action.ActionMirrorSyncPush:
				push := templates.ActionContent2Commits(act)
				repoLink := act.GetRepoLink()

				for _, commit := range push.Commits {
					if len(desc) != 0 {
						desc += "\n\n"
					}
					desc += fmt.Sprintf("<a href=\"%s\">%s</a>\n%s",
						html.EscapeString(fmt.Sprintf("%s/commit/%s", act.GetRepoLink(), commit.Sha1)),
						commit.Sha1,
						templates.RenderCommitMessage(ctx, commit.Message, repoLink, nil),
					)
				}

				if push.Len > 1 {
					link = &feeds.Link{Href: fmt.Sprintf("%s/%s", setting.AppSubURL, push.CompareURL)}
				} else if push.Len == 1 {
					link = &feeds.Link{Href: fmt.Sprintf("%s/commit/%s", act.GetRepoLink(), push.Commits[0].Sha1)}
				}

			case action.ActionCreateIssue, action.ActionCreatePullRequest:
				desc = strings.Join(act.GetIssueInfos(), "#")
				content = renderMarkdown(ctx, act, act.GetIssueContent())
			case action.ActionCommentIssue, action.ActionApprovePullRequest, action.ActionRejectPullRequest, action.ActionCommentPull:
				desc = act.GetIssueTitle()
				comment := act.GetIssueInfos()[1]
				if len(comment) != 0 {
					desc += "\n\n" + renderMarkdown(ctx, act, comment)
				}
			case action.ActionMergePullRequest:
				desc = act.GetIssueInfos()[1]
			case action.ActionCloseIssue, action.ActionReopenIssue, action.ActionClosePullRequest, action.ActionReopenPullRequest:
				desc = act.GetIssueTitle()
			case action.ActionPullReviewDismissed:
				desc = ctx.Tr("action.review_dismissed_reason") + "\n\n" + act.GetIssueInfos()[2]
			}
		}
		if len(content) == 0 {
			content = desc
		}

		items = append(items, &feeds.Item{
			Title:       title,
			Link:        link,
			Description: desc,
			Author: &feeds.Author{
				Name:  act.ActUser.DisplayName(),
				Email: act.ActUser.GetEmail(),
			},
			Id:      strconv.FormatInt(act.ID, 10),
			Created: act.CreatedUnix.AsTime(),
			Content: content,
		})
	}
	return
}

// GetFeedType return if it is a feed request and altered name and feed type.
func GetFeedType(name string, req *http.Request) (bool, string, string) {
	if strings.HasSuffix(name, ".rss") ||
		strings.Contains(req.Header.Get("Accept"), "application/rss+xml") {
		return true, strings.TrimSuffix(name, ".rss"), "rss"
	}

	if strings.HasSuffix(name, ".atom") ||
		strings.Contains(req.Header.Get("Accept"), "application/atom+xml") {
		return true, strings.TrimSuffix(name, ".atom"), "atom"
	}

	return false, name, ""
}
