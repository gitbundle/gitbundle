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
	"bytes"
	"context"
	"fmt"

	"github.com/gitbundle/gitbundle/git/command"
	"github.com/hashicorp/go-version"
)

// getGitVersion returns current Git version from shell.
func getGitVersion(ctx context.Context) (*version.Version, error) {
	cmd := command.New("version")
	output := &bytes.Buffer{}
	if err := cmd.Run(ctx, command.WithStdout(output)); err != nil {
		return nil, fmt.Errorf("failed to trigger version command: %w", err)
	}

	data := output.Bytes()

	fields := bytes.Fields(data)
	if len(fields) < 3 {
		return nil, fmt.Errorf("invalid git version output: %s", data)
	}

	var versionBytes []byte

	// Handle special case on Windows.
	i := bytes.Index(fields[2], []byte("windows"))
	if i >= 1 {
		versionBytes = fields[2][:i-1]
	} else {
		versionBytes = fields[2]
	}

	return version.NewVersion(string(versionBytes))
}
