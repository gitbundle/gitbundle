// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	migrations "github.com/gitbundle/server/pkg/migrations/manager"
	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
	repo_model "github.com/gitbundle/server/pkg/repo"
	release_model "github.com/gitbundle/server/pkg/repo/release"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestGiteaUploadRepo(t *testing.T) {
	// FIXME: Since no accesskey or user/password will trigger rate limit of github, just skip
	t.Skip()

	unittest.PrepareTestEnv(t)

	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1}).(*user_model.User)

	var (
		ctx        = context.Background()
		downloader = migrations.NewGithubDownloaderV3(ctx, "https://github.com", "", "", "", "go-xorm", "builder")
		repoName   = "builder-" + time.Now().Format("2006-01-02-15-04-05")
		uploader   = migrations.NewGiteaLocalUploader(graceful.GetManager().HammerContext(), user, user.Name, repoName)
	)

	err := migrateRepository(downloader, uploader, schemas.MigrateOptions{
		CloneAddr:    "https://github.com/go-xorm/builder",
		RepoName:     repoName,
		AuthUsername: "",

		Wiki:         true,
		Issues:       true,
		Milestones:   true,
		Labels:       true,
		Releases:     true,
		Comments:     true,
		PullRequests: true,
		Private:      true,
		Mirror:       false,
	}, nil)
	assert.NoError(t, err)

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{OwnerID: user.ID, Name: repoName}).(*repo_model.Repository)
	assert.True(t, repo.HasWiki())
	assert.EqualValues(t, repo_model.RepositoryReady, repo.Status)

	milestones, _, err := issues_model.GetMilestones(issues_model.GetMilestonesOption{
		RepoID: repo.ID,
		State:  structs.StateOpen,
	})
	assert.NoError(t, err)
	assert.Len(t, milestones, 1)

	milestones, _, err = issues_model.GetMilestones(issues_model.GetMilestonesOption{
		RepoID: repo.ID,
		State:  structs.StateClosed,
	})
	assert.NoError(t, err)
	assert.Empty(t, milestones)

	labels, err := issues_model.GetLabelsByRepoID(ctx, repo.ID, "", db.ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, labels, 12)

	releases, err := release_model.GetReleasesByRepoID(repo.ID, release_model.FindReleasesOptions{
		ListOptions: db.ListOptions{
			PageSize: 10,
			Page:     0,
		},
		IncludeTags: true,
	})
	assert.NoError(t, err)
	assert.Len(t, releases, 8)

	releases, err = release_model.GetReleasesByRepoID(repo.ID, release_model.FindReleasesOptions{
		ListOptions: db.ListOptions{
			PageSize: 10,
			Page:     0,
		},
		IncludeTags: false,
	})
	assert.NoError(t, err)
	assert.Len(t, releases, 1)

	issues, err := issues_model.Issues(&issues_model.IssuesOptions{
		RepoID:   repo.ID,
		IsPull:   util.OptionalBoolFalse,
		SortType: "oldest",
	})
	assert.NoError(t, err)
	assert.Len(t, issues, 15)
	assert.NoError(t, issues[0].LoadDiscussComments())
	assert.Empty(t, issues[0].Comments)

	pulls, _, err := issues_model.PullRequests(repo.ID, &issues_model.PullRequestsOptions{
		SortType: "oldest",
	})
	assert.NoError(t, err)
	assert.Len(t, pulls, 30)
	assert.NoError(t, pulls[0].LoadIssue())
	assert.NoError(t, pulls[0].Issue.LoadDiscussComments())
	assert.Len(t, pulls[0].Issue.Comments, 2)
}

func TestGiteaUploadRemapLocalUser(t *testing.T) {
	unittest.PrepareTestEnv(t)
	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1}).(*user_model.User)
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)

	repoName := "migrated"
	uploaderPrivate := migrations.NewGiteaLocalUploader(context.Background(), doer, doer.Name, repoName)
	uploader := *((*GiteaLocalUploader)(unsafe.Pointer(&*uploaderPrivate)))

	// call remapLocalUser
	uploader.sameApp = true

	externalID := int64(1234567)
	externalName := "username"
	source := schemas.Release{
		PublisherID:   externalID,
		PublisherName: externalName,
	}

	//
	// The externalID does not match any existing user, everything
	// belongs to the doer
	//
	target := release_model.Release{}
	uploader.userMap = make(map[int64]int64)
	err := uploader.remapUser(&source, &target)
	assert.NoError(t, err)
	assert.EqualValues(t, doer.ID, target.GetUserID())

	//
	// The externalID matches a known user but the name does not match,
	// everything belongs to the doer
	//
	source.PublisherID = user.ID
	target = release_model.Release{}
	uploader.userMap = make(map[int64]int64)
	err = uploader.remapUser(&source, &target)
	assert.NoError(t, err)
	assert.EqualValues(t, doer.ID, target.GetUserID())

	//
	// The externalID and externalName match an existing user, everything
	// belongs to the existing user
	//
	source.PublisherName = user.Name
	target = release_model.Release{}
	uploader.userMap = make(map[int64]int64)
	err = uploader.remapUser(&source, &target)
	assert.NoError(t, err)
	assert.EqualValues(t, user.ID, target.GetUserID())
}

