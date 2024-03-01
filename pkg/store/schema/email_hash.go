// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type EmailHash struct {
	ent.Schema
}

func (EmailHash) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").MaxLen(32),
		field.String("email"),
	}
}

func (EmailHash) Edges() []ent.Edge {
	return nil
}

func (EmailHash) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
	}
}
