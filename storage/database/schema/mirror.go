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

type Mirror struct {
	ent.Schema
}

func (Mirror) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("interval").Optional(),
		field.Bool("enable_prune").Default(true),
		field.Time("updated_at").Optional(),
		field.Time("next_updated_at").Optional(),
		field.Bool("lfs_enabled").Default(false),
		field.Text("lfs_endpoint").Optional(),
	}
}

func (Mirror) Edges() []ent.Edge {
	return nil
}

func (Mirror) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id"),
		index.Fields("updated_at"),
		index.Fields("next_updated_at"),
	}
}
