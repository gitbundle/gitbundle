package atlashooks

import (
	"context"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

// In case a field was renamed in the ent/schema, Ent won't detect this change as renaming and will propose
// DropColumn and AddColumn changes in the diff stage. One way to get over this is to use the StorageKey option
// on the field and keep the old column name in the database table. However, using Atlas Diff hooks allow
// replacing the DropColumn and AddColumn changes with a RenameColumn change.
func DefaultDiffHook(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		changes, err := next.Diff(current, desired)
		if err != nil {
			return nil, err
		}

		return changes, nil
	})
}

// The Apply hook allows accessing and mutating the migration plan and its raw changes (SQL statements), but in
// addition to that it is also useful for executing custom SQL statements before or after the plan is applied.
// For example, changing a nullable column to non-nullable without a default value is not allowed by default.
// However, we can work around this using an Apply hook that UPDATEs all rows that contain NULL value in this column:
func DefaultApplyHook(next schema.Applier) schema.Applier {
	return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
		return next.Apply(ctx, conn, plan)
	})
}
