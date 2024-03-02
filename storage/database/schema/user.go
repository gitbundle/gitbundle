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
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")).MaxLen(255),
		field.String("lower_name").MaxLen(255),
		field.String("full_name").Optional().MaxLen(255),
		field.String("email").Match(regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)).MaxLen(255),
		field.Bool("keep_email_private").Optional(),
		field.String("email_notifications_preference").MaxLen(20),
		field.String("password").MaxLen(255).Sensitive(),
		field.String("password_hash_algo").MaxLen(255).Default("argon2").Sensitive(),
		field.Bool("must_change_password").Default(false),
		field.Int("login_type").Optional(),
		field.Int64("login_source").Default(0),
		field.String("login_name").Optional().MaxLen(255),
		field.Int("type").Optional(),
		field.String("location").Optional().MaxLen(255),
		field.String("website").Optional().MaxLen(255),
		field.String("rands").Optional().MaxLen(32).Sensitive(),
		field.String("salt").Optional().MaxLen(32).Sensitive(),
		field.String("language").Optional().MaxLen(5),
		field.String("description").Optional().MaxLen(255),
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
		field.String("avatar_email").MaxLen(255),
		// field.Bool("use_custom_avatar").Optional(),
		field.Int("num_followers").Optional(),
		field.Int("num_following").Default(0), // Fix: Is default 0 necessary?
		field.Int("num_stars").Optional(),
		field.Int("num_repos").Optional(),
		field.Int("num_teams").Optional(),
		field.Int("num_members").Optional(),
		field.Int("visibility").Default(0),
		field.Bool("repo_admin_change_team_access").Default(false), // not sure
		field.String("diff_view_style").Optional().MaxLen(255),
		field.String("theme").Optional().MaxLen(255),
		field.Bool("keep_activity_private").Default(false),
		field.Int("pid").Optional().Default(0), // parent namespace
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
