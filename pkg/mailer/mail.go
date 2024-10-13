// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"mime"
	"regexp"
	"strconv"
	"strings"
	texttmpl "text/template"
	"time"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/emoji"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/markup"
	"github.com/gitbundle/modules/markup/markdown"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/modules/translation"
	"github.com/gitbundle/server/pkg/action"
	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	repo_model "github.com/gitbundle/server/pkg/repo"
	user_model "github.com/gitbundle/server/pkg/user"

	"gopkg.in/gomail.v2"
)

const (
	mailAuthActivate       base.TplName = "auth/activate"
	mailAuthActivateEmail  base.TplName = "auth/activate_email"
	mailAuthResetPassword  base.TplName = "auth/reset_passwd"
	mailAuthRegisterNotify base.TplName = "auth/register_notify"
	mailNotifyCollaborator base.TplName = "notify/collaborator"
	mailRepoTransferNotify base.TplName = "notify/repo_transfer"

	// There's no actual limit for subject in RFC 5322
	mailMaxSubjectRunes = 256
)

var (
	bodyTemplates       *template.Template
	subjectTemplates    *texttmpl.Template
	subjectRemoveSpaces = regexp.MustCompile(`[\s]+`)
)

// InitMailRender initializes the mail renderer
func InitMailRender(subjectTpl *texttmpl.Template, bodyTpl *template.Template) {
	subjectTemplates = subjectTpl
	bodyTemplates = bodyTpl
}

// SendTestMail sends a test mail
func SendTestMail(email string) error {
	if setting.MailService == nil {
		// No mail service configured
		return nil
	}
	return gomail.Send(Sender, NewMessage([]string{email}, "GitBundle Test Email!", "GitBundle Test Email!").ToMessage())
}

// sendUserMail sends a mail to the user
func sendUserMail(language string, u *user_model.User, tpl base.TplName, code, subject, info string) {
	locale := translation.NewLocale(language)
	data := map[string]interface{}{
		"DisplayName":       u.DisplayName(),
		"ActiveCodeLives":   timeutil.MinutesToFriendly(setting.Service.ActiveCodeLives, language),
		"ResetPwdCodeLives": timeutil.MinutesToFriendly(setting.Service.ResetPwdCodeLives, language),
		"Code":              code,
		"Language":          locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var content bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&content, string(tpl), data); err != nil {
		log.Error("Template: %v", err)
		return
	}

	msg := NewMessage([]string{u.Email}, subject, content.String())
	msg.Info = fmt.Sprintf("UID: %d, %s", u.ID, info)

	SendAsync(msg)
}

// SendActivateAccountMail sends an activation mail to the user (new user registration)
func SendActivateAccountMail(locale translation.Locale, u *user_model.User) {
	if setting.MailService == nil {
		// No mail service configured
		return
	}
	sendUserMail(locale.Language(), u, mailAuthActivate, u.GenerateEmailActivateCode(u.Email), locale.Tr("mail.activate_account"), "activate account")
}

// SendResetPasswordMail sends a password reset mail to the user
func SendResetPasswordMail(u *user_model.User) {
	if setting.MailService == nil {
		// No mail service configured
		return
	}
	locale := translation.NewLocale(u.Language)
	sendUserMail(u.Language, u, mailAuthResetPassword, u.GenerateEmailActivateCode(u.Email), locale.Tr("mail.reset_password"), "recover account")
}

// SendActivateEmailMail sends confirmation email to confirm new email address
func SendActivateEmailMail(u *user_model.User, email *user_model.EmailAddress) {
	if setting.MailService == nil {
		// No mail service configured
		return
	}
	locale := translation.NewLocale(u.Language)
	data := map[string]interface{}{
		"DisplayName":     u.DisplayName(),
		"ActiveCodeLives": timeutil.MinutesToFriendly(setting.Service.ActiveCodeLives, locale.Language()),
		"Code":            u.GenerateEmailActivateCode(email.Email),
		"Email":           email.Email,
		"Language":        locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var content bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&content, string(mailAuthActivateEmail), data); err != nil {
		log.Error("Template: %v", err)
		return
	}

	msg := NewMessage([]string{email.Email}, locale.Tr("mail.activate_email"), content.String())
	msg.Info = fmt.Sprintf("UID: %d, activate email", u.ID)

	SendAsync(msg)
}

