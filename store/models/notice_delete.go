// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/notice"
	"github.com/gitbundle/gitbundle/store/models/predicate"
)

// NoticeDelete is the builder for deleting a Notice entity.
type NoticeDelete struct {
	config
	hooks    []Hook
	mutation *NoticeMutation
}

// Where appends a list predicates to the NoticeDelete builder.
func (nd *NoticeDelete) Where(ps ...predicate.Notice) *NoticeDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NoticeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NoticeDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NoticeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notice.Table, sqlgraph.NewFieldSpec(notice.FieldID, field.TypeInt))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NoticeDeleteOne is the builder for deleting a single Notice entity.
type NoticeDeleteOne struct {
	nd *NoticeDelete
}

// Where appends a list predicates to the NoticeDelete builder.
func (ndo *NoticeDeleteOne) Where(ps ...predicate.Notice) *NoticeDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NoticeDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NoticeDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}