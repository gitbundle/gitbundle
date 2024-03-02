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

type Oauth2ApplicationCode struct {
	ent.Schema
}

func (Oauth2ApplicationCode) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("grant_id").Optional(),
		field.String("code").Optional().MaxLen(255),
		field.String("code_challenge").Optional().MaxLen(255),
		field.String("code_challenge_method").Optional().MaxLen(255),
		field.String("redirect_uri").Optional().MaxLen(255),
		field.Time("valid_until").Optional(),
	}
}

func (Oauth2ApplicationCode) Edges() []ent.Edge {
	return nil
}

func (Oauth2ApplicationCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique(),
		index.Fields("valid_until"),
	}
}
