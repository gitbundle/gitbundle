// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Repo struct {
	ent.Schema
}

func (Repo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id"),
		field.String("owner_name").MaxLen(255),
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")).MaxLen(255),
		field.String("lower_name").MaxLen(255),
		field.Text("description").Optional(),
		field.String("website").Optional().MaxLen(2048),
		field.Int("original_service_type").Optional(),
		field.String("original_url").Optional().MaxLen(2048),
		field.String("default_branch").Optional().MaxLen(255),
		field.Int("num_watches").Optional(),
		field.Int("num_stars").Optional(),
		field.Int("num_forks").Optional(),
		field.Int("num_issues").Optional(),        // will sync from issue service
		field.Int("num_closed_issues").Optional(), // will sync from issue service
		field.Int("num_prs").Optional(),           // will sync from pr service
		field.Int("num_closed_prs").Optional(),    // will sync from pr service
		field.Int("num_milestones"),               // will sync from milestones service
		field.Int("num_closed_milestones"),        // will sync from milestones service
		field.Int("num_projects"),                 // will sync from project service
		field.Int("num_closed_projects"),          // will sync from project service
		field.Bool("is_private").Optional(),
		field.Bool("is_empty").Optional(),
		field.Bool("is_archived").Optional(),
		field.Bool("is_mirror").Optional(),
		field.Int("status").Default(0),
		field.Bool("is_forked").Default(false),
		field.Int64("fork_id").Optional(),
		field.Bool("is_template").Default(false),
		field.Int64("template_id").Optional(),
		field.Int64("size").Default(0),
		field.Bool("is_fsck_enabled").Default(true),
		field.Bool("close_issues_via_commit_in_any_branch").Default(false),
		field.JSON("topics", []string{}).Optional(),
		field.Int("trust_model").Optional(),
		field.String("avatar").Optional().MaxLen(64),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Repo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("repos").Unique(),
	}
}

func (Repo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "owner_name").Unique(),
		index.Fields("name"),
		index.Fields("lower_name"),
		index.Fields("original_service_type"),
		index.Fields("is_private"),
		index.Fields("is_empty"),
		index.Fields("is_archived"),
		index.Fields("is_mirror"),
		index.Fields("is_forked"),
		index.Fields("fork_id"),
		index.Fields("is_template"),
		index.Fields("template_id"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
