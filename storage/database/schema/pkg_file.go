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

type PkgFile struct {
	ent.Schema
}

func (PkgFile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("version_id"),
		field.Int64("blob_id"),
		field.String("name").MaxLen(255),
		field.String("lower_name").MaxLen(255),
		field.String("composite_key").Optional().MaxLen(255),
		field.Bool("is_lead").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

func (PkgFile) Edges() []ent.Edge {
	return nil
}

func (PkgFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("blob_id"),
		index.Fields("version_id", "lower_name", "composite_key").Unique(),
	}
}
