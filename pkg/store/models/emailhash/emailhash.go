// Code generated by ent, DO NOT EDIT.

package emailhash

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the emailhash type in the database.
	Label = "email_hash"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// Table holds the table name of the emailhash in the database.
	Table = "email_hashes"
)

// Columns holds all SQL columns for emailhash fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldEmail,
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
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
)

// OrderOption defines the ordering options for the EmailHash queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}