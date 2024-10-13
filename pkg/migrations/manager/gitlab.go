// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"

	"github.com/xanzy/go-gitlab"
)

var (
	_ schemas.Downloader        = &GitlabDownloader{}
	_ schemas.DownloaderFactory = &GitlabDownloaderFactory{}
)

func init() {
	RegisterDownloaderFactory(&GitlabDownloaderFactory{})
}

// GitlabDownloaderFactory defines a gitlab downloader factory
type GitlabDownloaderFactory struct{}

// New returns a Downloader related to this factory according MigrateOptions
func (f *GitlabDownloaderFactory) New(ctx context.Context, opts schemas.MigrateOptions) (schemas.Downloader, error) {
	u, err := url.Parse(opts.CloneAddr)
	if err != nil {
		return nil, err
	}

	baseURL := u.Scheme + "://" + u.Host
	repoNameSpace := strings.TrimPrefix(u.Path, "/")
	repoNameSpace = strings.TrimSuffix(repoNameSpace, ".git")

	log.Trace("Create gitlab downloader. BaseURL: %s RepoName: %s", baseURL, repoNameSpace)

	return NewGitlabDownloader(ctx, baseURL, repoNameSpace, opts.AuthUsername, opts.AuthPassword, opts.AuthToken)
}

// GitServiceType returns the type of git service
func (f *GitlabDownloaderFactory) GitServiceType() structs.GitServiceType {
	return structs.GitlabService
}

// GitlabDownloader implements a Downloader interface to get repository information
// from gitlab via go-gitlab
// - issueCount is incremented in GetIssues() to ensure PR and Issue numbers do not overlap,
// because Gitlab has individual Issue and Pull Request numbers.
type GitlabDownloader struct {
	schemas.NullDownloader
	ctx        context.Context
	client     *gitlab.Client
	repoID     int
	repoName   string
	issueCount int64
	maxPerPage int
}

// NewGitlabDownloader creates a gitlab Downloader via gitlab API
//
//	Use either a username/password, personal token entered into the username field, or anonymous/public access
//	Note: Public access only allows very basic access
func NewGitlabDownloader(ctx context.Context, baseURL, repoPath, username, password, token string) (*GitlabDownloader, error) {
	gitlabClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL), gitlab.WithHTTPClient(NewMigrationHTTPClient()))
	// Only use basic auth if token is blank and password is NOT
	// Basic auth will fail with empty strings, but empty token will allow anonymous public API usage
	if token == "" && password != "" {
		gitlabClient, err = gitlab.NewBasicAuthClient(username, password, gitlab.WithBaseURL(baseURL), gitlab.WithHTTPClient(NewMigrationHTTPClient()))
	}

	if err != nil {
		log.Trace("Error logging into gitlab: %v", err)
		return nil, err
	}

	// split namespace and subdirectory
	pathParts := strings.Split(strings.Trim(repoPath, "/"), "/")
	var resp *gitlab.Response
	u, _ := url.Parse(baseURL)
	for len(pathParts) >= 2 {
		_, resp, err = gitlabClient.Version.GetVersion()
		if err == nil || resp != nil && resp.StatusCode == http.StatusUnauthorized {
			err = nil // if no authentication given, this still should work
			break
		}

		u.Path = path.Join(u.Path, pathParts[0])
		baseURL = u.String()
		pathParts = pathParts[1:]
		_ = gitlab.WithBaseURL(baseURL)(gitlabClient)
		repoPath = strings.Join(pathParts, "/")
	}
	if err != nil {
		log.Trace("Error could not get gitlab version: %v", err)
		return nil, err
	}

	log.Trace("gitlab downloader: use BaseURL: '%s' and RepoPath: '%s'", baseURL, repoPath)

	// Grab and store project/repo ID here, due to issues using the URL escaped path
	gr, _, err := gitlabClient.Projects.GetProject(repoPath, nil, nil, gitlab.WithContext(ctx))
	if err != nil {
		log.Trace("Error retrieving project: %v", err)
		return nil, err
	}

	if gr == nil {
		log.Trace("Error getting project, project is nil")
		return nil, errors.New("Error getting project, project is nil")
	}

	return &GitlabDownloader{
		ctx:        ctx,
		client:     gitlabClient,
		repoID:     gr.ID,
		repoName:   gr.Name,
		maxPerPage: 100,
	}, nil
}

