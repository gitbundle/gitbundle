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

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.String("lower_name"),
		field.String("full_name").Optional(),
		field.String("email").Match(regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)),
		field.Bool("keep_email_private").Optional(),
		field.String("email_notifications_preference"),
		field.String("password").Sensitive(),
		field.String("password_hash_algo").Sensitive().Default("argon2"),
		field.Bool("must_change_password").Default(false),
		field.Int("login_type").Optional(),
		field.Int64("login_source").Default(0),
		field.String("login_name").Optional(),
		field.Int("type").Optional(),
		field.String("location").Optional(),
		field.String("website").Optional(),
		field.String("rands").Optional(),
		field.String("salt").Optional(),
		field.String("language").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Time("last_login_at").Optional(),
		field.Bool("last_repo_visibility").Optional(),
		field.Int("max_repo_creation").Default(-1),
		field.Bool("is_active").Optional(),
		field.Bool("is_admin").Optional(),
		field.Bool("is_restricted").Default(false),
		// field.Bool("allow_git_hook").Optional(),
		// field.Bool("allow_import_local").Optional(),
		field.Bool("allow_create_organization").Optional().Default(false), // only admin can set this perm
		field.Bool("prohibit_login").Default(false),
		field.String("avatar").MaxLen(2048),
		field.String("avatar_email"),
		// field.Bool("use_custom_avatar").Optional(),
		field.Int("num_followers").Optional(),
		field.Int("num_following").Default(0), // Fix: Is default 0 necessary?
		field.Int("num_stars").Optional(),
		field.Int("num_repos").Optional(),
		field.Int("num_teams").Optional(),
		field.Int("num_members").Optional(),
		field.Int("visibility").Default(0),
		field.Bool("repo_admin_change_team_access").Default(false), // not sure
		field.String("diff_view_style").Optional(),
		field.String("theme").Optional(),
		field.Bool("keep_activity_private").Default(false),
		field.Int("pid").Optional(), // parent namespace
		field.JSON("extensions", []string{}).Optional(),
		// field.JSON("resource_quotas").Optional(),
		field.Bool("enabled_packages").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repos", Repo.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("lower_name").Unique(),
		index.Fields("is_active"),
		index.Fields("last_login_at"),
		index.Fields("updated_at"),
		index.Fields("created_at"),
	}
}
