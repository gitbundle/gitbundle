// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package asymkey

import "github.com/gitbundle/server/pkg/db"

//    __________________  ________   ____  __.
//   /  _____/\______   \/  _____/  |    |/ _|____ ___.__.
//  /   \  ___ |     ___/   \  ___  |      <_/ __ <   |  |
//  \    \_\  \|    |   \    \_\  \ |    |  \  ___/\___  |
//   \______  /|____|    \______  / |____|__ \___  > ____|
//          \/                  \/          \/   \/\/
//  .___                              __
//  |   | _____ ______   ____________/  |_
//  |   |/     \\____ \ /  _ \_  __ \   __\
//  |   |  Y Y  \  |_> >  <_> )  | \/|  |
//  |___|__|_|  /   __/ \____/|__|   |__|
//            \/|__|

// This file contains functions related to the original import of a key

// GPGKeyImport the original import of key
type GPGKeyImport struct {
	KeyID   string `xorm:"pk CHAR(16) NOT NULL"`
	Content string `xorm:"TEXT NOT NULL"`
}

func init() {
	db.RegisterModel(new(GPGKeyImport))
}

// GetGPGImportByKeyID returns the import public armored key by given KeyID.
func GetGPGImportByKeyID(keyID string) (*GPGKeyImport, error) {
	key := new(GPGKeyImport)
	has, err := db.GetEngine(db.DefaultContext).ID(keyID).Get(key)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrGPGKeyImportNotExist{keyID}
	}
	return key, nil
}
