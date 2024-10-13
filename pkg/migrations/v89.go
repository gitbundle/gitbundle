// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import "xorm.io/xorm"

func addOriginalMigrationInfo(x *xorm.Engine) error {
	// Issue see models/issue.go
	type Issue struct {
		OriginalAuthor   string
		OriginalAuthorID int64
	}

	if err := x.Sync2(new(Issue)); err != nil {
		return err
	}

	// Issue see models/issue_comment.go
	type Comment struct {
		OriginalAuthor   string
		OriginalAuthorID int64
	}

	if err := x.Sync2(new(Comment)); err != nil {
		return err
	}

	// Issue see models/repo.go
	type Repository struct {
		OriginalURL string
	}

	return x.Sync2(new(Repository))
}
