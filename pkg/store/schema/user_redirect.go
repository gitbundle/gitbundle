// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserRedirect struct {
	ent.Schema
}

func (UserRedirect) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lower_name").Optional(),
		field.String("redirect_user_id").Optional(),
	}
}

func (UserRedirect) Edges() []ent.Edge {
	return nil
}

func (UserRedirect) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lower_name").Unique(),
	}
}
