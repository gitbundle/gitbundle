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
)

type ProtectedTag struct {
	ent.Schema
}

func (ProtectedTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.String("name_pattern").Optional().MaxLen(255),
		field.Text("allowlist_user_i_ds").Optional(),
		field.Text("allowlist_team_i_ds").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (ProtectedTag) Edges() []ent.Edge {
	return nil
}

func (ProtectedTag) Indexes() []ent.Index {
	return nil
}
