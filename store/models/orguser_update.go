// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/orguser"
	"github.com/gitbundle/gitbundle/store/models/predicate"
)

// OrgUserUpdate is the builder for updating OrgUser entities.
type OrgUserUpdate struct {
	config
	hooks    []Hook
	mutation *OrgUserMutation
}

// Where appends a list predicates to the OrgUserUpdate builder.
func (ouu *OrgUserUpdate) Where(ps ...predicate.OrgUser) *OrgUserUpdate {
	ouu.mutation.Where(ps...)
	return ouu
}

// SetUserID sets the "user_id" field.
func (ouu *OrgUserUpdate) SetUserID(i int64) *OrgUserUpdate {
	ouu.mutation.ResetUserID()
	ouu.mutation.SetUserID(i)
	return ouu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableUserID(i *int64) *OrgUserUpdate {
	if i != nil {
		ouu.SetUserID(*i)
	}
	return ouu
}

// AddUserID adds i to the "user_id" field.
func (ouu *OrgUserUpdate) AddUserID(i int64) *OrgUserUpdate {
	ouu.mutation.AddUserID(i)
	return ouu
}

// ClearUserID clears the value of the "user_id" field.
func (ouu *OrgUserUpdate) ClearUserID() *OrgUserUpdate {
	ouu.mutation.ClearUserID()
	return ouu
}

// SetOrgID sets the "org_id" field.
func (ouu *OrgUserUpdate) SetOrgID(i int64) *OrgUserUpdate {
	ouu.mutation.ResetOrgID()
	ouu.mutation.SetOrgID(i)
	return ouu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableOrgID(i *int64) *OrgUserUpdate {
	if i != nil {
		ouu.SetOrgID(*i)
	}
	return ouu
}

// AddOrgID adds i to the "org_id" field.
func (ouu *OrgUserUpdate) AddOrgID(i int64) *OrgUserUpdate {
	ouu.mutation.AddOrgID(i)
	return ouu
}

// ClearOrgID clears the value of the "org_id" field.
func (ouu *OrgUserUpdate) ClearOrgID() *OrgUserUpdate {
	ouu.mutation.ClearOrgID()
	return ouu
}

// SetIsPublic sets the "is_public" field.
func (ouu *OrgUserUpdate) SetIsPublic(b bool) *OrgUserUpdate {
	ouu.mutation.SetIsPublic(b)
	return ouu
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableIsPublic(b *bool) *OrgUserUpdate {
	if b != nil {
		ouu.SetIsPublic(*b)
	}
	return ouu
}

// ClearIsPublic clears the value of the "is_public" field.
func (ouu *OrgUserUpdate) ClearIsPublic() *OrgUserUpdate {
	ouu.mutation.ClearIsPublic()
	return ouu
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouu *OrgUserUpdate) Mutation() *OrgUserMutation {
	return ouu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ouu *OrgUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ouu.sqlSave, ouu.mutation, ouu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouu *OrgUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ouu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ouu *OrgUserUpdate) Exec(ctx context.Context) error {
	_, err := ouu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouu *OrgUserUpdate) ExecX(ctx context.Context) {
	if err := ouu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouu *OrgUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orguser.Table, orguser.Columns, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	if ps := ouu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouu.mutation.UserID(); ok {
		_spec.SetField(orguser.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ouu.mutation.AddedUserID(); ok {
		_spec.AddField(orguser.FieldUserID, field.TypeInt64, value)
	}
	if ouu.mutation.UserIDCleared() {
		_spec.ClearField(orguser.FieldUserID, field.TypeInt64)
	}
	if value, ok := ouu.mutation.OrgID(); ok {
		_spec.SetField(orguser.FieldOrgID, field.TypeInt64, value)
	}
	if value, ok := ouu.mutation.AddedOrgID(); ok {
		_spec.AddField(orguser.FieldOrgID, field.TypeInt64, value)
	}
	if ouu.mutation.OrgIDCleared() {
		_spec.ClearField(orguser.FieldOrgID, field.TypeInt64)
	}
	if value, ok := ouu.mutation.IsPublic(); ok {
		_spec.SetField(orguser.FieldIsPublic, field.TypeBool, value)
	}
	if ouu.mutation.IsPublicCleared() {
		_spec.ClearField(orguser.FieldIsPublic, field.TypeBool)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ouu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ouu.mutation.done = true
	return n, nil
}

// OrgUserUpdateOne is the builder for updating a single OrgUser entity.
type OrgUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgUserMutation
}

// SetUserID sets the "user_id" field.
func (ouuo *OrgUserUpdateOne) SetUserID(i int64) *OrgUserUpdateOne {
	ouuo.mutation.ResetUserID()
	ouuo.mutation.SetUserID(i)
	return ouuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableUserID(i *int64) *OrgUserUpdateOne {
	if i != nil {
		ouuo.SetUserID(*i)
	}
	return ouuo
}

// AddUserID adds i to the "user_id" field.
func (ouuo *OrgUserUpdateOne) AddUserID(i int64) *OrgUserUpdateOne {
	ouuo.mutation.AddUserID(i)
	return ouuo
}

// ClearUserID clears the value of the "user_id" field.
func (ouuo *OrgUserUpdateOne) ClearUserID() *OrgUserUpdateOne {
	ouuo.mutation.ClearUserID()
	return ouuo
}

// SetOrgID sets the "org_id" field.
func (ouuo *OrgUserUpdateOne) SetOrgID(i int64) *OrgUserUpdateOne {
	ouuo.mutation.ResetOrgID()
	ouuo.mutation.SetOrgID(i)
	return ouuo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableOrgID(i *int64) *OrgUserUpdateOne {
	if i != nil {
		ouuo.SetOrgID(*i)
	}
	return ouuo
}

// AddOrgID adds i to the "org_id" field.
func (ouuo *OrgUserUpdateOne) AddOrgID(i int64) *OrgUserUpdateOne {
	ouuo.mutation.AddOrgID(i)
	return ouuo
}

// ClearOrgID clears the value of the "org_id" field.
func (ouuo *OrgUserUpdateOne) ClearOrgID() *OrgUserUpdateOne {
	ouuo.mutation.ClearOrgID()
	return ouuo
}

// SetIsPublic sets the "is_public" field.
func (ouuo *OrgUserUpdateOne) SetIsPublic(b bool) *OrgUserUpdateOne {
	ouuo.mutation.SetIsPublic(b)
	return ouuo
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableIsPublic(b *bool) *OrgUserUpdateOne {
	if b != nil {
		ouuo.SetIsPublic(*b)
	}
	return ouuo
}

// ClearIsPublic clears the value of the "is_public" field.
func (ouuo *OrgUserUpdateOne) ClearIsPublic() *OrgUserUpdateOne {
	ouuo.mutation.ClearIsPublic()
	return ouuo
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouuo *OrgUserUpdateOne) Mutation() *OrgUserMutation {
	return ouuo.mutation
}

// Where appends a list predicates to the OrgUserUpdate builder.
func (ouuo *OrgUserUpdateOne) Where(ps ...predicate.OrgUser) *OrgUserUpdateOne {
	ouuo.mutation.Where(ps...)
	return ouuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouuo *OrgUserUpdateOne) Select(field string, fields ...string) *OrgUserUpdateOne {
	ouuo.fields = append([]string{field}, fields...)
	return ouuo
}

// Save executes the query and returns the updated OrgUser entity.
func (ouuo *OrgUserUpdateOne) Save(ctx context.Context) (*OrgUser, error) {
	return withHooks(ctx, ouuo.sqlSave, ouuo.mutation, ouuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouuo *OrgUserUpdateOne) SaveX(ctx context.Context) *OrgUser {
	node, err := ouuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouuo *OrgUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ouuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouuo *OrgUserUpdateOne) ExecX(ctx context.Context) {
	if err := ouuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouuo *OrgUserUpdateOne) sqlSave(ctx context.Context) (_node *OrgUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(orguser.Table, orguser.Columns, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	id, ok := ouuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OrgUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orguser.FieldID)
		for _, f := range fields {
			if !orguser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != orguser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouuo.mutation.UserID(); ok {
		_spec.SetField(orguser.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ouuo.mutation.AddedUserID(); ok {
		_spec.AddField(orguser.FieldUserID, field.TypeInt64, value)
	}
	if ouuo.mutation.UserIDCleared() {
		_spec.ClearField(orguser.FieldUserID, field.TypeInt64)
	}
	if value, ok := ouuo.mutation.OrgID(); ok {
		_spec.SetField(orguser.FieldOrgID, field.TypeInt64, value)
	}
	if value, ok := ouuo.mutation.AddedOrgID(); ok {
		_spec.AddField(orguser.FieldOrgID, field.TypeInt64, value)
	}
	if ouuo.mutation.OrgIDCleared() {
		_spec.ClearField(orguser.FieldOrgID, field.TypeInt64)
	}
	if value, ok := ouuo.mutation.IsPublic(); ok {
		_spec.SetField(orguser.FieldIsPublic, field.TypeBool, value)
	}
	if ouuo.mutation.IsPublicCleared() {
		_spec.ClearField(orguser.FieldIsPublic, field.TypeBool)
	}
	_node = &OrgUser{config: ouuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouuo.mutation.done = true
	return _node, nil
}