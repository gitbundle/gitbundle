// Code generated by ent, DO NOT EDIT.

package hooktask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the hooktask type in the database.
	Label = "hook_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// FieldHookID holds the string denoting the hook_id field in the database.
	FieldHookID = "hook_id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldPayloadContent holds the string denoting the payload_content field in the database.
	FieldPayloadContent = "payload_content"
	// FieldEventType holds the string denoting the event_type field in the database.
	FieldEventType = "event_type"
	// FieldIsDelivered holds the string denoting the is_delivered field in the database.
	FieldIsDelivered = "is_delivered"
	// FieldDelivered holds the string denoting the delivered field in the database.
	FieldDelivered = "delivered"
	// FieldIsSucceed holds the string denoting the is_succeed field in the database.
	FieldIsSucceed = "is_succeed"
	// FieldRequestContent holds the string denoting the request_content field in the database.
	FieldRequestContent = "request_content"
	// FieldResponseContent holds the string denoting the response_content field in the database.
	FieldResponseContent = "response_content"
	// Table holds the table name of the hooktask in the database.
	Table = "hook_tasks"
)

// Columns holds all SQL columns for hooktask fields.
var Columns = []string{
	FieldID,
	FieldRepoID,
	FieldHookID,
	FieldUUID,
	FieldPayloadContent,
	FieldEventType,
	FieldIsDelivered,
	FieldDelivered,
	FieldIsSucceed,
	FieldRequestContent,
	FieldResponseContent,
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
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// EventTypeValidator is a validator for the "event_type" field. It is called by the builders before save.
	EventTypeValidator func(string) error
)

// OrderOption defines the ordering options for the HookTask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRepoID orders the results by the repo_id field.
func ByRepoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepoID, opts...).ToFunc()
}

// ByHookID orders the results by the hook_id field.
func ByHookID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHookID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByPayloadContent orders the results by the payload_content field.
func ByPayloadContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayloadContent, opts...).ToFunc()
}

// ByEventType orders the results by the event_type field.
func ByEventType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventType, opts...).ToFunc()
}

// ByIsDelivered orders the results by the is_delivered field.
func ByIsDelivered(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDelivered, opts...).ToFunc()
}

// ByDelivered orders the results by the delivered field.
func ByDelivered(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelivered, opts...).ToFunc()
}

// ByIsSucceed orders the results by the is_succeed field.
func ByIsSucceed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSucceed, opts...).ToFunc()
}

// ByRequestContent orders the results by the request_content field.
func ByRequestContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestContent, opts...).ToFunc()
}

// ByResponseContent orders the results by the response_content field.
func ByResponseContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResponseContent, opts...).ToFunc()
}