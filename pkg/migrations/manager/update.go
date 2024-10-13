// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"context"

	"github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/db"
	migratration_model "github.com/gitbundle/server/pkg/migrations"
	user_model "github.com/gitbundle/server/pkg/user"
)

// UpdateMigrationPosterID updates all migrated repositories' issues and comments posterID
func UpdateMigrationPosterID(ctx context.Context) error {
	for _, gitService := range structs.SupportedFullGitService {
		select {
		case <-ctx.Done():
			log.Warn("UpdateMigrationPosterID aborted before %s", gitService.Name())
			return db.ErrCancelledf("during UpdateMigrationPosterID before %s", gitService.Name())
		default:
		}
		if err := updateMigrationPosterIDByGitService(ctx, gitService); err != nil {
			log.Error("updateMigrationPosterIDByGitService failed: %v", err)
		}
	}
	return nil
}

func updateMigrationPosterIDByGitService(ctx context.Context, tp structs.GitServiceType) error {
	provider := tp.Name()
	if len(provider) == 0 {
		return nil
	}

	const batchSize = 100
	var start int
	for {
		select {
		case <-ctx.Done():
			log.Warn("UpdateMigrationPosterIDByGitService(%s) cancelled", tp.Name())
			return nil
		default:
		}

		users, err := user_model.FindExternalUsersByProvider(user_model.FindExternalUserOptions{
			Provider: provider,
			Start:    start,
			Limit:    batchSize,
		})
		if err != nil {
			return err
		}

		for _, user := range users {
			select {
			case <-ctx.Done():
				log.Warn("UpdateMigrationPosterIDByGitService(%s) cancelled", tp.Name())
				return nil
			default:
			}
			externalUserID := user.ExternalID
			if err := migratration_model.UpdateMigrationsByType(tp, externalUserID, user.UserID); err != nil {
				log.Error("UpdateMigrationsByType type %s external user id %v to local user id %v failed: %v", tp.Name(), user.ExternalID, user.UserID, err)
			}
		}

		if len(users) < batchSize {
			break
		}
		start += len(users)
	}
	return nil
}
