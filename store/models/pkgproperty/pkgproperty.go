// Code generated by ent, DO NOT EDIT.

package pkgproperty

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the pkgproperty type in the database.
	Label = "pkg_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRefType holds the string denoting the ref_type field in the database.
	FieldRefType = "ref_type"
	// FieldRefID holds the string denoting the ref_id field in the database.
	FieldRefID = "ref_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the pkgproperty in the database.
	Table = "pkg_properties"
)

// Columns holds all SQL columns for pkgproperty fields.
var Columns = []string{
	FieldID,
	FieldRefType,
	FieldRefID,
	FieldName,
	FieldValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the PkgProperty queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRefType orders the results by the ref_type field.
func ByRefType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefType, opts...).ToFunc()
}

// ByRefID orders the results by the ref_id field.
func ByRefID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}