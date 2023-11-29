// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/util"
	asymkey_model "github.com/gitbundle/server/pkg/asymkey"
	"github.com/gitbundle/server/pkg/auth"
	"github.com/gitbundle/server/pkg/db"
	git_model "github.com/gitbundle/server/pkg/git"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/organization"
	"github.com/gitbundle/server/pkg/perm"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	repo_model "github.com/gitbundle/server/pkg/repo"
	"github.com/gitbundle/server/pkg/unit"
	user_model "github.com/gitbundle/server/pkg/user"
	"github.com/gitbundle/server/pkg/webhook"
	webhook_service "github.com/gitbundle/server/pkg/webhook/manager"
)

// ToEmail convert models.EmailAddress to api.Email
func ToEmail(email *user_model.EmailAddress) *api.Email {
	return &api.Email{
		Email:    email.Email,
		Verified: email.IsActivated,
		Primary:  email.IsPrimary,
	}
}

// ToBranch convert a git.Commit and git.Branch to an api.Branch
func ToBranch(repo *repo_model.Repository, b *git.Branch, c *git.Commit, bp *git_model.ProtectedBranch, user *user_model.User, isRepoAdmin bool) (*api.Branch, error) {
	if bp == nil {
		var hasPerm bool
		var canPush bool
		var err error
		if user != nil {
			hasPerm, err = access_model.HasAccessUnit(db.DefaultContext, user, repo, unit.TypeCode, perm.AccessModeWrite)
			if err != nil {
				return nil, err
			}

			perms, err := access_model.GetUserRepoPermission(db.DefaultContext, repo, user)
			if err != nil {
				return nil, err
			}
			canPush = issues_model.CanMaintainerWriteToBranch(perms, b.Name, user)
		}

		return &api.Branch{
			Name:                b.Name,
			Commit:              ToPayloadCommit(repo, c),
			Protected:           false,
			RequiredApprovals:   0,
			EnableStatusCheck:   false,
			StatusCheckContexts: []string{},
			UserCanPush:         canPush,
			UserCanMerge:        hasPerm,
		}, nil
	}

	branch := &api.Branch{
		Name:                b.Name,
		Commit:              ToPayloadCommit(repo, c),
		Protected:           true,
		RequiredApprovals:   bp.RequiredApprovals,
		EnableStatusCheck:   bp.EnableStatusCheck,
		StatusCheckContexts: bp.StatusCheckContexts,
	}

	if isRepoAdmin {
		branch.EffectiveBranchProtectionName = bp.BranchName
	}

	if user != nil {
		permission, err := access_model.GetUserRepoPermission(db.DefaultContext, repo, user)
		if err != nil {
			return nil, err
		}
		branch.UserCanPush = bp.CanUserPush(user.ID)
		branch.UserCanMerge = git_model.IsUserMergeWhitelisted(db.DefaultContext, bp, user.ID, permission)
	}

	return branch, nil
}

