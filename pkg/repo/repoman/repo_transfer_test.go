// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repoman_test

import (
	"testing"

	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryTransfer(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 3}).(*user_model.User)
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 3}).(*repo_model.Repository)

	transfer, err := repoman.GetPendingRepositoryTransfer(repo)
	assert.NoError(t, err)
	assert.NotNil(t, transfer)

	// Cancel transfer
	assert.NoError(t, repoman.CancelRepositoryTransfer(repo))

	transfer, err = repoman.GetPendingRepositoryTransfer(repo)
	assert.Error(t, err)
	assert.Nil(t, transfer)
	assert.True(t, repoman.IsErrNoPendingTransfer(err))

	user2 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)

	assert.NoError(t, repoman.CreatePendingRepositoryTransfer(doer, user2, repo.ID, nil))

	transfer, err = repoman.GetPendingRepositoryTransfer(repo)
	assert.Nil(t, err)
	assert.NoError(t, transfer.LoadAttributes())
	assert.Equal(t, "user2", transfer.Recipient.Name)

	user6 := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)

	// Only transfer can be started at any given time
	err = repoman.CreatePendingRepositoryTransfer(doer, user6, repo.ID, nil)
	assert.Error(t, err)
	assert.True(t, repoman.IsErrRepoTransferInProgress(err))

	// Unknown user
	err = repoman.CreatePendingRepositoryTransfer(doer, &user_model.User{ID: 1000, LowerName: "user1000"}, repo.ID, nil)
	assert.Error(t, err)

	// Cancel transfer
	assert.NoError(t, repoman.CancelRepositoryTransfer(repo))
}
