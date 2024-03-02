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
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Team struct {
	ent.Schema
}

func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("org_id").Optional(),
		field.String("lower_name").Optional().MaxLen(255),
		field.String("name").Optional().MaxLen(255),
		field.String("description").Optional().MaxLen(255),
		field.Int("authorize").Optional(),
		field.Int("num_repos").Optional(),
		field.Int("num_members").Optional(),
		field.Bool("include_all_repos").Default(false),
		field.Bool("can_create_org_repo").Default(false),
	}
}

func (Team) Edges() []ent.Edge {
	return nil
}

func (Team) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("org_id"),
	}
}
