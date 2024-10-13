// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert_test

import (
	"testing"
	"time"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/convert"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestToCommitMeta(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	headRepo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1}).(*repo_model.Repository)
	sha1, _ := git.NewIDFromString("0000000000000000000000000000000000000000")
	signature := &git.Signature{Name: "Test Signature", Email: "test@email.com", When: time.Unix(0, 0)}
	tag := &git.Tag{
		Name:    "Test Tag",
		ID:      sha1,
		Object:  sha1,
		Type:    "Test Type",
		Tagger:  signature,
		Message: "Test Message",
	}

	commitMeta := convert.ToCommitMeta(headRepo, tag)

	assert.NotNil(t, commitMeta)
	assert.EqualValues(t, &api.CommitMeta{
		SHA:     "0000000000000000000000000000000000000000",
		URL:     util.URLJoin(headRepo.APIURL(), "git/commits", "0000000000000000000000000000000000000000"),
		Created: time.Unix(0, 0),
	}, commitMeta)
}