// SetContext set context
func (g *GitlabDownloader) SetContext(ctx context.Context) {
	g.ctx = ctx
}

// GetRepoInfo returns a repository information
func (g *GitlabDownloader) GetRepoInfo() (*schemas.Repository, error) {
	gr, _, err := g.client.Projects.GetProject(g.repoID, nil, nil, gitlab.WithContext(g.ctx))
	if err != nil {
		return nil, err
	}

	var private bool
	switch gr.Visibility {
	case gitlab.InternalVisibility:
		private = true
	case gitlab.PrivateVisibility:
		private = true
	}

	var owner string
	if gr.Owner == nil {
		log.Trace("gr.Owner is nil, trying to get owner from Namespace")
		if gr.Namespace != nil && gr.Namespace.Kind == "user" {
			owner = gr.Namespace.Path
		}
	} else {
		owner = gr.Owner.Username
	}

	// convert gitlab repo to stand Repo
	return &schemas.Repository{
		Owner:         owner,
		Name:          gr.Name,
		IsPrivate:     private,
		Description:   gr.Description,
		OriginalURL:   gr.WebURL,
		CloneURL:      gr.HTTPURLToRepo,
		DefaultBranch: gr.DefaultBranch,
	}, nil
}

// GetTopics return gitlab topics
func (g *GitlabDownloader) GetTopics() ([]string, error) {
	gr, _, err := g.client.Projects.GetProject(g.repoID, nil, nil, gitlab.WithContext(g.ctx))
	if err != nil {
		return nil, err
	}
	return gr.TagList, err
}

// GetMilestones returns milestones
func (g *GitlabDownloader) GetMilestones() ([]*schemas.Milestone, error) {
	perPage := g.maxPerPage
	state := "all"
	milestones := make([]*schemas.Milestone, 0, perPage)
	for i := 1; ; i++ {
		ms, _, err := g.client.Milestones.ListMilestones(g.repoID, &gitlab.ListMilestonesOptions{
			State: &state,
			ListOptions: gitlab.ListOptions{
				Page:    i,
				PerPage: perPage,
			},
		}, nil, gitlab.WithContext(g.ctx))
		if err != nil {
			return nil, err
		}

		for _, m := range ms {
			var desc string
			if m.Description != "" {
				desc = m.Description
			}
			state := "open"
			var closedAt *time.Time
			if m.State != "" {
				state = m.State
				if state == "closed" {
					closedAt = m.UpdatedAt
				}
			}

			var deadline *time.Time
			if m.DueDate != nil {
				deadlineParsed, err := time.Parse("2006-01-02", m.DueDate.String())
				if err != nil {
					log.Trace("Error parsing Milestone DueDate time")
					deadline = nil
				} else {
					deadline = &deadlineParsed
				}
			}

			milestones = append(milestones, &schemas.Milestone{
				Title:       m.Title,
				Description: desc,
				Deadline:    deadline,
				State:       state,
				Created:     *m.CreatedAt,
				Updated:     m.UpdatedAt,
				Closed:      closedAt,
			})
		}
		if len(ms) < perPage {
			break
		}
	}
	return milestones, nil
}

func (g *GitlabDownloader) normalizeColor(val string) string {
	val = strings.TrimLeft(val, "#")
	val = strings.ToLower(val)
	if len(val) == 3 {
		c := []rune(val)
		val = fmt.Sprintf("%c%c%c%c%c%c", c[0], c[0], c[1], c[1], c[2], c[2])
	}
	if len(val) != 6 {
		return ""
	}
	return val
}

// GetLabels returns labels
func (g *GitlabDownloader) GetLabels() ([]*schemas.Label, error) {
	perPage := g.maxPerPage
	labels := make([]*schemas.Label, 0, perPage)
	for i := 1; ; i++ {
		ls, _, err := g.client.Labels.ListLabels(g.repoID, &gitlab.ListLabelsOptions{ListOptions: gitlab.ListOptions{
			Page:    i,
			PerPage: perPage,
		}}, nil, gitlab.WithContext(g.ctx))
		if err != nil {
			return nil, err
		}
		for _, label := range ls {
			baseLabel := &schemas.Label{
				Name:        label.Name,
				Color:       g.normalizeColor(label.Color),
				Description: label.Description,
			}
			labels = append(labels, baseLabel)
		}
		if len(ls) < perPage {
			break
		}
	}
	return labels, nil
}

