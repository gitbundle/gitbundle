// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/commitstatusindex"
)

// CommitStatusIndexCreate is the builder for creating a CommitStatusIndex entity.
type CommitStatusIndexCreate struct {
	config
	mutation *CommitStatusIndexMutation
	hooks    []Hook
}

// SetRepoID sets the "repo_id" field.
func (csic *CommitStatusIndexCreate) SetRepoID(i int64) *CommitStatusIndexCreate {
	csic.mutation.SetRepoID(i)
	return csic
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (csic *CommitStatusIndexCreate) SetNillableRepoID(i *int64) *CommitStatusIndexCreate {
	if i != nil {
		csic.SetRepoID(*i)
	}
	return csic
}

// SetSha sets the "sha" field.
func (csic *CommitStatusIndexCreate) SetSha(s string) *CommitStatusIndexCreate {
	csic.mutation.SetSha(s)
	return csic
}

// SetNillableSha sets the "sha" field if the given value is not nil.
func (csic *CommitStatusIndexCreate) SetNillableSha(s *string) *CommitStatusIndexCreate {
	if s != nil {
		csic.SetSha(*s)
	}
	return csic
}

// SetMaxIndex sets the "max_index" field.
func (csic *CommitStatusIndexCreate) SetMaxIndex(i int64) *CommitStatusIndexCreate {
	csic.mutation.SetMaxIndex(i)
	return csic
}

// SetNillableMaxIndex sets the "max_index" field if the given value is not nil.
func (csic *CommitStatusIndexCreate) SetNillableMaxIndex(i *int64) *CommitStatusIndexCreate {
	if i != nil {
		csic.SetMaxIndex(*i)
	}
	return csic
}

// Mutation returns the CommitStatusIndexMutation object of the builder.
func (csic *CommitStatusIndexCreate) Mutation() *CommitStatusIndexMutation {
	return csic.mutation
}

// Save creates the CommitStatusIndex in the database.
func (csic *CommitStatusIndexCreate) Save(ctx context.Context) (*CommitStatusIndex, error) {
	return withHooks(ctx, csic.sqlSave, csic.mutation, csic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (csic *CommitStatusIndexCreate) SaveX(ctx context.Context) *CommitStatusIndex {
	v, err := csic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csic *CommitStatusIndexCreate) Exec(ctx context.Context) error {
	_, err := csic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csic *CommitStatusIndexCreate) ExecX(ctx context.Context) {
	if err := csic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csic *CommitStatusIndexCreate) check() error {
	if v, ok := csic.mutation.Sha(); ok {
		if err := commitstatusindex.ShaValidator(v); err != nil {
			return &ValidationError{Name: "sha", err: fmt.Errorf(`models: validator failed for field "CommitStatusIndex.sha": %w`, err)}
		}
	}
	return nil
}

func (csic *CommitStatusIndexCreate) sqlSave(ctx context.Context) (*CommitStatusIndex, error) {
	if err := csic.check(); err != nil {
		return nil, err
	}
	_node, _spec := csic.createSpec()
	if err := sqlgraph.CreateNode(ctx, csic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	csic.mutation.id = &_node.ID
	csic.mutation.done = true
	return _node, nil
}

func (csic *CommitStatusIndexCreate) createSpec() (*CommitStatusIndex, *sqlgraph.CreateSpec) {
	var (
		_node = &CommitStatusIndex{config: csic.config}
		_spec = sqlgraph.NewCreateSpec(commitstatusindex.Table, sqlgraph.NewFieldSpec(commitstatusindex.FieldID, field.TypeInt))
	)
	if value, ok := csic.mutation.RepoID(); ok {
		_spec.SetField(commitstatusindex.FieldRepoID, field.TypeInt64, value)
		_node.RepoID = value
	}
	if value, ok := csic.mutation.Sha(); ok {
		_spec.SetField(commitstatusindex.FieldSha, field.TypeString, value)
		_node.Sha = value
	}
	if value, ok := csic.mutation.MaxIndex(); ok {
		_spec.SetField(commitstatusindex.FieldMaxIndex, field.TypeInt64, value)
		_node.MaxIndex = value
	}
	return _node, _spec
}

// CommitStatusIndexCreateBulk is the builder for creating many CommitStatusIndex entities in bulk.
type CommitStatusIndexCreateBulk struct {
	config
	err      error
	builders []*CommitStatusIndexCreate
}

// Save creates the CommitStatusIndex entities in the database.
func (csicb *CommitStatusIndexCreateBulk) Save(ctx context.Context) ([]*CommitStatusIndex, error) {
	if csicb.err != nil {
		return nil, csicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(csicb.builders))
	nodes := make([]*CommitStatusIndex, len(csicb.builders))
	mutators := make([]Mutator, len(csicb.builders))
	for i := range csicb.builders {
		func(i int, root context.Context) {
			builder := csicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommitStatusIndexMutation)
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
					_, err = mutators[i+1].Mutate(root, csicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, csicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, csicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (csicb *CommitStatusIndexCreateBulk) SaveX(ctx context.Context) []*CommitStatusIndex {
	v, err := csicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csicb *CommitStatusIndexCreateBulk) Exec(ctx context.Context) error {
	_, err := csicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csicb *CommitStatusIndexCreateBulk) ExecX(ctx context.Context) {
	if err := csicb.Exec(ctx); err != nil {
		panic(err)
	}
}