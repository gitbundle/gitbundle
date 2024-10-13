// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package asymkey

import (
	"bytes"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/db"

	"github.com/42wim/sshsig"
)

// VerifySSHKey marks a SSH key as verified
func VerifySSHKey(ownerID int64, fingerprint, token, signature string) (string, error) {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return "", err
	}
	defer committer.Close()

	key := new(PublicKey)

	has, err := db.GetEngine(ctx).Where("owner_id = ? AND fingerprint = ?", ownerID, fingerprint).Get(key)
	if err != nil {
		return "", err
	} else if !has {
		return "", ErrKeyNotExist{}
	}

	if err := sshsig.Verify(bytes.NewBuffer([]byte(token)), []byte(signature), []byte(key.Content), "gitea"); err != nil {
		log.Error("Unable to validate token signature. Error: %v", err)
		return "", ErrSSHInvalidTokenSignature{
			Fingerprint: key.Fingerprint,
		}
	}

	key.Verified = true
	if _, err := db.GetEngine(ctx).ID(key.ID).Cols("verified").Update(key); err != nil {
		return "", err
	}

	if err := committer.Commit(); err != nil {
		return "", err
	}

	return key.Fingerprint, nil
}
