// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/png"
	"testing"

	repo_model "github.com/gitbundle/server/pkg/repo"
	repository "github.com/gitbundle/server/pkg/repo/repository_manager"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestUploadAvatar(t *testing.T) {
	// Generate image
	myImage := image.NewRGBA(image.Rect(0, 0, 1, 1))
	var buff bytes.Buffer
	png.Encode(&buff, myImage)

	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 10}).(*repo_model.Repository)

	err := repository.UploadAvatar(repo, buff.Bytes())
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%d-%x", 10, md5.Sum(buff.Bytes())), repo.Avatar)
}

func TestUploadBigAvatar(t *testing.T) {
	// Generate BIG image
	myImage := image.NewRGBA(image.Rect(0, 0, 5000, 1))
	var buff bytes.Buffer
	png.Encode(&buff, myImage)

	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 10}).(*repo_model.Repository)

	err := repository.UploadAvatar(repo, buff.Bytes())
	assert.Error(t, err)
}

func TestDeleteAvatar(t *testing.T) {
	// Generate image
	myImage := image.NewRGBA(image.Rect(0, 0, 1, 1))
	var buff bytes.Buffer
	png.Encode(&buff, myImage)

	assert.NoError(t, unittest.PrepareTestDatabase())
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 10}).(*repo_model.Repository)

	err := repository.UploadAvatar(repo, buff.Bytes())
	assert.NoError(t, err)

	err = repository.DeleteAvatar(repo)
	assert.NoError(t, err)

	assert.Equal(t, "", repo.Avatar)
}
