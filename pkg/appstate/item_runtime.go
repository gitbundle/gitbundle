// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package appstate

// RuntimeState contains app state for runtime, and we can save remote version for update checker here in future
type RuntimeState struct {
	LastAppPath string `json:"last_app_path"`
}

// Name returns the item name
func (a RuntimeState) Name() string {
	return "runtime-state"
}
