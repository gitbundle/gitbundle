// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package schemas

// Repository defines a standard repository information
type Repository struct {
	Name          string
	Owner         string
	IsPrivate     bool `yaml:"is_private"`
	IsMirror      bool `yaml:"is_mirror"`
	Description   string
	CloneURL      string `yaml:"clone_url"`
	OriginalURL   string `yaml:"original_url"`
	DefaultBranch string
}
