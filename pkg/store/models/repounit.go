// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/server/pkg/store/models/repounit"
)

// RepoUnit is the model entity for the RepoUnit schema.
type RepoUnit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// Config holds the value of the "config" field.
	Config string `json:"config,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RepoUnit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case repounit.FieldID, repounit.FieldRepoID, repounit.FieldType:
			values[i] = new(sql.NullInt64)
		case repounit.FieldConfig:
			values[i] = new(sql.NullString)
		case repounit.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RepoUnit fields.
func (ru *RepoUnit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repounit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ru.ID = int(value.Int64)
		case repounit.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				ru.RepoID = value.Int64
			}
		case repounit.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ru.Type = int(value.Int64)
			}
		case repounit.FieldConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value.Valid {
				ru.Config = value.String
			}
		case repounit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ru.CreatedAt = value.Time
			}
		default:
			ru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RepoUnit.
// This includes values selected through modifiers, order, etc.
func (ru *RepoUnit) Value(name string) (ent.Value, error) {
	return ru.selectValues.Get(name)
}

// Update returns a builder for updating this RepoUnit.
// Note that you need to call RepoUnit.Unwrap() before calling this method if this RepoUnit
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *RepoUnit) Update() *RepoUnitUpdateOne {
	return NewRepoUnitClient(ru.config).UpdateOne(ru)
}

// Unwrap unwraps the RepoUnit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *RepoUnit) Unwrap() *RepoUnit {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("models: RepoUnit is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *RepoUnit) String() string {
	var builder strings.Builder
	builder.WriteString("RepoUnit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ru.ID))
	builder.WriteString("repo_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.RepoID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ru.Type))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(ru.Config)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ru.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RepoUnits is a parsable slice of RepoUnit.
type RepoUnits []*RepoUnit