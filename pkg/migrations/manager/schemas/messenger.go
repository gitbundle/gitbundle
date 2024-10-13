// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package schemas

// Messenger is a formatting function similar to i18n.Tr
type Messenger func(key string, args ...interface{})

// NilMessenger represents an empty formatting function
func NilMessenger(string, ...interface{}) {}
