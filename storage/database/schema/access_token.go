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

type AccessToken struct {
	ent.Schema
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("name").Optional().MaxLen(255),
		field.String("token_hash").Optional().MaxLen(255).Sensitive(),
		field.String("token_salt").Optional().MaxLen(255).Sensitive(),
		field.String("token_last_eight").Optional().MaxLen(255).Sensitive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return nil
}

func (AccessToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("token_hash").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
