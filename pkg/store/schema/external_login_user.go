// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
		field.String("external_id"),
		field.Int64("user_id"),
		field.Int64("login_source_id"),
		// field.JSON("raw_data", ).Optional(),
		field.String("provider").Optional(),
		field.String("email").Optional(),
		field.String("name").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("nick_name").Optional(),
		field.String("description").Optional(),
		field.Text("avatar_url").Optional(),
		field.String("location").Optional(),
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
