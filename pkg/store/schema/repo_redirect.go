// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RepoRedirect struct {
	ent.Schema
}

func (RepoRedirect) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id").Optional(),
		field.String("lower_name"),
		field.Int64("redirect_repo_id").Optional(),
	}
}

func (RepoRedirect) Edges() []ent.Edge {
	return nil
}

func (RepoRedirect) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "lower_name").Unique(),
	}
}
