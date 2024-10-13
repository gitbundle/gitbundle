// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package files

import (
	"context"
	"strings"

	"github.com/gitbundle/server/pkg/git/gitdiff"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

// GetDiffPreview produces and returns diff result of a file which is not yet committed.
func GetDiffPreview(ctx context.Context, repo *repo_model.Repository, branch, treePath, content string) (*gitdiff.Diff, error) {
	if branch == "" {
		branch = repo.DefaultBranch
	}
	t, err := NewTemporaryUploadRepository(ctx, repo)
	if err != nil {
		return nil, err
	}
	defer t.Close()
	if err := t.Clone(branch); err != nil {
		return nil, err
	}
	if err := t.SetDefaultIndex(); err != nil {
		return nil, err
	}

	// Add the object to the database
	objectHash, err := t.HashObject(strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	// Add the object to the index
	if err := t.AddObjectToIndex("100644", objectHash, treePath); err != nil {
		return nil, err
	}
	return t.DiffIndex()
}
