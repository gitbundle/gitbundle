// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package database

// Config specifies the config for the database package.
type Config struct {
	Driver     string // one of sqlite, postgres, mysql
	Datasource string
	Debug      bool // debug database
}
