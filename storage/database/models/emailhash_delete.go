// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/emailhash"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// EmailHashDelete is the builder for deleting a EmailHash entity.
type EmailHashDelete struct {
	config
	hooks    []Hook
	mutation *EmailHashMutation
}

// Where appends a list predicates to the EmailHashDelete builder.
func (ehd *EmailHashDelete) Where(ps ...predicate.EmailHash) *EmailHashDelete {
	ehd.mutation.Where(ps...)
	return ehd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ehd *EmailHashDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ehd.sqlExec, ehd.mutation, ehd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ehd *EmailHashDelete) ExecX(ctx context.Context) int {
	n, err := ehd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ehd *EmailHashDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emailhash.Table, sqlgraph.NewFieldSpec(emailhash.FieldID, field.TypeInt))
	if ps := ehd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ehd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ehd.mutation.done = true
	return affected, err
}

// EmailHashDeleteOne is the builder for deleting a single EmailHash entity.
type EmailHashDeleteOne struct {
	ehd *EmailHashDelete
}

// Where appends a list predicates to the EmailHashDelete builder.
func (ehdo *EmailHashDeleteOne) Where(ps ...predicate.EmailHash) *EmailHashDeleteOne {
	ehdo.ehd.mutation.Where(ps...)
	return ehdo
}

// Exec executes the deletion query.
func (ehdo *EmailHashDeleteOne) Exec(ctx context.Context) error {
	n, err := ehdo.ehd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailhash.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ehdo *EmailHashDeleteOne) ExecX(ctx context.Context) {
	if err := ehdo.Exec(ctx); err != nil {
		panic(err)
	}
}