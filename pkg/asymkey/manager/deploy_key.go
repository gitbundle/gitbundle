// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package manager

import (
	asymkey_model "github.com/gitbundle/server/pkg/asymkey"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/repo/repoman"
	user_model "github.com/gitbundle/server/pkg/user"
)

// DeleteDeployKey deletes deploy key from its repository authorized_keys file if needed.
func DeleteDeployKey(doer *user_model.User, id int64) error {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()

	if err := repoman.DeleteDeployKey(ctx, doer, id); err != nil {
		return err
	}
	if err := committer.Commit(); err != nil {
		return err
	}

	return asymkey_model.RewriteAllPublicKeys()
}
