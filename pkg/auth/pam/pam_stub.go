// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !pam

package pam

import (
	"errors"
)

// Supported is false when built without PAM
var Supported = false

// Auth not supported lack of pam tag
func Auth(serviceName, userName, passwd string) (string, error) {
	return "", errors.New("PAM not supported")
}
