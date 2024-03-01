// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Mirror struct {
	ent.Schema
}

func (Mirror) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("interval").Optional(),
		field.Bool("enable_prune").Default(true),
		field.Time("updated_at").Optional(),
		field.Time("next_updated_at").Optional(),
		field.Bool("lfs_enabled").Default(false),
		field.Text("lfs_endpoint").Optional(),
	}
}

func (Mirror) Edges() []ent.Edge {
	return nil
}

func (Mirror) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id"),
		index.Fields("updated_at"),
		index.Fields("next_updated_at"),
	}
}
