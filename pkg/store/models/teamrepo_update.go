// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/server/pkg/store/models/predicate"
	"github.com/gitbundle/server/pkg/store/models/teamrepo"
)

// TeamRepoUpdate is the builder for updating TeamRepo entities.
type TeamRepoUpdate struct {
	config
	hooks    []Hook
	mutation *TeamRepoMutation
}

// Where appends a list predicates to the TeamRepoUpdate builder.
func (tru *TeamRepoUpdate) Where(ps ...predicate.TeamRepo) *TeamRepoUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetOrgID sets the "org_id" field.
func (tru *TeamRepoUpdate) SetOrgID(i int64) *TeamRepoUpdate {
	tru.mutation.ResetOrgID()
	tru.mutation.SetOrgID(i)
	return tru
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (tru *TeamRepoUpdate) SetNillableOrgID(i *int64) *TeamRepoUpdate {
	if i != nil {
		tru.SetOrgID(*i)
	}
	return tru
}

// AddOrgID adds i to the "org_id" field.
func (tru *TeamRepoUpdate) AddOrgID(i int64) *TeamRepoUpdate {
	tru.mutation.AddOrgID(i)
	return tru
}

// ClearOrgID clears the value of the "org_id" field.
func (tru *TeamRepoUpdate) ClearOrgID() *TeamRepoUpdate {
	tru.mutation.ClearOrgID()
	return tru
}

// SetTeamID sets the "team_id" field.
func (tru *TeamRepoUpdate) SetTeamID(i int64) *TeamRepoUpdate {
	tru.mutation.ResetTeamID()
	tru.mutation.SetTeamID(i)
	return tru
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (tru *TeamRepoUpdate) SetNillableTeamID(i *int64) *TeamRepoUpdate {
	if i != nil {
		tru.SetTeamID(*i)
	}
	return tru
}

// AddTeamID adds i to the "team_id" field.
func (tru *TeamRepoUpdate) AddTeamID(i int64) *TeamRepoUpdate {
	tru.mutation.AddTeamID(i)
	return tru
}

// ClearTeamID clears the value of the "team_id" field.
func (tru *TeamRepoUpdate) ClearTeamID() *TeamRepoUpdate {
	tru.mutation.ClearTeamID()
	return tru
}

// SetRepoID sets the "repo_id" field.
func (tru *TeamRepoUpdate) SetRepoID(i int64) *TeamRepoUpdate {
	tru.mutation.ResetRepoID()
	tru.mutation.SetRepoID(i)
	return tru
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (tru *TeamRepoUpdate) SetNillableRepoID(i *int64) *TeamRepoUpdate {
	if i != nil {
		tru.SetRepoID(*i)
	}
	return tru
}

// AddRepoID adds i to the "repo_id" field.
func (tru *TeamRepoUpdate) AddRepoID(i int64) *TeamRepoUpdate {
	tru.mutation.AddRepoID(i)
	return tru
}

// ClearRepoID clears the value of the "repo_id" field.
func (tru *TeamRepoUpdate) ClearRepoID() *TeamRepoUpdate {
	tru.mutation.ClearRepoID()
	return tru
}

// Mutation returns the TeamRepoMutation object of the builder.
func (tru *TeamRepoUpdate) Mutation() *TeamRepoMutation {
	return tru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TeamRepoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tru.sqlSave, tru.mutation, tru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TeamRepoUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TeamRepoUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TeamRepoUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tru *TeamRepoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(teamrepo.Table, teamrepo.Columns, sqlgraph.NewFieldSpec(teamrepo.FieldID, field.TypeInt))
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.OrgID(); ok {
		_spec.SetField(teamrepo.FieldOrgID, field.TypeInt64, value)
	}
	if value, ok := tru.mutation.AddedOrgID(); ok {
		_spec.AddField(teamrepo.FieldOrgID, field.TypeInt64, value)
	}
	if tru.mutation.OrgIDCleared() {
		_spec.ClearField(teamrepo.FieldOrgID, field.TypeInt64)
	}
	if value, ok := tru.mutation.TeamID(); ok {
		_spec.SetField(teamrepo.FieldTeamID, field.TypeInt64, value)
	}
	if value, ok := tru.mutation.AddedTeamID(); ok {
		_spec.AddField(teamrepo.FieldTeamID, field.TypeInt64, value)
	}
	if tru.mutation.TeamIDCleared() {
		_spec.ClearField(teamrepo.FieldTeamID, field.TypeInt64)
	}
	if value, ok := tru.mutation.RepoID(); ok {
		_spec.SetField(teamrepo.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := tru.mutation.AddedRepoID(); ok {
		_spec.AddField(teamrepo.FieldRepoID, field.TypeInt64, value)
	}
	if tru.mutation.RepoIDCleared() {
		_spec.ClearField(teamrepo.FieldRepoID, field.TypeInt64)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamrepo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tru.mutation.done = true
	return n, nil
}

// TeamRepoUpdateOne is the builder for updating a single TeamRepo entity.
type TeamRepoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamRepoMutation
}

// SetOrgID sets the "org_id" field.
func (truo *TeamRepoUpdateOne) SetOrgID(i int64) *TeamRepoUpdateOne {
	truo.mutation.ResetOrgID()
	truo.mutation.SetOrgID(i)
	return truo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (truo *TeamRepoUpdateOne) SetNillableOrgID(i *int64) *TeamRepoUpdateOne {
	if i != nil {
		truo.SetOrgID(*i)
	}
	return truo
}

// AddOrgID adds i to the "org_id" field.
func (truo *TeamRepoUpdateOne) AddOrgID(i int64) *TeamRepoUpdateOne {
	truo.mutation.AddOrgID(i)
	return truo
}

// ClearOrgID clears the value of the "org_id" field.
func (truo *TeamRepoUpdateOne) ClearOrgID() *TeamRepoUpdateOne {
	truo.mutation.ClearOrgID()
	return truo
}

// SetTeamID sets the "team_id" field.
func (truo *TeamRepoUpdateOne) SetTeamID(i int64) *TeamRepoUpdateOne {
	truo.mutation.ResetTeamID()
	truo.mutation.SetTeamID(i)
	return truo
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (truo *TeamRepoUpdateOne) SetNillableTeamID(i *int64) *TeamRepoUpdateOne {
	if i != nil {
		truo.SetTeamID(*i)
	}
	return truo
}

// AddTeamID adds i to the "team_id" field.
func (truo *TeamRepoUpdateOne) AddTeamID(i int64) *TeamRepoUpdateOne {
	truo.mutation.AddTeamID(i)
	return truo
}

// ClearTeamID clears the value of the "team_id" field.
func (truo *TeamRepoUpdateOne) ClearTeamID() *TeamRepoUpdateOne {
	truo.mutation.ClearTeamID()
	return truo
}

// SetRepoID sets the "repo_id" field.
func (truo *TeamRepoUpdateOne) SetRepoID(i int64) *TeamRepoUpdateOne {
	truo.mutation.ResetRepoID()
	truo.mutation.SetRepoID(i)
	return truo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (truo *TeamRepoUpdateOne) SetNillableRepoID(i *int64) *TeamRepoUpdateOne {
	if i != nil {
		truo.SetRepoID(*i)
	}
	return truo
}

// AddRepoID adds i to the "repo_id" field.
func (truo *TeamRepoUpdateOne) AddRepoID(i int64) *TeamRepoUpdateOne {
	truo.mutation.AddRepoID(i)
	return truo
}

// ClearRepoID clears the value of the "repo_id" field.
func (truo *TeamRepoUpdateOne) ClearRepoID() *TeamRepoUpdateOne {
	truo.mutation.ClearRepoID()
	return truo
}

// Mutation returns the TeamRepoMutation object of the builder.
func (truo *TeamRepoUpdateOne) Mutation() *TeamRepoMutation {
	return truo.mutation
}

// Where appends a list predicates to the TeamRepoUpdate builder.
func (truo *TeamRepoUpdateOne) Where(ps ...predicate.TeamRepo) *TeamRepoUpdateOne {
	truo.mutation.Where(ps...)
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TeamRepoUpdateOne) Select(field string, fields ...string) *TeamRepoUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TeamRepo entity.
func (truo *TeamRepoUpdateOne) Save(ctx context.Context) (*TeamRepo, error) {
	return withHooks(ctx, truo.sqlSave, truo.mutation, truo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TeamRepoUpdateOne) SaveX(ctx context.Context) *TeamRepo {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TeamRepoUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TeamRepoUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (truo *TeamRepoUpdateOne) sqlSave(ctx context.Context) (_node *TeamRepo, err error) {
	_spec := sqlgraph.NewUpdateSpec(teamrepo.Table, teamrepo.Columns, sqlgraph.NewFieldSpec(teamrepo.FieldID, field.TypeInt))
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "TeamRepo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teamrepo.FieldID)
		for _, f := range fields {
			if !teamrepo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != teamrepo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.OrgID(); ok {
		_spec.SetField(teamrepo.FieldOrgID, field.TypeInt64, value)
	}
	if value, ok := truo.mutation.AddedOrgID(); ok {
		_spec.AddField(teamrepo.FieldOrgID, field.TypeInt64, value)
	}
	if truo.mutation.OrgIDCleared() {
		_spec.ClearField(teamrepo.FieldOrgID, field.TypeInt64)
	}
	if value, ok := truo.mutation.TeamID(); ok {
		_spec.SetField(teamrepo.FieldTeamID, field.TypeInt64, value)
	}
	if value, ok := truo.mutation.AddedTeamID(); ok {
		_spec.AddField(teamrepo.FieldTeamID, field.TypeInt64, value)
	}
	if truo.mutation.TeamIDCleared() {
		_spec.ClearField(teamrepo.FieldTeamID, field.TypeInt64)
	}
	if value, ok := truo.mutation.RepoID(); ok {
		_spec.SetField(teamrepo.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := truo.mutation.AddedRepoID(); ok {
		_spec.AddField(teamrepo.FieldRepoID, field.TypeInt64, value)
	}
	if truo.mutation.RepoIDCleared() {
		_spec.ClearField(teamrepo.FieldRepoID, field.TypeInt64)
	}
	_node = &TeamRepo{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamrepo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	truo.mutation.done = true
	return _node, nil
}