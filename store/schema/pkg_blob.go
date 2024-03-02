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
