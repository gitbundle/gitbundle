// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package schemas

// Reaction represents a reaction to an issue/pr/comment.
type Reaction struct {
	UserID   int64  `yaml:"user_id" json:"user_id"`
	UserName string `yaml:"user_name" json:"user_name"`
	Content  string `json:"content"`
}

// GetExternalName ExternalUserMigrated interface
func (r *Reaction) GetExternalName() string { return r.UserName }

// GetExternalID ExternalUserMigrated interface
func (r *Reaction) GetExternalID() int64 { return r.UserID }
