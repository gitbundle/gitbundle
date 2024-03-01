// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserOpenid struct {
	ent.Schema
}

func (UserOpenid) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("uri").Optional(),
		field.Bool("show").Default(false),
	}
}

func (UserOpenid) Edges() []ent.Edge {
	return nil
}

func (UserOpenid) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uri").Unique(),
		index.Fields("user_id"),
	}
}
