// Copyright 2024 GitBundle, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gitbundle/gitbundle/git/command"
	"github.com/gitbundle/gitbundle/git/hook"
	"github.com/gitbundle/gitbundle/git/types"
	cache "github.com/gitbundle/gitbundle/storage/cache/adapter"

	"github.com/hashicorp/go-version"
)

// GitVersionRequired is the minimum Git version required
const GitVersionRequired = "2.0.0"

type Adapter struct {
	traceGit        bool
	lastCommitCache cache.Cache[CommitEntryKey, *types.Commit]
	githookFactory  hook.ClientFactory
}

func New(
	ctx context.Context,
	config types.Config,
	lastCommitCache cache.Cache[CommitEntryKey, *types.Commit],
	githookFactory hook.ClientFactory,
) (Adapter, error) {
	err := checkGitVersion(ctx)
	if err != nil {
		return Adapter{}, err
	}

	return Adapter{
		traceGit:        config.Trace,
		lastCommitCache: lastCommitCache,
		githookFactory:  githookFactory,
	}, nil
}

func checkGitVersion(ctx context.Context) error {
	absPath, err := exec.LookPath(command.GitExecutable)
	if err != nil {
		return fmt.Errorf("git not found: %w", err)
	}
	command.GitExecutable = absPath

	gitVersion, err := getGitVersion(ctx)
	if err != nil {
		return fmt.Errorf("unable to load git version: %w", err)
	}

	versionRequired, err := version.NewVersion(GitVersionRequired)
	if err != nil {
		return err
	}

	if gitVersion.LessThan(versionRequired) {
		moreHint := "get git: https://git-scm.com/download/"
		if runtime.GOOS == "linux" {
			// there are a lot of CentOS/RHEL users using old git, so we add a special hint for them
			if _, err = os.Stat("/etc/redhat-release"); err == nil {
				// ius.io is the recommended official(git-scm.com) method to install git
				moreHint = "get git: https://git-scm.com/download/linux and https://ius.io"
			}
		}
		return fmt.Errorf("installed git version %q is not supported, Gitea requires git version >= %q, %s", gitVersion.Original(), GitVersionRequired, moreHint)
	}

	return nil
}
