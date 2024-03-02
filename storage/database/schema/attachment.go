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
	"github.com/google/uuid"
)

type Attachment struct {
	ent.Schema
}

func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Int64("repo_id").Optional(),
		field.Int64("uploader_id").Optional(),
		// field.Int64("issue_id").Optional(), // will implemented in issue service
		// field.Int64("release_id").Optional(), // will implemented in release service
		// field.Int64("comment_id").Optional(), // will implemented in issue service
		field.String("name").Optional().MaxLen(255),
		field.Int64("download_count").Default(0),
		field.Int64("size").Default(0),
		field.Time("created_at").Default(time.Now),
	}
}

func (Attachment) Edges() []ent.Edge {
	return nil
}

func (Attachment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
		index.Fields("repo_id"),
		index.Fields("uploader_id"),
	}
}
