// Code generated by ent, DO NOT EDIT.

package orguser

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the orguser type in the database.
	Label = "org_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"
	// Table holds the table name of the orguser in the database.
	Table = "org_users"
)

// Columns holds all SQL columns for orguser fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldOrgID,
	FieldIsPublic,
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

// OrderOption defines the ordering options for the OrgUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByIsPublic orders the results by the is_public field.
func ByIsPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPublic, opts...).ToFunc()
}