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
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
	"github.com/gitbundle/gitbundle/storage/database/models/renamedbranch"
)

// RenamedBranchUpdate is the builder for updating RenamedBranch entities.
type RenamedBranchUpdate struct {
	config
	hooks    []Hook
	mutation *RenamedBranchMutation
}

// Where appends a list predicates to the RenamedBranchUpdate builder.
func (rbu *RenamedBranchUpdate) Where(ps ...predicate.RenamedBranch) *RenamedBranchUpdate {
	rbu.mutation.Where(ps...)
	return rbu
}

// SetRepoID sets the "repo_id" field.
func (rbu *RenamedBranchUpdate) SetRepoID(i int64) *RenamedBranchUpdate {
	rbu.mutation.ResetRepoID()
	rbu.mutation.SetRepoID(i)
	return rbu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (rbu *RenamedBranchUpdate) SetNillableRepoID(i *int64) *RenamedBranchUpdate {
	if i != nil {
		rbu.SetRepoID(*i)
	}
	return rbu
}

// AddRepoID adds i to the "repo_id" field.
func (rbu *RenamedBranchUpdate) AddRepoID(i int64) *RenamedBranchUpdate {
	rbu.mutation.AddRepoID(i)
	return rbu
}

// SetFrom sets the "from" field.
func (rbu *RenamedBranchUpdate) SetFrom(s string) *RenamedBranchUpdate {
	rbu.mutation.SetFrom(s)
	return rbu
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (rbu *RenamedBranchUpdate) SetNillableFrom(s *string) *RenamedBranchUpdate {
	if s != nil {
		rbu.SetFrom(*s)
	}
	return rbu
}

// ClearFrom clears the value of the "from" field.
func (rbu *RenamedBranchUpdate) ClearFrom() *RenamedBranchUpdate {
	rbu.mutation.ClearFrom()
	return rbu
}

// SetTo sets the "to" field.
func (rbu *RenamedBranchUpdate) SetTo(s string) *RenamedBranchUpdate {
	rbu.mutation.SetTo(s)
	return rbu
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (rbu *RenamedBranchUpdate) SetNillableTo(s *string) *RenamedBranchUpdate {
	if s != nil {
		rbu.SetTo(*s)
	}
	return rbu
}

// ClearTo clears the value of the "to" field.
func (rbu *RenamedBranchUpdate) ClearTo() *RenamedBranchUpdate {
	rbu.mutation.ClearTo()
	return rbu
}

// SetCreatedAt sets the "created_at" field.
func (rbu *RenamedBranchUpdate) SetCreatedAt(t time.Time) *RenamedBranchUpdate {
	rbu.mutation.SetCreatedAt(t)
	return rbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rbu *RenamedBranchUpdate) SetNillableCreatedAt(t *time.Time) *RenamedBranchUpdate {
	if t != nil {
		rbu.SetCreatedAt(*t)
	}
	return rbu
}

// Mutation returns the RenamedBranchMutation object of the builder.
func (rbu *RenamedBranchUpdate) Mutation() *RenamedBranchMutation {
	return rbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rbu *RenamedBranchUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rbu.sqlSave, rbu.mutation, rbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rbu *RenamedBranchUpdate) SaveX(ctx context.Context) int {
	affected, err := rbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rbu *RenamedBranchUpdate) Exec(ctx context.Context) error {
	_, err := rbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbu *RenamedBranchUpdate) ExecX(ctx context.Context) {
	if err := rbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rbu *RenamedBranchUpdate) check() error {
	if v, ok := rbu.mutation.From(); ok {
		if err := renamedbranch.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`models: validator failed for field "RenamedBranch.from": %w`, err)}
		}
	}
	if v, ok := rbu.mutation.To(); ok {
		if err := renamedbranch.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`models: validator failed for field "RenamedBranch.to": %w`, err)}
		}
	}
	return nil
}

