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

type PkgVersion struct {
	ent.Schema
}

func (PkgVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("package_id"),
		field.Int64("creator_id"),
		field.String("version").MaxLen(255),
		field.String("lower_version").MaxLen(255),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_internal").Default(false),
		// field.JSON("metadata").Optional(),
		field.Int64("download_count").Default(0),
	}
}

func (PkgVersion) Edges() []ent.Edge {
	return nil
}

func (PkgVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("package_id", "lower_version").Unique(),
		index.Fields("created_at"),
		index.Fields("is_internal"),
	}
}
