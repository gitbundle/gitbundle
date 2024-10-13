// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package access_token_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/access_token"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	token := &access_token.AccessToken{
		UID:  3,
		Name: "Token C",
	}
	assert.NoError(t, access_token.NewAccessToken(token))
	unittest.AssertExistsAndLoadBean(t, token)

	invalidToken := &access_token.AccessToken{
		ID:   token.ID, // duplicate
		UID:  2,
		Name: "Token F",
	}
	assert.Error(t, access_token.NewAccessToken(invalidToken))
}

func TestAccessTokenByNameExists(t *testing.T) {
	name := "Token Gitea"

	assert.NoError(t, unittest.PrepareTestDatabase())
	token := &access_token.AccessToken{
		UID:  3,
		Name: name,
	}

	// Check to make sure it doesn't exists already
	exist, err := access_token.AccessTokenByNameExists(token)
	assert.NoError(t, err)
	assert.False(t, exist)

	// Save it to the database
	assert.NoError(t, access_token.NewAccessToken(token))
	unittest.AssertExistsAndLoadBean(t, token)

	// This token must be found by name in the DB now
	exist, err = access_token.AccessTokenByNameExists(token)
	assert.NoError(t, err)
	assert.True(t, exist)

	user4Token := &access_token.AccessToken{
		UID:  4,
		Name: name,
	}

	// Name matches but different user ID, this shouldn't exists in the
	// database
	exist, err = access_token.AccessTokenByNameExists(user4Token)
	assert.NoError(t, err)
	assert.False(t, exist)
}

func TestGetAccessTokenBySHA(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	token, err := access_token.GetAccessTokenBySHA("d2c6c1ba3890b309189a8e618c72a162e4efbf36")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), token.UID)
	assert.Equal(t, "Token A", token.Name)
	assert.Equal(t, "2b3668e11cb82d3af8c6e4524fc7841297668f5008d1626f0ad3417e9fa39af84c268248b78c481daa7e5dc437784003494f", token.TokenHash)
	assert.Equal(t, "e4efbf36", token.TokenLastEight)

	_, err = access_token.GetAccessTokenBySHA("notahash")
	assert.Error(t, err)
	assert.True(t, access_token.IsErrAccessTokenNotExist(err))

	_, err = access_token.GetAccessTokenBySHA("")
	assert.Error(t, err)
	assert.True(t, access_token.IsErrAccessTokenEmpty(err))
}

func TestListAccessTokens(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	tokens, err := access_token.ListAccessTokens(access_token.ListAccessTokensOptions{UserID: 1})
	assert.NoError(t, err)
	if assert.Len(t, tokens, 2) {
		assert.Equal(t, int64(1), tokens[0].UID)
		assert.Equal(t, int64(1), tokens[1].UID)
		assert.Contains(t, []string{tokens[0].Name, tokens[1].Name}, "Token A")
		assert.Contains(t, []string{tokens[0].Name, tokens[1].Name}, "Token B")
	}

	tokens, err = access_token.ListAccessTokens(access_token.ListAccessTokensOptions{UserID: 2})
	assert.NoError(t, err)
	if assert.Len(t, tokens, 1) {
		assert.Equal(t, int64(2), tokens[0].UID)
		assert.Equal(t, "Token A", tokens[0].Name)
	}

	tokens, err = access_token.ListAccessTokens(access_token.ListAccessTokensOptions{UserID: 100})
	assert.NoError(t, err)
	assert.Empty(t, tokens)
}

func TestUpdateAccessToken(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	token, err := access_token.GetAccessTokenBySHA("4c6f36e6cf498e2a448662f915d932c09c5a146c")
	assert.NoError(t, err)
	token.Name = "Token Z"

	assert.NoError(t, access_token.UpdateAccessToken(token))
	unittest.AssertExistsAndLoadBean(t, token)
}

func TestDeleteAccessTokenByID(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	token, err := access_token.GetAccessTokenBySHA("4c6f36e6cf498e2a448662f915d932c09c5a146c")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), token.UID)

	assert.NoError(t, access_token.DeleteAccessTokenByID(token.ID, 1))
	unittest.AssertNotExistsBean(t, token)

	err = access_token.DeleteAccessTokenByID(100, 100)
	assert.Error(t, err)
	assert.True(t, access_token.IsErrAccessTokenNotExist(err))
}
