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

type LoginSource struct {
	ent.Schema
}

func (LoginSource) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Optional(),
		field.String("name").Optional().MaxLen(255),
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
