// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrgUser struct {
	ent.Schema
}

func (OrgUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int64("org_id").Optional(),
		field.Bool("is_public").Optional(),
	}
}

func (OrgUser) Edges() []ent.Edge {
	return nil
}

func (OrgUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "org_id").Unique(),
		index.Fields("is_public"),
	}
}
