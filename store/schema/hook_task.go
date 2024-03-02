// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type HookTask struct {
	ent.Schema
}

func (HookTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.Int64("hook_id").Optional(),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Text("payload_content").Optional(),
		field.String("event_type").Optional().MaxLen(255),
		field.Bool("is_delivered").Optional(),
		field.Int64("delivered").Optional(),
		field.Bool("is_succeed").Optional(),
		field.Text("request_content").Optional(),
		field.Text("response_content").Optional(),
	}
}

func (HookTask) Edges() []ent.Edge {
	return nil
}

func (HookTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id"),
	}
}
