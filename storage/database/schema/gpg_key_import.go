// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type GpgKeyImport struct {
	ent.Schema
}

func (GpgKeyImport) Fields() []ent.Field {
	return []ent.Field{
		field.String("key_id").MaxLen(16),
		field.Text("content"),
	}
}

func (GpgKeyImport) Edges() []ent.Edge {
	return nil
}

func (GpgKeyImport) Indexes() []ent.Index {
	return nil
}
