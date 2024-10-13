// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/builder"
	"xorm.io/xorm"
)

func fixTopicRepositoryCount(x *xorm.Engine) error {
	_, err := x.Exec(builder.Delete(builder.NotIn("`repo_id`", builder.Select("`id`").From("`repository`"))).From("`repo_topic`"))
	if err != nil {
		return err
	}

	_, err = x.Exec(builder.Update(
		builder.Eq{
			"`repo_count`": builder.Select("count(*)").From("`repo_topic`").Where(builder.Eq{
				"`repo_topic`.`topic_id`": builder.Expr("`topic`.`id`"),
			}),
		}).From("`topic`").Where(builder.Eq{"'1'": "1"}))
	return err
}