// SendRegisterNotifyMail triggers a notify e-mail by admin created a account.
func SendRegisterNotifyMail(u *user_model.User) {
	if setting.MailService == nil || !u.IsActive {
		// No mail service configured OR user is inactive
		return
	}
	locale := translation.NewLocale(u.Language)

	data := map[string]interface{}{
		"DisplayName": u.DisplayName(),
		"Username":    u.Name,
		"Language":    locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var content bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&content, string(mailAuthRegisterNotify), data); err != nil {
		log.Error("Template: %v", err)
		return
	}

	msg := NewMessage([]string{u.Email}, locale.Tr("mail.register_notify"), content.String())
	msg.Info = fmt.Sprintf("UID: %d, registration notify", u.ID)

	SendAsync(msg)
}

// SendCollaboratorMail sends mail notification to new collaborator.
func SendCollaboratorMail(u, doer *user_model.User, repo *repo_model.Repository) {
	if setting.MailService == nil || !u.IsActive {
		// No mail service configured OR the user is inactive
		return
	}
	locale := translation.NewLocale(u.Language)
	repoName := repo.FullName()

	subject := locale.Tr("mail.repo.collaborator.added.subject", doer.DisplayName(), repoName)
	data := map[string]interface{}{
		"Subject":  subject,
		"RepoName": repoName,
		"Link":     repo.HTMLURL(),
		"Language": locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var content bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&content, string(mailNotifyCollaborator), data); err != nil {
		log.Error("Template: %v", err)
		return
	}

	msg := NewMessage([]string{u.Email}, subject, content.String())
	msg.Info = fmt.Sprintf("UID: %d, add collaborator", u.ID)

	SendAsync(msg)
}

