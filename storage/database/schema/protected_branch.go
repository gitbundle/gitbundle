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

type ProtectedBranch struct {
	ent.Schema
}

func (ProtectedBranch) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Optional(),
		field.String("branch_name").Optional().MaxLen(255),
		field.Bool("can_push").Default(false),
		field.Bool("enable_whitelist"),
		field.Text("whitelist_user_i_ds").Optional(),
		field.Text("whitelist_team_i_ds").Optional(),
		field.Bool("enable_merge_whitelist").Default(false),
		field.Bool("whitelist_deploy_keys").Default(false),
		field.Text("merge_whitelist_user_i_ds").Optional(),
		field.Text("merge_whitelist_team_i_ds").Optional(),
		field.Bool("enable_status_check").Default(false),
		field.Text("status_check_contexts").Optional(),
		field.Bool("enable_approvals_whitelist").Default(false),
		field.Text("approvals_whitelist_user_i_ds").Optional(),
		field.Text("approvals_whitelist_team_i_ds").Optional(),
		field.Int64("required_approvals").Default(0),
		field.Bool("block_on_rejected_reviews").Default(false),
		field.Bool("block_on_official_review_requests").Default(false),
		field.Bool("block_on_outdated_branch").Default(false),
		field.Bool("dismiss_stale_approvals").Default(false),
		field.Bool("require_signed_commits").Default(false),
		field.Text("protected_file_patterns").Optional(),
		field.Text("unprotected_file_patterns").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (ProtectedBranch) Edges() []ent.Edge {
	return nil
}

func (ProtectedBranch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("repo_id", "branch_name").Unique(),
	}
}
