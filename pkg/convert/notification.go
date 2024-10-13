// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert

import (
	"net/url"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/perm"
	"github.com/gitbundle/server/pkg/repo/repoman"
)

// ToNotificationThread convert a Notification to api.NotificationThread
func ToNotificationThread(n *repoman.Notification) *api.NotificationThread {
	result := &api.NotificationThread{
		ID:        n.ID,
		Unread:    !(n.Status == repoman.NotificationStatusRead || n.Status == repoman.NotificationStatusPinned),
		Pinned:    n.Status == repoman.NotificationStatusPinned,
		UpdatedAt: n.UpdatedUnix.AsTime(),
		URL:       n.APIURL(),
	}

	// since user only get notifications when he has access to use minimal access mode
	if n.Repository != nil {
		result.Repository = ToRepo(n.Repository, perm.AccessModeRead)

		// This permission is not correct and we should not be reporting it
		for repository := result.Repository; repository != nil; repository = repository.Parent {
			repository.Permissions = nil
		}
	}

	// handle Subject
	switch n.Source {
	case repoman.NotificationSourceIssue:
		result.Subject = &api.NotificationSubject{Type: api.NotifySubjectIssue}
		if n.Issue != nil {
			result.Subject.Title = n.Issue.Title
			result.Subject.URL = n.Issue.APIURL()
			result.Subject.HTMLURL = n.Issue.HTMLURL()
			result.Subject.State = n.Issue.State()
			comment, err := n.Issue.GetLastComment()
			if err == nil && comment != nil {
				result.Subject.LatestCommentURL = comment.APIURL()
				result.Subject.LatestCommentHTMLURL = comment.HTMLURL()
			}
		}
	case repoman.NotificationSourcePullRequest:
		result.Subject = &api.NotificationSubject{Type: api.NotifySubjectPull}
		if n.Issue != nil {
			result.Subject.Title = n.Issue.Title
			result.Subject.URL = n.Issue.APIURL()
			result.Subject.HTMLURL = n.Issue.HTMLURL()
			result.Subject.State = n.Issue.State()
			comment, err := n.Issue.GetLastComment()
			if err == nil && comment != nil {
				result.Subject.LatestCommentURL = comment.APIURL()
				result.Subject.LatestCommentHTMLURL = comment.HTMLURL()
			}

			pr, _ := n.Issue.GetPullRequest()
			if pr != nil && pr.HasMerged {
				result.Subject.State = "merged"
			}
		}
	case repoman.NotificationSourceCommit:
		url := n.Repository.HTMLURL() + "/commit/" + url.PathEscape(n.CommitID)
		result.Subject = &api.NotificationSubject{
			Type:    api.NotifySubjectCommit,
			Title:   n.CommitID,
			URL:     url,
			HTMLURL: url,
		}
	case repoman.NotificationSourceRepository:
		result.Subject = &api.NotificationSubject{
			Type:  api.NotifySubjectRepository,
			Title: n.Repository.FullName(),
			// FIXME: this is a relative URL, rather useless and inconsistent, but keeping for backwards compat
			URL:     n.Repository.Link(),
			HTMLURL: n.Repository.HTMLURL(),
		}
	}

	return result
}

// ToNotifications convert list of Notification to api.NotificationThread list
func ToNotifications(nl repoman.NotificationList) []*api.NotificationThread {
	result := make([]*api.NotificationThread, 0, len(nl))
	for _, n := range nl {
		result = append(result, ToNotificationThread(n))
	}
	return result
}
