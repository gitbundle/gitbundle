// Code generated by ent, DO NOT EDIT.

package reporedirect

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the reporedirect type in the database.
	Label = "repo_redirect"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldLowerName holds the string denoting the lower_name field in the database.
	FieldLowerName = "lower_name"
	// FieldRedirectRepoID holds the string denoting the redirect_repo_id field in the database.
	FieldRedirectRepoID = "redirect_repo_id"
	// Table holds the table name of the reporedirect in the database.
	Table = "repo_redirects"
)

// Columns holds all SQL columns for reporedirect fields.
var Columns = []string{
	FieldID,
	FieldOwnerID,
	FieldLowerName,
	FieldRedirectRepoID,
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

// OrderOption defines the ordering options for the RepoRedirect queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByLowerName orders the results by the lower_name field.
func ByLowerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLowerName, opts...).ToFunc()
}

// ByRedirectRepoID orders the results by the redirect_repo_id field.
func ByRedirectRepoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirectRepoID, opts...).ToFunc()
}