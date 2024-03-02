// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Team struct {
	ent.Schema
}

func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("org_id").Optional(),
		field.String("lower_name").Optional().MaxLen(255),
		field.String("name").Optional().MaxLen(255),
		field.String("description").Optional().MaxLen(255),
		field.Int("authorize").Optional(),
		field.Int("num_repos").Optional(),
		field.Int("num_members").Optional(),
		field.Bool("include_all_repos").Default(false),
		field.Bool("can_create_org_repo").Default(false),
	}
}

func (Team) Edges() []ent.Edge {
	return nil
}

func (Team) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("org_id"),
	}
}
