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

type GpgKey struct {
	ent.Schema
}

func (GpgKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.String("key_id").MaxLen(16),
		field.String("primary_key_id").Optional().MaxLen(16),
		field.Text("content"),
		field.Time("created_at").Default(time.Now),
		field.Time("expired_at").Optional(),
		// field.Time("added_at").Optional(),
		field.Text("emails").Optional(),
		field.Bool("verified").Default(false),
		field.Bool("can_sign").Optional(),
		field.Bool("can_encrypt_comms").Optional(),
		field.Bool("can_encrypt_storage").Optional(),
		field.Bool("can_certify").Optional(),
	}
}

func (GpgKey) Edges() []ent.Edge{
	return nil
}

func (GpgKey) Indexes() []ent.Index{
	return []ent.Index{
		index.Fields("owner_id"),
		index.Fields("key_id"),
	}
}
