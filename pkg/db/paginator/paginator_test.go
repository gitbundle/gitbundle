// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package paginator_test

import (
	"testing"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestPaginator(t *testing.T) {
	cases := []struct {
		db.Paginator
		Skip  int
		Take  int
		Start int
		End   int
	}{
		{
			Paginator: &db.ListOptions{Page: -1, PageSize: -1},
			Skip:      0,
			Take:      setting.API.DefaultPagingNum,
			Start:     0,
			End:       setting.API.DefaultPagingNum,
		},
		{
			Paginator: &db.ListOptions{Page: 2, PageSize: 10},
			Skip:      10,
			Take:      10,
			Start:     10,
			End:       20,
		},
		{
			Paginator: db.NewAbsoluteListOptions(-1, -1),
			Skip:      0,
			Take:      setting.API.DefaultPagingNum,
			Start:     0,
			End:       setting.API.DefaultPagingNum,
		},
		{
			Paginator: db.NewAbsoluteListOptions(2, 10),
			Skip:      2,
			Take:      10,
			Start:     2,
			End:       12,
		},
	}

	for _, c := range cases {
		skip, take := c.Paginator.GetSkipTake()
		start, end := c.Paginator.GetStartEnd()

		assert.Equal(t, c.Skip, skip)
		assert.Equal(t, c.Take, take)
		assert.Equal(t, c.Start, start)
		assert.Equal(t, c.End, end)
	}
}
