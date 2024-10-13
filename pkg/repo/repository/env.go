// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/gitbundle/modules/setting"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"
)

// env keys for git hooks need
const (
	EnvRepoName     = "GITBUNDLE_REPO_NAME"
	EnvRepoUsername = "GITBUNDLE_REPO_USER_NAME"
	EnvRepoID       = "GITBUNDLE_REPO_ID"
	EnvRepoIsWiki   = "GITBUNDLE_REPO_IS_WIKI"
	EnvPusherName   = "GITBUNDLE_PUSHER_NAME"
	EnvPusherEmail  = "GITBUNDLE_PUSHER_EMAIL"
	EnvPusherID     = "GITBUNDLE_PUSHER_ID"
	EnvKeyID        = "GITBUNDLE_KEY_ID" // public key ID
	EnvDeployKeyID  = "GITBUNDLE_DEPLOY_KEY_ID"
	EnvPRID         = "GITBUNDLE_PR_ID"
	EnvIsInternal   = "GITBUNDLE_INTERNAL_PUSH"
	EnvAppURL       = "GITBUNDLE_ROOT_URL"
)

// InternalPushingEnvironment returns an os environment to switch off hooks on push
// It is recommended to avoid using this unless you are pushing within a transaction
// or if you absolutely are sure that post-receive and pre-receive will do nothing
// We provide the full pushing-environment for other hook providers
func InternalPushingEnvironment(doer *user_model.User, repo *repo_model.Repository) []string {
	return append(PushingEnvironment(doer, repo),
		EnvIsInternal+"=true",
	)
}

// PushingEnvironment returns an os environment to allow hooks to work on push
func PushingEnvironment(doer *user_model.User, repo *repo_model.Repository) []string {
	return FullPushingEnvironment(doer, doer, repo, repo.Name, 0)
}

// FullPushingEnvironment returns an os environment to allow hooks to work on push
func FullPushingEnvironment(author, committer *user_model.User, repo *repo_model.Repository, repoName string, prID int64) []string {
	isWiki := "false"
	if strings.HasSuffix(repoName, ".wiki") {
		isWiki = "true"
	}

	authorSig := author.NewGitSig()
	committerSig := committer.NewGitSig()

	environ := append(os.Environ(),
		"GIT_AUTHOR_NAME="+authorSig.Name,
		"GIT_AUTHOR_EMAIL="+authorSig.Email,
		"GIT_COMMITTER_NAME="+committerSig.Name,
		"GIT_COMMITTER_EMAIL="+committerSig.Email,
		EnvRepoName+"="+repoName,
		EnvRepoUsername+"="+repo.OwnerName,
		EnvRepoIsWiki+"="+isWiki,
		EnvPusherName+"="+committer.Name,
		EnvPusherID+"="+fmt.Sprintf("%d", committer.ID),
		EnvRepoID+"="+fmt.Sprintf("%d", repo.ID),
		EnvPRID+"="+fmt.Sprintf("%d", prID),
		EnvAppURL+"="+setting.AppURL,
		"SSH_ORIGINAL_COMMAND=gitea-internal",
	)

	if !committer.KeepEmailPrivate {
		environ = append(environ, EnvPusherEmail+"="+committer.Email)
	}

	return environ
}
