// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gitbundle/gitbundle/storage/database/models/pkgversion"
)

// PkgVersion is the model entity for the PkgVersion schema.
type PkgVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID int64 `json:"package_id,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID int64 `json:"creator_id,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// LowerVersion holds the value of the "lower_version" field.
	LowerVersion string `json:"lower_version,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// IsInternal holds the value of the "is_internal" field.
	IsInternal bool `json:"is_internal,omitempty"`
	// DownloadCount holds the value of the "download_count" field.
	DownloadCount int64 `json:"download_count,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PkgVersion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pkgversion.FieldIsInternal:
			values[i] = new(sql.NullBool)
		case pkgversion.FieldID, pkgversion.FieldPackageID, pkgversion.FieldCreatorID, pkgversion.FieldDownloadCount:
			values[i] = new(sql.NullInt64)
		case pkgversion.FieldVersion, pkgversion.FieldLowerVersion:
			values[i] = new(sql.NullString)
		case pkgversion.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PkgVersion fields.
func (pv *PkgVersion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pkgversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = int(value.Int64)
		case pkgversion.FieldPackageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				pv.PackageID = value.Int64
			}
		case pkgversion.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				pv.CreatorID = value.Int64
			}
		case pkgversion.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pv.Version = value.String
			}
		case pkgversion.FieldLowerVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lower_version", values[i])
			} else if value.Valid {
				pv.LowerVersion = value.String
			}
		case pkgversion.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pv.CreatedAt = value.Time
			}
		case pkgversion.FieldIsInternal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_internal", values[i])
			} else if value.Valid {
				pv.IsInternal = value.Bool
			}
		case pkgversion.FieldDownloadCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field download_count", values[i])
			} else if value.Valid {
				pv.DownloadCount = value.Int64
			}
		default:
			pv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PkgVersion.
// This includes values selected through modifiers, order, etc.
func (pv *PkgVersion) Value(name string) (ent.Value, error) {
	return pv.selectValues.Get(name)
}

// Update returns a builder for updating this PkgVersion.
// Note that you need to call PkgVersion.Unwrap() before calling this method if this PkgVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PkgVersion) Update() *PkgVersionUpdateOne {
	return NewPkgVersionClient(pv.config).UpdateOne(pv)
}

// Unwrap unwraps the PkgVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PkgVersion) Unwrap() *PkgVersion {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("models: PkgVersion is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PkgVersion) String() string {
	var builder strings.Builder
	builder.WriteString("PkgVersion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("package_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.PackageID))
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.CreatorID))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pv.Version)
	builder.WriteString(", ")
	builder.WriteString("lower_version=")
	builder.WriteString(pv.LowerVersion)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_internal=")
	builder.WriteString(fmt.Sprintf("%v", pv.IsInternal))
	builder.WriteString(", ")
	builder.WriteString("download_count=")
	builder.WriteString(fmt.Sprintf("%v", pv.DownloadCount))
	builder.WriteByte(')')
	return builder.String()
}

// PkgVersions is a parsable slice of PkgVersion.
type PkgVersions []*PkgVersion