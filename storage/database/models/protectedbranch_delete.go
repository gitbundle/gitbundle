// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
	"github.com/gitbundle/gitbundle/storage/database/models/protectedbranch"
)

// ProtectedBranchDelete is the builder for deleting a ProtectedBranch entity.
type ProtectedBranchDelete struct {
	config
	hooks    []Hook
	mutation *ProtectedBranchMutation
}

// Where appends a list predicates to the ProtectedBranchDelete builder.
func (pbd *ProtectedBranchDelete) Where(ps ...predicate.ProtectedBranch) *ProtectedBranchDelete {
	pbd.mutation.Where(ps...)
	return pbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pbd *ProtectedBranchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pbd.sqlExec, pbd.mutation, pbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pbd *ProtectedBranchDelete) ExecX(ctx context.Context) int {
	n, err := pbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pbd *ProtectedBranchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(protectedbranch.Table, sqlgraph.NewFieldSpec(protectedbranch.FieldID, field.TypeInt))
	if ps := pbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pbd.mutation.done = true
	return affected, err
}

// ProtectedBranchDeleteOne is the builder for deleting a single ProtectedBranch entity.
type ProtectedBranchDeleteOne struct {
	pbd *ProtectedBranchDelete
}

// Where appends a list predicates to the ProtectedBranchDelete builder.
func (pbdo *ProtectedBranchDeleteOne) Where(ps ...predicate.ProtectedBranch) *ProtectedBranchDeleteOne {
	pbdo.pbd.mutation.Where(ps...)
	return pbdo
}

// Exec executes the deletion query.
func (pbdo *ProtectedBranchDeleteOne) Exec(ctx context.Context) error {
	n, err := pbdo.pbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{protectedbranch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pbdo *ProtectedBranchDeleteOne) ExecX(ctx context.Context) {
	if err := pbdo.Exec(ctx); err != nil {
		panic(err)
	}
}