func (rbu *RenamedBranchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(renamedbranch.Table, renamedbranch.Columns, sqlgraph.NewFieldSpec(renamedbranch.FieldID, field.TypeInt))
	if ps := rbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rbu.mutation.RepoID(); ok {
		_spec.SetField(renamedbranch.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := rbu.mutation.AddedRepoID(); ok {
		_spec.AddField(renamedbranch.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := rbu.mutation.From(); ok {
		_spec.SetField(renamedbranch.FieldFrom, field.TypeString, value)
	}
	if rbu.mutation.FromCleared() {
		_spec.ClearField(renamedbranch.FieldFrom, field.TypeString)
	}
	if value, ok := rbu.mutation.To(); ok {
		_spec.SetField(renamedbranch.FieldTo, field.TypeString, value)
	}
	if rbu.mutation.ToCleared() {
		_spec.ClearField(renamedbranch.FieldTo, field.TypeString)
	}
	if value, ok := rbu.mutation.CreatedAt(); ok {
		_spec.SetField(renamedbranch.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{renamedbranch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rbu.mutation.done = true
	return n, nil
}

// RenamedBranchUpdateOne is the builder for updating a single RenamedBranch entity.
type RenamedBranchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RenamedBranchMutation
}

// SetRepoID sets the "repo_id" field.
func (rbuo *RenamedBranchUpdateOne) SetRepoID(i int64) *RenamedBranchUpdateOne {
	rbuo.mutation.ResetRepoID()
	rbuo.mutation.SetRepoID(i)
	return rbuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (rbuo *RenamedBranchUpdateOne) SetNillableRepoID(i *int64) *RenamedBranchUpdateOne {
	if i != nil {
		rbuo.SetRepoID(*i)
	}
	return rbuo
}

// AddRepoID adds i to the "repo_id" field.
func (rbuo *RenamedBranchUpdateOne) AddRepoID(i int64) *RenamedBranchUpdateOne {
	rbuo.mutation.AddRepoID(i)
	return rbuo
}

// SetFrom sets the "from" field.
func (rbuo *RenamedBranchUpdateOne) SetFrom(s string) *RenamedBranchUpdateOne {
	rbuo.mutation.SetFrom(s)
	return rbuo
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (rbuo *RenamedBranchUpdateOne) SetNillableFrom(s *string) *RenamedBranchUpdateOne {
	if s != nil {
		rbuo.SetFrom(*s)
	}
	return rbuo
}

// ClearFrom clears the value of the "from" field.
func (rbuo *RenamedBranchUpdateOne) ClearFrom() *RenamedBranchUpdateOne {
	rbuo.mutation.ClearFrom()
	return rbuo
}

// SetTo sets the "to" field.
func (rbuo *RenamedBranchUpdateOne) SetTo(s string) *RenamedBranchUpdateOne {
	rbuo.mutation.SetTo(s)
	return rbuo
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (rbuo *RenamedBranchUpdateOne) SetNillableTo(s *string) *RenamedBranchUpdateOne {
	if s != nil {
		rbuo.SetTo(*s)
	}
	return rbuo
}

// ClearTo clears the value of the "to" field.
func (rbuo *RenamedBranchUpdateOne) ClearTo() *RenamedBranchUpdateOne {
	rbuo.mutation.ClearTo()
	return rbuo
}

// SetCreatedAt sets the "created_at" field.
func (rbuo *RenamedBranchUpdateOne) SetCreatedAt(t time.Time) *RenamedBranchUpdateOne {
	rbuo.mutation.SetCreatedAt(t)
	return rbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rbuo *RenamedBranchUpdateOne) SetNillableCreatedAt(t *time.Time) *RenamedBranchUpdateOne {
	if t != nil {
		rbuo.SetCreatedAt(*t)
	}
	return rbuo
}

// Mutation returns the RenamedBranchMutation object of the builder.
func (rbuo *RenamedBranchUpdateOne) Mutation() *RenamedBranchMutation {
	return rbuo.mutation
}

// Where appends a list predicates to the RenamedBranchUpdate builder.
func (rbuo *RenamedBranchUpdateOne) Where(ps ...predicate.RenamedBranch) *RenamedBranchUpdateOne {
	rbuo.mutation.Where(ps...)
	return rbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rbuo *RenamedBranchUpdateOne) Select(field string, fields ...string) *RenamedBranchUpdateOne {
	rbuo.fields = append([]string{field}, fields...)
	return rbuo
}

// Save executes the query and returns the updated RenamedBranch entity.
func (rbuo *RenamedBranchUpdateOne) Save(ctx context.Context) (*RenamedBranch, error) {
	return withHooks(ctx, rbuo.sqlSave, rbuo.mutation, rbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rbuo *RenamedBranchUpdateOne) SaveX(ctx context.Context) *RenamedBranch {
	node, err := rbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rbuo *RenamedBranchUpdateOne) Exec(ctx context.Context) error {
	_, err := rbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbuo *RenamedBranchUpdateOne) ExecX(ctx context.Context) {
	if err := rbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rbuo *RenamedBranchUpdateOne) check() error {
	if v, ok := rbuo.mutation.From(); ok {
		if err := renamedbranch.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`models: validator failed for field "RenamedBranch.from": %w`, err)}
		}
	}
	if v, ok := rbuo.mutation.To(); ok {
		if err := renamedbranch.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`models: validator failed for field "RenamedBranch.to": %w`, err)}
		}
	}
	return nil
}

func (rbuo *RenamedBranchUpdateOne) sqlSave(ctx context.Context) (_node *RenamedBranch, err error) {
	if err := rbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(renamedbranch.Table, renamedbranch.Columns, sqlgraph.NewFieldSpec(renamedbranch.FieldID, field.TypeInt))
	id, ok := rbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "RenamedBranch.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, renamedbranch.FieldID)
		for _, f := range fields {
			if !renamedbranch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != renamedbranch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rbuo.mutation.RepoID(); ok {
		_spec.SetField(renamedbranch.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := rbuo.mutation.AddedRepoID(); ok {
		_spec.AddField(renamedbranch.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := rbuo.mutation.From(); ok {
		_spec.SetField(renamedbranch.FieldFrom, field.TypeString, value)
	}
	if rbuo.mutation.FromCleared() {
		_spec.ClearField(renamedbranch.FieldFrom, field.TypeString)
	}
	if value, ok := rbuo.mutation.To(); ok {
		_spec.SetField(renamedbranch.FieldTo, field.TypeString, value)
	}
	if rbuo.mutation.ToCleared() {
		_spec.ClearField(renamedbranch.FieldTo, field.TypeString)
	}
	if value, ok := rbuo.mutation.CreatedAt(); ok {
		_spec.SetField(renamedbranch.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &RenamedBranch{config: rbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{renamedbranch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rbuo.mutation.done = true
	return _node, nil
}