// Code generated by ent, DO NOT EDIT.

package pkgproperty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLTE(FieldID, id))
}

// RefType applies equality check predicate on the "ref_type" field. It's identical to RefTypeEQ.
func RefType(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldRefType, v))
}

// RefID applies equality check predicate on the "ref_id" field. It's identical to RefIDEQ.
func RefID(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldRefID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldName, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldValue, v))
}

// RefTypeEQ applies the EQ predicate on the "ref_type" field.
func RefTypeEQ(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldRefType, v))
}

// RefTypeNEQ applies the NEQ predicate on the "ref_type" field.
func RefTypeNEQ(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNEQ(FieldRefType, v))
}

// RefTypeIn applies the In predicate on the "ref_type" field.
func RefTypeIn(vs ...int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldIn(FieldRefType, vs...))
}

// RefTypeNotIn applies the NotIn predicate on the "ref_type" field.
func RefTypeNotIn(vs ...int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNotIn(FieldRefType, vs...))
}

// RefTypeGT applies the GT predicate on the "ref_type" field.
func RefTypeGT(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGT(FieldRefType, v))
}

// RefTypeGTE applies the GTE predicate on the "ref_type" field.
func RefTypeGTE(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGTE(FieldRefType, v))
}

// RefTypeLT applies the LT predicate on the "ref_type" field.
func RefTypeLT(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLT(FieldRefType, v))
}

// RefTypeLTE applies the LTE predicate on the "ref_type" field.
func RefTypeLTE(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLTE(FieldRefType, v))
}

// RefIDEQ applies the EQ predicate on the "ref_id" field.
func RefIDEQ(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldRefID, v))
}

// RefIDNEQ applies the NEQ predicate on the "ref_id" field.
func RefIDNEQ(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNEQ(FieldRefID, v))
}

// RefIDIn applies the In predicate on the "ref_id" field.
func RefIDIn(vs ...int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldIn(FieldRefID, vs...))
}

// RefIDNotIn applies the NotIn predicate on the "ref_id" field.
func RefIDNotIn(vs ...int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNotIn(FieldRefID, vs...))
}

// RefIDGT applies the GT predicate on the "ref_id" field.
func RefIDGT(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGT(FieldRefID, v))
}

// RefIDGTE applies the GTE predicate on the "ref_id" field.
func RefIDGTE(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGTE(FieldRefID, v))
}

// RefIDLT applies the LT predicate on the "ref_id" field.
func RefIDLT(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLT(FieldRefID, v))
}

// RefIDLTE applies the LTE predicate on the "ref_id" field.
func RefIDLTE(v int64) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLTE(FieldRefID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldContainsFold(FieldName, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.PkgProperty {
	return predicate.PkgProperty(sql.FieldContainsFold(FieldValue, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PkgProperty) predicate.PkgProperty {
	return predicate.PkgProperty(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PkgProperty) predicate.PkgProperty {
	return predicate.PkgProperty(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PkgProperty) predicate.PkgProperty {
	return predicate.PkgProperty(sql.NotPredicates(p))
}