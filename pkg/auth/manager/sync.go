// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/auth"
	"github.com/gitbundle/server/pkg/db"
)

// SyncExternalUsers is used to synchronize users with external authorization source
func SyncExternalUsers(ctx context.Context, updateExisting bool) error {
	log.Trace("Doing: SyncExternalUsers")

	ls, err := auth.Sources()
	if err != nil {
		log.Error("SyncExternalUsers: %v", err)
		return err
	}

	for _, s := range ls {
		if !s.IsActive || !s.IsSyncEnabled {
			continue
		}
		select {
		case <-ctx.Done():
			log.Warn("SyncExternalUsers: Cancelled before update of %s", s.Name)
			return db.ErrCancelledf("Before update of %s", s.Name)
		default:
		}

		if syncable, ok := s.Cfg.(SynchronizableSource); ok {
			err := syncable.Sync(ctx, updateExisting)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
