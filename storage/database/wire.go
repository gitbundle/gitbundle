// Copyright 2024 GitBundle, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"context"
	"fmt"
	"os"

	"github.com/gitbundle/gitbundle/storage/database/atlashooks"
	"github.com/gitbundle/gitbundle/storage/database/models"
	"github.com/gitbundle/gitbundle/storage/database/models/migrate"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var WireSet = wire.NewSet(
	ProvideDatabaseClient,
)

// ProvideDatabaseClient provides a ent database connection client.
// The database client should be closed by caller.
func ProvideDatabaseClient(ctx context.Context, config Config) (*models.Client, error) {
	client, err := models.Open(config.Driver, config.Datasource)
	if err != nil {
		return nil, fmt.Errorf("database: failed opening connection: %v", err)
	}

	if config.Debug {
		//Print the sql to be executed.
		if err := client.Schema.WriteTo(
			ctx,
			os.Stdout,
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
			migrate.WithGlobalUniqueID(true),
			// Hook into Atlas Diff process.
			schema.WithDiffHook(
				atlashooks.DefaultDiffHook,
			),
			// Hook into Atlas Apply process.
			schema.WithApplyHook(
				// Default apply hook should be executed at last
				atlashooks.DefaultApplyHook,
			),
		); err != nil {
			return nil, fmt.Errorf("database: failed printing schema changes: %v", err)
		}
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true),
		// Hook into Atlas Diff process.
		schema.WithDiffHook(
			atlashooks.DefaultDiffHook,
		),
		// Hook into Atlas Apply process.
		schema.WithApplyHook(
			// Default apply hook should be executed at last
			atlashooks.DefaultApplyHook,
		),
	); err != nil {
		return nil, fmt.Errorf("database: failed creating schema resources: %v", err)
	}

	return client, nil
}
