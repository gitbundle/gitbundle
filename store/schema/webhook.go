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

type Webhook struct {
	ent.Schema
}

func (Webhook) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("org_id").Optional(),
		field.Bool("is_system_webhook").Optional(),
		field.Text("url").Optional(),
		field.String("http_method").Optional().MaxLen(255),
		field.Int("content_type").Optional(),
		field.Text("secret").Optional(),
		field.Text("events").Optional(),
		field.Bool("is_active").Optional(),
		field.String("type").Optional().MaxLen(16),
		field.Text("meta").Optional(),
		field.Int("last_status").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Webhook) Edges() []ent.Edge {
	return nil
}

func (Webhook) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("org_id"),
		index.Fields("repo_id"),
		index.Fields("is_active"),
	}
}
