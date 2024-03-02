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

type Collaboration struct {
	ent.Schema
}

func (Collaboration) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.Int64("user_id"),
		field.Int("mode"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Collaboration) Edges() []ent.Edge {
	return nil
}

func (Collaboration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "user_id").Unique(),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