func (g *GitlabDownloader) convertGitlabRelease(rel *gitlab.Release) *schemas.Release {
	var zero int
	r := &schemas.Release{
		TagName:         rel.TagName,
		TargetCommitish: rel.Commit.ID,
		Name:            rel.Name,
		Body:            rel.Description,
		Created:         *rel.CreatedAt,
		PublisherID:     int64(rel.Author.ID),
		PublisherName:   rel.Author.Username,
	}

	httpClient := NewMigrationHTTPClient()

	for k, asset := range rel.Assets.Links {
		r.Assets = append(r.Assets, &schemas.ReleaseAsset{
			ID:            int64(asset.ID),
			Name:          asset.Name,
			ContentType:   &rel.Assets.Sources[k].Format,
			Size:          &zero,
			DownloadCount: &zero,
			DownloadFunc: func() (io.ReadCloser, error) {
				link, _, err := g.client.ReleaseLinks.GetReleaseLink(g.repoID, rel.TagName, asset.ID, gitlab.WithContext(g.ctx))
				if err != nil {
					return nil, err
				}

				req, err := http.NewRequest("GET", link.URL, nil)
				if err != nil {
					return nil, err
				}
				req = req.WithContext(g.ctx)
				resp, err := httpClient.Do(req)
				if err != nil {
					return nil, err
				}

				// resp.Body is closed by the uploader
				return resp.Body, nil
			},
		})
	}
	return r
}

// GetReleases returns releases
func (g *GitlabDownloader) GetReleases() ([]*schemas.Release, error) {
	perPage := g.maxPerPage
	releases := make([]*schemas.Release, 0, perPage)
	for i := 1; ; i++ {
		ls, _, err := g.client.Releases.ListReleases(g.repoID, &gitlab.ListReleasesOptions{
			Page:    i,
			PerPage: perPage,
		}, nil, gitlab.WithContext(g.ctx))
		if err != nil {
			return nil, err
		}

		for _, release := range ls {
			releases = append(releases, g.convertGitlabRelease(release))
		}
		if len(ls) < perPage {
			break
		}
	}
	return releases, nil
}

type gitlabIssueContext struct {
	IsMergeRequest bool
}

// GetIssues returns issues according start and limit
//
//	Note: issue label description and colors are not supported by the go-gitlab library at this time
func (g *GitlabDownloader) GetIssues(page, perPage int) ([]*schemas.Issue, bool, error) {
	state := "all"
	sort := "asc"

	if perPage > g.maxPerPage {
		perPage = g.maxPerPage
	}

	opt := &gitlab.ListProjectIssuesOptions{
		State: &state,
		Sort:  &sort,
		ListOptions: gitlab.ListOptions{
			PerPage: perPage,
			Page:    page,
		},
	}

	allIssues := make([]*schemas.Issue, 0, perPage)

	issues, _, err := g.client.Issues.ListProjectIssues(g.repoID, opt, nil, gitlab.WithContext(g.ctx))
	if err != nil {
		return nil, false, fmt.Errorf("error while listing issues: %v", err)
	}
	for _, issue := range issues {

		labels := make([]*schemas.Label, 0, len(issue.Labels))
		for _, l := range issue.Labels {
			labels = append(labels, &schemas.Label{
				Name: l,
			})
		}

		var milestone string
		if issue.Milestone != nil {
			milestone = issue.Milestone.Title
		}

		var reactions []*schemas.Reaction
		awardPage := 1
		for {
			awards, _, err := g.client.AwardEmoji.ListIssueAwardEmoji(g.repoID, issue.IID, &gitlab.ListAwardEmojiOptions{Page: awardPage, PerPage: perPage}, gitlab.WithContext(g.ctx))
			if err != nil {
				return nil, false, fmt.Errorf("error while listing issue awards: %v", err)
			}

			for i := range awards {
				reactions = append(reactions, g.awardToReaction(awards[i]))
			}

			if len(awards) < perPage {
				break
			}

			awardPage++
		}

		allIssues = append(allIssues, &schemas.Issue{
			Title:        issue.Title,
			Number:       int64(issue.IID),
			PosterID:     int64(issue.Author.ID),
			PosterName:   issue.Author.Username,
			Content:      issue.Description,
			Milestone:    milestone,
			State:        issue.State,
			Created:      *issue.CreatedAt,
			Labels:       labels,
			Reactions:    reactions,
			Closed:       issue.ClosedAt,
			IsLocked:     issue.DiscussionLocked,
			Updated:      *issue.UpdatedAt,
			ForeignIndex: int64(issue.IID),
			Context:      gitlabIssueContext{IsMergeRequest: false},
		})

		// increment issueCount, to be used in GetPullRequests()
		g.issueCount++
	}

	return allIssues, len(issues) < perPage, nil
}

