// Code generated by ent, DO NOT EDIT.

package teamuser

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the teamuser type in the database.
	Label = "team_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldTeamID holds the string denoting the team_id field in the database.
	FieldTeamID = "team_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the teamuser in the database.
	Table = "team_users"
)

// Columns holds all SQL columns for teamuser fields.
var Columns = []string{
	FieldID,
	FieldOrgID,
	FieldTeamID,
	FieldUserID,
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

// OrderOption defines the ordering options for the TeamUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByTeamID orders the results by the team_id field.
func ByTeamID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}