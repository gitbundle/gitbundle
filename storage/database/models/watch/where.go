// Code generated by ent, DO NOT EDIT.

package watch

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldUserID, v))
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldRepoID, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldMode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Watch {
	return predicate.Watch(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Watch {
	return predicate.Watch(sql.FieldNotNull(FieldUserID))
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldRepoID, v))
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldRepoID, v))
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldRepoID, vs...))
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...int64) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldRepoID, vs...))
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldRepoID, v))
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldRepoID, v))
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldRepoID, v))
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v int64) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldRepoID, v))
}

// RepoIDIsNil applies the IsNil predicate on the "repo_id" field.
func RepoIDIsNil() predicate.Watch {
	return predicate.Watch(sql.FieldIsNull(FieldRepoID))
}

// RepoIDNotNil applies the NotNil predicate on the "repo_id" field.
func RepoIDNotNil() predicate.Watch {
	return predicate.Watch(sql.FieldNotNull(FieldRepoID))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...int16) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...int16) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v int16) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldMode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Watch {
	return predicate.Watch(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Watch {
	return predicate.Watch(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Watch {
	return predicate.Watch(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Watch) predicate.Watch {
	return predicate.Watch(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Watch) predicate.Watch {
	return predicate.Watch(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Watch) predicate.Watch {
	return predicate.Watch(sql.NotPredicates(p))
}