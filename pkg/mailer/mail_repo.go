// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"fmt"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/translation"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/organization"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"
)

// SendRepoTransferNotifyMail triggers a notification e-mail when a pending repository transfer was created
func SendRepoTransferNotifyMail(doer, newOwner *user_model.User, repo *repo_model.Repository) error {
	if setting.MailService == nil {
		// No mail service configured
		return nil
	}

	if newOwner.IsOrganization() {
		users, err := organization.GetUsersWhoCanCreateOrgRepo(db.DefaultContext, newOwner.ID)
		if err != nil {
			return err
		}

		langMap := make(map[string][]string)
		for _, user := range users {
			if !user.IsActive {
				// don't send emails to inactive users
				continue
			}
			langMap[user.Language] = append(langMap[user.Language], user.Email)
		}

		for lang, tos := range langMap {
			if err := sendRepoTransferNotifyMailPerLang(lang, newOwner, doer, tos, repo); err != nil {
				return err
			}
		}

		return nil
	}

	return sendRepoTransferNotifyMailPerLang(newOwner.Language, newOwner, doer, []string{newOwner.Email}, repo)
}

// sendRepoTransferNotifyMail triggers a notification e-mail when a pending repository transfer was created for each language
func sendRepoTransferNotifyMailPerLang(lang string, newOwner, doer *user_model.User, emails []string, repo *repo_model.Repository) error {
	var (
		locale  = translation.NewLocale(lang)
		content bytes.Buffer
	)

	destination := locale.Tr("mail.repo.transfer.to_you")
	subject := locale.Tr("mail.repo.transfer.subject_to_you", doer.DisplayName(), repo.FullName())
	if newOwner.IsOrganization() {
		destination = newOwner.DisplayName()
		subject = locale.Tr("mail.repo.transfer.subject_to", doer.DisplayName(), repo.FullName(), destination)
	}

	data := map[string]interface{}{
		"Doer":        doer,
		"User":        repo.Owner,
		"Repo":        repo.FullName(),
		"Link":        repo.HTMLURL(),
		"Subject":     subject,
		"Language":    locale.Language(),
		"Destination": destination,
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	if err := bodyTemplates.ExecuteTemplate(&content, string(mailRepoTransferNotify), data); err != nil {
		return err
	}

	msg := NewMessage(emails, subject, content.String())
	msg.Info = fmt.Sprintf("UID: %d, repository pending transfer notification", newOwner.ID)

	SendAsync(msg)
	return nil
}
