// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/store/models/foreignreference"
)

// ForeignReference is the model entity for the ForeignReference schema.
type ForeignReference struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id,omitempty"`
	// LocalIndex holds the value of the "local_index" field.
	LocalIndex int64 `json:"local_index,omitempty"`
	// ForeignIndex holds the value of the "foreign_index" field.
	ForeignIndex string `json:"foreign_index,omitempty"`
	// Type holds the value of the "type" field.
	Type         string `json:"type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ForeignReference) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case foreignreference.FieldID, foreignreference.FieldRepoID, foreignreference.FieldLocalIndex:
			values[i] = new(sql.NullInt64)
		case foreignreference.FieldForeignIndex, foreignreference.FieldType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ForeignReference fields.
func (fr *ForeignReference) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case foreignreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fr.ID = int(value.Int64)
		case foreignreference.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				fr.RepoID = value.Int64
			}
		case foreignreference.FieldLocalIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field local_index", values[i])
			} else if value.Valid {
				fr.LocalIndex = value.Int64
			}
		case foreignreference.FieldForeignIndex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field foreign_index", values[i])
			} else if value.Valid {
				fr.ForeignIndex = value.String
			}
		case foreignreference.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				fr.Type = value.String
			}
		default:
			fr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ForeignReference.
// This includes values selected through modifiers, order, etc.
func (fr *ForeignReference) Value(name string) (ent.Value, error) {
	return fr.selectValues.Get(name)
}

// Update returns a builder for updating this ForeignReference.
// Note that you need to call ForeignReference.Unwrap() before calling this method if this ForeignReference
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *ForeignReference) Update() *ForeignReferenceUpdateOne {
	return NewForeignReferenceClient(fr.config).UpdateOne(fr)
}

// Unwrap unwraps the ForeignReference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *ForeignReference) Unwrap() *ForeignReference {
	_tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("models: ForeignReference is not a transactional entity")
	}
	fr.config.driver = _tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *ForeignReference) String() string {
	var builder strings.Builder
	builder.WriteString("ForeignReference(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fr.ID))
	builder.WriteString("repo_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.RepoID))
	builder.WriteString(", ")
	builder.WriteString("local_index=")
	builder.WriteString(fmt.Sprintf("%v", fr.LocalIndex))
	builder.WriteString(", ")
	builder.WriteString("foreign_index=")
	builder.WriteString(fr.ForeignIndex)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fr.Type)
	builder.WriteByte(')')
	return builder.String()
}

// ForeignReferences is a parsable slice of ForeignReference.
type ForeignReferences []*ForeignReference