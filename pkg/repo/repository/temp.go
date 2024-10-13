// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/util"
)

// LocalCopyPath returns the local repository temporary copy path.
func LocalCopyPath() string {
	if filepath.IsAbs(setting.Repository.Local.LocalCopyPath) {
		return setting.Repository.Local.LocalCopyPath
	}
	return path.Join(setting.AppDataPath, setting.Repository.Local.LocalCopyPath)
}

// CreateTemporaryPath creates a temporary path
func CreateTemporaryPath(prefix string) (string, error) {
	if err := os.MkdirAll(LocalCopyPath(), os.ModePerm); err != nil {
		log.Error("Unable to create localcopypath directory: %s (%v)", LocalCopyPath(), err)
		return "", fmt.Errorf("Failed to create localcopypath directory %s: %v", LocalCopyPath(), err)
	}
	basePath, err := os.MkdirTemp(LocalCopyPath(), prefix+".git")
	if err != nil {
		log.Error("Unable to create temporary directory: %s-*.git (%v)", prefix, err)
		return "", fmt.Errorf("Failed to create dir %s-*.git: %v", prefix, err)

	}
	return basePath, nil
}

// RemoveTemporaryPath removes the temporary path
func RemoveTemporaryPath(basePath string) error {
	if _, err := os.Stat(basePath); !os.IsNotExist(err) {
		return util.RemoveAll(basePath)
	}
	return nil
}
