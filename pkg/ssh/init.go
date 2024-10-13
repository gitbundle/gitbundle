// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package ssh

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
)

func Init() error {
	if setting.SSH.Disabled {
		return nil
	}

	if setting.SSH.StartBuiltinServer {
		Listen(setting.SSH.ListenHost, setting.SSH.ListenPort, setting.SSH.ServerCiphers, setting.SSH.ServerKeyExchanges, setting.SSH.ServerMACs)
		log.Info("SSH server started on %s. Cipher list (%v), key exchange algorithms (%v), MACs (%v)",
			net.JoinHostPort(setting.SSH.ListenHost, strconv.Itoa(setting.SSH.ListenPort)),
			setting.SSH.ServerCiphers, setting.SSH.ServerKeyExchanges, setting.SSH.ServerMACs,
		)
		return nil
	}

	builtinUnused()

	// FIXME: why 0o644 for a directory .....
	if err := os.MkdirAll(setting.SSH.KeyTestPath, 0o644); err != nil {
		return fmt.Errorf("failed to create directory %q for ssh key test: %w", setting.SSH.KeyTestPath, err)
	}

	if len(setting.SSH.TrustedUserCAKeys) > 0 && setting.SSH.AuthorizedPrincipalsEnabled {
		caKeysFileName := setting.SSH.TrustedUserCAKeysFile
		caKeysFileDir := filepath.Dir(caKeysFileName)

		err := os.MkdirAll(caKeysFileDir, 0o700) // SSH.RootPath by default (That is `~/.ssh` in most cases)
		if err != nil {
			return fmt.Errorf("failed to create directory %q for ssh trusted ca keys: %w", caKeysFileDir, err)
		}

		if err := os.WriteFile(caKeysFileName, []byte(strings.Join(setting.SSH.TrustedUserCAKeys, "\n")), 0o600); err != nil {
			return fmt.Errorf("failed to write ssh trusted ca keys to %q: %w", caKeysFileName, err)
		}
	}

	return nil
}