// GetComments returns comments according issueNumber
// TODO: figure out how to transfer comment reactions
func (g *GitlabDownloader) GetComments(commentable schemas.Commentable) ([]*schemas.Comment, bool, error) {
	context, ok := commentable.GetContext().(gitlabIssueContext)
	if !ok {
		return nil, false, fmt.Errorf("unexpected context: %+v", commentable.GetContext())
	}

	allComments := make([]*schemas.Comment, 0, g.maxPerPage)

	page := 1

	for {
		var comments []*gitlab.Discussion
		var resp *gitlab.Response
		var err error
		if !context.IsMergeRequest {
			comments, resp, err = g.client.Discussions.ListIssueDiscussions(g.repoID, int(commentable.GetForeignIndex()), &gitlab.ListIssueDiscussionsOptions{
				Page:    page,
				PerPage: g.maxPerPage,
			}, nil, gitlab.WithContext(g.ctx))
		} else {
			comments, resp, err = g.client.Discussions.ListMergeRequestDiscussions(g.repoID, int(commentable.GetForeignIndex()), &gitlab.ListMergeRequestDiscussionsOptions{
				Page:    page,
				PerPage: g.maxPerPage,
			}, nil, gitlab.WithContext(g.ctx))
		}

		if err != nil {
			return nil, false, fmt.Errorf("error while listing comments: %v %v", g.repoID, err)
		}
		for _, comment := range comments {
			// Flatten comment threads
			if !comment.IndividualNote {
				for _, note := range comment.Notes {
					allComments = append(allComments, &schemas.Comment{
						IssueIndex:  commentable.GetLocalIndex(),
						Index:       int64(note.ID),
						PosterID:    int64(note.Author.ID),
						PosterName:  note.Author.Username,
						PosterEmail: note.Author.Email,
						Content:     note.Body,
						Created:     *note.CreatedAt,
					})
				}
			} else {
				c := comment.Notes[0]
				allComments = append(allComments, &schemas.Comment{
					IssueIndex:  commentable.GetLocalIndex(),
					Index:       int64(c.ID),
					PosterID:    int64(c.Author.ID),
					PosterName:  c.Author.Username,
					PosterEmail: c.Author.Email,
					Content:     c.Body,
					Created:     *c.CreatedAt,
				})
			}
		}
		if resp.NextPage == 0 {
			break
		}
		page = resp.NextPage
	}
	return allComments, true, nil
}

