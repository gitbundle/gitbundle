// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package externalaccount

import (
	"strings"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/server/pkg/auth"
	migration_model "github.com/gitbundle/server/pkg/migrations"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/markbates/goth"
)

func toExternalLoginUser(user *user_model.User, gothUser goth.User) (*user_model.ExternalLoginUser, error) {
	authSource, err := auth.GetActiveOAuth2SourceByName(gothUser.Provider)
	if err != nil {
		return nil, err
	}
	return &user_model.ExternalLoginUser{
		ExternalID:        gothUser.UserID,
		UserID:            user.ID,
		LoginSourceID:     authSource.ID,
		RawData:           gothUser.RawData,
		Provider:          gothUser.Provider,
		Email:             gothUser.Email,
		Name:              gothUser.Name,
		FirstName:         gothUser.FirstName,
		LastName:          gothUser.LastName,
		NickName:          gothUser.NickName,
		Description:       gothUser.Description,
		AvatarURL:         gothUser.AvatarURL,
		Location:          gothUser.Location,
		AccessToken:       gothUser.AccessToken,
		AccessTokenSecret: gothUser.AccessTokenSecret,
		RefreshToken:      gothUser.RefreshToken,
		ExpiresAt:         gothUser.ExpiresAt,
	}, nil
}

// LinkAccountToUser link the gothUser to the user
func LinkAccountToUser(user *user_model.User, gothUser goth.User) error {
	externalLoginUser, err := toExternalLoginUser(user, gothUser)
	if err != nil {
		return err
	}

	if err := user_model.LinkExternalToUser(user, externalLoginUser); err != nil {
		return err
	}

	externalID := externalLoginUser.ExternalID

	var tp structs.GitServiceType
	for _, s := range structs.SupportedFullGitService {
		if strings.EqualFold(s.Name(), gothUser.Provider) {
			tp = s
			break
		}
	}

	if tp.Name() != "" {
		return migration_model.UpdateMigrationsByType(tp, externalID, user.ID)
	}

	return nil
}

// UpdateExternalUser updates external user's information
func UpdateExternalUser(user *user_model.User, gothUser goth.User) error {
	externalLoginUser, err := toExternalLoginUser(user, gothUser)
	if err != nil {
		return err
	}

	return user_model.UpdateExternalUserByExternalID(externalLoginUser)
}
