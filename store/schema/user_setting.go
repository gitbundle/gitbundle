// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserSetting struct {
	ent.Schema
}

func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.String("setting_key").Optional().MaxLen(255),
		field.Text("setting_value").Optional(),
	}
}

func (UserSetting) Edges() []ent.Edge {
	return nil
}

func (UserSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "setting_key").Unique(),
	}
}
