// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func fixMigratedRepositoryServiceType(x *xorm.Engine) error {
	// structs.GithubService:
	// GithubService = 2
	_, err := x.Exec("UPDATE repository SET original_service_type = ? WHERE original_url LIKE 'https://github.com/%'", 2)
	return err
}
