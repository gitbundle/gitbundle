// Code generated by ent, DO NOT EDIT.

package userredirect

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/server/pkg/store/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLTE(FieldID, id))
}

// LowerName applies equality check predicate on the "lower_name" field. It's identical to LowerNameEQ.
func LowerName(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldLowerName, v))
}

// RedirectUserID applies equality check predicate on the "redirect_user_id" field. It's identical to RedirectUserIDEQ.
func RedirectUserID(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldRedirectUserID, v))
}

// LowerNameEQ applies the EQ predicate on the "lower_name" field.
func LowerNameEQ(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldLowerName, v))
}

// LowerNameNEQ applies the NEQ predicate on the "lower_name" field.
func LowerNameNEQ(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNEQ(FieldLowerName, v))
}

// LowerNameIn applies the In predicate on the "lower_name" field.
func LowerNameIn(vs ...int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldIn(FieldLowerName, vs...))
}

// LowerNameNotIn applies the NotIn predicate on the "lower_name" field.
func LowerNameNotIn(vs ...int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNotIn(FieldLowerName, vs...))
}

// LowerNameGT applies the GT predicate on the "lower_name" field.
func LowerNameGT(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGT(FieldLowerName, v))
}

// LowerNameGTE applies the GTE predicate on the "lower_name" field.
func LowerNameGTE(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGTE(FieldLowerName, v))
}

// LowerNameLT applies the LT predicate on the "lower_name" field.
func LowerNameLT(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLT(FieldLowerName, v))
}

// LowerNameLTE applies the LTE predicate on the "lower_name" field.
func LowerNameLTE(v int64) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLTE(FieldLowerName, v))
}

// LowerNameIsNil applies the IsNil predicate on the "lower_name" field.
func LowerNameIsNil() predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldIsNull(FieldLowerName))
}

// LowerNameNotNil applies the NotNil predicate on the "lower_name" field.
func LowerNameNotNil() predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNotNull(FieldLowerName))
}

// RedirectUserIDEQ applies the EQ predicate on the "redirect_user_id" field.
func RedirectUserIDEQ(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEQ(FieldRedirectUserID, v))
}

// RedirectUserIDNEQ applies the NEQ predicate on the "redirect_user_id" field.
func RedirectUserIDNEQ(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNEQ(FieldRedirectUserID, v))
}

// RedirectUserIDIn applies the In predicate on the "redirect_user_id" field.
func RedirectUserIDIn(vs ...string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldIn(FieldRedirectUserID, vs...))
}

// RedirectUserIDNotIn applies the NotIn predicate on the "redirect_user_id" field.
func RedirectUserIDNotIn(vs ...string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNotIn(FieldRedirectUserID, vs...))
}

// RedirectUserIDGT applies the GT predicate on the "redirect_user_id" field.
func RedirectUserIDGT(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGT(FieldRedirectUserID, v))
}

// RedirectUserIDGTE applies the GTE predicate on the "redirect_user_id" field.
func RedirectUserIDGTE(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldGTE(FieldRedirectUserID, v))
}

// RedirectUserIDLT applies the LT predicate on the "redirect_user_id" field.
func RedirectUserIDLT(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLT(FieldRedirectUserID, v))
}

// RedirectUserIDLTE applies the LTE predicate on the "redirect_user_id" field.
func RedirectUserIDLTE(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldLTE(FieldRedirectUserID, v))
}

// RedirectUserIDContains applies the Contains predicate on the "redirect_user_id" field.
func RedirectUserIDContains(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldContains(FieldRedirectUserID, v))
}

// RedirectUserIDHasPrefix applies the HasPrefix predicate on the "redirect_user_id" field.
func RedirectUserIDHasPrefix(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldHasPrefix(FieldRedirectUserID, v))
}

// RedirectUserIDHasSuffix applies the HasSuffix predicate on the "redirect_user_id" field.
func RedirectUserIDHasSuffix(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldHasSuffix(FieldRedirectUserID, v))
}

// RedirectUserIDIsNil applies the IsNil predicate on the "redirect_user_id" field.
func RedirectUserIDIsNil() predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldIsNull(FieldRedirectUserID))
}

// RedirectUserIDNotNil applies the NotNil predicate on the "redirect_user_id" field.
func RedirectUserIDNotNil() predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldNotNull(FieldRedirectUserID))
}

// RedirectUserIDEqualFold applies the EqualFold predicate on the "redirect_user_id" field.
func RedirectUserIDEqualFold(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldEqualFold(FieldRedirectUserID, v))
}

// RedirectUserIDContainsFold applies the ContainsFold predicate on the "redirect_user_id" field.
func RedirectUserIDContainsFold(v string) predicate.UserRedirect {
	return predicate.UserRedirect(sql.FieldContainsFold(FieldRedirectUserID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserRedirect) predicate.UserRedirect {
	return predicate.UserRedirect(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserRedirect) predicate.UserRedirect {
	return predicate.UserRedirect(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserRedirect) predicate.UserRedirect {
	return predicate.UserRedirect(sql.NotPredicates(p))
}