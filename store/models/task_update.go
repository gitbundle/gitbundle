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
	"github.com/gitbundle/gitbundle/store/models/predicate"
	"github.com/gitbundle/gitbundle/store/models/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDoerID sets the "doer_id" field.
func (tu *TaskUpdate) SetDoerID(i int64) *TaskUpdate {
	tu.mutation.ResetDoerID()
	tu.mutation.SetDoerID(i)
	return tu
}

// SetNillableDoerID sets the "doer_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDoerID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetDoerID(*i)
	}
	return tu
}

// AddDoerID adds i to the "doer_id" field.
func (tu *TaskUpdate) AddDoerID(i int64) *TaskUpdate {
	tu.mutation.AddDoerID(i)
	return tu
}

// ClearDoerID clears the value of the "doer_id" field.
func (tu *TaskUpdate) ClearDoerID() *TaskUpdate {
	tu.mutation.ClearDoerID()
	return tu
}

// SetOwnerID sets the "owner_id" field.
func (tu *TaskUpdate) SetOwnerID(i int64) *TaskUpdate {
	tu.mutation.ResetOwnerID()
	tu.mutation.SetOwnerID(i)
	return tu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOwnerID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetOwnerID(*i)
	}
	return tu
}

// AddOwnerID adds i to the "owner_id" field.
func (tu *TaskUpdate) AddOwnerID(i int64) *TaskUpdate {
	tu.mutation.AddOwnerID(i)
	return tu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (tu *TaskUpdate) ClearOwnerID() *TaskUpdate {
	tu.mutation.ClearOwnerID()
	return tu
}

// SetRepoID sets the "repo_id" field.
func (tu *TaskUpdate) SetRepoID(i int64) *TaskUpdate {
	tu.mutation.ResetRepoID()
	tu.mutation.SetRepoID(i)
	return tu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRepoID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetRepoID(*i)
	}
	return tu
}

// AddRepoID adds i to the "repo_id" field.
func (tu *TaskUpdate) AddRepoID(i int64) *TaskUpdate {
	tu.mutation.AddRepoID(i)
	return tu
}

// ClearRepoID clears the value of the "repo_id" field.
func (tu *TaskUpdate) ClearRepoID() *TaskUpdate {
	tu.mutation.ClearRepoID()
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(i int) *TaskUpdate {
	tu.mutation.ResetType()
	tu.mutation.SetType(i)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableType(i *int) *TaskUpdate {
	if i != nil {
		tu.SetType(*i)
	}
	return tu
}

// AddType adds i to the "type" field.
func (tu *TaskUpdate) AddType(i int) *TaskUpdate {
	tu.mutation.AddType(i)
	return tu
}

// ClearType clears the value of the "type" field.
func (tu *TaskUpdate) ClearType() *TaskUpdate {
	tu.mutation.ClearType()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(i int) *TaskUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(i)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(i *int) *TaskUpdate {
	if i != nil {
		tu.SetStatus(*i)
	}
	return tu
}

// AddStatus adds i to the "status" field.
func (tu *TaskUpdate) AddStatus(i int) *TaskUpdate {
	tu.mutation.AddStatus(i)
	return tu
}

// ClearStatus clears the value of the "status" field.
func (tu *TaskUpdate) ClearStatus() *TaskUpdate {
	tu.mutation.ClearStatus()
	return tu
}

// SetStartedAt sets the "started_at" field.
func (tu *TaskUpdate) SetStartedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetStartedAt(t)
	return tu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStartedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetStartedAt(*t)
	}
	return tu
}

// ClearStartedAt clears the value of the "started_at" field.
func (tu *TaskUpdate) ClearStartedAt() *TaskUpdate {
	tu.mutation.ClearStartedAt()
	return tu
}

// SetEndedAt sets the "ended_at" field.
func (tu *TaskUpdate) SetEndedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetEndedAt(t)
	return tu
}

// SetNillableEndedAt sets the "ended_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableEndedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetEndedAt(*t)
	}
	return tu
}

// ClearEndedAt clears the value of the "ended_at" field.
func (tu *TaskUpdate) ClearEndedAt() *TaskUpdate {
	tu.mutation.ClearEndedAt()
	return tu
}

// SetPayloadContent sets the "payload_content" field.
func (tu *TaskUpdate) SetPayloadContent(s string) *TaskUpdate {
	tu.mutation.SetPayloadContent(s)
	return tu
}

// SetNillablePayloadContent sets the "payload_content" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePayloadContent(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPayloadContent(*s)
	}
	return tu
}

// ClearPayloadContent clears the value of the "payload_content" field.
func (tu *TaskUpdate) ClearPayloadContent() *TaskUpdate {
	tu.mutation.ClearPayloadContent()
	return tu
}

