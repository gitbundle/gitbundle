// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RepoArchiver struct {
	ent.Schema
}

func (RepoArchiver) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int("type").Optional(),
		field.Int("status").Optional(),
		field.String("commit_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (RepoArchiver) Edges() []ent.Edge {
	return nil
}

func (RepoArchiver) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "type", "commit_id").Unique(),
		index.Fields("created_at"),
	}
}
