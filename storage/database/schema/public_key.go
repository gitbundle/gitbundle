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

type PublicKey struct {
	ent.Schema
}

func (PublicKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.String("name").MaxLen(255),
		field.String("fingerprint").MaxLen(255),
		field.Text("content"),
		field.Int("mode").Default(2),
		field.Int("type").Default(1),
		field.Int("login_source_id").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Bool("verified").Default(false),
	}
}

func (PublicKey) Edges() []ent.Edge {
	return nil
}

func (PublicKey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fingerprint"),
		index.Fields("owner_id"),
	}
}
