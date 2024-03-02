// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProtectedTag struct {
	ent.Schema
}

func (ProtectedTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.String("name_pattern").Optional().MaxLen(255),
		field.Text("allowlist_user_i_ds").Optional(),
		field.Text("allowlist_team_i_ds").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (ProtectedTag) Edges() []ent.Edge {
	return nil
}

func (ProtectedTag) Indexes() []ent.Index {
	return nil
}
