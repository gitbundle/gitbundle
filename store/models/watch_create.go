// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/watch"
)

// WatchCreate is the builder for creating a Watch entity.
type WatchCreate struct {
	config
	mutation *WatchMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (wc *WatchCreate) SetUserID(i int64) *WatchCreate {
	wc.mutation.SetUserID(i)
	return wc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wc *WatchCreate) SetNillableUserID(i *int64) *WatchCreate {
	if i != nil {
		wc.SetUserID(*i)
	}
	return wc
}

// SetRepoID sets the "repo_id" field.
func (wc *WatchCreate) SetRepoID(i int64) *WatchCreate {
	wc.mutation.SetRepoID(i)
	return wc
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (wc *WatchCreate) SetNillableRepoID(i *int64) *WatchCreate {
	if i != nil {
		wc.SetRepoID(*i)
	}
	return wc
}

// SetMode sets the "mode" field.
func (wc *WatchCreate) SetMode(i int16) *WatchCreate {
	wc.mutation.SetMode(i)
	return wc
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (wc *WatchCreate) SetNillableMode(i *int16) *WatchCreate {
	if i != nil {
		wc.SetMode(*i)
	}
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WatchCreate) SetCreatedAt(t time.Time) *WatchCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WatchCreate) SetNillableCreatedAt(t *time.Time) *WatchCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WatchCreate) SetUpdatedAt(t time.Time) *WatchCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WatchCreate) SetNillableUpdatedAt(t *time.Time) *WatchCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// Mutation returns the WatchMutation object of the builder.
func (wc *WatchCreate) Mutation() *WatchMutation {
	return wc.mutation
}

// Save creates the Watch in the database.
func (wc *WatchCreate) Save(ctx context.Context) (*Watch, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WatchCreate) SaveX(ctx context.Context) *Watch {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WatchCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WatchCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WatchCreate) defaults() {
	if _, ok := wc.mutation.Mode(); !ok {
		v := watch.DefaultMode
		wc.mutation.SetMode(v)
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := watch.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WatchCreate) check() error {
	if _, ok := wc.mutation.Mode(); !ok {
		return &ValidationError{Name: "mode", err: errors.New(`models: missing required field "Watch.mode"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "Watch.created_at"`)}
	}
	return nil
}

func (wc *WatchCreate) sqlSave(ctx context.Context) (*Watch, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WatchCreate) createSpec() (*Watch, *sqlgraph.CreateSpec) {
	var (
		_node = &Watch{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(watch.Table, sqlgraph.NewFieldSpec(watch.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.UserID(); ok {
		_spec.SetField(watch.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := wc.mutation.RepoID(); ok {
		_spec.SetField(watch.FieldRepoID, field.TypeInt64, value)
		_node.RepoID = value
	}
	if value, ok := wc.mutation.Mode(); ok {
		_spec.SetField(watch.FieldMode, field.TypeInt16, value)
		_node.Mode = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(watch.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(watch.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WatchCreateBulk is the builder for creating many Watch entities in bulk.
type WatchCreateBulk struct {
	config
	err      error
	builders []*WatchCreate
}

// Save creates the Watch entities in the database.
func (wcb *WatchCreateBulk) Save(ctx context.Context) ([]*Watch, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Watch, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WatchMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WatchCreateBulk) SaveX(ctx context.Context) []*Watch {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WatchCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WatchCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}