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

type Notification struct {
	ent.Schema
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("repo_id"),
		field.Int16("status"),
		field.Int16("source"),
		// field.Int64("issue_id"), // Should sync from issue service?
		field.String("commit_id").Optional().MaxLen(255),
		// field.Int64("comment_id").Optional(),//Should sync from comment service?
		field.Int("updated_by"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Notification) Edges() []ent.Edge {
	return nil
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source"),
		index.Fields("commit_id"),
		index.Fields("repo_id"),
		// index.Fields("issue_id"),
		index.Fields("updated_by"),
		index.Fields("updated_at"),
		index.Fields("created_at"),
		index.Fields("user_id"),
		index.Fields("status"),
	}
}
