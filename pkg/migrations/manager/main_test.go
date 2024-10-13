// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	_ "unsafe"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func assertTimeEqual(t *testing.T, expected, actual time.Time) {
	assert.Equal(t, expected.UTC(), actual.UTC())
}

func assertTimePtrEqual(t *testing.T, expected, actual *time.Time) {
	if expected == nil {
		assert.Nil(t, actual)
	} else {
		assert.NotNil(t, actual)
		assertTimeEqual(t, *expected, *actual)
	}
}

func assertCommentEqual(t *testing.T, expected, actual *schemas.Comment) {
	assert.Equal(t, expected.IssueIndex, actual.IssueIndex)
	assert.Equal(t, expected.PosterID, actual.PosterID)
	assert.Equal(t, expected.PosterName, actual.PosterName)
	assert.Equal(t, expected.PosterEmail, actual.PosterEmail)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimeEqual(t, expected.Updated, actual.Updated)
	assert.Equal(t, expected.Content, actual.Content)
	assertReactionsEqual(t, expected.Reactions, actual.Reactions)
}

func assertCommentsEqual(t *testing.T, expected, actual []*schemas.Comment) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertCommentEqual(t, expected[i], actual[i])
		}
	}
}

func assertLabelEqual(t *testing.T, expected, actual *schemas.Label) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Color, actual.Color)
	assert.Equal(t, expected.Description, actual.Description)
}

func assertLabelsEqual(t *testing.T, expected, actual []*schemas.Label) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertLabelEqual(t, expected[i], actual[i])
		}
	}
}

func assertMilestoneEqual(t *testing.T, expected, actual *schemas.Milestone) {
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Description, actual.Description)
	assertTimePtrEqual(t, expected.Deadline, actual.Deadline)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimePtrEqual(t, expected.Updated, actual.Updated)
	assertTimePtrEqual(t, expected.Closed, actual.Closed)
	assert.Equal(t, expected.State, actual.State)
}

func assertMilestonesEqual(t *testing.T, expected, actual []*schemas.Milestone) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertMilestoneEqual(t, expected[i], actual[i])
		}
	}
}

func assertIssueEqual(t *testing.T, expected, actual *schemas.Issue) {
	assert.Equal(t, expected.Number, actual.Number)
	assert.Equal(t, expected.PosterID, actual.PosterID)
	assert.Equal(t, expected.PosterName, actual.PosterName)
	assert.Equal(t, expected.PosterEmail, actual.PosterEmail)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Content, actual.Content)
	assert.Equal(t, expected.Ref, actual.Ref)
	assert.Equal(t, expected.Milestone, actual.Milestone)
	assert.Equal(t, expected.State, actual.State)
	assert.Equal(t, expected.IsLocked, actual.IsLocked)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimeEqual(t, expected.Updated, actual.Updated)
	assertTimePtrEqual(t, expected.Closed, actual.Closed)
	assertLabelsEqual(t, expected.Labels, actual.Labels)
	assertReactionsEqual(t, expected.Reactions, actual.Reactions)
	assert.ElementsMatch(t, expected.Assignees, actual.Assignees)
}

func assertIssuesEqual(t *testing.T, expected, actual []*schemas.Issue) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertIssueEqual(t, expected[i], actual[i])
		}
	}
}

func assertPullRequestEqual(t *testing.T, expected, actual *schemas.PullRequest) {
	assert.Equal(t, expected.Number, actual.Number)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.PosterID, actual.PosterID)
	assert.Equal(t, expected.PosterName, actual.PosterName)
	assert.Equal(t, expected.PosterEmail, actual.PosterEmail)
	assert.Equal(t, expected.Content, actual.Content)
	assert.Equal(t, expected.Milestone, actual.Milestone)
	assert.Equal(t, expected.State, actual.State)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimeEqual(t, expected.Updated, actual.Updated)
	assertTimePtrEqual(t, expected.Closed, actual.Closed)
	assertLabelsEqual(t, expected.Labels, actual.Labels)
	assert.Equal(t, expected.PatchURL, actual.PatchURL)
	assert.Equal(t, expected.Merged, actual.Merged)
	assertTimePtrEqual(t, expected.MergedTime, actual.MergedTime)
	assert.Equal(t, expected.MergeCommitSHA, actual.MergeCommitSHA)
	assertPullRequestBranchEqual(t, expected.Head, actual.Head)
	assertPullRequestBranchEqual(t, expected.Base, actual.Base)
	assert.ElementsMatch(t, expected.Assignees, actual.Assignees)
	assert.Equal(t, expected.IsLocked, actual.IsLocked)
	assertReactionsEqual(t, expected.Reactions, actual.Reactions)
}

