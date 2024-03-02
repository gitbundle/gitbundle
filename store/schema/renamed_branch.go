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

type RenamedBranch struct {
	ent.Schema
}

func (RenamedBranch) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.String("from").Optional().MaxLen(255),
		field.String("to").Optional().MaxLen(255),
		field.Time("created_at").Default(time.Now),
	}
}

func (RenamedBranch) Edges() []ent.Edge {
	return nil
}

func (RenamedBranch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id"),
	}
}
