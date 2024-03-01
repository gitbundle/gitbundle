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

type RepoTransfer struct {
	ent.Schema
}

func (RepoTransfer) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("doer_id").Optional(),
		field.Int64("recipient_id").Optional(),
		field.Int64("repo_id").Optional(),
		field.Text("team_i_ds").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (RepoTransfer) Edges() []ent.Edge {
	return nil
}

func (RepoTransfer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
