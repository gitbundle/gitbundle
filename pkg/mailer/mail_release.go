// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"context"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/markup"
	"github.com/gitbundle/modules/markup/markdown"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/translation"
	repo_model "github.com/gitbundle/server/pkg/repo"
	release_model "github.com/gitbundle/server/pkg/repo/release"
	user_model "github.com/gitbundle/server/pkg/user"
)

const (
	tplNewReleaseMail base.TplName = "release"
)

// MailNewRelease send new release notify to all all repo watchers.
func MailNewRelease(ctx context.Context, rel *release_model.Release) {
	if setting.MailService == nil {
		// No mail service configured
		return
	}

	watcherIDList, err := repo_model.GetRepoWatchersIDs(ctx, rel.RepoID)
	if err != nil {
		log.Error("GetRepoWatchersIDs(%d): %v", rel.RepoID, err)
		return
	}

	recipients, err := user_model.GetMaileableUsersByIDs(watcherIDList, false)
	if err != nil {
		log.Error("user_model.GetMaileableUsersByIDs: %v", err)
		return
	}

	langMap := make(map[string][]string)
	for _, user := range recipients {
		if user.ID != rel.PublisherID {
			langMap[user.Language] = append(langMap[user.Language], user.Email)
		}
	}

	for lang, tos := range langMap {
		mailNewRelease(ctx, lang, tos, rel)
	}
}

func mailNewRelease(ctx context.Context, lang string, tos []string, rel *release_model.Release) {
	locale := translation.NewLocale(lang)

	var err error
	rel.RenderedNote, err = markdown.RenderString(&markup.RenderContext{
		Ctx:       ctx,
		URLPrefix: rel.Repo.Link(),
		Metas:     rel.Repo.ComposeMetas(),
	}, rel.Note)
	if err != nil {
		log.Error("markdown.RenderString(%d): %v", rel.RepoID, err)
		return
	}

	subject := locale.Tr("mail.release.new.subject", rel.TagName, rel.Repo.FullName())
	mailMeta := map[string]interface{}{
		"Release":  rel,
		"Subject":  subject,
		"Language": locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var mailBody bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&mailBody, string(tplNewReleaseMail), mailMeta); err != nil {
		log.Error("ExecuteTemplate [%s]: %v", string(tplNewReleaseMail)+"/body", err)
		return
	}

	msgs := make([]*Message, 0, len(tos))
	publisherName := rel.Publisher.DisplayName()
	relURL := "<" + rel.HTMLURL() + ">"
	for _, to := range tos {
		msg := NewMessageFrom([]string{to}, publisherName, setting.MailService.FromEmail, subject, mailBody.String())
		msg.Info = subject
		msg.SetHeader("Message-ID", relURL)
		msgs = append(msgs, msg)
	}

	SendAsyncs(msgs)
}
