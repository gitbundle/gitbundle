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

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("doer_id").Optional(),
		field.Int64("owner_id").Optional(),
		field.Int64("repo_id").Optional(),
		field.Int("type").Optional(),
		field.Int("status").Optional(),
		field.Time("started_at").Optional(),
		field.Time("ended_at").Optional(),
		field.Text("payload_content").Optional(),
		field.Text("message").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Task) Edges() []ent.Edge {
	return nil
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("doer_id"),
		index.Fields("owner_id"),
		index.Fields("repo_id"),
		index.Fields("status"),
	}
}
