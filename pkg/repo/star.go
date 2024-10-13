// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"context"

	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

// Star represents a starred repo by an user.
type Star struct {
	ID          int64              `xorm:"pk autoincr"`
	UID         int64              `xorm:"UNIQUE(s)"`
	RepoID      int64              `xorm:"UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

func init() {
	db.RegisterModel(new(Star))
}

// StarRepo or unstar repository.
func StarRepo(userID, repoID int64, star bool) error {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return err
	}
	defer committer.Close()
	staring := IsStaring(ctx, userID, repoID)

	if star {
		if staring {
			return nil
		}

		if err := db.Insert(ctx, &Star{UID: userID, RepoID: repoID}); err != nil {
			return err
		}
		if _, err := db.Exec(ctx, "UPDATE `repository` SET num_stars = num_stars + 1 WHERE id = ?", repoID); err != nil {
			return err
		}
		if _, err := db.Exec(ctx, "UPDATE `user` SET num_stars = num_stars + 1 WHERE id = ?", userID); err != nil {
			return err
		}
	} else {
		if !staring {
			return nil
		}

		if _, err := db.DeleteByBean(ctx, &Star{UID: userID, RepoID: repoID}); err != nil {
			return err
		}
		if _, err := db.Exec(ctx, "UPDATE `repository` SET num_stars = num_stars - 1 WHERE id = ?", repoID); err != nil {
			return err
		}
		if _, err := db.Exec(ctx, "UPDATE `user` SET num_stars = num_stars - 1 WHERE id = ?", userID); err != nil {
			return err
		}
	}

	return committer.Commit()
}

// IsStaring checks if user has starred given repository.
func IsStaring(ctx context.Context, userID, repoID int64) bool {
	has, _ := db.GetEngine(ctx).Get(&Star{UID: userID, RepoID: repoID})
	return has
}

// GetStargazers returns the users that starred the repo.
func GetStargazers(repo *Repository, opts db.ListOptions) ([]*user_model.User, error) {
	sess := db.GetEngine(db.DefaultContext).Where("star.repo_id = ?", repo.ID).
		Join("LEFT", "star", "`user`.id = star.uid")
	if opts.Page > 0 {
		sess = db.SetSessionPagination(sess, &opts)

		users := make([]*user_model.User, 0, opts.PageSize)
		return users, sess.Find(&users)
	}

	users := make([]*user_model.User, 0, 8)
	return users, sess.Find(&users)
}
