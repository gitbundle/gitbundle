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

type CommitStatus struct {
	ent.Schema
}

func (CommitStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("index").Optional(),
		field.Int64("repo_id").Optional(),
		field.String("state").MaxLen(7),
		field.String("sha").MaxLen(64),
		field.Text("target_url").Optional(),
		field.Text("description").Optional(),
		field.String("context_hash").MaxLen(40),
		field.Text("context").Optional(),
		field.Int64("creator_id").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (CommitStatus) Edges() []ent.Edge {
	return nil
}

func (CommitStatus) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("context_hash"),
		index.Fields("index", "repo_id", "sha").Unique(),
		index.Fields("updated_at"),
	}
}
