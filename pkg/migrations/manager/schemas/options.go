// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package schemas

import "github.com/gitbundle/api/pkg/structs"

// MigrateOptions defines the way a repository gets migrated
// this is for internal usage by migrations module and func who interact with it
type MigrateOptions struct {
	// required: true
	CloneAddr             string `json:"clone_addr" binding:"Required"`
	CloneAddrEncrypted    string `json:"clone_addr_encrypted,omitempty"`
	AuthUsername          string `json:"auth_username"`
	AuthPassword          string `json:"-"`
	AuthPasswordEncrypted string `json:"auth_password_encrypted,omitempty"`
	AuthToken             string `json:"-"`
	AuthTokenEncrypted    string `json:"auth_token_encrypted,omitempty"`
	// required: true
	UID int `json:"uid" binding:"Required"`
	// required: true
	RepoName        string `json:"repo_name" binding:"Required"`
	Mirror          bool   `json:"mirror"`
	LFS             bool   `json:"lfs"`
	LFSEndpoint     string `json:"lfs_endpoint"`
	Private         bool   `json:"private"`
	Description     string `json:"description"`
	OriginalURL     string
	GitServiceType  structs.GitServiceType
	Wiki            bool
	Issues          bool
	Milestones      bool
	Labels          bool
	Releases        bool
	Comments        bool
	PullRequests    bool
	ReleaseAssets   bool
	MigrateToRepoID int64
	MirrorInterval  string `json:"mirror_interval"`
}
