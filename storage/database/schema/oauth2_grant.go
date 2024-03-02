// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
