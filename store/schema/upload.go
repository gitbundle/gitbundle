// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Upload struct {
	ent.Schema
}

func (Upload) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.String("name").Optional().MaxLen(255),
	}
}

func (Upload) Edges() []ent.Edge {
	return nil
}

func (Upload) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
	}
}
