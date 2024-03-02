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

type Notification struct {
	ent.Schema
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("repo_id"),
		field.Int16("status"),
		field.Int16("source"),
		// field.Int64("issue_id"), // Should sync from issue service?
		field.String("commit_id").Optional().MaxLen(255),
		// field.Int64("comment_id").Optional(),//Should sync from comment service?
		field.Int("updated_by"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Notification) Edges() []ent.Edge {
	return nil
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source"),
		index.Fields("commit_id"),
		index.Fields("repo_id"),
		// index.Fields("issue_id"),
		index.Fields("updated_by"),
		index.Fields("updated_at"),
		index.Fields("created_at"),
		index.Fields("user_id"),
		index.Fields("status"),
	}
}
