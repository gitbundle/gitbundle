// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/orguser"
)

// OrgUserCreate is the builder for creating a OrgUser entity.
type OrgUserCreate struct {
	config
	mutation *OrgUserMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ouc *OrgUserCreate) SetUserID(i int64) *OrgUserCreate {
	ouc.mutation.SetUserID(i)
	return ouc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableUserID(i *int64) *OrgUserCreate {
	if i != nil {
		ouc.SetUserID(*i)
	}
	return ouc
}

// SetOrgID sets the "org_id" field.
func (ouc *OrgUserCreate) SetOrgID(i int64) *OrgUserCreate {
	ouc.mutation.SetOrgID(i)
	return ouc
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableOrgID(i *int64) *OrgUserCreate {
	if i != nil {
		ouc.SetOrgID(*i)
	}
	return ouc
}

// SetIsPublic sets the "is_public" field.
func (ouc *OrgUserCreate) SetIsPublic(b bool) *OrgUserCreate {
	ouc.mutation.SetIsPublic(b)
	return ouc
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableIsPublic(b *bool) *OrgUserCreate {
	if b != nil {
		ouc.SetIsPublic(*b)
	}
	return ouc
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouc *OrgUserCreate) Mutation() *OrgUserMutation {
	return ouc.mutation
}

// Save creates the OrgUser in the database.
func (ouc *OrgUserCreate) Save(ctx context.Context) (*OrgUser, error) {
	return withHooks(ctx, ouc.sqlSave, ouc.mutation, ouc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ouc *OrgUserCreate) SaveX(ctx context.Context) *OrgUser {
	v, err := ouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ouc *OrgUserCreate) Exec(ctx context.Context) error {
	_, err := ouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouc *OrgUserCreate) ExecX(ctx context.Context) {
	if err := ouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouc *OrgUserCreate) check() error {
	return nil
}

func (ouc *OrgUserCreate) sqlSave(ctx context.Context) (*OrgUser, error) {
	if err := ouc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ouc.mutation.id = &_node.ID
	ouc.mutation.done = true
	return _node, nil
}

func (ouc *OrgUserCreate) createSpec() (*OrgUser, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgUser{config: ouc.config}
		_spec = sqlgraph.NewCreateSpec(orguser.Table, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	)
	if value, ok := ouc.mutation.UserID(); ok {
		_spec.SetField(orguser.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := ouc.mutation.OrgID(); ok {
		_spec.SetField(orguser.FieldOrgID, field.TypeInt64, value)
		_node.OrgID = value
	}
	if value, ok := ouc.mutation.IsPublic(); ok {
		_spec.SetField(orguser.FieldIsPublic, field.TypeBool, value)
		_node.IsPublic = value
	}
	return _node, _spec
}

// OrgUserCreateBulk is the builder for creating many OrgUser entities in bulk.
type OrgUserCreateBulk struct {
	config
	err      error
	builders []*OrgUserCreate
}

// Save creates the OrgUser entities in the database.
func (oucb *OrgUserCreateBulk) Save(ctx context.Context) ([]*OrgUser, error) {
	if oucb.err != nil {
		return nil, oucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oucb.builders))
	nodes := make([]*OrgUser, len(oucb.builders))
	mutators := make([]Mutator, len(oucb.builders))
	for i := range oucb.builders {
		func(i int, root context.Context) {
			builder := oucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgUserMutation)
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
					_, err = mutators[i+1].Mutate(root, oucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) SaveX(ctx context.Context) []*OrgUser {
	v, err := oucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oucb *OrgUserCreateBulk) Exec(ctx context.Context) error {
	_, err := oucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) ExecX(ctx context.Context) {
	if err := oucb.Exec(ctx); err != nil {
		panic(err)
	}
}