// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"errors"

	"github.com/google/go-github/v45/github"
)

// ErrRepoNotCreated returns the error that repository not created
var ErrRepoNotCreated = errors.New("repository is not created yet")

// IsRateLimitError returns true if the err is github.RateLimitError
func IsRateLimitError(err error) bool {
	_, ok := err.(*github.RateLimitError)
	return ok
}

// IsTwoFactorAuthError returns true if the err is github.TwoFactorAuthError
func IsTwoFactorAuthError(err error) bool {
	_, ok := err.(*github.TwoFactorAuthError)
	return ok
}
