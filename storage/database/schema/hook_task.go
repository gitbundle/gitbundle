// Copyright 2024 GitBundle, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
