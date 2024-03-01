// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Attachment struct {
	ent.Schema
}

func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Optional().MaxLen(80),
		field.Int64("repo_id").Optional(),
		field.Int64("uploader_id").Optional(),
		// field.Int64("issue_id").Optional(), // will implemented in issue service
		// field.Int64("release_id").Optional(), // will implemented in release service
		// field.Int64("comment_id").Optional(), // will implemented in issue service
		field.String("name").Optional(),
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
