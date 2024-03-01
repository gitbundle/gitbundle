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

type LanguageStat struct {
	ent.Schema
}

func (LanguageStat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.String("commit_id").Optional(),
		field.Bool("is_primary").Optional(),
		field.String("language").MaxLen(50),
		field.Int64("size"),
		field.Time("created_at").Default(time.Now),
	}
}

func (LanguageStat) Edges() []ent.Edge {
	return nil
}

func (LanguageStat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("repo_id", "language").Unique(),
	}
}
