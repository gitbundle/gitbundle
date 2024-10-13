// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build !bindata

package options

import (
	"github.com/gitbundle/modules/setting"
)

// Dir returns all files from bindata or custom directory.
func Dir(name string) ([]string, error) {
	return setting.Dir(name)
}

// Locale reads the content of a specific locale from bindata or custom path.
func Locale(name string) ([]byte, error) {
	return setting.Locale(name)
}

// Readme reads the content of a specific readme from bindata or custom path.
func Readme(name string) ([]byte, error) {
	return setting.Readme(name)
}

// Gitignore reads the content of a gitignore locale from bindata or custom path.
func Gitignore(name string) ([]byte, error) {
	return setting.Gitignore(name)
}

// License reads the content of a specific license from bindata or custom path.
func License(name string) ([]byte, error) {
	return setting.License(name)
}

// Labels reads the content of a specific labels from static or custom path.
func Labels(name string) ([]byte, error) {
	return setting.Labels(name)
}

// IsDynamic will return false when using embedded data (-tags bindata)
func IsDynamic() bool {
	return true
}
