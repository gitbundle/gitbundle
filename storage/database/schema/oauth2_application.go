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

type Oauth2Application struct {
	ent.Schema
}

func (Oauth2Application) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("name").Optional().MaxLen(255),
		field.String("client_id").Optional().MaxLen(255),
		field.String("client_secret").Optional().MaxLen(255),
		field.Text("redirect_uris").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Oauth2Application) Edges() []ent.Edge {
	return nil
}

func (Oauth2Application) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("client_id").Unique(),
		index.Fields("updated_at"),
		index.Fields("created_at"),
		index.Fields("user_id"),
	}
}
