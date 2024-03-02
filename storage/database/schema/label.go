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

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("org_id").Optional(),
		field.String("name").Optional().MaxLen(255),
		field.String("description").Optional().MaxLen(255),
		field.String("color").Optional().MaxLen(7),
		field.Int64("num_issues").Optional(),
		field.Int64("num_closed_issues").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Label) Edges() []ent.Edge {
	return nil
}

func (Label) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("org_id"),
		index.Fields("updated_at"),
		index.Fields("repo_id"),
	}
}
