// Code generated by ent, DO NOT EDIT.

package emailaddress

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/store/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldUserID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldEmail, v))
}

// LowerEmail applies equality check predicate on the "lower_email" field. It's identical to LowerEmailEQ.
func LowerEmail(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldLowerEmail, v))
}

// IsActivated applies equality check predicate on the "is_activated" field. It's identical to IsActivatedEQ.
func IsActivated(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldIsActivated, v))
}

// IsPrimary applies equality check predicate on the "is_primary" field. It's identical to IsPrimaryEQ.
func IsPrimary(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldIsPrimary, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLTE(FieldUserID, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldContainsFold(FieldEmail, v))
}

// LowerEmailEQ applies the EQ predicate on the "lower_email" field.
func LowerEmailEQ(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldLowerEmail, v))
}

// LowerEmailNEQ applies the NEQ predicate on the "lower_email" field.
func LowerEmailNEQ(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldLowerEmail, v))
}

// LowerEmailIn applies the In predicate on the "lower_email" field.
func LowerEmailIn(vs ...string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldIn(FieldLowerEmail, vs...))
}

// LowerEmailNotIn applies the NotIn predicate on the "lower_email" field.
func LowerEmailNotIn(vs ...string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNotIn(FieldLowerEmail, vs...))
}

// LowerEmailGT applies the GT predicate on the "lower_email" field.
func LowerEmailGT(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGT(FieldLowerEmail, v))
}

// LowerEmailGTE applies the GTE predicate on the "lower_email" field.
func LowerEmailGTE(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldGTE(FieldLowerEmail, v))
}

// LowerEmailLT applies the LT predicate on the "lower_email" field.
func LowerEmailLT(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLT(FieldLowerEmail, v))
}

// LowerEmailLTE applies the LTE predicate on the "lower_email" field.
func LowerEmailLTE(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldLTE(FieldLowerEmail, v))
}

// LowerEmailContains applies the Contains predicate on the "lower_email" field.
func LowerEmailContains(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldContains(FieldLowerEmail, v))
}

// LowerEmailHasPrefix applies the HasPrefix predicate on the "lower_email" field.
func LowerEmailHasPrefix(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldHasPrefix(FieldLowerEmail, v))
}

// LowerEmailHasSuffix applies the HasSuffix predicate on the "lower_email" field.
func LowerEmailHasSuffix(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldHasSuffix(FieldLowerEmail, v))
}

// LowerEmailEqualFold applies the EqualFold predicate on the "lower_email" field.
func LowerEmailEqualFold(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEqualFold(FieldLowerEmail, v))
}

// LowerEmailContainsFold applies the ContainsFold predicate on the "lower_email" field.
func LowerEmailContainsFold(v string) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldContainsFold(FieldLowerEmail, v))
}

// IsActivatedEQ applies the EQ predicate on the "is_activated" field.
func IsActivatedEQ(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldIsActivated, v))
}

// IsActivatedNEQ applies the NEQ predicate on the "is_activated" field.
func IsActivatedNEQ(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldIsActivated, v))
}

// IsActivatedIsNil applies the IsNil predicate on the "is_activated" field.
func IsActivatedIsNil() predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldIsNull(FieldIsActivated))
}

// IsActivatedNotNil applies the NotNil predicate on the "is_activated" field.
func IsActivatedNotNil() predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNotNull(FieldIsActivated))
}

// IsPrimaryEQ applies the EQ predicate on the "is_primary" field.
func IsPrimaryEQ(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldEQ(FieldIsPrimary, v))
}

// IsPrimaryNEQ applies the NEQ predicate on the "is_primary" field.
func IsPrimaryNEQ(v bool) predicate.EmailAddress {
	return predicate.EmailAddress(sql.FieldNEQ(FieldIsPrimary, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailAddress) predicate.EmailAddress {
	return predicate.EmailAddress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailAddress) predicate.EmailAddress {
	return predicate.EmailAddress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmailAddress) predicate.EmailAddress {
	return predicate.EmailAddress(sql.NotPredicates(p))
}