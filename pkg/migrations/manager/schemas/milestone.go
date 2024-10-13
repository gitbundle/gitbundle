// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package schemas

import "time"

// Milestone defines a standard milestone
type Milestone struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	Created     time.Time  `json:"created"`
	Updated     *time.Time `json:"updated"`
	Closed      *time.Time `json:"closed"`
	State       string     `json:"state"` // open, closed
}
