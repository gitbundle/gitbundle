// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/store/models/pkgblobupload"
)

// PkgBlobUploadCreate is the builder for creating a PkgBlobUpload entity.
type PkgBlobUploadCreate struct {
	config
	mutation *PkgBlobUploadMutation
	hooks    []Hook
}

// SetBlobUploadID sets the "blob_upload_id" field.
func (pbuc *PkgBlobUploadCreate) SetBlobUploadID(s string) *PkgBlobUploadCreate {
	pbuc.mutation.SetBlobUploadID(s)
	return pbuc
}

// SetBytesReceived sets the "bytes_received" field.
func (pbuc *PkgBlobUploadCreate) SetBytesReceived(i int64) *PkgBlobUploadCreate {
	pbuc.mutation.SetBytesReceived(i)
	return pbuc
}

// SetHashStateBytes sets the "hash_state_bytes" field.
func (pbuc *PkgBlobUploadCreate) SetHashStateBytes(b []byte) *PkgBlobUploadCreate {
	pbuc.mutation.SetHashStateBytes(b)
	return pbuc
}

// SetCreatedAt sets the "created_at" field.
func (pbuc *PkgBlobUploadCreate) SetCreatedAt(t time.Time) *PkgBlobUploadCreate {
	pbuc.mutation.SetCreatedAt(t)
	return pbuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbuc *PkgBlobUploadCreate) SetNillableCreatedAt(t *time.Time) *PkgBlobUploadCreate {
	if t != nil {
		pbuc.SetCreatedAt(*t)
	}
	return pbuc
}

// SetUpdatedAt sets the "updated_at" field.
func (pbuc *PkgBlobUploadCreate) SetUpdatedAt(t time.Time) *PkgBlobUploadCreate {
	pbuc.mutation.SetUpdatedAt(t)
	return pbuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pbuc *PkgBlobUploadCreate) SetNillableUpdatedAt(t *time.Time) *PkgBlobUploadCreate {
	if t != nil {
		pbuc.SetUpdatedAt(*t)
	}
	return pbuc
}

// Mutation returns the PkgBlobUploadMutation object of the builder.
func (pbuc *PkgBlobUploadCreate) Mutation() *PkgBlobUploadMutation {
	return pbuc.mutation
}

// Save creates the PkgBlobUpload in the database.
func (pbuc *PkgBlobUploadCreate) Save(ctx context.Context) (*PkgBlobUpload, error) {
	pbuc.defaults()
	return withHooks(ctx, pbuc.sqlSave, pbuc.mutation, pbuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pbuc *PkgBlobUploadCreate) SaveX(ctx context.Context) *PkgBlobUpload {
	v, err := pbuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbuc *PkgBlobUploadCreate) Exec(ctx context.Context) error {
	_, err := pbuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbuc *PkgBlobUploadCreate) ExecX(ctx context.Context) {
	if err := pbuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbuc *PkgBlobUploadCreate) defaults() {
	if _, ok := pbuc.mutation.CreatedAt(); !ok {
		v := pkgblobupload.DefaultCreatedAt()
		pbuc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pbuc *PkgBlobUploadCreate) check() error {
	if _, ok := pbuc.mutation.BlobUploadID(); !ok {
		return &ValidationError{Name: "blob_upload_id", err: errors.New(`models: missing required field "PkgBlobUpload.blob_upload_id"`)}
	}
	if v, ok := pbuc.mutation.BlobUploadID(); ok {
		if err := pkgblobupload.BlobUploadIDValidator(v); err != nil {
			return &ValidationError{Name: "blob_upload_id", err: fmt.Errorf(`models: validator failed for field "PkgBlobUpload.blob_upload_id": %w`, err)}
		}
	}
	if _, ok := pbuc.mutation.BytesReceived(); !ok {
		return &ValidationError{Name: "bytes_received", err: errors.New(`models: missing required field "PkgBlobUpload.bytes_received"`)}
	}
	if _, ok := pbuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "PkgBlobUpload.created_at"`)}
	}
	return nil
}

func (pbuc *PkgBlobUploadCreate) sqlSave(ctx context.Context) (*PkgBlobUpload, error) {
	if err := pbuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pbuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pbuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pbuc.mutation.id = &_node.ID
	pbuc.mutation.done = true
	return _node, nil
}

func (pbuc *PkgBlobUploadCreate) createSpec() (*PkgBlobUpload, *sqlgraph.CreateSpec) {
	var (
		_node = &PkgBlobUpload{config: pbuc.config}
		_spec = sqlgraph.NewCreateSpec(pkgblobupload.Table, sqlgraph.NewFieldSpec(pkgblobupload.FieldID, field.TypeInt))
	)
	if value, ok := pbuc.mutation.BlobUploadID(); ok {
		_spec.SetField(pkgblobupload.FieldBlobUploadID, field.TypeString, value)
		_node.BlobUploadID = value
	}
	if value, ok := pbuc.mutation.BytesReceived(); ok {
		_spec.SetField(pkgblobupload.FieldBytesReceived, field.TypeInt64, value)
		_node.BytesReceived = value
	}
	if value, ok := pbuc.mutation.HashStateBytes(); ok {
		_spec.SetField(pkgblobupload.FieldHashStateBytes, field.TypeBytes, value)
		_node.HashStateBytes = value
	}
	if value, ok := pbuc.mutation.CreatedAt(); ok {
		_spec.SetField(pkgblobupload.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pbuc.mutation.UpdatedAt(); ok {
		_spec.SetField(pkgblobupload.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// PkgBlobUploadCreateBulk is the builder for creating many PkgBlobUpload entities in bulk.
type PkgBlobUploadCreateBulk struct {
	config
	err      error
	builders []*PkgBlobUploadCreate
}

// Save creates the PkgBlobUpload entities in the database.
func (pbucb *PkgBlobUploadCreateBulk) Save(ctx context.Context) ([]*PkgBlobUpload, error) {
	if pbucb.err != nil {
		return nil, pbucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pbucb.builders))
	nodes := make([]*PkgBlobUpload, len(pbucb.builders))
	mutators := make([]Mutator, len(pbucb.builders))
	for i := range pbucb.builders {
		func(i int, root context.Context) {
			builder := pbucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PkgBlobUploadMutation)
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
					_, err = mutators[i+1].Mutate(root, pbucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pbucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pbucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pbucb *PkgBlobUploadCreateBulk) SaveX(ctx context.Context) []*PkgBlobUpload {
	v, err := pbucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbucb *PkgBlobUploadCreateBulk) Exec(ctx context.Context) error {
	_, err := pbucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbucb *PkgBlobUploadCreateBulk) ExecX(ctx context.Context) {
	if err := pbucb.Exec(ctx); err != nil {
		panic(err)
	}
}