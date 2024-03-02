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

type TwoFactor struct {
	ent.Schema
}

func (TwoFactor) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("secret").Optional().MaxLen(255).Sensitive(),
		field.String("scratch_salt").Optional().MaxLen(255).Sensitive(),
		field.String("scratch_hash").Optional().MaxLen(255).Sensitive(),
		field.String("last_used_passcode").Optional().MaxLen(10).Sensitive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (TwoFactor) Edges() []ent.Edge {
	return nil
}

func (TwoFactor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
