// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/externalloginuser"
)

// ExternalLoginUserCreate is the builder for creating a ExternalLoginUser entity.
type ExternalLoginUserCreate struct {
	config
	mutation *ExternalLoginUserMutation
	hooks    []Hook
}

// SetExternalID sets the "external_id" field.
func (eluc *ExternalLoginUserCreate) SetExternalID(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetExternalID(s)
	return eluc
}

// SetUserID sets the "user_id" field.
func (eluc *ExternalLoginUserCreate) SetUserID(i int64) *ExternalLoginUserCreate {
	eluc.mutation.SetUserID(i)
	return eluc
}

// SetLoginSourceID sets the "login_source_id" field.
func (eluc *ExternalLoginUserCreate) SetLoginSourceID(i int64) *ExternalLoginUserCreate {
	eluc.mutation.SetLoginSourceID(i)
	return eluc
}

// SetProvider sets the "provider" field.
func (eluc *ExternalLoginUserCreate) SetProvider(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetProvider(s)
	return eluc
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableProvider(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetProvider(*s)
	}
	return eluc
}

// SetEmail sets the "email" field.
func (eluc *ExternalLoginUserCreate) SetEmail(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetEmail(s)
	return eluc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableEmail(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetEmail(*s)
	}
	return eluc
}

// SetName sets the "name" field.
func (eluc *ExternalLoginUserCreate) SetName(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetName(s)
	return eluc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableName(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetName(*s)
	}
	return eluc
}

// SetFirstName sets the "first_name" field.
func (eluc *ExternalLoginUserCreate) SetFirstName(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetFirstName(s)
	return eluc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableFirstName(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetFirstName(*s)
	}
	return eluc
}

// SetLastName sets the "last_name" field.
func (eluc *ExternalLoginUserCreate) SetLastName(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetLastName(s)
	return eluc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableLastName(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetLastName(*s)
	}
	return eluc
}

// SetNickName sets the "nick_name" field.
func (eluc *ExternalLoginUserCreate) SetNickName(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetNickName(s)
	return eluc
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableNickName(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetNickName(*s)
	}
	return eluc
}

// SetDescription sets the "description" field.
func (eluc *ExternalLoginUserCreate) SetDescription(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetDescription(s)
	return eluc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableDescription(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetDescription(*s)
	}
	return eluc
}

// SetAvatarURL sets the "avatar_url" field.
func (eluc *ExternalLoginUserCreate) SetAvatarURL(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetAvatarURL(s)
	return eluc
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableAvatarURL(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetAvatarURL(*s)
	}
	return eluc
}

// SetLocation sets the "location" field.
func (eluc *ExternalLoginUserCreate) SetLocation(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetLocation(s)
	return eluc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableLocation(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetLocation(*s)
	}
	return eluc
}

// SetAccessToken sets the "access_token" field.
func (eluc *ExternalLoginUserCreate) SetAccessToken(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetAccessToken(s)
	return eluc
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableAccessToken(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetAccessToken(*s)
	}
	return eluc
}

// SetAccessTokenSecret sets the "access_token_secret" field.
func (eluc *ExternalLoginUserCreate) SetAccessTokenSecret(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetAccessTokenSecret(s)
	return eluc
}

// SetNillableAccessTokenSecret sets the "access_token_secret" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableAccessTokenSecret(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetAccessTokenSecret(*s)
	}
	return eluc
}

// SetRefreshToken sets the "refresh_token" field.
func (eluc *ExternalLoginUserCreate) SetRefreshToken(s string) *ExternalLoginUserCreate {
	eluc.mutation.SetRefreshToken(s)
	return eluc
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableRefreshToken(s *string) *ExternalLoginUserCreate {
	if s != nil {
		eluc.SetRefreshToken(*s)
	}
	return eluc
}

// SetExpiredAt sets the "expired_at" field.
func (eluc *ExternalLoginUserCreate) SetExpiredAt(t time.Time) *ExternalLoginUserCreate {
	eluc.mutation.SetExpiredAt(t)
	return eluc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (eluc *ExternalLoginUserCreate) SetNillableExpiredAt(t *time.Time) *ExternalLoginUserCreate {
	if t != nil {
		eluc.SetExpiredAt(*t)
	}
	return eluc
}

// Mutation returns the ExternalLoginUserMutation object of the builder.
func (eluc *ExternalLoginUserCreate) Mutation() *ExternalLoginUserMutation {
	return eluc.mutation
}

// Save creates the ExternalLoginUser in the database.
func (eluc *ExternalLoginUserCreate) Save(ctx context.Context) (*ExternalLoginUser, error) {
	return withHooks(ctx, eluc.sqlSave, eluc.mutation, eluc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eluc *ExternalLoginUserCreate) SaveX(ctx context.Context) *ExternalLoginUser {
	v, err := eluc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eluc *ExternalLoginUserCreate) Exec(ctx context.Context) error {
	_, err := eluc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eluc *ExternalLoginUserCreate) ExecX(ctx context.Context) {
	if err := eluc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eluc *ExternalLoginUserCreate) check() error {
	if _, ok := eluc.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New(`models: missing required field "ExternalLoginUser.external_id"`)}
	}
	if v, ok := eluc.mutation.ExternalID(); ok {
		if err := externalloginuser.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.external_id": %w`, err)}
		}
	}
	if _, ok := eluc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`models: missing required field "ExternalLoginUser.user_id"`)}
	}
	if _, ok := eluc.mutation.LoginSourceID(); !ok {
		return &ValidationError{Name: "login_source_id", err: errors.New(`models: missing required field "ExternalLoginUser.login_source_id"`)}
	}
	if v, ok := eluc.mutation.Provider(); ok {
		if err := externalloginuser.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.provider": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.Email(); ok {
		if err := externalloginuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.email": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.Name(); ok {
		if err := externalloginuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.name": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.FirstName(); ok {
		if err := externalloginuser.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.first_name": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.LastName(); ok {
		if err := externalloginuser.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.last_name": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.NickName(); ok {
		if err := externalloginuser.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.nick_name": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.Description(); ok {
		if err := externalloginuser.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.description": %w`, err)}
		}
	}
	if v, ok := eluc.mutation.Location(); ok {
		if err := externalloginuser.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`models: validator failed for field "ExternalLoginUser.location": %w`, err)}
		}
	}
	return nil
}

