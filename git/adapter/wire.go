// Copyright 2024 GitBundle, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"fmt"
	"time"

	"github.com/gitbundle/gitbundle/git/enum"
	"github.com/gitbundle/gitbundle/git/types"
	cache "github.com/gitbundle/gitbundle/storage/cache/adapter"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var WireSet = wire.NewSet(
	ProvideLastCommitCache,
)

func ProvideLastCommitCache(
	config types.Config,
	redisClient redis.UniversalClient,
) (cache.Cache[CommitEntryKey, *types.Commit], error) {
	cacheDuration := config.LastCommitCache.Duration

	// no need to cache if it's too short
	if cacheDuration < time.Second {
		return NoLastCommitCache(), nil
	}

	switch config.LastCommitCache.Mode {
	case enum.LastCommitCacheModeNone:
		return NoLastCommitCache(), nil
	case enum.LastCommitCacheModeInMemory:
		return NewInMemoryLastCommitCache(cacheDuration), nil
	case enum.LastCommitCacheModeRedis:
		return NewRedisLastCommitCache(redisClient, cacheDuration)
	default:
		return nil, fmt.Errorf("unknown last commit cache mode provided: %q", config.LastCommitCache.Mode)
	}
}