// GetPullRequests returns pull requests according page and perPage
func (g *GitlabDownloader) GetPullRequests(page, perPage int) ([]*schemas.PullRequest, bool, error) {
	if perPage > g.maxPerPage {
		perPage = g.maxPerPage
	}

	opt := &gitlab.ListProjectMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: perPage,
			Page:    page,
		},
	}

	allPRs := make([]*schemas.PullRequest, 0, perPage)

	prs, _, err := g.client.MergeRequests.ListProjectMergeRequests(g.repoID, opt, nil, gitlab.WithContext(g.ctx))
	if err != nil {
		return nil, false, fmt.Errorf("error while listing merge requests: %v", err)
	}
	for _, pr := range prs {

		labels := make([]*schemas.Label, 0, len(pr.Labels))
		for _, l := range pr.Labels {
			labels = append(labels, &schemas.Label{
				Name: l,
			})
		}

		var merged bool
		if pr.State == "merged" {
			merged = true
			pr.State = "closed"
		}

		mergeTime := pr.MergedAt
		if merged && pr.MergedAt == nil {
			mergeTime = pr.UpdatedAt
		}

		closeTime := pr.ClosedAt
		if merged && pr.ClosedAt == nil {
			closeTime = pr.UpdatedAt
		}

		var locked bool
		if pr.State == "locked" {
			locked = true
		}

		var milestone string
		if pr.Milestone != nil {
			milestone = pr.Milestone.Title
		}

		var reactions []*schemas.Reaction
		awardPage := 1
		for {
			awards, _, err := g.client.AwardEmoji.ListMergeRequestAwardEmoji(g.repoID, pr.IID, &gitlab.ListAwardEmojiOptions{Page: awardPage, PerPage: perPage}, gitlab.WithContext(g.ctx))
			if err != nil {
				return nil, false, fmt.Errorf("error while listing merge requests awards: %v", err)
			}

			for i := range awards {
				reactions = append(reactions, g.awardToReaction(awards[i]))
			}

			if len(awards) < perPage {
				break
			}

			awardPage++
		}

		// Add the PR ID to the Issue Count because PR and Issues share ID space in Gitea
		newPRNumber := g.issueCount + int64(pr.IID)

		allPRs = append(allPRs, &schemas.PullRequest{
			Title:          pr.Title,
			Number:         newPRNumber,
			PosterName:     pr.Author.Username,
			PosterID:       int64(pr.Author.ID),
			Content:        pr.Description,
			Milestone:      milestone,
			State:          pr.State,
			Created:        *pr.CreatedAt,
			Closed:         closeTime,
			Labels:         labels,
			Merged:         merged,
			MergeCommitSHA: pr.MergeCommitSHA,
			MergedTime:     mergeTime,
			IsLocked:       locked,
			Reactions:      reactions,
			Head: schemas.PullRequestBranch{
				Ref:       pr.SourceBranch,
				SHA:       pr.SHA,
				RepoName:  g.repoName,
				OwnerName: pr.Author.Username,
				CloneURL:  pr.WebURL,
			},
			Base: schemas.PullRequestBranch{
				Ref:       pr.TargetBranch,
				SHA:       pr.DiffRefs.BaseSha,
				RepoName:  g.repoName,
				OwnerName: pr.Author.Username,
			},
			PatchURL:     pr.WebURL + ".patch",
			ForeignIndex: int64(pr.IID),
			Context:      gitlabIssueContext{IsMergeRequest: true},
		})
	}

	return allPRs, len(prs) < perPage, nil
}

// GetReviews returns pull requests review
func (g *GitlabDownloader) GetReviews(reviewable schemas.Reviewable) ([]*schemas.Review, error) {
	approvals, resp, err := g.client.MergeRequestApprovals.GetConfiguration(g.repoID, int(reviewable.GetForeignIndex()), gitlab.WithContext(g.ctx))
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			log.Error(fmt.Sprintf("GitlabDownloader: while migrating a error occurred: '%s'", err.Error()))
			return []*schemas.Review{}, nil
		}
		return nil, err
	}

	var createdAt time.Time
	if approvals.CreatedAt != nil {
		createdAt = *approvals.CreatedAt
	} else if approvals.UpdatedAt != nil {
		createdAt = *approvals.UpdatedAt
	} else {
		createdAt = time.Now()
	}

	reviews := make([]*schemas.Review, 0, len(approvals.ApprovedBy))
	for _, user := range approvals.ApprovedBy {
		reviews = append(reviews, &schemas.Review{
			IssueIndex:   reviewable.GetLocalIndex(),
			ReviewerID:   int64(user.User.ID),
			ReviewerName: user.User.Username,
			CreatedAt:    createdAt,
			// All we get are approvals
			State: schemas.ReviewStateApproved,
		})
	}

	return reviews, nil
}

func (g *GitlabDownloader) awardToReaction(award *gitlab.AwardEmoji) *schemas.Reaction {
	return &schemas.Reaction{
		UserID:   int64(award.User.ID),
		UserName: award.User.Username,
		Content:  award.Name,
	}
}
