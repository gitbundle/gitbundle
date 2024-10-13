// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package mailer_test

import (
	"context"
	"path/filepath"
	"testing"
	_ "unsafe"

	"github.com/gitbundle/server/pkg/action"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/mailer"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	_ "github.com/gitbundle/server/pkg/access_token"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
	})
}

type mailCommentContext struct {
	context.Context
	Issue      *issues_model.Issue
	Doer       *user_model.User
	ActionType action.ActionType
	Content    string
	Comment    *issues_model.Comment
}

//go:linkname composeIssueCommentMessages github.com/gitbundle/server/pkg/mailer.composeIssueCommentMessages
func composeIssueCommentMessages(ctx *mailCommentContext, lang string, recipients []*user_model.User, fromMention bool, info string) ([]*mailer.Message, error)

//go:linkname generateAdditionalHeaders github.com/gitbundle/server/pkg/mailer.generateAdditionalHeaders
func generateAdditionalHeaders(ctx *mailCommentContext, reason string, recipient *user_model.User) map[string]string

//go:linkname createReference github.com/gitbundle/server/pkg/mailer.createReference
func createReference(issue *issues_model.Issue, comment *issues_model.Comment, actionType action.ActionType) string
