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

type CommitStatus struct {
	ent.Schema
}

func (CommitStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("index").Optional(),
		field.Int64("repo_id").Optional(),
		field.String("state").MaxLen(7),
		field.String("sha").MaxLen(64),
		field.Text("target_url").Optional(),
		field.Text("description").Optional(),
		field.String("context_hash").MaxLen(40),
		field.Text("context").Optional(),
		field.Int64("creator_id").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (CommitStatus) Edges() []ent.Edge {
	return nil
}

func (CommitStatus) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("context_hash"),
		index.Fields("index", "repo_id", "sha").Unique(),
		index.Fields("updated_at"),
	}
}
