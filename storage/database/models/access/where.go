// Code generated by ent, DO NOT EDIT.

package access

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUserID, v))
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRepoID, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldMode, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Access {
	return predicate.Access(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Access {
	return predicate.Access(sql.FieldNotNull(FieldUserID))
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRepoID, v))
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldRepoID, v))
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldRepoID, vs...))
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...int64) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldRepoID, vs...))
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v int64) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldRepoID, v))
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v int64) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldRepoID, v))
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v int64) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldRepoID, v))
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v int64) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldRepoID, v))
}

// RepoIDIsNil applies the IsNil predicate on the "repo_id" field.
func RepoIDIsNil() predicate.Access {
	return predicate.Access(sql.FieldIsNull(FieldRepoID))
}

// RepoIDNotNil applies the NotNil predicate on the "repo_id" field.
func RepoIDNotNil() predicate.Access {
	return predicate.Access(sql.FieldNotNull(FieldRepoID))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v int) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v int) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v int) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v int) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldMode, v))
}

// ModeIsNil applies the IsNil predicate on the "mode" field.
func ModeIsNil() predicate.Access {
	return predicate.Access(sql.FieldIsNull(FieldMode))
}

// ModeNotNil applies the NotNil predicate on the "mode" field.
func ModeNotNil() predicate.Access {
	return predicate.Access(sql.FieldNotNull(FieldMode))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Access) predicate.Access {
	return predicate.Access(sql.NotPredicates(p))
}