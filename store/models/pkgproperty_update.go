// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/pkgproperty"
	"github.com/gitbundle/gitbundle/store/models/predicate"
)

// PkgPropertyUpdate is the builder for updating PkgProperty entities.
type PkgPropertyUpdate struct {
	config
	hooks    []Hook
	mutation *PkgPropertyMutation
}

// Where appends a list predicates to the PkgPropertyUpdate builder.
func (ppu *PkgPropertyUpdate) Where(ps ...predicate.PkgProperty) *PkgPropertyUpdate {
	ppu.mutation.Where(ps...)
	return ppu
}

// SetRefType sets the "ref_type" field.
func (ppu *PkgPropertyUpdate) SetRefType(i int64) *PkgPropertyUpdate {
	ppu.mutation.ResetRefType()
	ppu.mutation.SetRefType(i)
	return ppu
}

// SetNillableRefType sets the "ref_type" field if the given value is not nil.
func (ppu *PkgPropertyUpdate) SetNillableRefType(i *int64) *PkgPropertyUpdate {
	if i != nil {
		ppu.SetRefType(*i)
	}
	return ppu
}

// AddRefType adds i to the "ref_type" field.
func (ppu *PkgPropertyUpdate) AddRefType(i int64) *PkgPropertyUpdate {
	ppu.mutation.AddRefType(i)
	return ppu
}

// SetRefID sets the "ref_id" field.
func (ppu *PkgPropertyUpdate) SetRefID(i int64) *PkgPropertyUpdate {
	ppu.mutation.ResetRefID()
	ppu.mutation.SetRefID(i)
	return ppu
}

// SetNillableRefID sets the "ref_id" field if the given value is not nil.
func (ppu *PkgPropertyUpdate) SetNillableRefID(i *int64) *PkgPropertyUpdate {
	if i != nil {
		ppu.SetRefID(*i)
	}
	return ppu
}

// AddRefID adds i to the "ref_id" field.
func (ppu *PkgPropertyUpdate) AddRefID(i int64) *PkgPropertyUpdate {
	ppu.mutation.AddRefID(i)
	return ppu
}

// SetName sets the "name" field.
func (ppu *PkgPropertyUpdate) SetName(s string) *PkgPropertyUpdate {
	ppu.mutation.SetName(s)
	return ppu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ppu *PkgPropertyUpdate) SetNillableName(s *string) *PkgPropertyUpdate {
	if s != nil {
		ppu.SetName(*s)
	}
	return ppu
}

