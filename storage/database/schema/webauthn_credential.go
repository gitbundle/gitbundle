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

type WebauthnCredential struct {
	ent.Schema
}

func (WebauthnCredential) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
		field.String("lower_name").Optional().MaxLen(255),
		field.Int64("user_id").Optional(),
		field.Bytes("public_key").Optional(),
		field.String("attestation_type").Optional().MaxLen(255),
		field.Bytes("aaguid").Optional(),
		field.Int64("sign_count").Optional(),
		field.Bool("clone_warning").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Bytes("credential_id").Optional(),
	}
}

func (WebauthnCredential) Edges() []ent.Edge {
	return nil
}

func (WebauthnCredential) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("lower_name", "user_id").Unique(),
	}
}
