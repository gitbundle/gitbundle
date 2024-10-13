// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package x

import (
	"github.com/gitbundle/modules/nosql"
	"github.com/gitbundle/modules/setting"
	"github.com/redis/go-redis/v9"
)

var RedisClient *rdbClient

type rdbClient struct {
	redis.UniversalClient
}

func InitRedisClient() error {
	redisClient := nosql.GetManager().GetRedisClient(setting.Redis.Connection)
	RedisClient = &rdbClient{redisClient}
	return nil
}
