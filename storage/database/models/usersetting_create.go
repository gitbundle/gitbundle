// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/usersetting"
)

// UserSettingCreate is the builder for creating a UserSetting entity.
type UserSettingCreate struct {
	config
	mutation *UserSettingMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (usc *UserSettingCreate) SetUserID(i int64) *UserSettingCreate {
	usc.mutation.SetUserID(i)
	return usc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableUserID(i *int64) *UserSettingCreate {
	if i != nil {
		usc.SetUserID(*i)
	}
	return usc
}

// SetSettingKey sets the "setting_key" field.
func (usc *UserSettingCreate) SetSettingKey(s string) *UserSettingCreate {
	usc.mutation.SetSettingKey(s)
	return usc
}

// SetNillableSettingKey sets the "setting_key" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableSettingKey(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetSettingKey(*s)
	}
	return usc
}

// SetSettingValue sets the "setting_value" field.
func (usc *UserSettingCreate) SetSettingValue(s string) *UserSettingCreate {
	usc.mutation.SetSettingValue(s)
	return usc
}

// SetNillableSettingValue sets the "setting_value" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableSettingValue(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetSettingValue(*s)
	}
	return usc
}

// Mutation returns the UserSettingMutation object of the builder.
func (usc *UserSettingCreate) Mutation() *UserSettingMutation {
	return usc.mutation
}

// Save creates the UserSetting in the database.
func (usc *UserSettingCreate) Save(ctx context.Context) (*UserSetting, error) {
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSettingCreate) SaveX(ctx context.Context) *UserSetting {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSettingCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSettingCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSettingCreate) check() error {
	if v, ok := usc.mutation.SettingKey(); ok {
		if err := usersetting.SettingKeyValidator(v); err != nil {
			return &ValidationError{Name: "setting_key", err: fmt.Errorf(`models: validator failed for field "UserSetting.setting_key": %w`, err)}
		}
	}
	return nil
}

func (usc *UserSettingCreate) sqlSave(ctx context.Context) (*UserSetting, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSettingCreate) createSpec() (*UserSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSetting{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersetting.Table, sqlgraph.NewFieldSpec(usersetting.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.UserID(); ok {
		_spec.SetField(usersetting.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := usc.mutation.SettingKey(); ok {
		_spec.SetField(usersetting.FieldSettingKey, field.TypeString, value)
		_node.SettingKey = value
	}
	if value, ok := usc.mutation.SettingValue(); ok {
		_spec.SetField(usersetting.FieldSettingValue, field.TypeString, value)
		_node.SettingValue = value
	}
	return _node, _spec
}

// UserSettingCreateBulk is the builder for creating many UserSetting entities in bulk.
type UserSettingCreateBulk struct {
	config
	err      error
	builders []*UserSettingCreate
}

// Save creates the UserSetting entities in the database.
func (uscb *UserSettingCreateBulk) Save(ctx context.Context) ([]*UserSetting, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSetting, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) SaveX(ctx context.Context) []*UserSetting {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}