// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/notice"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// NoticeUpdate is the builder for updating Notice entities.
type NoticeUpdate struct {
	config
	hooks    []Hook
	mutation *NoticeMutation
}

// Where appends a list predicates to the NoticeUpdate builder.
func (nu *NoticeUpdate) Where(ps ...predicate.Notice) *NoticeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetType sets the "type" field.
func (nu *NoticeUpdate) SetType(i int) *NoticeUpdate {
	nu.mutation.ResetType()
	nu.mutation.SetType(i)
	return nu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nu *NoticeUpdate) SetNillableType(i *int) *NoticeUpdate {
	if i != nil {
		nu.SetType(*i)
	}
	return nu
}

// AddType adds i to the "type" field.
func (nu *NoticeUpdate) AddType(i int) *NoticeUpdate {
	nu.mutation.AddType(i)
	return nu
}

// ClearType clears the value of the "type" field.
func (nu *NoticeUpdate) ClearType() *NoticeUpdate {
	nu.mutation.ClearType()
	return nu
}

// SetDescription sets the "description" field.
func (nu *NoticeUpdate) SetDescription(s string) *NoticeUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nu *NoticeUpdate) SetNillableDescription(s *string) *NoticeUpdate {
	if s != nil {
		nu.SetDescription(*s)
	}
	return nu
}

// ClearDescription clears the value of the "description" field.
func (nu *NoticeUpdate) ClearDescription() *NoticeUpdate {
	nu.mutation.ClearDescription()
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NoticeUpdate) SetCreatedAt(t time.Time) *NoticeUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NoticeUpdate) SetNillableCreatedAt(t *time.Time) *NoticeUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// Mutation returns the NoticeMutation object of the builder.
func (nu *NoticeUpdate) Mutation() *NoticeMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NoticeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NoticeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NoticeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NoticeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NoticeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notice.Table, notice.Columns, sqlgraph.NewFieldSpec(notice.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.SetField(notice.FieldType, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedType(); ok {
		_spec.AddField(notice.FieldType, field.TypeInt, value)
	}
	if nu.mutation.TypeCleared() {
		_spec.ClearField(notice.FieldType, field.TypeInt)
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.SetField(notice.FieldDescription, field.TypeString, value)
	}
	if nu.mutation.DescriptionCleared() {
		_spec.ClearField(notice.FieldDescription, field.TypeString)
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(notice.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NoticeUpdateOne is the builder for updating a single Notice entity.
type NoticeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NoticeMutation
}

// SetType sets the "type" field.
func (nuo *NoticeUpdateOne) SetType(i int) *NoticeUpdateOne {
	nuo.mutation.ResetType()
	nuo.mutation.SetType(i)
	return nuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nuo *NoticeUpdateOne) SetNillableType(i *int) *NoticeUpdateOne {
	if i != nil {
		nuo.SetType(*i)
	}
	return nuo
}

// AddType adds i to the "type" field.
func (nuo *NoticeUpdateOne) AddType(i int) *NoticeUpdateOne {
	nuo.mutation.AddType(i)
	return nuo
}

// ClearType clears the value of the "type" field.
func (nuo *NoticeUpdateOne) ClearType() *NoticeUpdateOne {
	nuo.mutation.ClearType()
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NoticeUpdateOne) SetDescription(s string) *NoticeUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nuo *NoticeUpdateOne) SetNillableDescription(s *string) *NoticeUpdateOne {
	if s != nil {
		nuo.SetDescription(*s)
	}
	return nuo
}

// ClearDescription clears the value of the "description" field.
func (nuo *NoticeUpdateOne) ClearDescription() *NoticeUpdateOne {
	nuo.mutation.ClearDescription()
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NoticeUpdateOne) SetCreatedAt(t time.Time) *NoticeUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NoticeUpdateOne) SetNillableCreatedAt(t *time.Time) *NoticeUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// Mutation returns the NoticeMutation object of the builder.
func (nuo *NoticeUpdateOne) Mutation() *NoticeMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NoticeUpdate builder.
func (nuo *NoticeUpdateOne) Where(ps ...predicate.Notice) *NoticeUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NoticeUpdateOne) Select(field string, fields ...string) *NoticeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notice entity.
func (nuo *NoticeUpdateOne) Save(ctx context.Context) (*Notice, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NoticeUpdateOne) SaveX(ctx context.Context) *Notice {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NoticeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NoticeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NoticeUpdateOne) sqlSave(ctx context.Context) (_node *Notice, err error) {
	_spec := sqlgraph.NewUpdateSpec(notice.Table, notice.Columns, sqlgraph.NewFieldSpec(notice.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Notice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notice.FieldID)
		for _, f := range fields {
			if !notice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != notice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.SetField(notice.FieldType, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedType(); ok {
		_spec.AddField(notice.FieldType, field.TypeInt, value)
	}
	if nuo.mutation.TypeCleared() {
		_spec.ClearField(notice.FieldType, field.TypeInt)
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.SetField(notice.FieldDescription, field.TypeString, value)
	}
	if nuo.mutation.DescriptionCleared() {
		_spec.ClearField(notice.FieldDescription, field.TypeString)
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(notice.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Notice{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}