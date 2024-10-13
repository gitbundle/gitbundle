// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build pam

package pam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPamAuth(t *testing.T) {
	result, err := Auth("gitea", "user1", "false-pwd")
	assert.Error(t, err)
	assert.EqualError(t, err, "Authentication failure")
	assert.Len(t, result, 0)
}
