// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CommitStatusIndex struct {
	ent.Schema
}

func (CommitStatusIndex) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.String("sha").Optional().MaxLen(255),
		field.Int64("max_index").Optional(),
	}
}

func (CommitStatusIndex) Edges() []ent.Edge {
	return nil
}

func (CommitStatusIndex) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "sha").Unique(),
		index.Fields("max_index"),
	}
}