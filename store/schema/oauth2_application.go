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
