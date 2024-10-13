// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package foreignreference

import (
	"github.com/gitbundle/server/pkg/db"
)

// Type* are valid values for the Type field of ForeignReference
const (
	TypeIssue         = "issue"
	TypePullRequest   = "pull_request"
	TypeComment       = "comment"
	TypeReview        = "review"
	TypeReviewComment = "review_comment"
	TypeRelease       = "release"
)

// ForeignReference represents external references
type ForeignReference struct {
	// RepoID is the first column in all indices. now we only need 2 indices: (repo, local) and (repo, foreign, type)
	RepoID       int64  `xorm:"UNIQUE(repo_foreign_type) INDEX(repo_local)" `
	LocalIndex   int64  `xorm:"INDEX(repo_local)"` // the resource key inside GitBundle, it can be IssueIndex, or some model ID.
	ForeignIndex string `xorm:"INDEX UNIQUE(repo_foreign_type)"`
	Type         string `xorm:"VARCHAR(16) INDEX UNIQUE(repo_foreign_type)"`
}

func init() {
	db.RegisterModel(new(ForeignReference))
}
