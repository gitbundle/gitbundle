// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package stats

import (
	"fmt"

	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/queue"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

// statsQueue represents a queue to handle repository stats updates
var statsQueue queue.UniqueQueue

// handle passed PR IDs and test the PRs
func handle(data ...queue.Data) []queue.Data {
	for _, datum := range data {
		opts := datum.(int64)
		if err := indexer.Index(opts); err != nil {
			log.Error("stats queue indexer.Index(%d) failed: %v", opts, err)
		}
	}
	return nil
}

func initStatsQueue() error {
	statsQueue = queue.CreateUniqueQueue("repo_stats_update", handle, int64(0))
	if statsQueue == nil {
		return fmt.Errorf("Unable to create repo_stats_update Queue")
	}

	go graceful.GetManager().RunWithShutdownFns(statsQueue.Run)

	return nil
}

// UpdateRepoIndexer update a repository's entries in the indexer
func UpdateRepoIndexer(repo *repo_model.Repository) error {
	if err := statsQueue.Push(repo.ID); err != nil {
		if err != queue.ErrAlreadyInQueue {
			return err
		}
		log.Debug("Repo ID: %d already queued", repo.ID)
	}
	return nil
}
