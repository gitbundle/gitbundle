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

type Notice struct {
	ent.Schema
}

func (Notice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Optional(),
		field.Text("description").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Notice) Edges() []ent.Edge{
	return nil
}

func (Notice) Indexes() []ent.Index{
	return []ent.Index{
		index.Fields("created_at"),
	}
}