// SetValue sets the "value" field.
func (ppu *PkgPropertyUpdate) SetValue(s string) *PkgPropertyUpdate {
	ppu.mutation.SetValue(s)
	return ppu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (ppu *PkgPropertyUpdate) SetNillableValue(s *string) *PkgPropertyUpdate {
	if s != nil {
		ppu.SetValue(*s)
	}
	return ppu
}

// Mutation returns the PkgPropertyMutation object of the builder.
func (ppu *PkgPropertyUpdate) Mutation() *PkgPropertyMutation {
	return ppu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ppu *PkgPropertyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ppu.sqlSave, ppu.mutation, ppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *PkgPropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *PkgPropertyUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *PkgPropertyUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppu *PkgPropertyUpdate) check() error {
	if v, ok := ppu.mutation.Name(); ok {
		if err := pkgproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "PkgProperty.name": %w`, err)}
		}
	}
	return nil
}

func (ppu *PkgPropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ppu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pkgproperty.Table, pkgproperty.Columns, sqlgraph.NewFieldSpec(pkgproperty.FieldID, field.TypeInt))
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppu.mutation.RefType(); ok {
		_spec.SetField(pkgproperty.FieldRefType, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.AddedRefType(); ok {
		_spec.AddField(pkgproperty.FieldRefType, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.RefID(); ok {
		_spec.SetField(pkgproperty.FieldRefID, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.AddedRefID(); ok {
		_spec.AddField(pkgproperty.FieldRefID, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.Name(); ok {
		_spec.SetField(pkgproperty.FieldName, field.TypeString, value)
	}
	if value, ok := ppu.mutation.Value(); ok {
		_spec.SetField(pkgproperty.FieldValue, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkgproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ppu.mutation.done = true
	return n, nil
}

// PkgPropertyUpdateOne is the builder for updating a single PkgProperty entity.
type PkgPropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PkgPropertyMutation
}

// SetRefType sets the "ref_type" field.
func (ppuo *PkgPropertyUpdateOne) SetRefType(i int64) *PkgPropertyUpdateOne {
	ppuo.mutation.ResetRefType()
	ppuo.mutation.SetRefType(i)
	return ppuo
}

// SetNillableRefType sets the "ref_type" field if the given value is not nil.
func (ppuo *PkgPropertyUpdateOne) SetNillableRefType(i *int64) *PkgPropertyUpdateOne {
	if i != nil {
		ppuo.SetRefType(*i)
	}
	return ppuo
}

// AddRefType adds i to the "ref_type" field.
func (ppuo *PkgPropertyUpdateOne) AddRefType(i int64) *PkgPropertyUpdateOne {
	ppuo.mutation.AddRefType(i)
	return ppuo
}

// SetRefID sets the "ref_id" field.
func (ppuo *PkgPropertyUpdateOne) SetRefID(i int64) *PkgPropertyUpdateOne {
	ppuo.mutation.ResetRefID()
	ppuo.mutation.SetRefID(i)
	return ppuo
}

// SetNillableRefID sets the "ref_id" field if the given value is not nil.
func (ppuo *PkgPropertyUpdateOne) SetNillableRefID(i *int64) *PkgPropertyUpdateOne {
	if i != nil {
		ppuo.SetRefID(*i)
	}
	return ppuo
}

// AddRefID adds i to the "ref_id" field.
func (ppuo *PkgPropertyUpdateOne) AddRefID(i int64) *PkgPropertyUpdateOne {
	ppuo.mutation.AddRefID(i)
	return ppuo
}

// SetName sets the "name" field.
func (ppuo *PkgPropertyUpdateOne) SetName(s string) *PkgPropertyUpdateOne {
	ppuo.mutation.SetName(s)
	return ppuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ppuo *PkgPropertyUpdateOne) SetNillableName(s *string) *PkgPropertyUpdateOne {
	if s != nil {
		ppuo.SetName(*s)
	}
	return ppuo
}

// SetValue sets the "value" field.
func (ppuo *PkgPropertyUpdateOne) SetValue(s string) *PkgPropertyUpdateOne {
	ppuo.mutation.SetValue(s)
	return ppuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (ppuo *PkgPropertyUpdateOne) SetNillableValue(s *string) *PkgPropertyUpdateOne {
	if s != nil {
		ppuo.SetValue(*s)
	}
	return ppuo
}

// Mutation returns the PkgPropertyMutation object of the builder.
func (ppuo *PkgPropertyUpdateOne) Mutation() *PkgPropertyMutation {
	return ppuo.mutation
}

// Where appends a list predicates to the PkgPropertyUpdate builder.
func (ppuo *PkgPropertyUpdateOne) Where(ps ...predicate.PkgProperty) *PkgPropertyUpdateOne {
	ppuo.mutation.Where(ps...)
	return ppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ppuo *PkgPropertyUpdateOne) Select(field string, fields ...string) *PkgPropertyUpdateOne {
	ppuo.fields = append([]string{field}, fields...)
	return ppuo
}

// Save executes the query and returns the updated PkgProperty entity.
func (ppuo *PkgPropertyUpdateOne) Save(ctx context.Context) (*PkgProperty, error) {
	return withHooks(ctx, ppuo.sqlSave, ppuo.mutation, ppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *PkgPropertyUpdateOne) SaveX(ctx context.Context) *PkgProperty {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *PkgPropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *PkgPropertyUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppuo *PkgPropertyUpdateOne) check() error {
	if v, ok := ppuo.mutation.Name(); ok {
		if err := pkgproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "PkgProperty.name": %w`, err)}
		}
	}
	return nil
}

func (ppuo *PkgPropertyUpdateOne) sqlSave(ctx context.Context) (_node *PkgProperty, err error) {
	if err := ppuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pkgproperty.Table, pkgproperty.Columns, sqlgraph.NewFieldSpec(pkgproperty.FieldID, field.TypeInt))
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "PkgProperty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pkgproperty.FieldID)
		for _, f := range fields {
			if !pkgproperty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != pkgproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppuo.mutation.RefType(); ok {
		_spec.SetField(pkgproperty.FieldRefType, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.AddedRefType(); ok {
		_spec.AddField(pkgproperty.FieldRefType, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.RefID(); ok {
		_spec.SetField(pkgproperty.FieldRefID, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.AddedRefID(); ok {
		_spec.AddField(pkgproperty.FieldRefID, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.Name(); ok {
		_spec.SetField(pkgproperty.FieldName, field.TypeString, value)
	}
	if value, ok := ppuo.mutation.Value(); ok {
		_spec.SetField(pkgproperty.FieldValue, field.TypeString, value)
	}
	_node = &PkgProperty{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkgproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ppuo.mutation.done = true
	return _node, nil
}