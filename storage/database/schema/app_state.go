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
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type AppState struct {
	ent.Schema
}

func (AppState) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("revision").Optional(),
		field.Text("content").Optional(),
	}
}

func (AppState) Edges() []ent.Edge {
	return nil
}

func (AppState) Indexes() []ent.Index {
	return nil
}
