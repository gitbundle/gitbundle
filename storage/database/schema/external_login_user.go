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

type ExternalLoginUser struct {
	ent.Schema
}

func (ExternalLoginUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").MaxLen(255),
		field.Int64("user_id"),
		field.Int64("login_source_id"),
		// field.JSON("raw_data", ).Optional(),
		field.String("provider").Optional().MaxLen(25),
		field.String("email").Optional().MaxLen(255),
		field.String("name").Optional().MaxLen(255),
		field.String("first_name").Optional().MaxLen(255),
		field.String("last_name").Optional().MaxLen(255),
		field.String("nick_name").Optional().MaxLen(255),
		field.String("description").Optional().MaxLen(255),
		field.Text("avatar_url").Optional(),
		field.String("location").Optional().MaxLen(255),
		field.Text("access_token").Optional(),
		field.Text("access_token_secret").Optional(),
		field.Text("refresh_token").Optional(),
		field.Time("expired_at").Optional(),
	}
}

func (ExternalLoginUser) Edges() []ent.Edge {
	return nil
}

func (ExternalLoginUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_id", "login_source_id"),
		index.Fields("user_id"),
		index.Fields("provider"),
	}
}
