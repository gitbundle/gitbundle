// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/appstate"
)

// AppState is the model entity for the AppState schema.
type AppState struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Revision holds the value of the "revision" field.
	Revision int64 `json:"revision,omitempty"`
	// Content holds the value of the "content" field.
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appstate.FieldID, appstate.FieldRevision:
			values[i] = new(sql.NullInt64)
		case appstate.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppState fields.
func (as *AppState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appstate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = int(value.Int64)
		case appstate.FieldRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				as.Revision = value.Int64
			}
		case appstate.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				as.Content = value.String
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppState.
// This includes values selected through modifiers, order, etc.
func (as *AppState) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// Update returns a builder for updating this AppState.
// Note that you need to call AppState.Unwrap() before calling this method if this AppState
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AppState) Update() *AppStateUpdateOne {
	return NewAppStateClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AppState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AppState) Unwrap() *AppState {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("models: AppState is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AppState) String() string {
	var builder strings.Builder
	builder.WriteString("AppState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("revision=")
	builder.WriteString(fmt.Sprintf("%v", as.Revision))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(as.Content)
	builder.WriteByte(')')
	return builder.String()
}

// AppStates is a parsable slice of AppState.
type AppStates []*AppState