func TestGiteaUploadRemapExternalUser(t *testing.T) {
	unittest.PrepareTestEnv(t)
	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1}).(*user_model.User)

	repoName := "migrated"
	uploaderPrivate := migrations.NewGiteaLocalUploader(context.Background(), doer, doer.Name, repoName)
	uploader := *((*GiteaLocalUploader)(unsafe.Pointer(&*uploaderPrivate)))
	uploader.gitServiceType = structs.GitBundleService
	// call remapExternalUser
	uploader.sameApp = false

	externalID := int64(1234567)
	externalName := "username"
	source := schemas.Release{
		PublisherID:   externalID,
		PublisherName: externalName,
	}

	//
	// When there is no user linked to the external ID, the migrated data is authored
	// by the doer
	//
	uploader.userMap = make(map[int64]int64)
	target := release_model.Release{}
	err := uploader.remapUser(&source, &target)
	assert.NoError(t, err)
	assert.EqualValues(t, doer.ID, target.GetUserID())

	//
	// Link the external ID to an existing user
	//
	linkedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	externalLoginUser := &user_model.ExternalLoginUser{
		ExternalID:    strconv.FormatInt(externalID, 10),
		UserID:        linkedUser.ID,
		LoginSourceID: 0,
		Provider:      structs.GitBundleService.Name(),
	}
	err = user_model.LinkExternalToUser(linkedUser, externalLoginUser)
	assert.NoError(t, err)

	//
	// When a user is linked to the external ID, it becomes the author of
	// the migrated data
	//
	uploader.userMap = make(map[int64]int64)
	target = release_model.Release{}
	err = uploader.remapUser(&source, &target)
	assert.NoError(t, err)
	assert.EqualValues(t, linkedUser.ID, target.GetUserID())
}

