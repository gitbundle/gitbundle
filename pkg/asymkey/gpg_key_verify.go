// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package asymkey

import (
	"strconv"
	"time"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

//   __________________  ________   ____  __.
//  /  _____/\______   \/  _____/  |    |/ _|____ ___.__.
// /   \  ___ |     ___/   \  ___  |      <_/ __ <   |  |
// \    \_\  \|    |   \    \_\  \ |    |  \  ___/\___  |
//  \______  /|____|    \______  / |____|__ \___  > ____|
//         \/                  \/          \/   \/\/
// ____   ____           .__  _____
// \   \ /   /___________|__|/ ____\__.__.
//  \   Y   // __ \_  __ \  \   __<   |  |
//   \     /\  ___/|  | \/  ||  |  \___  |
//    \___/  \___  >__|  |__||__|  / ____|
//               \/                \/

// This file provides functions relating verifying gpg keys

// VerifyGPGKey marks a GPG key as verified
func VerifyGPGKey(ownerID int64, keyID, token, signature string) (string, error) {
	ctx, committer, err := db.TxContext()
	if err != nil {
		return "", err
	}
	defer committer.Close()

	key := new(GPGKey)

	has, err := db.GetEngine(ctx).Where("owner_id = ? AND key_id = ?", ownerID, keyID).Get(key)
	if err != nil {
		return "", err
	} else if !has {
		return "", ErrGPGKeyNotExist{}
	}

	sig, err := extractSignature(signature)
	if err != nil {
		return "", ErrGPGInvalidTokenSignature{
			ID:      key.KeyID,
			Wrapped: err,
		}
	}

	signer, err := hashAndVerifyWithSubKeys(sig, token, key)
	if err != nil {
		return "", ErrGPGInvalidTokenSignature{
			ID:      key.KeyID,
			Wrapped: err,
		}
	}
	if signer == nil {
		signer, err = hashAndVerifyWithSubKeys(sig, token+"\n", key)

		if err != nil {
			return "", ErrGPGInvalidTokenSignature{
				ID:      key.KeyID,
				Wrapped: err,
			}
		}
	}
	if signer == nil {
		signer, err = hashAndVerifyWithSubKeys(sig, token+"\n\n", key)
		if err != nil {
			return "", ErrGPGInvalidTokenSignature{
				ID:      key.KeyID,
				Wrapped: err,
			}
		}
	}

	if signer == nil {
		log.Error("Unable to validate token signature. Error: %v", err)
		return "", ErrGPGInvalidTokenSignature{
			ID: key.KeyID,
		}
	}

	if signer.PrimaryKeyID != key.KeyID && signer.KeyID != key.KeyID {
		return "", ErrGPGKeyNotExist{}
	}

	key.Verified = true
	if _, err := db.GetEngine(ctx).ID(key.ID).SetExpr("verified", true).Update(new(GPGKey)); err != nil {
		return "", err
	}

	if err := committer.Commit(); err != nil {
		return "", err
	}

	return key.KeyID, nil
}

// VerificationToken returns token for the user that will be valid in minutes (time)
func VerificationToken(user *user_model.User, minutes int) string {
	return base.EncodeSha256(
		time.Now().Truncate(1*time.Minute).Add(time.Duration(minutes)*time.Minute).Format(time.RFC1123Z) + ":" +
			user.CreatedUnix.FormatLong() + ":" +
			user.Name + ":" +
			user.Email + ":" +
			strconv.FormatInt(user.ID, 10))
}
