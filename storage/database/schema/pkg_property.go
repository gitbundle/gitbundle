// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PkgProperty struct {
	ent.Schema
}

func (PkgProperty) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("ref_type"),
		field.Int64("ref_id"),
		field.String("name").MaxLen(255),
		field.Text("value"),
	}
}

func (PkgProperty) Edges() []ent.Edge {
	return nil
}

func (PkgProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ref_type"),
		index.Fields("ref_id"),
		index.Fields("name"),
	}
}
