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

type Topic struct {
	ent.Schema
}

func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(50),
		field.Int("repo_count").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Topic) Edges() []ent.Edge {
	return nil
}

func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("updated_at"),
		index.Fields("created_at"),
	}
}
