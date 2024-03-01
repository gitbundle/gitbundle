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

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("org_id").Optional(),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.String("color").Optional().MaxLen(7),
		field.Int64("num_issues").Optional(),
		field.Int64("num_closed_issues").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Label) Edges() []ent.Edge {
	return nil
}

func (Label) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("org_id"),
		index.Fields("updated_at"),
		index.Fields("repo_id"),
	}
}
