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
	"github.com/gitbundle/gitbundle/store/models/attachment"
	"github.com/gitbundle/gitbundle/store/models/predicate"
	"github.com/google/uuid"
)

// AttachmentUpdate is the builder for updating Attachment entities.
type AttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where appends a list predicates to the AttachmentUpdate builder.
func (au *AttachmentUpdate) Where(ps ...predicate.Attachment) *AttachmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUUID sets the "uuid" field.
func (au *AttachmentUpdate) SetUUID(u uuid.UUID) *AttachmentUpdate {
	au.mutation.SetUUID(u)
	return au
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableUUID(u *uuid.UUID) *AttachmentUpdate {
	if u != nil {
		au.SetUUID(*u)
	}
	return au
}

// SetRepoID sets the "repo_id" field.
func (au *AttachmentUpdate) SetRepoID(i int64) *AttachmentUpdate {
	au.mutation.ResetRepoID()
	au.mutation.SetRepoID(i)
	return au
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableRepoID(i *int64) *AttachmentUpdate {
	if i != nil {
		au.SetRepoID(*i)
	}
	return au
}

// AddRepoID adds i to the "repo_id" field.
func (au *AttachmentUpdate) AddRepoID(i int64) *AttachmentUpdate {
	au.mutation.AddRepoID(i)
	return au
}

// ClearRepoID clears the value of the "repo_id" field.
func (au *AttachmentUpdate) ClearRepoID() *AttachmentUpdate {
	au.mutation.ClearRepoID()
	return au
}

// SetUploaderID sets the "uploader_id" field.
func (au *AttachmentUpdate) SetUploaderID(i int64) *AttachmentUpdate {
	au.mutation.ResetUploaderID()
	au.mutation.SetUploaderID(i)
	return au
}

// SetNillableUploaderID sets the "uploader_id" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableUploaderID(i *int64) *AttachmentUpdate {
	if i != nil {
		au.SetUploaderID(*i)
	}
	return au
}

// AddUploaderID adds i to the "uploader_id" field.
func (au *AttachmentUpdate) AddUploaderID(i int64) *AttachmentUpdate {
	au.mutation.AddUploaderID(i)
	return au
}

// ClearUploaderID clears the value of the "uploader_id" field.
func (au *AttachmentUpdate) ClearUploaderID() *AttachmentUpdate {
	au.mutation.ClearUploaderID()
	return au
}

// SetName sets the "name" field.
func (au *AttachmentUpdate) SetName(s string) *AttachmentUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableName(s *string) *AttachmentUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *AttachmentUpdate) ClearName() *AttachmentUpdate {
	au.mutation.ClearName()
	return au
}

// SetDownloadCount sets the "download_count" field.
func (au *AttachmentUpdate) SetDownloadCount(i int64) *AttachmentUpdate {
	au.mutation.ResetDownloadCount()
	au.mutation.SetDownloadCount(i)
	return au
}

// SetNillableDownloadCount sets the "download_count" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableDownloadCount(i *int64) *AttachmentUpdate {
	if i != nil {
		au.SetDownloadCount(*i)
	}
	return au
}

// AddDownloadCount adds i to the "download_count" field.
func (au *AttachmentUpdate) AddDownloadCount(i int64) *AttachmentUpdate {
	au.mutation.AddDownloadCount(i)
	return au
}

