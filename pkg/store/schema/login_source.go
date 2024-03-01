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

type LoginSource struct {
	ent.Schema
}

func (LoginSource) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Optional(),
		field.String("name").Optional(),
		field.Bool("is_active").Default(false),
		field.Bool("is_sync_enabled").Default(false),
		field.Text("config").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (LoginSource) Edges() []ent.Edge {
	return nil
}

func (LoginSource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_sync_enabled"),
		index.Fields("updated_at"),
		index.Fields("name"),
		index.Fields("is_active"),
		index.Fields("created_at"),
	}
}
