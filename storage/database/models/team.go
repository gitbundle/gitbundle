// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/team"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OrgID holds the value of the "org_id" field.
	OrgID int64 `json:"org_id,omitempty"`
	// LowerName holds the value of the "lower_name" field.
	LowerName string `json:"lower_name,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Authorize holds the value of the "authorize" field.
	Authorize int `json:"authorize,omitempty"`
	// NumRepos holds the value of the "num_repos" field.
	NumRepos int `json:"num_repos,omitempty"`
	// NumMembers holds the value of the "num_members" field.
	NumMembers int `json:"num_members,omitempty"`
	// IncludeAllRepos holds the value of the "include_all_repos" field.
	IncludeAllRepos bool `json:"include_all_repos,omitempty"`
	// CanCreateOrgRepo holds the value of the "can_create_org_repo" field.
	CanCreateOrgRepo bool `json:"can_create_org_repo,omitempty"`
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldIncludeAllRepos, team.FieldCanCreateOrgRepo:
			values[i] = new(sql.NullBool)
		case team.FieldID, team.FieldOrgID, team.FieldAuthorize, team.FieldNumRepos, team.FieldNumMembers:
			values[i] = new(sql.NullInt64)
		case team.FieldLowerName, team.FieldName, team.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				t.OrgID = value.Int64
			}
		case team.FieldLowerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lower_name", values[i])
			} else if value.Valid {
				t.LowerName = value.String
			}
		case team.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case team.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case team.FieldAuthorize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field authorize", values[i])
			} else if value.Valid {
				t.Authorize = int(value.Int64)
			}
		case team.FieldNumRepos:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_repos", values[i])
			} else if value.Valid {
				t.NumRepos = int(value.Int64)
			}
		case team.FieldNumMembers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_members", values[i])
			} else if value.Valid {
				t.NumMembers = int(value.Int64)
			}
		case team.FieldIncludeAllRepos:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field include_all_repos", values[i])
			} else if value.Valid {
				t.IncludeAllRepos = value.Bool
			}
		case team.FieldCanCreateOrgRepo:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_create_org_repo", values[i])
			} else if value.Valid {
				t.CanCreateOrgRepo = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Team.
// This includes values selected through modifiers, order, etc.
func (t *Team) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return NewTeamClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("models: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", t.OrgID))
	builder.WriteString(", ")
	builder.WriteString("lower_name=")
	builder.WriteString(t.LowerName)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("authorize=")
	builder.WriteString(fmt.Sprintf("%v", t.Authorize))
	builder.WriteString(", ")
	builder.WriteString("num_repos=")
	builder.WriteString(fmt.Sprintf("%v", t.NumRepos))
	builder.WriteString(", ")
	builder.WriteString("num_members=")
	builder.WriteString(fmt.Sprintf("%v", t.NumMembers))
	builder.WriteString(", ")
	builder.WriteString("include_all_repos=")
	builder.WriteString(fmt.Sprintf("%v", t.IncludeAllRepos))
	builder.WriteString(", ")
	builder.WriteString("can_create_org_repo=")
	builder.WriteString(fmt.Sprintf("%v", t.CanCreateOrgRepo))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team