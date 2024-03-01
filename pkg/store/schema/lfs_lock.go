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

type LfsLock struct {
	ent.Schema
}

func (LfsLock) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.Int64("owner_id"),
		field.Text("path").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (LfsLock) Edges() []ent.Edge {
	return nil
}

func (LfsLock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
		index.Fields("repo_id"),
	}
}
