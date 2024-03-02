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

type Oauth2Grant struct {
	ent.Schema
}

func (Oauth2Grant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int64("application_id").Optional(),
		field.Int("counter"),
		field.Text("scope").Optional(),
		field.Text("nonce").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Oauth2Grant) Edges() []ent.Edge {
	return nil
}

func (Oauth2Grant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "application_id").Unique(),
	}
}