// SetSize sets the "size" field.
func (au *AttachmentUpdate) SetSize(i int64) *AttachmentUpdate {
	au.mutation.ResetSize()
	au.mutation.SetSize(i)
	return au
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableSize(i *int64) *AttachmentUpdate {
	if i != nil {
		au.SetSize(*i)
	}
	return au
}

// AddSize adds i to the "size" field.
func (au *AttachmentUpdate) AddSize(i int64) *AttachmentUpdate {
	au.mutation.AddSize(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AttachmentUpdate) SetCreatedAt(t time.Time) *AttachmentUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableCreatedAt(t *time.Time) *AttachmentUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// Mutation returns the AttachmentMutation object of the builder.
func (au *AttachmentUpdate) Mutation() *AttachmentMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttachmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttachmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttachmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttachmentUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := attachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "Attachment.name": %w`, err)}
		}
	}
	return nil
}

func (au *AttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attachment.Table, attachment.Columns, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.SetField(attachment.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := au.mutation.RepoID(); ok {
		_spec.SetField(attachment.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedRepoID(); ok {
		_spec.AddField(attachment.FieldRepoID, field.TypeInt64, value)
	}
	if au.mutation.RepoIDCleared() {
		_spec.ClearField(attachment.FieldRepoID, field.TypeInt64)
	}
	if value, ok := au.mutation.UploaderID(); ok {
		_spec.SetField(attachment.FieldUploaderID, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUploaderID(); ok {
		_spec.AddField(attachment.FieldUploaderID, field.TypeInt64, value)
	}
	if au.mutation.UploaderIDCleared() {
		_spec.ClearField(attachment.FieldUploaderID, field.TypeInt64)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(attachment.FieldName, field.TypeString, value)
	}
	if au.mutation.NameCleared() {
		_spec.ClearField(attachment.FieldName, field.TypeString)
	}
	if value, ok := au.mutation.DownloadCount(); ok {
		_spec.SetField(attachment.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedDownloadCount(); ok {
		_spec.AddField(attachment.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := au.mutation.Size(); ok {
		_spec.SetField(attachment.FieldSize, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedSize(); ok {
		_spec.AddField(attachment.FieldSize, field.TypeInt64, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttachmentUpdateOne is the builder for updating a single Attachment entity.
type AttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentMutation
}

// SetUUID sets the "uuid" field.
func (auo *AttachmentUpdateOne) SetUUID(u uuid.UUID) *AttachmentUpdateOne {
	auo.mutation.SetUUID(u)
	return auo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableUUID(u *uuid.UUID) *AttachmentUpdateOne {
	if u != nil {
		auo.SetUUID(*u)
	}
	return auo
}

// SetRepoID sets the "repo_id" field.
func (auo *AttachmentUpdateOne) SetRepoID(i int64) *AttachmentUpdateOne {
	auo.mutation.ResetRepoID()
	auo.mutation.SetRepoID(i)
	return auo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableRepoID(i *int64) *AttachmentUpdateOne {
	if i != nil {
		auo.SetRepoID(*i)
	}
	return auo
}

// AddRepoID adds i to the "repo_id" field.
func (auo *AttachmentUpdateOne) AddRepoID(i int64) *AttachmentUpdateOne {
	auo.mutation.AddRepoID(i)
	return auo
}

// ClearRepoID clears the value of the "repo_id" field.
func (auo *AttachmentUpdateOne) ClearRepoID() *AttachmentUpdateOne {
	auo.mutation.ClearRepoID()
	return auo
}

// SetUploaderID sets the "uploader_id" field.
func (auo *AttachmentUpdateOne) SetUploaderID(i int64) *AttachmentUpdateOne {
	auo.mutation.ResetUploaderID()
	auo.mutation.SetUploaderID(i)
	return auo
}

// SetNillableUploaderID sets the "uploader_id" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableUploaderID(i *int64) *AttachmentUpdateOne {
	if i != nil {
		auo.SetUploaderID(*i)
	}
	return auo
}

// AddUploaderID adds i to the "uploader_id" field.
func (auo *AttachmentUpdateOne) AddUploaderID(i int64) *AttachmentUpdateOne {
	auo.mutation.AddUploaderID(i)
	return auo
}

// ClearUploaderID clears the value of the "uploader_id" field.
func (auo *AttachmentUpdateOne) ClearUploaderID() *AttachmentUpdateOne {
	auo.mutation.ClearUploaderID()
	return auo
}

// SetName sets the "name" field.
func (auo *AttachmentUpdateOne) SetName(s string) *AttachmentUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableName(s *string) *AttachmentUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *AttachmentUpdateOne) ClearName() *AttachmentUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetDownloadCount sets the "download_count" field.
func (auo *AttachmentUpdateOne) SetDownloadCount(i int64) *AttachmentUpdateOne {
	auo.mutation.ResetDownloadCount()
	auo.mutation.SetDownloadCount(i)
	return auo
}

// SetNillableDownloadCount sets the "download_count" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableDownloadCount(i *int64) *AttachmentUpdateOne {
	if i != nil {
		auo.SetDownloadCount(*i)
	}
	return auo
}

// AddDownloadCount adds i to the "download_count" field.
func (auo *AttachmentUpdateOne) AddDownloadCount(i int64) *AttachmentUpdateOne {
	auo.mutation.AddDownloadCount(i)
	return auo
}

// SetSize sets the "size" field.
func (auo *AttachmentUpdateOne) SetSize(i int64) *AttachmentUpdateOne {
	auo.mutation.ResetSize()
	auo.mutation.SetSize(i)
	return auo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableSize(i *int64) *AttachmentUpdateOne {
	if i != nil {
		auo.SetSize(*i)
	}
	return auo
}

// AddSize adds i to the "size" field.
func (auo *AttachmentUpdateOne) AddSize(i int64) *AttachmentUpdateOne {
	auo.mutation.AddSize(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AttachmentUpdateOne) SetCreatedAt(t time.Time) *AttachmentUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableCreatedAt(t *time.Time) *AttachmentUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// Mutation returns the AttachmentMutation object of the builder.
func (auo *AttachmentUpdateOne) Mutation() *AttachmentMutation {
	return auo.mutation
}

// Where appends a list predicates to the AttachmentUpdate builder.
func (auo *AttachmentUpdateOne) Where(ps ...predicate.Attachment) *AttachmentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttachmentUpdateOne) Select(field string, fields ...string) *AttachmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attachment entity.
func (auo *AttachmentUpdateOne) Save(ctx context.Context) (*Attachment, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttachmentUpdateOne) SaveX(ctx context.Context) *Attachment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttachmentUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := attachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "Attachment.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AttachmentUpdateOne) sqlSave(ctx context.Context) (_node *Attachment, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attachment.Table, attachment.Columns, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Attachment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for _, f := range fields {
			if !attachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.SetField(attachment.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := auo.mutation.RepoID(); ok {
		_spec.SetField(attachment.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedRepoID(); ok {
		_spec.AddField(attachment.FieldRepoID, field.TypeInt64, value)
	}
	if auo.mutation.RepoIDCleared() {
		_spec.ClearField(attachment.FieldRepoID, field.TypeInt64)
	}
	if value, ok := auo.mutation.UploaderID(); ok {
		_spec.SetField(attachment.FieldUploaderID, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUploaderID(); ok {
		_spec.AddField(attachment.FieldUploaderID, field.TypeInt64, value)
	}
	if auo.mutation.UploaderIDCleared() {
		_spec.ClearField(attachment.FieldUploaderID, field.TypeInt64)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(attachment.FieldName, field.TypeString, value)
	}
	if auo.mutation.NameCleared() {
		_spec.ClearField(attachment.FieldName, field.TypeString)
	}
	if value, ok := auo.mutation.DownloadCount(); ok {
		_spec.SetField(attachment.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedDownloadCount(); ok {
		_spec.AddField(attachment.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.Size(); ok {
		_spec.SetField(attachment.FieldSize, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedSize(); ok {
		_spec.AddField(attachment.FieldSize, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Attachment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}