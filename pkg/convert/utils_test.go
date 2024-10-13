// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/convert"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestToCorrectPageSize(t *testing.T) {
	assert.EqualValues(t, 30, convert.ToCorrectPageSize(0))
	assert.EqualValues(t, 30, convert.ToCorrectPageSize(-10))
	assert.EqualValues(t, 20, convert.ToCorrectPageSize(20))
	assert.EqualValues(t, 50, convert.ToCorrectPageSize(100))
}

func TestToGitServiceType(t *testing.T) {
	tc := []struct {
		typ  string
		enum int
	}{{
		typ: "github", enum: 2,
	}, {
		typ: "gitea", enum: 3,
	}, {
		typ: "gitlab", enum: 4,
	}, {
		typ: "gogs", enum: 5,
	}, {
		typ: "trash", enum: 1,
	}}
	for _, test := range tc {
		assert.EqualValues(t, test.enum, convert.ToGitServiceType(test.typ))
	}
}
