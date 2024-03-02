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

type PublicKey struct {
	ent.Schema
}

func (PublicKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.String("name").MaxLen(255),
		field.String("fingerprint").MaxLen(255),
		field.Text("content"),
		field.Int("mode").Default(2),
		field.Int("type").Default(1),
		field.Int("login_source_id").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Bool("verified").Default(false),
	}
}

func (PublicKey) Edges() []ent.Edge {
	return nil
}

func (PublicKey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fingerprint"),
		index.Fields("owner_id"),
	}
}
