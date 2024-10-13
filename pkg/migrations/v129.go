// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func purgeUnusedDependencies(x *xorm.Engine) error {
	if _, err := x.Exec("DELETE FROM issue_dependency WHERE issue_id NOT IN (SELECT id FROM issue)"); err != nil {
		return err
	}
	_, err := x.Exec("DELETE FROM issue_dependency WHERE dependency_id NOT IN (SELECT id FROM issue)")
	return err
}
