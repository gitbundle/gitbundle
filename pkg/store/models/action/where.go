// Code generated by ent, DO NOT EDIT.

package action

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/server/pkg/store/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldUserID, v))
}

// OpType applies equality check predicate on the "op_type" field. It's identical to OpTypeEQ.
func OpType(v int) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldOpType, v))
}

// ActUserID applies equality check predicate on the "act_user_id" field. It's identical to ActUserIDEQ.
func ActUserID(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldActUserID, v))
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldRepoID, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldIsDeleted, v))
}

// RefName applies equality check predicate on the "ref_name" field. It's identical to RefNameEQ.
func RefName(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldRefName, v))
}

// IsPrivate applies equality check predicate on the "is_private" field. It's identical to IsPrivateEQ.
func IsPrivate(v bool) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldIsPrivate, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldUserID))
}

// OpTypeEQ applies the EQ predicate on the "op_type" field.
func OpTypeEQ(v int) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldOpType, v))
}

// OpTypeNEQ applies the NEQ predicate on the "op_type" field.
func OpTypeNEQ(v int) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldOpType, v))
}

// OpTypeIn applies the In predicate on the "op_type" field.
func OpTypeIn(vs ...int) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldOpType, vs...))
}

// OpTypeNotIn applies the NotIn predicate on the "op_type" field.
func OpTypeNotIn(vs ...int) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldOpType, vs...))
}

// OpTypeGT applies the GT predicate on the "op_type" field.
func OpTypeGT(v int) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldOpType, v))
}

// OpTypeGTE applies the GTE predicate on the "op_type" field.
func OpTypeGTE(v int) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldOpType, v))
}

// OpTypeLT applies the LT predicate on the "op_type" field.
func OpTypeLT(v int) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldOpType, v))
}

// OpTypeLTE applies the LTE predicate on the "op_type" field.
func OpTypeLTE(v int) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldOpType, v))
}

// OpTypeIsNil applies the IsNil predicate on the "op_type" field.
func OpTypeIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldOpType))
}

// OpTypeNotNil applies the NotNil predicate on the "op_type" field.
func OpTypeNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldOpType))
}

// ActUserIDEQ applies the EQ predicate on the "act_user_id" field.
func ActUserIDEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldActUserID, v))
}

// ActUserIDNEQ applies the NEQ predicate on the "act_user_id" field.
func ActUserIDNEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldActUserID, v))
}

// ActUserIDIn applies the In predicate on the "act_user_id" field.
func ActUserIDIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldActUserID, vs...))
}

// ActUserIDNotIn applies the NotIn predicate on the "act_user_id" field.
func ActUserIDNotIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldActUserID, vs...))
}

// ActUserIDGT applies the GT predicate on the "act_user_id" field.
func ActUserIDGT(v int64) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldActUserID, v))
}

// ActUserIDGTE applies the GTE predicate on the "act_user_id" field.
func ActUserIDGTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldActUserID, v))
}

// ActUserIDLT applies the LT predicate on the "act_user_id" field.
func ActUserIDLT(v int64) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldActUserID, v))
}

// ActUserIDLTE applies the LTE predicate on the "act_user_id" field.
func ActUserIDLTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldActUserID, v))
}

// ActUserIDIsNil applies the IsNil predicate on the "act_user_id" field.
func ActUserIDIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldActUserID))
}

// ActUserIDNotNil applies the NotNil predicate on the "act_user_id" field.
func ActUserIDNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldActUserID))
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldRepoID, v))
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldRepoID, v))
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldRepoID, vs...))
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...int64) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldRepoID, vs...))
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v int64) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldRepoID, v))
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldRepoID, v))
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v int64) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldRepoID, v))
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v int64) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldRepoID, v))
}

// RepoIDIsNil applies the IsNil predicate on the "repo_id" field.
func RepoIDIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldRepoID))
}

// RepoIDNotNil applies the NotNil predicate on the "repo_id" field.
func RepoIDNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldRepoID))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldIsDeleted, v))
}

// RefNameEQ applies the EQ predicate on the "ref_name" field.
func RefNameEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldRefName, v))
}

// RefNameNEQ applies the NEQ predicate on the "ref_name" field.
func RefNameNEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldRefName, v))
}

// RefNameIn applies the In predicate on the "ref_name" field.
func RefNameIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldRefName, vs...))
}

// RefNameNotIn applies the NotIn predicate on the "ref_name" field.
func RefNameNotIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldRefName, vs...))
}

// RefNameGT applies the GT predicate on the "ref_name" field.
func RefNameGT(v string) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldRefName, v))
}

// RefNameGTE applies the GTE predicate on the "ref_name" field.
func RefNameGTE(v string) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldRefName, v))
}

// RefNameLT applies the LT predicate on the "ref_name" field.
func RefNameLT(v string) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldRefName, v))
}

// RefNameLTE applies the LTE predicate on the "ref_name" field.
func RefNameLTE(v string) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldRefName, v))
}

// RefNameContains applies the Contains predicate on the "ref_name" field.
func RefNameContains(v string) predicate.Action {
	return predicate.Action(sql.FieldContains(FieldRefName, v))
}

// RefNameHasPrefix applies the HasPrefix predicate on the "ref_name" field.
func RefNameHasPrefix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasPrefix(FieldRefName, v))
}

// RefNameHasSuffix applies the HasSuffix predicate on the "ref_name" field.
func RefNameHasSuffix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasSuffix(FieldRefName, v))
}

// RefNameIsNil applies the IsNil predicate on the "ref_name" field.
func RefNameIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldRefName))
}

// RefNameNotNil applies the NotNil predicate on the "ref_name" field.
func RefNameNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldRefName))
}

// RefNameEqualFold applies the EqualFold predicate on the "ref_name" field.
func RefNameEqualFold(v string) predicate.Action {
	return predicate.Action(sql.FieldEqualFold(FieldRefName, v))
}

// RefNameContainsFold applies the ContainsFold predicate on the "ref_name" field.
func RefNameContainsFold(v string) predicate.Action {
	return predicate.Action(sql.FieldContainsFold(FieldRefName, v))
}

// IsPrivateEQ applies the EQ predicate on the "is_private" field.
func IsPrivateEQ(v bool) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldIsPrivate, v))
}

// IsPrivateNEQ applies the NEQ predicate on the "is_private" field.
func IsPrivateNEQ(v bool) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldIsPrivate, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Action {
	return predicate.Action(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasSuffix(FieldContent, v))
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.Action {
	return predicate.Action(sql.FieldIsNull(FieldContent))
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.Action {
	return predicate.Action(sql.FieldNotNull(FieldContent))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Action {
	return predicate.Action(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Action {
	return predicate.Action(sql.FieldContainsFold(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Action) predicate.Action {
	return predicate.Action(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Action) predicate.Action {
	return predicate.Action(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Action) predicate.Action {
	return predicate.Action(sql.NotPredicates(p))
}