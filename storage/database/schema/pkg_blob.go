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

type PkgBlob struct {
	ent.Schema
}

func (PkgBlob) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("size"),
		field.String("hash_md5").MaxLen(32),
		field.String("hash_sha1").MaxLen(40),
		field.String("hash_sha256").MaxLen(64),
		field.String("hash_sha512").MaxLen(128),
		field.Time("created_at").Default(time.Now),
	}
}

func (PkgBlob) Edges() []ent.Edge {
	return nil
}

func (PkgBlob) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash_md5").Unique(),
		index.Fields("hash_sha256").Unique(),
		index.Fields("hash_sha512").Unique(),
		index.Fields("hash_sha1").Unique(),
		index.Fields("created_at"),
	}
}
