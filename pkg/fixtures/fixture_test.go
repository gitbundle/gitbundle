// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package fixtures_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/fixtures"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
)

func TestFixtureGeneration(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	test := func(gen func() (string, error), name string) {
		expected, err := gen()
		if !assert.NoError(t, err) {
			return
		}
		bytes, err := os.ReadFile(filepath.Join(unittest.FixturesDir(), name+".yml"))
		if !assert.NoError(t, err) {
			return
		}
		data := string(util.NormalizeEOL(bytes))
		assert.True(t, data == expected, "Differences detected for %s.yml", name)
	}

	test(fixtures.GetYamlFixturesAccess, "access")
}
