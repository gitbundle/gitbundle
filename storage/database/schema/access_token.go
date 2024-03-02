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

type AccessToken struct {
	ent.Schema
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("name").Optional().MaxLen(255),
		field.String("token_hash").Optional().MaxLen(255).Sensitive(),
		field.String("token_salt").Optional().MaxLen(255).Sensitive(),
		field.String("token_last_eight").Optional().MaxLen(255).Sensitive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return nil
}

func (AccessToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("token_hash").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
