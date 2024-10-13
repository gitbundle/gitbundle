// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"path/filepath"
	"testing"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/unittest"
	webhook_service "github.com/gitbundle/server/pkg/webhook/manager"
)

func TestMain(m *testing.M) {
	setting.LoadForTest()
	setting.NewQueueService()
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", "..", "..", ".."),
		SetUp:         webhook_service.Init,
	})
}
