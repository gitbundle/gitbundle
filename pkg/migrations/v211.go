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

func createForeignReferenceTable(x *xorm.Engine) error {
	type ForeignReference struct {
		// RepoID is the first column in all indices. now we only need 2 indices: (repo, local) and (repo, foreign, type)
		RepoID       int64  `xorm:"UNIQUE(repo_foreign_type) INDEX(repo_local)" `
		LocalIndex   int64  `xorm:"INDEX(repo_local)"` // the resource key inside Gitea, it can be IssueIndex, or some model ID.
		ForeignIndex string `xorm:"INDEX UNIQUE(repo_foreign_type)"`
		Type         string `xorm:"VARCHAR(16) INDEX UNIQUE(repo_foreign_type)"`
	}

	if err := x.Sync2(new(ForeignReference)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}
	return nil
}
