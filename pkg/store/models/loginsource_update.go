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
	"github.com/gitbundle/server/pkg/store/models/loginsource"
	"github.com/gitbundle/server/pkg/store/models/predicate"
)

// LoginSourceUpdate is the builder for updating LoginSource entities.
type LoginSourceUpdate struct {
	config
	hooks    []Hook
	mutation *LoginSourceMutation
}

// Where appends a list predicates to the LoginSourceUpdate builder.
func (lsu *LoginSourceUpdate) Where(ps ...predicate.LoginSource) *LoginSourceUpdate {
	lsu.mutation.Where(ps...)
	return lsu
}

// SetType sets the "type" field.
func (lsu *LoginSourceUpdate) SetType(i int) *LoginSourceUpdate {
	lsu.mutation.ResetType()
	lsu.mutation.SetType(i)
	return lsu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableType(i *int) *LoginSourceUpdate {
	if i != nil {
		lsu.SetType(*i)
	}
	return lsu
}

// AddType adds i to the "type" field.
func (lsu *LoginSourceUpdate) AddType(i int) *LoginSourceUpdate {
	lsu.mutation.AddType(i)
	return lsu
}

// ClearType clears the value of the "type" field.
func (lsu *LoginSourceUpdate) ClearType() *LoginSourceUpdate {
	lsu.mutation.ClearType()
	return lsu
}

// SetName sets the "name" field.
func (lsu *LoginSourceUpdate) SetName(s string) *LoginSourceUpdate {
	lsu.mutation.SetName(s)
	return lsu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableName(s *string) *LoginSourceUpdate {
	if s != nil {
		lsu.SetName(*s)
	}
	return lsu
}

// ClearName clears the value of the "name" field.
func (lsu *LoginSourceUpdate) ClearName() *LoginSourceUpdate {
	lsu.mutation.ClearName()
	return lsu
}

// SetIsActive sets the "is_active" field.
func (lsu *LoginSourceUpdate) SetIsActive(b bool) *LoginSourceUpdate {
	lsu.mutation.SetIsActive(b)
	return lsu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableIsActive(b *bool) *LoginSourceUpdate {
	if b != nil {
		lsu.SetIsActive(*b)
	}
	return lsu
}

// SetIsSyncEnabled sets the "is_sync_enabled" field.
func (lsu *LoginSourceUpdate) SetIsSyncEnabled(b bool) *LoginSourceUpdate {
	lsu.mutation.SetIsSyncEnabled(b)
	return lsu
}

// SetNillableIsSyncEnabled sets the "is_sync_enabled" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableIsSyncEnabled(b *bool) *LoginSourceUpdate {
	if b != nil {
		lsu.SetIsSyncEnabled(*b)
	}
	return lsu
}

// SetConfig sets the "config" field.
func (lsu *LoginSourceUpdate) SetConfig(s string) *LoginSourceUpdate {
	lsu.mutation.SetConfig(s)
	return lsu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableConfig(s *string) *LoginSourceUpdate {
	if s != nil {
		lsu.SetConfig(*s)
	}
	return lsu
}

// ClearConfig clears the value of the "config" field.
func (lsu *LoginSourceUpdate) ClearConfig() *LoginSourceUpdate {
	lsu.mutation.ClearConfig()
	return lsu
}

// SetCreatedAt sets the "created_at" field.
func (lsu *LoginSourceUpdate) SetCreatedAt(t time.Time) *LoginSourceUpdate {
	lsu.mutation.SetCreatedAt(t)
	return lsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableCreatedAt(t *time.Time) *LoginSourceUpdate {
	if t != nil {
		lsu.SetCreatedAt(*t)
	}
	return lsu
}

// SetUpdatedAt sets the "updated_at" field.
func (lsu *LoginSourceUpdate) SetUpdatedAt(t time.Time) *LoginSourceUpdate {
	lsu.mutation.SetUpdatedAt(t)
	return lsu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lsu *LoginSourceUpdate) SetNillableUpdatedAt(t *time.Time) *LoginSourceUpdate {
	if t != nil {
		lsu.SetUpdatedAt(*t)
	}
	return lsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (lsu *LoginSourceUpdate) ClearUpdatedAt() *LoginSourceUpdate {
	lsu.mutation.ClearUpdatedAt()
	return lsu
}

// Mutation returns the LoginSourceMutation object of the builder.
func (lsu *LoginSourceUpdate) Mutation() *LoginSourceMutation {
	return lsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lsu *LoginSourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lsu.sqlSave, lsu.mutation, lsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsu *LoginSourceUpdate) SaveX(ctx context.Context) int {
	affected, err := lsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lsu *LoginSourceUpdate) Exec(ctx context.Context) error {
	_, err := lsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsu *LoginSourceUpdate) ExecX(ctx context.Context) {
	if err := lsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsu *LoginSourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(loginsource.Table, loginsource.Columns, sqlgraph.NewFieldSpec(loginsource.FieldID, field.TypeInt))
	if ps := lsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsu.mutation.GetType(); ok {
		_spec.SetField(loginsource.FieldType, field.TypeInt, value)
	}
	if value, ok := lsu.mutation.AddedType(); ok {
		_spec.AddField(loginsource.FieldType, field.TypeInt, value)
	}
	if lsu.mutation.TypeCleared() {
		_spec.ClearField(loginsource.FieldType, field.TypeInt)
	}
	if value, ok := lsu.mutation.Name(); ok {
		_spec.SetField(loginsource.FieldName, field.TypeString, value)
	}
	if lsu.mutation.NameCleared() {
		_spec.ClearField(loginsource.FieldName, field.TypeString)
	}
	if value, ok := lsu.mutation.IsActive(); ok {
		_spec.SetField(loginsource.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := lsu.mutation.IsSyncEnabled(); ok {
		_spec.SetField(loginsource.FieldIsSyncEnabled, field.TypeBool, value)
	}
	if value, ok := lsu.mutation.Config(); ok {
		_spec.SetField(loginsource.FieldConfig, field.TypeString, value)
	}
	if lsu.mutation.ConfigCleared() {
		_spec.ClearField(loginsource.FieldConfig, field.TypeString)
	}
	if value, ok := lsu.mutation.CreatedAt(); ok {
		_spec.SetField(loginsource.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lsu.mutation.UpdatedAt(); ok {
		_spec.SetField(loginsource.FieldUpdatedAt, field.TypeTime, value)
	}
	if lsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(loginsource.FieldUpdatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginsource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lsu.mutation.done = true
	return n, nil
}

// LoginSourceUpdateOne is the builder for updating a single LoginSource entity.
type LoginSourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginSourceMutation
}

// SetType sets the "type" field.
func (lsuo *LoginSourceUpdateOne) SetType(i int) *LoginSourceUpdateOne {
	lsuo.mutation.ResetType()
	lsuo.mutation.SetType(i)
	return lsuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableType(i *int) *LoginSourceUpdateOne {
	if i != nil {
		lsuo.SetType(*i)
	}
	return lsuo
}

// AddType adds i to the "type" field.
func (lsuo *LoginSourceUpdateOne) AddType(i int) *LoginSourceUpdateOne {
	lsuo.mutation.AddType(i)
	return lsuo
}

// ClearType clears the value of the "type" field.
func (lsuo *LoginSourceUpdateOne) ClearType() *LoginSourceUpdateOne {
	lsuo.mutation.ClearType()
	return lsuo
}

// SetName sets the "name" field.
func (lsuo *LoginSourceUpdateOne) SetName(s string) *LoginSourceUpdateOne {
	lsuo.mutation.SetName(s)
	return lsuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableName(s *string) *LoginSourceUpdateOne {
	if s != nil {
		lsuo.SetName(*s)
	}
	return lsuo
}

// ClearName clears the value of the "name" field.
func (lsuo *LoginSourceUpdateOne) ClearName() *LoginSourceUpdateOne {
	lsuo.mutation.ClearName()
	return lsuo
}

// SetIsActive sets the "is_active" field.
func (lsuo *LoginSourceUpdateOne) SetIsActive(b bool) *LoginSourceUpdateOne {
	lsuo.mutation.SetIsActive(b)
	return lsuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableIsActive(b *bool) *LoginSourceUpdateOne {
	if b != nil {
		lsuo.SetIsActive(*b)
	}
	return lsuo
}

// SetIsSyncEnabled sets the "is_sync_enabled" field.
func (lsuo *LoginSourceUpdateOne) SetIsSyncEnabled(b bool) *LoginSourceUpdateOne {
	lsuo.mutation.SetIsSyncEnabled(b)
	return lsuo
}

// SetNillableIsSyncEnabled sets the "is_sync_enabled" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableIsSyncEnabled(b *bool) *LoginSourceUpdateOne {
	if b != nil {
		lsuo.SetIsSyncEnabled(*b)
	}
	return lsuo
}

// SetConfig sets the "config" field.
func (lsuo *LoginSourceUpdateOne) SetConfig(s string) *LoginSourceUpdateOne {
	lsuo.mutation.SetConfig(s)
	return lsuo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableConfig(s *string) *LoginSourceUpdateOne {
	if s != nil {
		lsuo.SetConfig(*s)
	}
	return lsuo
}

// ClearConfig clears the value of the "config" field.
func (lsuo *LoginSourceUpdateOne) ClearConfig() *LoginSourceUpdateOne {
	lsuo.mutation.ClearConfig()
	return lsuo
}

// SetCreatedAt sets the "created_at" field.
func (lsuo *LoginSourceUpdateOne) SetCreatedAt(t time.Time) *LoginSourceUpdateOne {
	lsuo.mutation.SetCreatedAt(t)
	return lsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableCreatedAt(t *time.Time) *LoginSourceUpdateOne {
	if t != nil {
		lsuo.SetCreatedAt(*t)
	}
	return lsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (lsuo *LoginSourceUpdateOne) SetUpdatedAt(t time.Time) *LoginSourceUpdateOne {
	lsuo.mutation.SetUpdatedAt(t)
	return lsuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lsuo *LoginSourceUpdateOne) SetNillableUpdatedAt(t *time.Time) *LoginSourceUpdateOne {
	if t != nil {
		lsuo.SetUpdatedAt(*t)
	}
	return lsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (lsuo *LoginSourceUpdateOne) ClearUpdatedAt() *LoginSourceUpdateOne {
	lsuo.mutation.ClearUpdatedAt()
	return lsuo
}

// Mutation returns the LoginSourceMutation object of the builder.
func (lsuo *LoginSourceUpdateOne) Mutation() *LoginSourceMutation {
	return lsuo.mutation
}

// Where appends a list predicates to the LoginSourceUpdate builder.
func (lsuo *LoginSourceUpdateOne) Where(ps ...predicate.LoginSource) *LoginSourceUpdateOne {
	lsuo.mutation.Where(ps...)
	return lsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lsuo *LoginSourceUpdateOne) Select(field string, fields ...string) *LoginSourceUpdateOne {
	lsuo.fields = append([]string{field}, fields...)
	return lsuo
}

// Save executes the query and returns the updated LoginSource entity.
func (lsuo *LoginSourceUpdateOne) Save(ctx context.Context) (*LoginSource, error) {
	return withHooks(ctx, lsuo.sqlSave, lsuo.mutation, lsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsuo *LoginSourceUpdateOne) SaveX(ctx context.Context) *LoginSource {
	node, err := lsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lsuo *LoginSourceUpdateOne) Exec(ctx context.Context) error {
	_, err := lsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsuo *LoginSourceUpdateOne) ExecX(ctx context.Context) {
	if err := lsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsuo *LoginSourceUpdateOne) sqlSave(ctx context.Context) (_node *LoginSource, err error) {
	_spec := sqlgraph.NewUpdateSpec(loginsource.Table, loginsource.Columns, sqlgraph.NewFieldSpec(loginsource.FieldID, field.TypeInt))
	id, ok := lsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "LoginSource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginsource.FieldID)
		for _, f := range fields {
			if !loginsource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != loginsource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsuo.mutation.GetType(); ok {
		_spec.SetField(loginsource.FieldType, field.TypeInt, value)
	}
	if value, ok := lsuo.mutation.AddedType(); ok {
		_spec.AddField(loginsource.FieldType, field.TypeInt, value)
	}
	if lsuo.mutation.TypeCleared() {
		_spec.ClearField(loginsource.FieldType, field.TypeInt)
	}
	if value, ok := lsuo.mutation.Name(); ok {
		_spec.SetField(loginsource.FieldName, field.TypeString, value)
	}
	if lsuo.mutation.NameCleared() {
		_spec.ClearField(loginsource.FieldName, field.TypeString)
	}
	if value, ok := lsuo.mutation.IsActive(); ok {
		_spec.SetField(loginsource.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := lsuo.mutation.IsSyncEnabled(); ok {
		_spec.SetField(loginsource.FieldIsSyncEnabled, field.TypeBool, value)
	}
	if value, ok := lsuo.mutation.Config(); ok {
		_spec.SetField(loginsource.FieldConfig, field.TypeString, value)
	}
	if lsuo.mutation.ConfigCleared() {
		_spec.ClearField(loginsource.FieldConfig, field.TypeString)
	}
	if value, ok := lsuo.mutation.CreatedAt(); ok {
		_spec.SetField(loginsource.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(loginsource.FieldUpdatedAt, field.TypeTime, value)
	}
	if lsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(loginsource.FieldUpdatedAt, field.TypeTime)
	}
	_node = &LoginSource{config: lsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginsource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lsuo.mutation.done = true
	return _node, nil
}