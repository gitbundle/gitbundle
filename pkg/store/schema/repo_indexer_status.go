// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RepoIndexerStatus struct {
	ent.Schema
}

func (RepoIndexerStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.String("commit_sha").Optional().MaxLen(40),
		field.Int("indexer_type"),
	}
}

func (RepoIndexerStatus) Edges() []ent.Edge {
	return nil
}

func (RepoIndexerStatus) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "indexer_type"),
	}
}
