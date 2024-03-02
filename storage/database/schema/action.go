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

type Action struct {
	ent.Schema
}

func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional(),
		field.Int("op_type").Optional(),
		field.Int64("act_user_id").Optional(),
		field.Int64("repo_id").Optional(),
		// field.Int64("comment_id").Optional(),
		field.Bool("is_deleted").Default(false),
		field.String("ref_name").Optional().MaxLen(255),
		field.Bool("is_private").Default(false),
		field.Text("content").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Action) Edges() []ent.Edge {
	return nil
}

func (Action) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "user_id", "is_deleted"),
		index.Fields("act_user_id", "repo_id", "created_at", "user_id", "is_deleted"),
		// index.Fields("comment_id"),
	}
}