// ToBranchProtection convert a ProtectedBranch to api.BranchProtection
func ToBranchProtection(bp *git_model.ProtectedBranch) *api.BranchProtection {
	pushWhitelistUsernames, err := user_model.GetUserNamesByIDs(bp.WhitelistUserIDs)
	if err != nil {
		log.Error("GetUserNamesByIDs (WhitelistUserIDs): %v", err)
	}
	mergeWhitelistUsernames, err := user_model.GetUserNamesByIDs(bp.MergeWhitelistUserIDs)
	if err != nil {
		log.Error("GetUserNamesByIDs (MergeWhitelistUserIDs): %v", err)
	}
	approvalsWhitelistUsernames, err := user_model.GetUserNamesByIDs(bp.ApprovalsWhitelistUserIDs)
	if err != nil {
		log.Error("GetUserNamesByIDs (ApprovalsWhitelistUserIDs): %v", err)
	}
	pushWhitelistTeams, err := organization.GetTeamNamesByID(bp.WhitelistTeamIDs)
	if err != nil {
		log.Error("GetTeamNamesByID (WhitelistTeamIDs): %v", err)
	}
	mergeWhitelistTeams, err := organization.GetTeamNamesByID(bp.MergeWhitelistTeamIDs)
	if err != nil {
		log.Error("GetTeamNamesByID (MergeWhitelistTeamIDs): %v", err)
	}
	approvalsWhitelistTeams, err := organization.GetTeamNamesByID(bp.ApprovalsWhitelistTeamIDs)
	if err != nil {
		log.Error("GetTeamNamesByID (ApprovalsWhitelistTeamIDs): %v", err)
	}

	return &api.BranchProtection{
		BranchName:                    bp.BranchName,
		EnablePush:                    bp.CanPush,
		EnablePushWhitelist:           bp.EnableWhitelist,
		PushWhitelistUsernames:        pushWhitelistUsernames,
		PushWhitelistTeams:            pushWhitelistTeams,
		PushWhitelistDeployKeys:       bp.WhitelistDeployKeys,
		EnableMergeWhitelist:          bp.EnableMergeWhitelist,
		MergeWhitelistUsernames:       mergeWhitelistUsernames,
		MergeWhitelistTeams:           mergeWhitelistTeams,
		EnableStatusCheck:             bp.EnableStatusCheck,
		StatusCheckContexts:           bp.StatusCheckContexts,
		RequiredApprovals:             bp.RequiredApprovals,
		EnableApprovalsWhitelist:      bp.EnableApprovalsWhitelist,
		ApprovalsWhitelistUsernames:   approvalsWhitelistUsernames,
		ApprovalsWhitelistTeams:       approvalsWhitelistTeams,
		BlockOnRejectedReviews:        bp.BlockOnRejectedReviews,
		BlockOnOfficialReviewRequests: bp.BlockOnOfficialReviewRequests,
		BlockOnOutdatedBranch:         bp.BlockOnOutdatedBranch,
		DismissStaleApprovals:         bp.DismissStaleApprovals,
		RequireSignedCommits:          bp.RequireSignedCommits,
		ProtectedFilePatterns:         bp.ProtectedFilePatterns,
		UnprotectedFilePatterns:       bp.UnprotectedFilePatterns,
		Created:                       bp.CreatedUnix.AsTime(),
		Updated:                       bp.UpdatedUnix.AsTime(),
	}
}

// ToTag convert a git.Tag to an api.Tag
func ToTag(repo *repo_model.Repository, t *git.Tag) *api.Tag {
	return &api.Tag{
		Name:       t.Name,
		Message:    strings.TrimSpace(t.Message),
		ID:         t.ID.String(),
		Commit:     ToCommitMeta(repo, t),
		ZipballURL: util.URLJoin(repo.HTMLURL(), "archive", t.Name+".zip"),
		TarballURL: util.URLJoin(repo.HTMLURL(), "archive", t.Name+".tar.gz"),
	}
}

// ToVerification convert a git.Commit.Signature to an api.PayloadCommitVerification
func ToVerification(c *git.Commit) *api.PayloadCommitVerification {
	verif := asymkey_model.ParseCommitWithSignature(c)
	commitVerification := &api.PayloadCommitVerification{
		Verified: verif.Verified,
		Reason:   verif.Reason,
	}
	if c.Signature != nil {
		commitVerification.Signature = c.Signature.Signature
		commitVerification.Payload = c.Signature.Payload
	}
	if verif.SigningUser != nil {
		commitVerification.Signer = &api.PayloadUser{
			Name:  verif.SigningUser.Name,
			Email: verif.SigningUser.Email,
		}
	}
	return commitVerification
}

// ToPublicKey convert asymkey_model.PublicKey to api.PublicKey
func ToPublicKey(apiLink string, key *asymkey_model.PublicKey) *api.PublicKey {
	return &api.PublicKey{
		ID:          key.ID,
		Key:         key.Content,
		URL:         fmt.Sprintf("%s%d", apiLink, key.ID),
		Title:       key.Name,
		Fingerprint: key.Fingerprint,
		Created:     key.CreatedUnix.AsTime(),
	}
}

