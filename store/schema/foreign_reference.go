// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ForeignReference struct {
	ent.Schema
}

func (ForeignReference) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("local_index").Optional(),
		field.String("foreign_index").Optional().MaxLen(255),
		field.String("type").Optional().MaxLen(16),
	}
}

func (ForeignReference) Edges() []ent.Edge {
	return nil
}

func (ForeignReference) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("repo_id", "foreign_index", "type").Unique(),
		index.Fields("repo_id", "local_index"),
	}
}
