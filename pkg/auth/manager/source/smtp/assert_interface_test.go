// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package smtp_test

import (
	auth_model "github.com/gitbundle/server/pkg/auth"
	auth "github.com/gitbundle/server/pkg/auth/manager"
	"github.com/gitbundle/server/pkg/auth/manager/source/smtp"
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.PasswordAuthenticator
	auth_model.Config
	auth_model.SkipVerifiable
	auth_model.HasTLSer
	auth_model.UseTLSer
	auth_model.SourceSettable
}

var _ sourceInterface = &smtp.Source{}
