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

type RepoUnit struct {
	ent.Schema
}

func (RepoUnit) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int("type").Optional(),
		field.Text("config").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (RepoUnit) Edges() []ent.Edge {
	return nil
}

func (RepoUnit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "type"),
		index.Fields("created_at"),
	}
}
