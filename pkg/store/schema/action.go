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

type Action struct {
	ent.Schema
}

func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int("op_type").Optional(),
		field.Int64("act_user_id").Optional(),
		field.Int64("repo_id").Optional(),
		// field.Int64("comment_id").Optional(),
		field.Bool("is_deleted").Default(false),
		field.String("ref_name").Optional(),
		field.Bool("is_private").Default(false),
		field.Text("content").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Action) Edges() []ent.Edge {
	return nil
}

func (Action) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "user_id", "is_deleted"),
		index.Fields("act_user_id", "repo_id", "created_at", "user_id", "is_deleted"),
		// index.Fields("comment_id"),
	}
}
