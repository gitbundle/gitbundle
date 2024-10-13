// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package appstate

import (
	"github.com/gitbundle/modules/json"

	"github.com/yuin/goldmark/util"
)

// StateStore is the interface to get/set app state items
type StateStore interface {
	Get(item StateItem) error
	Set(item StateItem) error
}

// StateItem provides the name for a state item. the name will be used to generate filenames, etc
type StateItem interface {
	Name() string
}

// AppState contains the state items for the app
var AppState2 StateStore

// Init initialize AppState interface
func Init() error {
	AppState2 = &DBStore{}
	return nil
}

// DBStore can be used to store app state items in local filesystem
type DBStore struct{}

// Get reads the state item
func (f *DBStore) Get(item StateItem) error {
	content, err := GetAppStateContent(item.Name())
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}
	return json.Unmarshal(util.StringToReadOnlyBytes(content), item)
}

// Set saves the state item
func (f *DBStore) Set(item StateItem) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return SaveAppStateContent(item.Name(), util.BytesToReadOnlyString(b))
}
