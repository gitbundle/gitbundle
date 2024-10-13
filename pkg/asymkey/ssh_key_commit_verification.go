// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package asymkey

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/42wim/sshsig"
)

// ParseCommitWithSSHSignature check if signature is good against keystore.
func ParseCommitWithSSHSignature(c *git.Commit, committer *user_model.User) *CommitVerification {
	// Now try to associate the signature with the committer, if present
	if committer.ID != 0 {
		keys, err := ListPublicKeys(committer.ID, db.ListOptions{})
		if err != nil { // Skipping failed to get ssh keys of user
			log.Error("ListPublicKeys: %v", err)
			return &CommitVerification{
				CommittingUser: committer,
				Verified:       false,
				Reason:         "gpg.error.failed_retrieval_gpg_keys",
			}
		}

		committerEmailAddresses, err := user_model.GetEmailAddresses(committer.ID)
		if err != nil {
			log.Error("GetEmailAddresses: %v", err)
		}

		activated := false
		for _, e := range committerEmailAddresses {
			if e.IsActivated && strings.EqualFold(e.Email, c.Committer.Email) {
				activated = true
				break
			}
		}

		for _, k := range keys {
			if k.Verified && activated {
				commitVerification := verifySSHCommitVerification(c.Signature.Signature, c.Signature.Payload, k, committer, committer, c.Committer.Email)
				if commitVerification != nil {
					return commitVerification
				}
			}
		}
	}

	return &CommitVerification{
		CommittingUser: committer,
		Verified:       false,
		Reason:         NoKeyFound,
	}
}

func verifySSHCommitVerification(sig, payload string, k *PublicKey, committer, signer *user_model.User, email string) *CommitVerification {
	if err := sshsig.Verify(bytes.NewBuffer([]byte(payload)), []byte(sig), []byte(k.Content), "git"); err != nil {
		return nil
	}

	return &CommitVerification{ // Everything is ok
		CommittingUser: committer,
		Verified:       true,
		Reason:         fmt.Sprintf("%s / %s", signer.Name, k.Fingerprint),
		SigningUser:    signer,
		SigningSSHKey:  k,
		SigningEmail:   email,
	}
}