// ToGPGKey converts models.GPGKey to api.GPGKey
func ToGPGKey(key *asymkey_model.GPGKey) *api.GPGKey {
	subkeys := make([]*api.GPGKey, len(key.SubsKey))
	for id, k := range key.SubsKey {
		subkeys[id] = &api.GPGKey{
			ID:                k.ID,
			PrimaryKeyID:      k.PrimaryKeyID,
			KeyID:             k.KeyID,
			PublicKey:         k.Content,
			Created:           k.CreatedUnix.AsTime(),
			Expires:           k.ExpiredUnix.AsTime(),
			CanSign:           k.CanSign,
			CanEncryptComms:   k.CanEncryptComms,
			CanEncryptStorage: k.CanEncryptStorage,
			CanCertify:        k.CanSign,
			Verified:          k.Verified,
		}
	}
	emails := make([]*api.GPGKeyEmail, len(key.Emails))
	for i, e := range key.Emails {
		emails[i] = ToGPGKeyEmail(e)
	}
	return &api.GPGKey{
		ID:                key.ID,
		PrimaryKeyID:      key.PrimaryKeyID,
		KeyID:             key.KeyID,
		PublicKey:         key.Content,
		Created:           key.CreatedUnix.AsTime(),
		Expires:           key.ExpiredUnix.AsTime(),
		Emails:            emails,
		SubsKey:           subkeys,
		CanSign:           key.CanSign,
		CanEncryptComms:   key.CanEncryptComms,
		CanEncryptStorage: key.CanEncryptStorage,
		CanCertify:        key.CanSign,
		Verified:          key.Verified,
	}
}

// ToGPGKeyEmail convert models.EmailAddress to api.GPGKeyEmail
func ToGPGKeyEmail(email *user_model.EmailAddress) *api.GPGKeyEmail {
	return &api.GPGKeyEmail{
		Email:    email.Email,
		Verified: email.IsActivated,
	}
}

// ToHook convert models.Webhook to api.Hook
func ToHook(repoLink string, w *webhook.Webhook) *api.Hook {
	config := map[string]string{
		"url":          w.URL,
		"content_type": w.ContentType.Name(),
	}
	if w.Type == webhook.SLACK {
		s := webhook_service.GetSlackHook(w)
		config["channel"] = s.Channel
		config["username"] = s.Username
		config["icon_url"] = s.IconURL
		config["color"] = s.Color
	}

	return &api.Hook{
		ID:      w.ID,
		Type:    string(w.Type),
		URL:     fmt.Sprintf("%s/settings/hooks/%d", repoLink, w.ID),
		Active:  w.IsActive,
		Config:  config,
		Events:  w.EventsArray(),
		Updated: w.UpdatedUnix.AsTime(),
		Created: w.CreatedUnix.AsTime(),
	}
}

// ToGitHook convert git.Hook to api.GitHook
func ToGitHook(h *git.Hook) *api.GitHook {
	return &api.GitHook{
		Name:     h.Name(),
		IsActive: h.IsActive,
		Content:  h.Content,
	}
}

// ToDeployKey convert asymkey_model.DeployKey to api.DeployKey
func ToDeployKey(apiLink string, key *asymkey_model.DeployKey) *api.DeployKey {
	return &api.DeployKey{
		ID:          key.ID,
		KeyID:       key.KeyID,
		Key:         key.Content,
		Fingerprint: key.Fingerprint,
		URL:         fmt.Sprintf("%s%d", apiLink, key.ID),
		Title:       key.Name,
		Created:     key.CreatedUnix.AsTime(),
		ReadOnly:    key.Mode == perm.AccessModeRead, // All deploy keys are read-only.
	}
}

// ToOrganization convert user_model.User to api.Organization
func ToOrganization(org *organization.Organization) *api.Organization {
	return &api.Organization{
		ID:                        org.ID,
		AvatarURL:                 org.AsUser().AvatarLink(),
		UserName:                  org.Name,
		FullName:                  org.FullName,
		Description:               org.Description,
		Website:                   org.Website,
		Location:                  org.Location,
		Visibility:                org.Visibility.String(),
		RepoAdminChangeTeamAccess: org.RepoAdminChangeTeamAccess,
	}
}