// SetMessage sets the "message" field.
func (tu *TaskUpdate) SetMessage(s string) *TaskUpdate {
	tu.mutation.SetMessage(s)
	return tu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableMessage(s *string) *TaskUpdate {
	if s != nil {
		tu.SetMessage(*s)
	}
	return tu
}

// ClearMessage clears the value of the "message" field.
func (tu *TaskUpdate) ClearMessage() *TaskUpdate {
	tu.mutation.ClearMessage()
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.DoerID(); ok {
		_spec.SetField(task.FieldDoerID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDoerID(); ok {
		_spec.AddField(task.FieldDoerID, field.TypeInt64, value)
	}
	if tu.mutation.DoerIDCleared() {
		_spec.ClearField(task.FieldDoerID, field.TypeInt64)
	}
	if value, ok := tu.mutation.OwnerID(); ok {
		_spec.SetField(task.FieldOwnerID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedOwnerID(); ok {
		_spec.AddField(task.FieldOwnerID, field.TypeInt64, value)
	}
	if tu.mutation.OwnerIDCleared() {
		_spec.ClearField(task.FieldOwnerID, field.TypeInt64)
	}
	if value, ok := tu.mutation.RepoID(); ok {
		_spec.SetField(task.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedRepoID(); ok {
		_spec.AddField(task.FieldRepoID, field.TypeInt64, value)
	}
	if tu.mutation.RepoIDCleared() {
		_spec.ClearField(task.FieldRepoID, field.TypeInt64)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedType(); ok {
		_spec.AddField(task.FieldType, field.TypeInt, value)
	}
	if tu.mutation.TypeCleared() {
		_spec.ClearField(task.FieldType, field.TypeInt)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.AddField(task.FieldStatus, field.TypeInt, value)
	}
	if tu.mutation.StatusCleared() {
		_spec.ClearField(task.FieldStatus, field.TypeInt)
	}
	if value, ok := tu.mutation.StartedAt(); ok {
		_spec.SetField(task.FieldStartedAt, field.TypeTime, value)
	}
	if tu.mutation.StartedAtCleared() {
		_spec.ClearField(task.FieldStartedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.EndedAt(); ok {
		_spec.SetField(task.FieldEndedAt, field.TypeTime, value)
	}
	if tu.mutation.EndedAtCleared() {
		_spec.ClearField(task.FieldEndedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.PayloadContent(); ok {
		_spec.SetField(task.FieldPayloadContent, field.TypeString, value)
	}
	if tu.mutation.PayloadContentCleared() {
		_spec.ClearField(task.FieldPayloadContent, field.TypeString)
	}
	if value, ok := tu.mutation.Message(); ok {
		_spec.SetField(task.FieldMessage, field.TypeString, value)
	}
	if tu.mutation.MessageCleared() {
		_spec.ClearField(task.FieldMessage, field.TypeString)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetDoerID sets the "doer_id" field.
func (tuo *TaskUpdateOne) SetDoerID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetDoerID()
	tuo.mutation.SetDoerID(i)
	return tuo
}

// SetNillableDoerID sets the "doer_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDoerID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetDoerID(*i)
	}
	return tuo
}

// AddDoerID adds i to the "doer_id" field.
func (tuo *TaskUpdateOne) AddDoerID(i int64) *TaskUpdateOne {
	tuo.mutation.AddDoerID(i)
	return tuo
}

// ClearDoerID clears the value of the "doer_id" field.
func (tuo *TaskUpdateOne) ClearDoerID() *TaskUpdateOne {
	tuo.mutation.ClearDoerID()
	return tuo
}

// SetOwnerID sets the "owner_id" field.
func (tuo *TaskUpdateOne) SetOwnerID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetOwnerID()
	tuo.mutation.SetOwnerID(i)
	return tuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOwnerID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetOwnerID(*i)
	}
	return tuo
}

// AddOwnerID adds i to the "owner_id" field.
func (tuo *TaskUpdateOne) AddOwnerID(i int64) *TaskUpdateOne {
	tuo.mutation.AddOwnerID(i)
	return tuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (tuo *TaskUpdateOne) ClearOwnerID() *TaskUpdateOne {
	tuo.mutation.ClearOwnerID()
	return tuo
}

// SetRepoID sets the "repo_id" field.
func (tuo *TaskUpdateOne) SetRepoID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetRepoID()
	tuo.mutation.SetRepoID(i)
	return tuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRepoID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetRepoID(*i)
	}
	return tuo
}

// AddRepoID adds i to the "repo_id" field.
func (tuo *TaskUpdateOne) AddRepoID(i int64) *TaskUpdateOne {
	tuo.mutation.AddRepoID(i)
	return tuo
}

// ClearRepoID clears the value of the "repo_id" field.
func (tuo *TaskUpdateOne) ClearRepoID() *TaskUpdateOne {
	tuo.mutation.ClearRepoID()
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(i int) *TaskUpdateOne {
	tuo.mutation.ResetType()
	tuo.mutation.SetType(i)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableType(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetType(*i)
	}
	return tuo
}

// AddType adds i to the "type" field.
func (tuo *TaskUpdateOne) AddType(i int) *TaskUpdateOne {
	tuo.mutation.AddType(i)
	return tuo
}

// ClearType clears the value of the "type" field.
func (tuo *TaskUpdateOne) ClearType() *TaskUpdateOne {
	tuo.mutation.ClearType()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(i int) *TaskUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(i)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetStatus(*i)
	}
	return tuo
}

// AddStatus adds i to the "status" field.
func (tuo *TaskUpdateOne) AddStatus(i int) *TaskUpdateOne {
	tuo.mutation.AddStatus(i)
	return tuo
}

// ClearStatus clears the value of the "status" field.
func (tuo *TaskUpdateOne) ClearStatus() *TaskUpdateOne {
	tuo.mutation.ClearStatus()
	return tuo
}

// SetStartedAt sets the "started_at" field.
func (tuo *TaskUpdateOne) SetStartedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetStartedAt(t)
	return tuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStartedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetStartedAt(*t)
	}
	return tuo
}

// ClearStartedAt clears the value of the "started_at" field.
func (tuo *TaskUpdateOne) ClearStartedAt() *TaskUpdateOne {
	tuo.mutation.ClearStartedAt()
	return tuo
}

// SetEndedAt sets the "ended_at" field.
func (tuo *TaskUpdateOne) SetEndedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetEndedAt(t)
	return tuo
}

// SetNillableEndedAt sets the "ended_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableEndedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetEndedAt(*t)
	}
	return tuo
}

// ClearEndedAt clears the value of the "ended_at" field.
func (tuo *TaskUpdateOne) ClearEndedAt() *TaskUpdateOne {
	tuo.mutation.ClearEndedAt()
	return tuo
}

// SetPayloadContent sets the "payload_content" field.
func (tuo *TaskUpdateOne) SetPayloadContent(s string) *TaskUpdateOne {
	tuo.mutation.SetPayloadContent(s)
	return tuo
}

// SetNillablePayloadContent sets the "payload_content" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePayloadContent(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPayloadContent(*s)
	}
	return tuo
}

// ClearPayloadContent clears the value of the "payload_content" field.
func (tuo *TaskUpdateOne) ClearPayloadContent() *TaskUpdateOne {
	tuo.mutation.ClearPayloadContent()
	return tuo
}

// SetMessage sets the "message" field.
func (tuo *TaskUpdateOne) SetMessage(s string) *TaskUpdateOne {
	tuo.mutation.SetMessage(s)
	return tuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableMessage(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetMessage(*s)
	}
	return tuo
}

// ClearMessage clears the value of the "message" field.
func (tuo *TaskUpdateOne) ClearMessage() *TaskUpdateOne {
	tuo.mutation.ClearMessage()
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.DoerID(); ok {
		_spec.SetField(task.FieldDoerID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDoerID(); ok {
		_spec.AddField(task.FieldDoerID, field.TypeInt64, value)
	}
	if tuo.mutation.DoerIDCleared() {
		_spec.ClearField(task.FieldDoerID, field.TypeInt64)
	}
	if value, ok := tuo.mutation.OwnerID(); ok {
		_spec.SetField(task.FieldOwnerID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedOwnerID(); ok {
		_spec.AddField(task.FieldOwnerID, field.TypeInt64, value)
	}
	if tuo.mutation.OwnerIDCleared() {
		_spec.ClearField(task.FieldOwnerID, field.TypeInt64)
	}
	if value, ok := tuo.mutation.RepoID(); ok {
		_spec.SetField(task.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedRepoID(); ok {
		_spec.AddField(task.FieldRepoID, field.TypeInt64, value)
	}
	if tuo.mutation.RepoIDCleared() {
		_spec.ClearField(task.FieldRepoID, field.TypeInt64)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedType(); ok {
		_spec.AddField(task.FieldType, field.TypeInt, value)
	}
	if tuo.mutation.TypeCleared() {
		_spec.ClearField(task.FieldType, field.TypeInt)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.AddField(task.FieldStatus, field.TypeInt, value)
	}
	if tuo.mutation.StatusCleared() {
		_spec.ClearField(task.FieldStatus, field.TypeInt)
	}
	if value, ok := tuo.mutation.StartedAt(); ok {
		_spec.SetField(task.FieldStartedAt, field.TypeTime, value)
	}
	if tuo.mutation.StartedAtCleared() {
		_spec.ClearField(task.FieldStartedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.EndedAt(); ok {
		_spec.SetField(task.FieldEndedAt, field.TypeTime, value)
	}
	if tuo.mutation.EndedAtCleared() {
		_spec.ClearField(task.FieldEndedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.PayloadContent(); ok {
		_spec.SetField(task.FieldPayloadContent, field.TypeString, value)
	}
	if tuo.mutation.PayloadContentCleared() {
		_spec.ClearField(task.FieldPayloadContent, field.TypeString)
	}
	if value, ok := tuo.mutation.Message(); ok {
		_spec.SetField(task.FieldMessage, field.TypeString, value)
	}
	if tuo.mutation.MessageCleared() {
		_spec.ClearField(task.FieldMessage, field.TypeString)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}