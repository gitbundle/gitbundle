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
	"github.com/gitbundle/gitbundle/store/models/lfsmetaobject"
	"github.com/gitbundle/gitbundle/store/models/predicate"
)

// LfsMetaObjectUpdate is the builder for updating LfsMetaObject entities.
type LfsMetaObjectUpdate struct {
	config
	hooks    []Hook
	mutation *LfsMetaObjectMutation
}

// Where appends a list predicates to the LfsMetaObjectUpdate builder.
func (lmou *LfsMetaObjectUpdate) Where(ps ...predicate.LfsMetaObject) *LfsMetaObjectUpdate {
	lmou.mutation.Where(ps...)
	return lmou
}

// SetOid sets the "oid" field.
func (lmou *LfsMetaObjectUpdate) SetOid(s string) *LfsMetaObjectUpdate {
	lmou.mutation.SetOid(s)
	return lmou
}

// SetNillableOid sets the "oid" field if the given value is not nil.
func (lmou *LfsMetaObjectUpdate) SetNillableOid(s *string) *LfsMetaObjectUpdate {
	if s != nil {
		lmou.SetOid(*s)
	}
	return lmou
}

// SetSize sets the "size" field.
func (lmou *LfsMetaObjectUpdate) SetSize(i int64) *LfsMetaObjectUpdate {
	lmou.mutation.ResetSize()
	lmou.mutation.SetSize(i)
	return lmou
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (lmou *LfsMetaObjectUpdate) SetNillableSize(i *int64) *LfsMetaObjectUpdate {
	if i != nil {
		lmou.SetSize(*i)
	}
	return lmou
}

// AddSize adds i to the "size" field.
func (lmou *LfsMetaObjectUpdate) AddSize(i int64) *LfsMetaObjectUpdate {
	lmou.mutation.AddSize(i)
	return lmou
}

// SetRepoID sets the "repo_id" field.
func (lmou *LfsMetaObjectUpdate) SetRepoID(i int64) *LfsMetaObjectUpdate {
	lmou.mutation.ResetRepoID()
	lmou.mutation.SetRepoID(i)
	return lmou
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (lmou *LfsMetaObjectUpdate) SetNillableRepoID(i *int64) *LfsMetaObjectUpdate {
	if i != nil {
		lmou.SetRepoID(*i)
	}
	return lmou
}

// AddRepoID adds i to the "repo_id" field.
func (lmou *LfsMetaObjectUpdate) AddRepoID(i int64) *LfsMetaObjectUpdate {
	lmou.mutation.AddRepoID(i)
	return lmou
}

// SetCreatedAt sets the "created_at" field.
func (lmou *LfsMetaObjectUpdate) SetCreatedAt(t time.Time) *LfsMetaObjectUpdate {
	lmou.mutation.SetCreatedAt(t)
	return lmou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lmou *LfsMetaObjectUpdate) SetNillableCreatedAt(t *time.Time) *LfsMetaObjectUpdate {
	if t != nil {
		lmou.SetCreatedAt(*t)
	}
	return lmou
}

// Mutation returns the LfsMetaObjectMutation object of the builder.
func (lmou *LfsMetaObjectUpdate) Mutation() *LfsMetaObjectMutation {
	return lmou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmou *LfsMetaObjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lmou.sqlSave, lmou.mutation, lmou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmou *LfsMetaObjectUpdate) SaveX(ctx context.Context) int {
	affected, err := lmou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmou *LfsMetaObjectUpdate) Exec(ctx context.Context) error {
	_, err := lmou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmou *LfsMetaObjectUpdate) ExecX(ctx context.Context) {
	if err := lmou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmou *LfsMetaObjectUpdate) check() error {
	if v, ok := lmou.mutation.Oid(); ok {
		if err := lfsmetaobject.OidValidator(v); err != nil {
			return &ValidationError{Name: "oid", err: fmt.Errorf(`models: validator failed for field "LfsMetaObject.oid": %w`, err)}
		}
	}
	return nil
}

func (lmou *LfsMetaObjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lmou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lfsmetaobject.Table, lfsmetaobject.Columns, sqlgraph.NewFieldSpec(lfsmetaobject.FieldID, field.TypeInt))
	if ps := lmou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmou.mutation.Oid(); ok {
		_spec.SetField(lfsmetaobject.FieldOid, field.TypeString, value)
	}
	if value, ok := lmou.mutation.Size(); ok {
		_spec.SetField(lfsmetaobject.FieldSize, field.TypeInt64, value)
	}
	if value, ok := lmou.mutation.AddedSize(); ok {
		_spec.AddField(lfsmetaobject.FieldSize, field.TypeInt64, value)
	}
	if value, ok := lmou.mutation.RepoID(); ok {
		_spec.SetField(lfsmetaobject.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := lmou.mutation.AddedRepoID(); ok {
		_spec.AddField(lfsmetaobject.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := lmou.mutation.CreatedAt(); ok {
		_spec.SetField(lfsmetaobject.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lmou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lfsmetaobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lmou.mutation.done = true
	return n, nil
}

// LfsMetaObjectUpdateOne is the builder for updating a single LfsMetaObject entity.
type LfsMetaObjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LfsMetaObjectMutation
}

// SetOid sets the "oid" field.
func (lmouo *LfsMetaObjectUpdateOne) SetOid(s string) *LfsMetaObjectUpdateOne {
	lmouo.mutation.SetOid(s)
	return lmouo
}

// SetNillableOid sets the "oid" field if the given value is not nil.
func (lmouo *LfsMetaObjectUpdateOne) SetNillableOid(s *string) *LfsMetaObjectUpdateOne {
	if s != nil {
		lmouo.SetOid(*s)
	}
	return lmouo
}

// SetSize sets the "size" field.
func (lmouo *LfsMetaObjectUpdateOne) SetSize(i int64) *LfsMetaObjectUpdateOne {
	lmouo.mutation.ResetSize()
	lmouo.mutation.SetSize(i)
	return lmouo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (lmouo *LfsMetaObjectUpdateOne) SetNillableSize(i *int64) *LfsMetaObjectUpdateOne {
	if i != nil {
		lmouo.SetSize(*i)
	}
	return lmouo
}

// AddSize adds i to the "size" field.
func (lmouo *LfsMetaObjectUpdateOne) AddSize(i int64) *LfsMetaObjectUpdateOne {
	lmouo.mutation.AddSize(i)
	return lmouo
}

// SetRepoID sets the "repo_id" field.
func (lmouo *LfsMetaObjectUpdateOne) SetRepoID(i int64) *LfsMetaObjectUpdateOne {
	lmouo.mutation.ResetRepoID()
	lmouo.mutation.SetRepoID(i)
	return lmouo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (lmouo *LfsMetaObjectUpdateOne) SetNillableRepoID(i *int64) *LfsMetaObjectUpdateOne {
	if i != nil {
		lmouo.SetRepoID(*i)
	}
	return lmouo
}

// AddRepoID adds i to the "repo_id" field.
func (lmouo *LfsMetaObjectUpdateOne) AddRepoID(i int64) *LfsMetaObjectUpdateOne {
	lmouo.mutation.AddRepoID(i)
	return lmouo
}

// SetCreatedAt sets the "created_at" field.
func (lmouo *LfsMetaObjectUpdateOne) SetCreatedAt(t time.Time) *LfsMetaObjectUpdateOne {
	lmouo.mutation.SetCreatedAt(t)
	return lmouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lmouo *LfsMetaObjectUpdateOne) SetNillableCreatedAt(t *time.Time) *LfsMetaObjectUpdateOne {
	if t != nil {
		lmouo.SetCreatedAt(*t)
	}
	return lmouo
}

// Mutation returns the LfsMetaObjectMutation object of the builder.
func (lmouo *LfsMetaObjectUpdateOne) Mutation() *LfsMetaObjectMutation {
	return lmouo.mutation
}

// Where appends a list predicates to the LfsMetaObjectUpdate builder.
func (lmouo *LfsMetaObjectUpdateOne) Where(ps ...predicate.LfsMetaObject) *LfsMetaObjectUpdateOne {
	lmouo.mutation.Where(ps...)
	return lmouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmouo *LfsMetaObjectUpdateOne) Select(field string, fields ...string) *LfsMetaObjectUpdateOne {
	lmouo.fields = append([]string{field}, fields...)
	return lmouo
}

// Save executes the query and returns the updated LfsMetaObject entity.
func (lmouo *LfsMetaObjectUpdateOne) Save(ctx context.Context) (*LfsMetaObject, error) {
	return withHooks(ctx, lmouo.sqlSave, lmouo.mutation, lmouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmouo *LfsMetaObjectUpdateOne) SaveX(ctx context.Context) *LfsMetaObject {
	node, err := lmouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmouo *LfsMetaObjectUpdateOne) Exec(ctx context.Context) error {
	_, err := lmouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmouo *LfsMetaObjectUpdateOne) ExecX(ctx context.Context) {
	if err := lmouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmouo *LfsMetaObjectUpdateOne) check() error {
	if v, ok := lmouo.mutation.Oid(); ok {
		if err := lfsmetaobject.OidValidator(v); err != nil {
			return &ValidationError{Name: "oid", err: fmt.Errorf(`models: validator failed for field "LfsMetaObject.oid": %w`, err)}
		}
	}
	return nil
}

func (lmouo *LfsMetaObjectUpdateOne) sqlSave(ctx context.Context) (_node *LfsMetaObject, err error) {
	if err := lmouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lfsmetaobject.Table, lfsmetaobject.Columns, sqlgraph.NewFieldSpec(lfsmetaobject.FieldID, field.TypeInt))
	id, ok := lmouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "LfsMetaObject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lfsmetaobject.FieldID)
		for _, f := range fields {
			if !lfsmetaobject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != lfsmetaobject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmouo.mutation.Oid(); ok {
		_spec.SetField(lfsmetaobject.FieldOid, field.TypeString, value)
	}
	if value, ok := lmouo.mutation.Size(); ok {
		_spec.SetField(lfsmetaobject.FieldSize, field.TypeInt64, value)
	}
	if value, ok := lmouo.mutation.AddedSize(); ok {
		_spec.AddField(lfsmetaobject.FieldSize, field.TypeInt64, value)
	}
	if value, ok := lmouo.mutation.RepoID(); ok {
		_spec.SetField(lfsmetaobject.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := lmouo.mutation.AddedRepoID(); ok {
		_spec.AddField(lfsmetaobject.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := lmouo.mutation.CreatedAt(); ok {
		_spec.SetField(lfsmetaobject.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &LfsMetaObject{config: lmouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lfsmetaobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lmouo.mutation.done = true
	return _node, nil
}