func assertPullRequestsEqual(t *testing.T, expected, actual []*schemas.PullRequest) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertPullRequestEqual(t, expected[i], actual[i])
		}
	}
}

func assertPullRequestBranchEqual(t *testing.T, expected, actual schemas.PullRequestBranch) {
	assert.Equal(t, expected.CloneURL, actual.CloneURL)
	assert.Equal(t, expected.Ref, actual.Ref)
	assert.Equal(t, expected.SHA, actual.SHA)
	assert.Equal(t, expected.RepoName, actual.RepoName)
	assert.Equal(t, expected.OwnerName, actual.OwnerName)
}

func assertReactionEqual(t *testing.T, expected, actual *schemas.Reaction) {
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.UserName, actual.UserName)
	assert.Equal(t, expected.Content, actual.Content)
}

func assertReactionsEqual(t *testing.T, expected, actual []*schemas.Reaction) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertReactionEqual(t, expected[i], actual[i])
		}
	}
}

func assertReleaseAssetEqual(t *testing.T, expected, actual *schemas.ReleaseAsset) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ContentType, actual.ContentType)
	assert.Equal(t, expected.Size, actual.Size)
	assert.Equal(t, expected.DownloadCount, actual.DownloadCount)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimeEqual(t, expected.Updated, actual.Updated)
	assert.Equal(t, expected.DownloadURL, actual.DownloadURL)
}

func assertReleaseAssetsEqual(t *testing.T, expected, actual []*schemas.ReleaseAsset) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertReleaseAssetEqual(t, expected[i], actual[i])
		}
	}
}

func assertReleaseEqual(t *testing.T, expected, actual *schemas.Release) {
	assert.Equal(t, expected.TagName, actual.TagName)
	assert.Equal(t, expected.TargetCommitish, actual.TargetCommitish)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Body, actual.Body)
	assert.Equal(t, expected.Draft, actual.Draft)
	assert.Equal(t, expected.Prerelease, actual.Prerelease)
	assert.Equal(t, expected.PublisherID, actual.PublisherID)
	assert.Equal(t, expected.PublisherName, actual.PublisherName)
	assert.Equal(t, expected.PublisherEmail, actual.PublisherEmail)
	assertReleaseAssetsEqual(t, expected.Assets, actual.Assets)
	assertTimeEqual(t, expected.Created, actual.Created)
	assertTimeEqual(t, expected.Published, actual.Published)
}

func assertReleasesEqual(t *testing.T, expected, actual []*schemas.Release) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertReleaseEqual(t, expected[i], actual[i])
		}
	}
}

func assertRepositoryEqual(t *testing.T, expected, actual *schemas.Repository) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Owner, actual.Owner)
	assert.Equal(t, expected.IsPrivate, actual.IsPrivate)
	assert.Equal(t, expected.IsMirror, actual.IsMirror)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.CloneURL, actual.CloneURL)
	assert.Equal(t, expected.OriginalURL, actual.OriginalURL)
	assert.Equal(t, expected.DefaultBranch, actual.DefaultBranch)
}

