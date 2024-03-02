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

type PkgVersion struct {
	ent.Schema
}

func (PkgVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("package_id"),
		field.Int64("creator_id"),
		field.String("version").MaxLen(255),
		field.String("lower_version").MaxLen(255),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_internal").Default(false),
		// field.JSON("metadata").Optional(),
		field.Int64("download_count").Default(0),
	}
}

func (PkgVersion) Edges() []ent.Edge {
	return nil
}

func (PkgVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("package_id", "lower_version").Unique(),
		index.Fields("created_at"),
		index.Fields("is_internal"),
	}
}
