// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert

import (
	"time"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/util"
	repo_model "github.com/gitbundle/server/pkg/repo"
	wiki_service "github.com/gitbundle/server/pkg/repo/wiki"
)

// ToWikiCommit convert a git commit into a WikiCommit
func ToWikiCommit(commit *git.Commit) *api.WikiCommit {
	return &api.WikiCommit{
		ID: commit.ID.String(),
		Author: &api.CommitUser{
			Identity: api.Identity{
				Name:  commit.Author.Name,
				Email: commit.Author.Email,
			},
			Date: commit.Author.When.UTC().Format(time.RFC3339),
		},
		Committer: &api.CommitUser{
			Identity: api.Identity{
				Name:  commit.Committer.Name,
				Email: commit.Committer.Email,
			},
			Date: commit.Committer.When.UTC().Format(time.RFC3339),
		},
		Message: commit.CommitMessage,
	}
}

// ToWikiCommitList convert a list of git commits into a WikiCommitList
func ToWikiCommitList(commits []*git.Commit, total int64) *api.WikiCommitList {
	result := make([]*api.WikiCommit, len(commits))
	for i := range commits {
		result[i] = ToWikiCommit(commits[i])
	}
	return &api.WikiCommitList{
		WikiCommits: result,
		Count:       total,
	}
}

// ToWikiPageMetaData converts meta information to a WikiPageMetaData
func ToWikiPageMetaData(title string, lastCommit *git.Commit, repo *repo_model.Repository) *api.WikiPageMetaData {
	suburl := wiki_service.NameToSubURL(title)
	return &api.WikiPageMetaData{
		Title:      title,
		HTMLURL:    util.URLJoin(repo.HTMLURL(), "wiki", suburl),
		SubURL:     suburl,
		LastCommit: ToWikiCommit(lastCommit),
	}
}
