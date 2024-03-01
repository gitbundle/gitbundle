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

type LfsMetaObject struct {
	ent.Schema
}

func (LfsMetaObject) Fields() []ent.Field {
	return []ent.Field{
		field.String("oid"),
		field.Int64("size"),
		field.Int64("repo_id"),
		field.Time("created_at").Default(time.Now),
	}
}

func (LfsMetaObject) Edges() []ent.Edge {
	return nil
}

func (LfsMetaObject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("oid", "repo_id").Unique(),
	}
}
