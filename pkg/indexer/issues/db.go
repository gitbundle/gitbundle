// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package issues

import (
	"context"

	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
)

// DBIndexer implements Indexer interface to use database's like search
type DBIndexer struct{}

// Init dummy function
func (i *DBIndexer) Init() (bool, error) {
	return false, nil
}

// SetAvailabilityChangeCallback dummy function
func (i *DBIndexer) SetAvailabilityChangeCallback(callback func(bool)) {
}

// Ping checks if database is available
func (i *DBIndexer) Ping() bool {
	return db.GetEngine(db.DefaultContext).Ping() != nil
}

// Index dummy function
func (i *DBIndexer) Index(issue []*IndexerData) error {
	return nil
}

// Delete dummy function
func (i *DBIndexer) Delete(ids ...int64) error {
	return nil
}

// Close dummy function
func (i *DBIndexer) Close() {
}

// Search dummy function
func (i *DBIndexer) Search(ctx context.Context, kw string, repoIDs []int64, limit, start int) (*SearchResult, error) {
	total, ids, err := issues_model.SearchIssueIDsByKeyword(ctx, kw, repoIDs, limit, start)
	if err != nil {
		return nil, err
	}
	result := SearchResult{
		Total: total,
		Hits:  make([]Match, 0, limit),
	}
	for _, id := range ids {
		result.Hits = append(result.Hits, Match{
			ID: id,
		})
	}
	return &result, nil
}
