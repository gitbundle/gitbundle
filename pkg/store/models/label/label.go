// Code generated by ent, DO NOT EDIT.

package label

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the label type in the database.
	Label = "label"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldNumIssues holds the string denoting the num_issues field in the database.
	FieldNumIssues = "num_issues"
	// FieldNumClosedIssues holds the string denoting the num_closed_issues field in the database.
	FieldNumClosedIssues = "num_closed_issues"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the label in the database.
	Table = "labels"
)

// Columns holds all SQL columns for label fields.
var Columns = []string{
	FieldID,
	FieldRepoID,
	FieldOrgID,
	FieldName,
	FieldDescription,
	FieldColor,
	FieldNumIssues,
	FieldNumClosedIssues,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// ColorValidator is a validator for the "color" field. It is called by the builders before save.
	ColorValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Label queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRepoID orders the results by the repo_id field.
func ByRepoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepoID, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByNumIssues orders the results by the num_issues field.
func ByNumIssues(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumIssues, opts...).ToFunc()
}

// ByNumClosedIssues orders the results by the num_closed_issues field.
func ByNumClosedIssues(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumClosedIssues, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}