func TestGiteaUploadUpdateGitForPullRequest(t *testing.T) {
	unittest.PrepareTestEnv(t)

	//
	// fromRepo master
	//
	fromRepo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	baseRef := "master"
	assert.NoError(t, git.InitRepository(git.DefaultContext, fromRepo.RepoPath(), false))
	err := git.NewCommand(git.DefaultContext, "symbolic-ref", "HEAD", git.BranchPrefix+baseRef).Run(&git.RunOpts{Dir: fromRepo.RepoPath()})
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(filepath.Join(fromRepo.RepoPath(), "README.md"), []byte(fmt.Sprintf("# Testing Repository\n\nOriginally created in: %s", fromRepo.RepoPath())), 0o644))
	assert.NoError(t, git.AddChanges(fromRepo.RepoPath(), true))
	signature := git.Signature{
		Email: "test@example.com",
		Name:  "test",
		When:  time.Now(),
	}
	assert.NoError(t, git.CommitChanges(fromRepo.RepoPath(), git.CommitChangesOptions{
		Committer: &signature,
		Author:    &signature,
		Message:   "Initial Commit",
	}))
	fromGitRepo, err := git.OpenRepository(git.DefaultContext, fromRepo.RepoPath())
	assert.NoError(t, err)
	defer fromGitRepo.Close()
	baseSHA, err := fromGitRepo.GetBranchCommitID(baseRef)
	assert.NoError(t, err)

	//
	// fromRepo branch1
	//
	headRef := "branch1"
	_, _, err = git.NewCommand(git.DefaultContext, "checkout", "-b", headRef).RunStdString(&git.RunOpts{Dir: fromRepo.RepoPath()})
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(filepath.Join(fromRepo.RepoPath(), "README.md"), []byte("SOMETHING"), 0o644))
	assert.NoError(t, git.AddChanges(fromRepo.RepoPath(), true))
	signature.When = time.Now()
	assert.NoError(t, git.CommitChanges(fromRepo.RepoPath(), git.CommitChangesOptions{
		Committer: &signature,
		Author:    &signature,
		Message:   "Pull request",
	}))
	assert.NoError(t, err)
	headSHA, err := fromGitRepo.GetBranchCommitID(headRef)
	assert.NoError(t, err)

	fromRepoOwner := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: fromRepo.OwnerID}).(*user_model.User)

	//
	// forkRepo branch2
	//
	forkHeadRef := "branch2"
	forkRepo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 8}).(*repo_model.Repository)
	assert.NoError(t, git.CloneWithArgs(git.DefaultContext, fromRepo.RepoPath(), forkRepo.RepoPath(), []string{}, git.CloneRepoOptions{
		Branch: headRef,
	}))
	_, _, err = git.NewCommand(git.DefaultContext, "checkout", "-b", forkHeadRef).RunStdString(&git.RunOpts{Dir: forkRepo.RepoPath()})
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(filepath.Join(forkRepo.RepoPath(), "README.md"), []byte(fmt.Sprintf("# branch2 %s", forkRepo.RepoPath())), 0o644))
	assert.NoError(t, git.AddChanges(forkRepo.RepoPath(), true))
	assert.NoError(t, git.CommitChanges(forkRepo.RepoPath(), git.CommitChangesOptions{
		Committer: &signature,
		Author:    &signature,
		Message:   "branch2 commit",
	}))
	forkGitRepo, err := git.OpenRepository(git.DefaultContext, forkRepo.RepoPath())
	assert.NoError(t, err)
	defer forkGitRepo.Close()
	forkHeadSHA, err := forkGitRepo.GetBranchCommitID(forkHeadRef)
	assert.NoError(t, err)

	toRepoName := "migrated"
	uploaderPrivate := migrations.NewGiteaLocalUploader(context.Background(), fromRepoOwner, fromRepoOwner.Name, toRepoName)
	uploader := *((*GiteaLocalUploader)(unsafe.Pointer(&*uploaderPrivate)))
	uploader.gitServiceType = structs.GitBundleService
	assert.NoError(t, uploader.CreateRepo(&schemas.Repository{
		Description: "description",
		OriginalURL: fromRepo.RepoPath(),
		CloneURL:    fromRepo.RepoPath(),
		IsPrivate:   false,
		IsMirror:    true,
	}, schemas.MigrateOptions{
		GitServiceType: structs.GitBundleService,
		Private:        false,
		Mirror:         true,
	}))

	for _, testCase := range []struct {
		name          string
		head          string
		assertContent func(t *testing.T, content string)
		pr            schemas.PullRequest
	}{
		{
			name: "fork, good Head.SHA",
			head: fmt.Sprintf("%s/%s", forkRepo.OwnerName, forkHeadRef),
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  forkRepo.RepoPath(),
					Ref:       forkHeadRef,
					SHA:       forkHeadSHA,
					RepoName:  forkRepo.Name,
					OwnerName: forkRepo.OwnerName,
				},
			},
		},
		{
			name: "fork, invalid Head.Ref",
			head: "unknown repository",
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  forkRepo.RepoPath(),
					Ref:       "INVALID",
					SHA:       forkHeadSHA,
					RepoName:  forkRepo.Name,
					OwnerName: forkRepo.OwnerName,
				},
			},
			assertContent: func(t *testing.T, content string) {
				assert.Contains(t, content, "Fetch branch from")
			},
		},
		{
			name: "invalid fork CloneURL",
			head: "unknown repository",
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  "UNLIKELY",
					Ref:       forkHeadRef,
					SHA:       forkHeadSHA,
					RepoName:  forkRepo.Name,
					OwnerName: "WRONG",
				},
			},
			assertContent: func(t *testing.T, content string) {
				assert.Contains(t, content, "AddRemote failed")
			},
		},
		{
			name: "no fork, good Head.SHA",
			head: headRef,
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       headRef,
					SHA:       headSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
			},
		},
		{
			name: "no fork, empty Head.SHA",
			head: headRef,
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       headRef,
					SHA:       "",
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
			},
			assertContent: func(t *testing.T, content string) {
				assert.Contains(t, content, "Empty reference, removing")
				assert.NotContains(t, content, "Cannot remove local head")
			},
		},
		{
			name: "no fork, invalid Head.SHA",
			head: headRef,
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       headRef,
					SHA:       "brokenSHA",
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
			},
			assertContent: func(t *testing.T, content string) {
				assert.Contains(t, content, "Deprecated local head")
				assert.Contains(t, content, "Cannot remove local head")
			},
		},
		{
			name: "no fork, not found Head.SHA",
			head: headRef,
			pr: schemas.PullRequest{
				PatchURL: "",
				Number:   1,
				State:    "open",
				Base: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       baseRef,
					SHA:       baseSHA,
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
				Head: schemas.PullRequestBranch{
					CloneURL:  fromRepo.RepoPath(),
					Ref:       headRef,
					SHA:       "2697b352310fcd01cbd1f3dbd43b894080027f68",
					RepoName:  fromRepo.Name,
					OwnerName: fromRepo.OwnerName,
				},
			},
			assertContent: func(t *testing.T, content string) {
				assert.Contains(t, content, "Deprecated local head")
				assert.NotContains(t, content, "Cannot remove local head")
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			logger, ok := log.NamedLoggers.Load(log.DEFAULT)
			assert.True(t, ok)
			logger.SetLogger("buffer", "buffer", "{}")
			defer logger.DelLogger("buffer")

			head, err := uploader.updateGitForPullRequest(&testCase.pr)
			assert.NoError(t, err)
			assert.EqualValues(t, testCase.head, head)
			if testCase.assertContent != nil {
				fence := fmt.Sprintf(">>>>>>>>>>>>>FENCE %s<<<<<<<<<<<<<<<", testCase.name)
				log.Error(fence)
				var content string
				for i := 0; i < 5000; i++ {
					content, err = logger.GetLoggerProviderContent("buffer")
					assert.NoError(t, err)
					if strings.Contains(content, fence) {
						break
					}
					time.Sleep(1 * time.Millisecond)
				}
				testCase.assertContent(t, content)
			}
		})
	}
}
