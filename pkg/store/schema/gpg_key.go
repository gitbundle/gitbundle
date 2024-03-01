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

type GpgKey struct {
	ent.Schema
}

func (GpgKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.String("key_id").MaxLen(16),
		field.String("primary_key_id").Optional().MaxLen(16),
		field.Text("content"),
		field.Time("created_at").Default(time.Now),
		field.Time("expired_at").Optional(),
		// field.Time("added_at").Optional(),
		field.Text("emails").Optional(),
		field.Bool("verified").Default(false),
		field.Bool("can_sign").Optional(),
		field.Bool("can_encrypt_comms").Optional(),
		field.Bool("can_encrypt_storage").Optional(),
		field.Bool("can_certify").Optional(),
	}
}

func (GpgKey) Edges() []ent.Edge{
	return nil
}

func (GpgKey) Indexes() []ent.Index{
	return []ent.Index{
		index.Fields("owner_id"),
		index.Fields("key_id"),
	}
}