// ToTeam convert models.Team to api.Team
func ToTeam(team *organization.Team, loadOrg ...bool) (*api.Team, error) {
	teams, err := ToTeams([]*organization.Team{team}, len(loadOrg) != 0 && loadOrg[0])
	if err != nil || len(teams) == 0 {
		return nil, err
	}
	return teams[0], nil
}

// ToTeams convert models.Team list to api.Team list
func ToTeams(teams []*organization.Team, loadOrgs bool) ([]*api.Team, error) {
	if len(teams) == 0 || teams[0] == nil {
		return nil, nil
	}

	cache := make(map[int64]*api.Organization)
	apiTeams := make([]*api.Team, len(teams))
	for i := range teams {
		if err := teams[i].GetUnits(); err != nil {
			return nil, err
		}

		apiTeams[i] = &api.Team{
			ID:                      teams[i].ID,
			Name:                    teams[i].Name,
			Description:             teams[i].Description,
			IncludesAllRepositories: teams[i].IncludesAllRepositories,
			CanCreateOrgRepo:        teams[i].CanCreateOrgRepo,
			Permission:              teams[i].AccessMode.String(),
			Units:                   teams[i].GetUnitNames(),
			UnitsMap:                teams[i].GetUnitsMap(),
		}

		if loadOrgs {
			apiOrg, ok := cache[teams[i].OrgID]
			if !ok {
				org, err := organization.GetOrgByID(db.DefaultContext, teams[i].OrgID)
				if err != nil {
					return nil, err
				}
				apiOrg = ToOrganization(org)
				cache[teams[i].OrgID] = apiOrg
			}
			apiTeams[i].Organization = apiOrg
		}
	}
	return apiTeams, nil
}

// ToAnnotatedTag convert git.Tag to api.AnnotatedTag
func ToAnnotatedTag(repo *repo_model.Repository, t *git.Tag, c *git.Commit) *api.AnnotatedTag {
	return &api.AnnotatedTag{
		Tag:          t.Name,
		SHA:          t.ID.String(),
		Object:       ToAnnotatedTagObject(repo, c),
		Message:      t.Message,
		URL:          util.URLJoin(repo.APIURL(), "git/tags", t.ID.String()),
		Tagger:       ToCommitUser(t.Tagger),
		Verification: ToVerification(c),
	}
}

// ToAnnotatedTagObject convert a git.Commit to an api.AnnotatedTagObject
func ToAnnotatedTagObject(repo *repo_model.Repository, commit *git.Commit) *api.AnnotatedTagObject {
	return &api.AnnotatedTagObject{
		SHA:  commit.ID.String(),
		Type: string(git.ObjectCommit),
		URL:  util.URLJoin(repo.APIURL(), "git/commits", commit.ID.String()),
	}
}

// ToTopicResponse convert from models.Topic to api.TopicResponse
func ToTopicResponse(topic *repo_model.Topic) *api.TopicResponse {
	return &api.TopicResponse{
		ID:        topic.ID,
		Name:      topic.Name,
		RepoCount: topic.RepoCount,
		Created:   topic.CreatedUnix.AsTime(),
		Updated:   topic.UpdatedUnix.AsTime(),
	}
}

// ToOAuth2Application convert from auth.OAuth2Application to api.OAuth2Application
func ToOAuth2Application(app *auth.OAuth2Application) *api.OAuth2Application {
	return &api.OAuth2Application{
		ID:           app.ID,
		Name:         app.Name,
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
		RedirectURIs: app.RedirectURIs,
		Created:      app.CreatedUnix.AsTime(),
	}
}

// ToLFSLock convert a LFSLock to api.LFSLock
func ToLFSLock(l *git_model.LFSLock) *api.LFSLock {
	u, err := user_model.GetUserByID(l.OwnerID)
	if err != nil {
		return nil
	}
	return &api.LFSLock{
		ID:       strconv.FormatInt(l.ID, 10),
		Path:     l.Path,
		LockedAt: l.Created.Round(time.Second),
		Owner: &api.LFSLockOwner{
			Name: u.DisplayName(),
		},
	}
}