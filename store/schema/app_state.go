// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type AppState struct {
	ent.Schema
}

func (AppState) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("revision").Optional(),
		field.Text("content").Optional(),
	}
}

func (AppState) Edges() []ent.Edge {
	return nil
}

func (AppState) Indexes() []ent.Index {
	return nil
}
