// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/server/pkg/store/models/label"
)

// Label is the model entity for the Label schema.
type Label struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id,omitempty"`
	// OrgID holds the value of the "org_id" field.
	OrgID int64 `json:"org_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// NumIssues holds the value of the "num_issues" field.
	NumIssues int64 `json:"num_issues,omitempty"`
	// NumClosedIssues holds the value of the "num_closed_issues" field.
	NumClosedIssues int64 `json:"num_closed_issues,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Label) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case label.FieldID, label.FieldRepoID, label.FieldOrgID, label.FieldNumIssues, label.FieldNumClosedIssues:
			values[i] = new(sql.NullInt64)
		case label.FieldName, label.FieldDescription, label.FieldColor:
			values[i] = new(sql.NullString)
		case label.FieldCreatedAt, label.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Label fields.
func (l *Label) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case label.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case label.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				l.RepoID = value.Int64
			}
		case label.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				l.OrgID = value.Int64
			}
		case label.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case label.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = value.String
			}
		case label.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				l.Color = value.String
			}
		case label.FieldNumIssues:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_issues", values[i])
			} else if value.Valid {
				l.NumIssues = value.Int64
			}
		case label.FieldNumClosedIssues:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_closed_issues", values[i])
			} else if value.Valid {
				l.NumClosedIssues = value.Int64
			}
		case label.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case label.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Label.
// This includes values selected through modifiers, order, etc.
func (l *Label) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this Label.
// Note that you need to call Label.Unwrap() before calling this method if this Label
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Label) Update() *LabelUpdateOne {
	return NewLabelClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Label entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Label) Unwrap() *Label {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("models: Label is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Label) String() string {
	var builder strings.Builder
	builder.WriteString("Label(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("repo_id=")
	builder.WriteString(fmt.Sprintf("%v", l.RepoID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", l.OrgID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(l.Description)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(l.Color)
	builder.WriteString(", ")
	builder.WriteString("num_issues=")
	builder.WriteString(fmt.Sprintf("%v", l.NumIssues))
	builder.WriteString(", ")
	builder.WriteString("num_closed_issues=")
	builder.WriteString(fmt.Sprintf("%v", l.NumClosedIssues))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Labels is a parsable slice of Label.
type Labels []*Label