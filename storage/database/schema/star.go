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

type Star struct {
	ent.Schema
}

func (Star) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int64("repo_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Star) Edges() []ent.Edge {
	return nil
}

func (Star) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "repo_id").Unique(),
		index.Fields("created_at"),
	}
}
