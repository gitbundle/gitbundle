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

type PkgBlobUpload struct {
	ent.Schema
}

func (PkgBlobUpload) Fields() []ent.Field {
	return []ent.Field{
		field.String("blob_upload_id").MaxLen(255), //id
		field.Int64("bytes_received"),
		field.Bytes("hash_state_bytes").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (PkgBlobUpload) Edges() []ent.Edge {
	return nil
}

func (PkgBlobUpload) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at"),
	}
}