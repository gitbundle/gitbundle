// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/emailhash"
)

// EmailHashCreate is the builder for creating a EmailHash entity.
type EmailHashCreate struct {
	config
	mutation *EmailHashMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (ehc *EmailHashCreate) SetHash(s string) *EmailHashCreate {
	ehc.mutation.SetHash(s)
	return ehc
}

// SetEmail sets the "email" field.
func (ehc *EmailHashCreate) SetEmail(s string) *EmailHashCreate {
	ehc.mutation.SetEmail(s)
	return ehc
}

// Mutation returns the EmailHashMutation object of the builder.
func (ehc *EmailHashCreate) Mutation() *EmailHashMutation {
	return ehc.mutation
}

// Save creates the EmailHash in the database.
func (ehc *EmailHashCreate) Save(ctx context.Context) (*EmailHash, error) {
	return withHooks(ctx, ehc.sqlSave, ehc.mutation, ehc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ehc *EmailHashCreate) SaveX(ctx context.Context) *EmailHash {
	v, err := ehc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehc *EmailHashCreate) Exec(ctx context.Context) error {
	_, err := ehc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehc *EmailHashCreate) ExecX(ctx context.Context) {
	if err := ehc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ehc *EmailHashCreate) check() error {
	if _, ok := ehc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`models: missing required field "EmailHash.hash"`)}
	}
	if v, ok := ehc.mutation.Hash(); ok {
		if err := emailhash.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`models: validator failed for field "EmailHash.hash": %w`, err)}
		}
	}
	if _, ok := ehc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`models: missing required field "EmailHash.email"`)}
	}
	if v, ok := ehc.mutation.Email(); ok {
		if err := emailhash.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "EmailHash.email": %w`, err)}
		}
	}
	return nil
}

func (ehc *EmailHashCreate) sqlSave(ctx context.Context) (*EmailHash, error) {
	if err := ehc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ehc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ehc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ehc.mutation.id = &_node.ID
	ehc.mutation.done = true
	return _node, nil
}

func (ehc *EmailHashCreate) createSpec() (*EmailHash, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailHash{config: ehc.config}
		_spec = sqlgraph.NewCreateSpec(emailhash.Table, sqlgraph.NewFieldSpec(emailhash.FieldID, field.TypeInt))
	)
	if value, ok := ehc.mutation.Hash(); ok {
		_spec.SetField(emailhash.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := ehc.mutation.Email(); ok {
		_spec.SetField(emailhash.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	return _node, _spec
}

// EmailHashCreateBulk is the builder for creating many EmailHash entities in bulk.
type EmailHashCreateBulk struct {
	config
	err      error
	builders []*EmailHashCreate
}

// Save creates the EmailHash entities in the database.
func (ehcb *EmailHashCreateBulk) Save(ctx context.Context) ([]*EmailHash, error) {
	if ehcb.err != nil {
		return nil, ehcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ehcb.builders))
	nodes := make([]*EmailHash, len(ehcb.builders))
	mutators := make([]Mutator, len(ehcb.builders))
	for i := range ehcb.builders {
		func(i int, root context.Context) {
			builder := ehcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailHashMutation)
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
					_, err = mutators[i+1].Mutate(root, ehcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ehcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ehcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ehcb *EmailHashCreateBulk) SaveX(ctx context.Context) []*EmailHash {
	v, err := ehcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehcb *EmailHashCreateBulk) Exec(ctx context.Context) error {
	_, err := ehcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehcb *EmailHashCreateBulk) ExecX(ctx context.Context) {
	if err := ehcb.Exec(ctx); err != nil {
		panic(err)
	}
}