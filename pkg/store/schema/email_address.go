// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type EmailAddress struct {
	ent.Schema
}

func (EmailAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.String("email"),
		field.String("lower_email"),
		field.Bool("is_activated").Optional(),
		field.Bool("is_primary").Default(false),
	}
}

func (EmailAddress) Edges() []ent.Edge {
	return nil
}

func (EmailAddress) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
		index.Fields("lower_email").Unique(),
		index.Fields("user_id"),
	}
}
