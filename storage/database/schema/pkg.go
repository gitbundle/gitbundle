// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Pkg struct {
	ent.Schema
}

func (Pkg) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.Int64("repo_id").Optional(),
		field.String("type").MaxLen(255),
		field.String("name").MaxLen(255),
		field.String("lower_name").MaxLen(255),
		field.Bool("semver_compatible").Default(false),
	}
}

func (Pkg) Edges() []ent.Edge {
	return nil
}

func (Pkg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "type", "lower_name").Unique(),
		index.Fields("repo_id"),
	}
}
