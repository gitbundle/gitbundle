// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"context"

	"github.com/gitbundle/server/pkg/migrations/manager/schemas"
)

var _ schemas.Downloader = &PlainGitDownloader{}

// PlainGitDownloader implements a Downloader interface to clone git from a http/https URL
type PlainGitDownloader struct {
	schemas.NullDownloader
	ownerName string
	repoName  string
	remoteURL string
}

// NewPlainGitDownloader creates a git Downloader
func NewPlainGitDownloader(ownerName, repoName, remoteURL string) *PlainGitDownloader {
	return &PlainGitDownloader{
		ownerName: ownerName,
		repoName:  repoName,
		remoteURL: remoteURL,
	}
}

// SetContext set context
func (g *PlainGitDownloader) SetContext(ctx context.Context) {
}

// GetRepoInfo returns a repository information
func (g *PlainGitDownloader) GetRepoInfo() (*schemas.Repository, error) {
	// convert github repo to stand Repo
	return &schemas.Repository{
		Owner:    g.ownerName,
		Name:     g.repoName,
		CloneURL: g.remoteURL,
	}, nil
}

// GetTopics return empty string slice
func (g PlainGitDownloader) GetTopics() ([]string, error) {
	return []string{}, nil
}