func assertReviewEqual(t *testing.T, expected, actual *schemas.Review) {
	assert.Equal(t, expected.ID, actual.ID, "ID")
	assert.Equal(t, expected.IssueIndex, actual.IssueIndex, "IsssueIndex")
	assert.Equal(t, expected.ReviewerID, actual.ReviewerID, "ReviewerID")
	assert.Equal(t, expected.ReviewerName, actual.ReviewerName, "ReviewerName")
	assert.Equal(t, expected.Official, actual.Official, "Official")
	assert.Equal(t, expected.CommitID, actual.CommitID, "CommitID")
	assert.Equal(t, expected.Content, actual.Content, "Content")
	assert.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, 10*time.Second)
	assert.Equal(t, expected.State, actual.State, "State")
	assertReviewCommentsEqual(t, expected.Comments, actual.Comments)
}

func assertReviewsEqual(t *testing.T, expected, actual []*schemas.Review) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertReviewEqual(t, expected[i], actual[i])
		}
	}
}

func assertReviewCommentEqual(t *testing.T, expected, actual *schemas.ReviewComment) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.InReplyTo, actual.InReplyTo)
	assert.Equal(t, expected.Content, actual.Content)
	assert.Equal(t, expected.TreePath, actual.TreePath)
	assert.Equal(t, expected.DiffHunk, actual.DiffHunk)
	assert.Equal(t, expected.Position, actual.Position)
	assert.Equal(t, expected.Line, actual.Line)
	assert.Equal(t, expected.CommitID, actual.CommitID)
	assert.Equal(t, expected.PosterID, actual.PosterID)
	assertReactionsEqual(t, expected.Reactions, actual.Reactions)
	assertTimeEqual(t, expected.CreatedAt, actual.CreatedAt)
	assertTimeEqual(t, expected.UpdatedAt, actual.UpdatedAt)
}

func assertReviewCommentsEqual(t *testing.T, expected, actual []*schemas.ReviewComment) {
	if assert.Len(t, actual, len(expected)) {
		for i := range expected {
			assertReviewCommentEqual(t, expected[i], actual[i])
		}
	}
}

type onedevIssueContext struct {
	IsPullRequest bool
}

type gitlabIssueContext struct {
	IsMergeRequest bool
}

//go:linkname migrateRepository github.com/gitbundle/server/pkg/migrations/manager.migrateRepository
func migrateRepository(downloader schemas.Downloader, uploader schemas.Uploader, opts schemas.MigrateOptions, messenger schemas.Messenger) error

type GiteaLocalUploader struct {
	ctx            context.Context
	doer           *user_model.User
	repoOwner      string
	repoName       string
	repo           *repo_model.Repository
	labels         map[string]*issues_model.Label
	milestones     map[string]int64
	issues         map[int64]*issues_model.Issue
	gitRepo        *git.Repository
	prHeadCache    map[string]struct{}
	sameApp        bool
	userMap        map[int64]int64 // external user id mapping to user id
	prCache        map[int64]*issues_model.PullRequest
	gitServiceType structs.GitServiceType
}

//go:linkname (*GiteaLocalUploader).remapUser github.com/gitbundle/server/pkg/migrations/manager.(*GiteaLocalUploader).remapUser
func (g *GiteaLocalUploader) remapUser(source user_model.ExternalUserMigrated, target user_model.ExternalUserRemappable) error

//go:linkname (*GiteaLocalUploader).CreateRepo github.com/gitbundle/server/pkg/migrations/manager.(*GiteaLocalUploader).CreateRepo
func (g *GiteaLocalUploader) CreateRepo(repo *schemas.Repository, opts schemas.MigrateOptions) error

//go:linkname (*GiteaLocalUploader).updateGitForPullRequest github.com/gitbundle/server/pkg/migrations/manager.(*GiteaLocalUploader).updateGitForPullRequest
func (g *GiteaLocalUploader) updateGitForPullRequest(pr *schemas.PullRequest) (head string, err error)

type GitlabDownloader struct {
	schemas.NullDownloader
	ctx        context.Context
	client     *gitlab.Client
	repoID     int
	repoName   string
	issueCount int64
	maxPerPage int
}

//go:linkname (*GitlabDownloader).GetReviews github.com/gitbundle/server/pkg/migrations/manager.(*GitlabDownloader).GetReviews
func (g *GitlabDownloader) GetReviews(reviewable schemas.Reviewable) ([]*schemas.Review, error)
