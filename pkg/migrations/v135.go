// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"xorm.io/xorm"
)

func addOrgIDLabelColumn(x *xorm.Engine) error {
	type Label struct {
		OrgID int64 `xorm:"INDEX"`
	}

	if err := x.Sync2(new(Label)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}
	return nil
}
