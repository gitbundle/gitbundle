// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/predicate"
	"github.com/gitbundle/gitbundle/store/models/repoindexerstatus"
)

// RepoIndexerStatusUpdate is the builder for updating RepoIndexerStatus entities.
type RepoIndexerStatusUpdate struct {
	config
	hooks    []Hook
	mutation *RepoIndexerStatusMutation
}

// Where appends a list predicates to the RepoIndexerStatusUpdate builder.
func (risu *RepoIndexerStatusUpdate) Where(ps ...predicate.RepoIndexerStatus) *RepoIndexerStatusUpdate {
	risu.mutation.Where(ps...)
	return risu
}

// SetRepoID sets the "repo_id" field.
func (risu *RepoIndexerStatusUpdate) SetRepoID(i int64) *RepoIndexerStatusUpdate {
	risu.mutation.ResetRepoID()
	risu.mutation.SetRepoID(i)
	return risu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (risu *RepoIndexerStatusUpdate) SetNillableRepoID(i *int64) *RepoIndexerStatusUpdate {
	if i != nil {
		risu.SetRepoID(*i)
	}
	return risu
}

// AddRepoID adds i to the "repo_id" field.
func (risu *RepoIndexerStatusUpdate) AddRepoID(i int64) *RepoIndexerStatusUpdate {
	risu.mutation.AddRepoID(i)
	return risu
}

// ClearRepoID clears the value of the "repo_id" field.
func (risu *RepoIndexerStatusUpdate) ClearRepoID() *RepoIndexerStatusUpdate {
	risu.mutation.ClearRepoID()
	return risu
}

// SetCommitSha sets the "commit_sha" field.
func (risu *RepoIndexerStatusUpdate) SetCommitSha(s string) *RepoIndexerStatusUpdate {
	risu.mutation.SetCommitSha(s)
	return risu
}

// SetNillableCommitSha sets the "commit_sha" field if the given value is not nil.
func (risu *RepoIndexerStatusUpdate) SetNillableCommitSha(s *string) *RepoIndexerStatusUpdate {
	if s != nil {
		risu.SetCommitSha(*s)
	}
	return risu
}

// ClearCommitSha clears the value of the "commit_sha" field.
func (risu *RepoIndexerStatusUpdate) ClearCommitSha() *RepoIndexerStatusUpdate {
	risu.mutation.ClearCommitSha()
	return risu
}

// SetIndexerType sets the "indexer_type" field.
func (risu *RepoIndexerStatusUpdate) SetIndexerType(i int) *RepoIndexerStatusUpdate {
	risu.mutation.ResetIndexerType()
	risu.mutation.SetIndexerType(i)
	return risu
}

// SetNillableIndexerType sets the "indexer_type" field if the given value is not nil.
func (risu *RepoIndexerStatusUpdate) SetNillableIndexerType(i *int) *RepoIndexerStatusUpdate {
	if i != nil {
		risu.SetIndexerType(*i)
	}
	return risu
}

// AddIndexerType adds i to the "indexer_type" field.
func (risu *RepoIndexerStatusUpdate) AddIndexerType(i int) *RepoIndexerStatusUpdate {
	risu.mutation.AddIndexerType(i)
	return risu
}

// Mutation returns the RepoIndexerStatusMutation object of the builder.
func (risu *RepoIndexerStatusUpdate) Mutation() *RepoIndexerStatusMutation {
	return risu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (risu *RepoIndexerStatusUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, risu.sqlSave, risu.mutation, risu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (risu *RepoIndexerStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := risu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (risu *RepoIndexerStatusUpdate) Exec(ctx context.Context) error {
	_, err := risu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (risu *RepoIndexerStatusUpdate) ExecX(ctx context.Context) {
	if err := risu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (risu *RepoIndexerStatusUpdate) check() error {
	if v, ok := risu.mutation.CommitSha(); ok {
		if err := repoindexerstatus.CommitShaValidator(v); err != nil {
			return &ValidationError{Name: "commit_sha", err: fmt.Errorf(`models: validator failed for field "RepoIndexerStatus.commit_sha": %w`, err)}
		}
	}
	return nil
}

func (risu *RepoIndexerStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := risu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(repoindexerstatus.Table, repoindexerstatus.Columns, sqlgraph.NewFieldSpec(repoindexerstatus.FieldID, field.TypeInt))
	if ps := risu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := risu.mutation.RepoID(); ok {
		_spec.SetField(repoindexerstatus.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := risu.mutation.AddedRepoID(); ok {
		_spec.AddField(repoindexerstatus.FieldRepoID, field.TypeInt64, value)
	}
	if risu.mutation.RepoIDCleared() {
		_spec.ClearField(repoindexerstatus.FieldRepoID, field.TypeInt64)
	}
	if value, ok := risu.mutation.CommitSha(); ok {
		_spec.SetField(repoindexerstatus.FieldCommitSha, field.TypeString, value)
	}
	if risu.mutation.CommitShaCleared() {
		_spec.ClearField(repoindexerstatus.FieldCommitSha, field.TypeString)
	}
	if value, ok := risu.mutation.IndexerType(); ok {
		_spec.SetField(repoindexerstatus.FieldIndexerType, field.TypeInt, value)
	}
	if value, ok := risu.mutation.AddedIndexerType(); ok {
		_spec.AddField(repoindexerstatus.FieldIndexerType, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, risu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repoindexerstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	risu.mutation.done = true
	return n, nil
}

// RepoIndexerStatusUpdateOne is the builder for updating a single RepoIndexerStatus entity.
type RepoIndexerStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepoIndexerStatusMutation
}

// SetRepoID sets the "repo_id" field.
func (risuo *RepoIndexerStatusUpdateOne) SetRepoID(i int64) *RepoIndexerStatusUpdateOne {
	risuo.mutation.ResetRepoID()
	risuo.mutation.SetRepoID(i)
	return risuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (risuo *RepoIndexerStatusUpdateOne) SetNillableRepoID(i *int64) *RepoIndexerStatusUpdateOne {
	if i != nil {
		risuo.SetRepoID(*i)
	}
	return risuo
}

// AddRepoID adds i to the "repo_id" field.
func (risuo *RepoIndexerStatusUpdateOne) AddRepoID(i int64) *RepoIndexerStatusUpdateOne {
	risuo.mutation.AddRepoID(i)
	return risuo
}

// ClearRepoID clears the value of the "repo_id" field.
func (risuo *RepoIndexerStatusUpdateOne) ClearRepoID() *RepoIndexerStatusUpdateOne {
	risuo.mutation.ClearRepoID()
	return risuo
}

// SetCommitSha sets the "commit_sha" field.
func (risuo *RepoIndexerStatusUpdateOne) SetCommitSha(s string) *RepoIndexerStatusUpdateOne {
	risuo.mutation.SetCommitSha(s)
	return risuo
}

// SetNillableCommitSha sets the "commit_sha" field if the given value is not nil.
func (risuo *RepoIndexerStatusUpdateOne) SetNillableCommitSha(s *string) *RepoIndexerStatusUpdateOne {
	if s != nil {
		risuo.SetCommitSha(*s)
	}
	return risuo
}

// ClearCommitSha clears the value of the "commit_sha" field.
func (risuo *RepoIndexerStatusUpdateOne) ClearCommitSha() *RepoIndexerStatusUpdateOne {
	risuo.mutation.ClearCommitSha()
	return risuo
}

// SetIndexerType sets the "indexer_type" field.
func (risuo *RepoIndexerStatusUpdateOne) SetIndexerType(i int) *RepoIndexerStatusUpdateOne {
	risuo.mutation.ResetIndexerType()
	risuo.mutation.SetIndexerType(i)
	return risuo
}

// SetNillableIndexerType sets the "indexer_type" field if the given value is not nil.
func (risuo *RepoIndexerStatusUpdateOne) SetNillableIndexerType(i *int) *RepoIndexerStatusUpdateOne {
	if i != nil {
		risuo.SetIndexerType(*i)
	}
	return risuo
}

// AddIndexerType adds i to the "indexer_type" field.
func (risuo *RepoIndexerStatusUpdateOne) AddIndexerType(i int) *RepoIndexerStatusUpdateOne {
	risuo.mutation.AddIndexerType(i)
	return risuo
}

// Mutation returns the RepoIndexerStatusMutation object of the builder.
func (risuo *RepoIndexerStatusUpdateOne) Mutation() *RepoIndexerStatusMutation {
	return risuo.mutation
}

// Where appends a list predicates to the RepoIndexerStatusUpdate builder.
func (risuo *RepoIndexerStatusUpdateOne) Where(ps ...predicate.RepoIndexerStatus) *RepoIndexerStatusUpdateOne {
	risuo.mutation.Where(ps...)
	return risuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (risuo *RepoIndexerStatusUpdateOne) Select(field string, fields ...string) *RepoIndexerStatusUpdateOne {
	risuo.fields = append([]string{field}, fields...)
	return risuo
}

// Save executes the query and returns the updated RepoIndexerStatus entity.
func (risuo *RepoIndexerStatusUpdateOne) Save(ctx context.Context) (*RepoIndexerStatus, error) {
	return withHooks(ctx, risuo.sqlSave, risuo.mutation, risuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (risuo *RepoIndexerStatusUpdateOne) SaveX(ctx context.Context) *RepoIndexerStatus {
	node, err := risuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (risuo *RepoIndexerStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := risuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (risuo *RepoIndexerStatusUpdateOne) ExecX(ctx context.Context) {
	if err := risuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (risuo *RepoIndexerStatusUpdateOne) check() error {
	if v, ok := risuo.mutation.CommitSha(); ok {
		if err := repoindexerstatus.CommitShaValidator(v); err != nil {
			return &ValidationError{Name: "commit_sha", err: fmt.Errorf(`models: validator failed for field "RepoIndexerStatus.commit_sha": %w`, err)}
		}
	}
	return nil
}

func (risuo *RepoIndexerStatusUpdateOne) sqlSave(ctx context.Context) (_node *RepoIndexerStatus, err error) {
	if err := risuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(repoindexerstatus.Table, repoindexerstatus.Columns, sqlgraph.NewFieldSpec(repoindexerstatus.FieldID, field.TypeInt))
	id, ok := risuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "RepoIndexerStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := risuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repoindexerstatus.FieldID)
		for _, f := range fields {
			if !repoindexerstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != repoindexerstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := risuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := risuo.mutation.RepoID(); ok {
		_spec.SetField(repoindexerstatus.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := risuo.mutation.AddedRepoID(); ok {
		_spec.AddField(repoindexerstatus.FieldRepoID, field.TypeInt64, value)
	}
	if risuo.mutation.RepoIDCleared() {
		_spec.ClearField(repoindexerstatus.FieldRepoID, field.TypeInt64)
	}
	if value, ok := risuo.mutation.CommitSha(); ok {
		_spec.SetField(repoindexerstatus.FieldCommitSha, field.TypeString, value)
	}
	if risuo.mutation.CommitShaCleared() {
		_spec.ClearField(repoindexerstatus.FieldCommitSha, field.TypeString)
	}
	if value, ok := risuo.mutation.IndexerType(); ok {
		_spec.SetField(repoindexerstatus.FieldIndexerType, field.TypeInt, value)
	}
	if value, ok := risuo.mutation.AddedIndexerType(); ok {
		_spec.AddField(repoindexerstatus.FieldIndexerType, field.TypeInt, value)
	}
	_node = &RepoIndexerStatus{config: risuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, risuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repoindexerstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	risuo.mutation.done = true
	return _node, nil
}