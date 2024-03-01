// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type DeletedBranch struct {
	ent.Schema
}

func (DeletedBranch) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.String("name"),
		field.String("commit"),
		field.Int64("deleted_by_id").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

func (DeletedBranch) Edges() []ent.Edge {
	return nil
}

func (DeletedBranch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
		index.Fields("deleted_by_id"),
		index.Fields("repo_id", "name", "commit").Unique(),
	}
}
