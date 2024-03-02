// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").MaxLen(16),
		field.Bytes("data").Optional(),
		field.Time("expired_at").Optional(),
	}
}

func (Session) Edges() []ent.Edge {
	return nil
}

func (Session) Indexes() []ent.Index {
	return nil
}