func (eluc *ExternalLoginUserCreate) sqlSave(ctx context.Context) (*ExternalLoginUser, error) {
	if err := eluc.check(); err != nil {
		return nil, err
	}
	_node, _spec := eluc.createSpec()
	if err := sqlgraph.CreateNode(ctx, eluc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	eluc.mutation.id = &_node.ID
	eluc.mutation.done = true
	return _node, nil
}

func (eluc *ExternalLoginUserCreate) createSpec() (*ExternalLoginUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ExternalLoginUser{config: eluc.config}
		_spec = sqlgraph.NewCreateSpec(externalloginuser.Table, sqlgraph.NewFieldSpec(externalloginuser.FieldID, field.TypeInt))
	)
	if value, ok := eluc.mutation.ExternalID(); ok {
		_spec.SetField(externalloginuser.FieldExternalID, field.TypeString, value)
		_node.ExternalID = value
	}
	if value, ok := eluc.mutation.UserID(); ok {
		_spec.SetField(externalloginuser.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := eluc.mutation.LoginSourceID(); ok {
		_spec.SetField(externalloginuser.FieldLoginSourceID, field.TypeInt64, value)
		_node.LoginSourceID = value
	}
	if value, ok := eluc.mutation.Provider(); ok {
		_spec.SetField(externalloginuser.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := eluc.mutation.Email(); ok {
		_spec.SetField(externalloginuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := eluc.mutation.Name(); ok {
		_spec.SetField(externalloginuser.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := eluc.mutation.FirstName(); ok {
		_spec.SetField(externalloginuser.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := eluc.mutation.LastName(); ok {
		_spec.SetField(externalloginuser.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := eluc.mutation.NickName(); ok {
		_spec.SetField(externalloginuser.FieldNickName, field.TypeString, value)
		_node.NickName = value
	}
	if value, ok := eluc.mutation.Description(); ok {
		_spec.SetField(externalloginuser.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := eluc.mutation.AvatarURL(); ok {
		_spec.SetField(externalloginuser.FieldAvatarURL, field.TypeString, value)
		_node.AvatarURL = value
	}
	if value, ok := eluc.mutation.Location(); ok {
		_spec.SetField(externalloginuser.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := eluc.mutation.AccessToken(); ok {
		_spec.SetField(externalloginuser.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := eluc.mutation.AccessTokenSecret(); ok {
		_spec.SetField(externalloginuser.FieldAccessTokenSecret, field.TypeString, value)
		_node.AccessTokenSecret = value
	}
	if value, ok := eluc.mutation.RefreshToken(); ok {
		_spec.SetField(externalloginuser.FieldRefreshToken, field.TypeString, value)
		_node.RefreshToken = value
	}
	if value, ok := eluc.mutation.ExpiredAt(); ok {
		_spec.SetField(externalloginuser.FieldExpiredAt, field.TypeTime, value)
		_node.ExpiredAt = value
	}
	return _node, _spec
}

// ExternalLoginUserCreateBulk is the builder for creating many ExternalLoginUser entities in bulk.
type ExternalLoginUserCreateBulk struct {
	config
	err      error
	builders []*ExternalLoginUserCreate
}

// Save creates the ExternalLoginUser entities in the database.
func (elucb *ExternalLoginUserCreateBulk) Save(ctx context.Context) ([]*ExternalLoginUser, error) {
	if elucb.err != nil {
		return nil, elucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(elucb.builders))
	nodes := make([]*ExternalLoginUser, len(elucb.builders))
	mutators := make([]Mutator, len(elucb.builders))
	for i := range elucb.builders {
		func(i int, root context.Context) {
			builder := elucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExternalLoginUserMutation)
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
					_, err = mutators[i+1].Mutate(root, elucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, elucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, elucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (elucb *ExternalLoginUserCreateBulk) SaveX(ctx context.Context) []*ExternalLoginUser {
	v, err := elucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elucb *ExternalLoginUserCreateBulk) Exec(ctx context.Context) error {
	_, err := elucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elucb *ExternalLoginUserCreateBulk) ExecX(ctx context.Context) {
	if err := elucb.Exec(ctx); err != nil {
		panic(err)
	}
}