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
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("doer_id").Optional(),
		field.Int64("owner_id").Optional(),
		field.Int64("repo_id").Optional(),
		field.Int("type").Optional(),
		field.Int("status").Optional(),
		field.Time("started_at").Optional(),
		field.Time("ended_at").Optional(),
		field.Text("payload_content").Optional(),
		field.Text("message").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Task) Edges() []ent.Edge {
	return nil
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("doer_id"),
		index.Fields("owner_id"),
		index.Fields("repo_id"),
		index.Fields("status"),
	}
}
