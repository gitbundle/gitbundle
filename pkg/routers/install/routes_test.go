// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package install

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	routes := Routes()
	assert.NotNil(t, routes)
	assert.EqualValues(t, "/", routes.R.Routes()[0].Pattern)
	assert.Nil(t, routes.R.Routes()[0].SubRoutes)
	assert.Len(t, routes.R.Routes()[0].Handlers, 2)
}
