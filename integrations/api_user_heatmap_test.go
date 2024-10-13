// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package integrations

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/user_heatmap"

	"github.com/stretchr/testify/assert"
)

func TestUserHeatmap(t *testing.T) {
	defer prepareTestEnv(t)()
	adminUsername := "user1"
	normalUsername := "user2"
	token := getUserToken(t, adminUsername)

	fakeNow := time.Date(2011, 10, 20, 0, 0, 0, 0, time.Local)
	timeutil.Set(fakeNow)
	defer timeutil.Unset()

	urlStr := fmt.Sprintf("/api/v1/users/%s/heatmap?token=%s", normalUsername, token)
	req := NewRequest(t, "GET", urlStr)
	resp := MakeRequest(t, req, http.StatusOK)
	var heatmap []*user_heatmap.UserHeatmapData
	DecodeJSON(t, resp, &heatmap)
	var dummyheatmap []*user_heatmap.UserHeatmapData
	dummyheatmap = append(dummyheatmap, &user_heatmap.UserHeatmapData{Timestamp: 1603227600, Contributions: 1})

	assert.Equal(t, dummyheatmap, heatmap)
}
