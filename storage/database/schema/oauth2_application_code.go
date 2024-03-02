// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
