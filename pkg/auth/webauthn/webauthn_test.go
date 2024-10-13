// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package webauthn

import (
	"testing"

	"github.com/gitbundle/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	setting.Domain = "domain"
	setting.AppName = "AppName"
	setting.AppURL = "https://domain/"
	rpOrigin := "https://domain"

	Init()

	assert.Equal(t, setting.Domain, WebAuthn.Config.RPID)
	assert.Equal(t, setting.AppName, WebAuthn.Config.RPDisplayName)
	assert.Equal(t, rpOrigin, WebAuthn.Config.RPOrigin)
}
