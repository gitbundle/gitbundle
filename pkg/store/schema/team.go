// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Team struct {
	ent.Schema
}

func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("org_id").Optional(),
		field.Int64("team_id").Optional(),
		field.Int64("repo_id").Optional(),
	}
}

func (Team) Edges() []ent.Edge{
	return nil
}

func (Team) Indexes() []ent.Index{
	return []ent.Index{
		index.Fields("team_id", "repo_id").Unique(),
		index.Fields("org_id"),
	}
}
