// Copyright 2024 GitBundle, Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package database

import (
	"context"
	"fmt"
	"os"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"

	"github.com/gitbundle/gitbundle/storage/database/atlashooks"
	"github.com/gitbundle/gitbundle/storage/database/models"
	"github.com/gitbundle/gitbundle/storage/database/models/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var WireSet = wire.NewSet(
	ProvideDatabaseClient,
)

// ProvideDatabaseClient provides a ent database connection client.
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
