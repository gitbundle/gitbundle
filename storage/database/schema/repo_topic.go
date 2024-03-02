// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RepoTopic struct {
	ent.Schema
}

func (RepoTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id"),
		field.Int64("topic_id"),
	}
}

func (RepoTopic) Edges() []ent.Edge {
	return nil
}

func (RepoTopic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "topic_id"),
	}
}
