// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func increaseCredentialIDTo410(x *xorm.Engine) error {
	// no-op
	// v208 was completely wrong
	// So now we have to no-op again.

	return nil
}