func composeIssueCommentMessages(ctx *mailCommentContext, lang string, recipients []*user_model.User, fromMention bool, info string) ([]*Message, error) {
	var (
		subject string
		link    string
		prefix  string
		// Fall back subject for bad templates, make sure subject is never empty
		fallback       string
		reviewComments []*issues_model.Comment
	)

	commentType := issues_model.CommentTypeComment
	if ctx.Comment != nil {
		commentType = ctx.Comment.Type
		link = ctx.Issue.HTMLURL() + "#" + ctx.Comment.HashTag()
	} else {
		link = ctx.Issue.HTMLURL()
	}

	reviewType := issues_model.ReviewTypeComment
	if ctx.Comment != nil && ctx.Comment.Review != nil {
		reviewType = ctx.Comment.Review.Type
	}

	// This is the body of the new issue or comment, not the mail body
	body, err := markdown.RenderString(&markup.RenderContext{
		Ctx:       ctx,
		URLPrefix: ctx.Issue.Repo.HTMLURL(),
		Metas:     ctx.Issue.Repo.ComposeMetas(),
	}, ctx.Content)
	if err != nil {
		return nil, err
	}

	actType, actName, tplName := actionToTemplate(ctx.Issue, ctx.ActionType, commentType, reviewType)

	if actName != "new" {
		prefix = "Re: "
	}
	fallback = prefix + fallbackMailSubject(ctx.Issue)

	if ctx.Comment != nil && ctx.Comment.Review != nil {
		reviewComments = make([]*issues_model.Comment, 0, 10)
		for _, lines := range ctx.Comment.Review.CodeComments {
			for _, comments := range lines {
				reviewComments = append(reviewComments, comments...)
			}
		}
	}
	locale := translation.NewLocale(lang)

	mailMeta := map[string]interface{}{
		"FallbackSubject": fallback,
		"Body":            body,
		"Link":            link,
		"Issue":           ctx.Issue,
		"Comment":         ctx.Comment,
		"IsPull":          ctx.Issue.IsPull,
		"User":            ctx.Issue.Repo.MustOwner(),
		"Repo":            ctx.Issue.Repo.FullName(),
		"Doer":            ctx.Doer,
		"IsMention":       fromMention,
		"SubjectPrefix":   prefix,
		"ActionType":      actType,
		"ActionName":      actName,
		"ReviewComments":  reviewComments,
		"Language":        locale.Language(),
		// helper
		"i18n":      locale,
		"Str2html":  Str2html,
		"DotEscape": DotEscape,
	}

	var mailSubject bytes.Buffer
	if err := subjectTemplates.ExecuteTemplate(&mailSubject, string(tplName), mailMeta); err == nil {
		subject = sanitizeSubject(mailSubject.String())
		if subject == "" {
			subject = fallback
		}
	} else {
		log.Error("ExecuteTemplate [%s]: %v", tplName+"/subject", err)
	}

	subject = emoji.ReplaceAliases(subject)

	mailMeta["Subject"] = subject

	var mailBody bytes.Buffer

	if err := bodyTemplates.ExecuteTemplate(&mailBody, string(tplName), mailMeta); err != nil {
		log.Error("ExecuteTemplate [%s]: %v", string(tplName)+"/body", err)
	}

	// Make sure to compose independent messages to avoid leaking user emails
	msgID := createReference(ctx.Issue, ctx.Comment, ctx.ActionType)
	reference := createReference(ctx.Issue, nil, action.ActionType(0))

	msgs := make([]*Message, 0, len(recipients))
	for _, recipient := range recipients {
		msg := NewMessageFrom([]string{recipient.Email}, ctx.Doer.DisplayName(), setting.MailService.FromEmail, subject, mailBody.String())
		msg.Info = fmt.Sprintf("Subject: %s, %s", subject, info)

		msg.SetHeader("Message-ID", "<"+msgID+">")
		msg.SetHeader("In-Reply-To", "<"+reference+">")
		msg.SetHeader("References", "<"+reference+">")

		for key, value := range generateAdditionalHeaders(ctx, actType, recipient) {
			msg.SetHeader(key, value)
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func createReference(issue *issues_model.Issue, comment *issues_model.Comment, actionType action.ActionType) string {
	var path string
	if issue.IsPull {
		path = "pulls"
	} else {
		path = "issues"
	}

	var extra string
	if comment != nil {
		extra = fmt.Sprintf("/comment/%d", comment.ID)
	} else {
		switch actionType {
		case action.ActionCloseIssue, action.ActionClosePullRequest:
			extra = fmt.Sprintf("/close/%d", time.Now().UnixNano()/1e6)
		case action.ActionReopenIssue, action.ActionReopenPullRequest:
			extra = fmt.Sprintf("/reopen/%d", time.Now().UnixNano()/1e6)
		case action.ActionMergePullRequest:
			extra = fmt.Sprintf("/merge/%d", time.Now().UnixNano()/1e6)
		case action.ActionPullRequestReadyForReview:
			extra = fmt.Sprintf("/ready/%d", time.Now().UnixNano()/1e6)
		}
	}

	return fmt.Sprintf("%s/%s/%d%s@%s", issue.Repo.FullName(), path, issue.Index, extra, setting.Domain)
}

func generateAdditionalHeaders(ctx *mailCommentContext, reason string, recipient *user_model.User) map[string]string {
	repo := ctx.Issue.Repo

	return map[string]string{
		// https://datatracker.ietf.org/doc/html/rfc2919
		"List-ID": fmt.Sprintf("%s <%s.%s.%s>", repo.FullName(), repo.Name, repo.OwnerName, setting.Domain),

		// https://datatracker.ietf.org/doc/html/rfc2369
		"List-Archive": fmt.Sprintf("<%s>", repo.HTMLURL()),
		//"List-Post": https://github.com/go-gitea/gitea/pull/13585
		"List-Unsubscribe": ctx.Issue.HTMLURL(),

		"X-Mailer":                  "GitBundle",
		"X-Gitea-Reason":            reason,
		"X-Gitea-Sender":            ctx.Doer.DisplayName(),
		"X-Gitea-Recipient":         recipient.DisplayName(),
		"X-Gitea-Recipient-Address": recipient.Email,
		"X-Gitea-Repository":        repo.Name,
		"X-Gitea-Repository-Path":   repo.FullName(),
		"X-Gitea-Repository-Link":   repo.HTMLURL(),
		"X-Gitea-Issue-ID":          strconv.FormatInt(ctx.Issue.Index, 10),
		"X-Gitea-Issue-Link":        ctx.Issue.HTMLURL(),

		"X-GitHub-Reason":            reason,
		"X-GitHub-Sender":            ctx.Doer.DisplayName(),
		"X-GitHub-Recipient":         recipient.DisplayName(),
		"X-GitHub-Recipient-Address": recipient.Email,

		"X-GitLab-NotificationReason": reason,
		"X-GitLab-Project":            repo.Name,
		"X-GitLab-Project-Path":       repo.FullName(),
		"X-GitLab-Issue-IID":          strconv.FormatInt(ctx.Issue.Index, 10),
	}
}

func sanitizeSubject(subject string) string {
	runes := []rune(strings.TrimSpace(subjectRemoveSpaces.ReplaceAllLiteralString(subject, " ")))
	if len(runes) > mailMaxSubjectRunes {
		runes = runes[:mailMaxSubjectRunes]
	}
	// Encode non-ASCII characters
	return mime.QEncoding.Encode("utf-8", string(runes))
}

// SendIssueAssignedMail composes and sends issue assigned email
func SendIssueAssignedMail(issue *issues_model.Issue, doer *user_model.User, content string, comment *issues_model.Comment, recipients []*user_model.User) error {
	if setting.MailService == nil {
		// No mail service configured
		return nil
	}

	if err := issue.LoadRepo(db.DefaultContext); err != nil {
		log.Error("Unable to load repo [%d] for issue #%d [%d]. Error: %v", issue.RepoID, issue.Index, issue.ID, err)
		return err
	}

	langMap := make(map[string][]*user_model.User)
	for _, user := range recipients {
		if !user.IsActive {
			// don't send emails to inactive users
			continue
		}
		langMap[user.Language] = append(langMap[user.Language], user)
	}

	for lang, tos := range langMap {
		msgs, err := composeIssueCommentMessages(&mailCommentContext{
			Context:    context.TODO(), // TODO: use a correct context
			Issue:      issue,
			Doer:       doer,
			ActionType: action.ActionType(0),
			Content:    content,
			Comment:    comment,
		}, lang, tos, false, "issue assigned")
		if err != nil {
			return err
		}
		SendAsyncs(msgs)
	}
	return nil
}

// actionToTemplate returns the type and name of the action facing the user
// (slightly different from action.ActionType) and the name of the template to use (based on availability)
func actionToTemplate(issue *issues_model.Issue, actionType action.ActionType,
	commentType issues_model.CommentType, reviewType issues_model.ReviewType,
) (typeName, name, template string) {
	if issue.IsPull {
		typeName = "pull"
	} else {
		typeName = "issue"
	}
	switch actionType {
	case action.ActionCreateIssue, action.ActionCreatePullRequest:
		name = "new"
	case action.ActionCommentIssue, action.ActionCommentPull:
		name = "comment"
	case action.ActionCloseIssue, action.ActionClosePullRequest:
		name = "close"
	case action.ActionReopenIssue, action.ActionReopenPullRequest:
		name = "reopen"
	case action.ActionMergePullRequest:
		name = "merge"
	case action.ActionPullReviewDismissed:
		name = "review_dismissed"
	case action.ActionPullRequestReadyForReview:
		name = "ready_for_review"
	default:
		switch commentType {
		case issues_model.CommentTypeReview:
			switch reviewType {
			case issues_model.ReviewTypeApprove:
				name = "approve"
			case issues_model.ReviewTypeReject:
				name = "reject"
			default:
				name = "review"
			}
		case issues_model.CommentTypeCode:
			name = "code"
		case issues_model.CommentTypeAssignees:
			name = "assigned"
		case issues_model.CommentTypePullRequestPush:
			name = "push"
		default:
			name = "default"
		}
	}

	template = typeName + "/" + name
	ok := bodyTemplates.Lookup(template) != nil
	if !ok && typeName != "issue" {
		template = "issue/" + name
		ok = bodyTemplates.Lookup(template) != nil
	}
	if !ok {
		template = typeName + "/default"
		ok = bodyTemplates.Lookup(template) != nil
	}
	if !ok {
		template = "issue/default"
	}
	return
}
