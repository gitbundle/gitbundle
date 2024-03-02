// Copyright 2024 GitBundle, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
