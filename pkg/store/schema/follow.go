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

type Follow struct {
	ent.Schema
}

func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int64("follow_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Follow) Edges() []ent.Edge {
	return nil
}

func (Follow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "follow_id").Unique(),
		index.Fields("created_at"),
	}
}
