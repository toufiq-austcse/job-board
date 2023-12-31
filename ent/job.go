// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/job"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Status holds the value of the "status" field.
	Status *string `json:"status,omitempty"`
	// ApplyTo holds the value of the "apply_to" field.
	ApplyTo string `json:"apply_to,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CompanyID holds the value of the "company_id" field.
	CompanyID    int `json:"company_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldID, job.FieldCompanyID:
			values[i] = new(sql.NullInt64)
		case job.FieldTitle, job.FieldSlug, job.FieldStatus, job.FieldApplyTo, job.FieldDescription:
			values[i] = new(sql.NullString)
		case job.FieldCreatedAt, job.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = int(value.Int64)
		case job.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case job.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				j.UpdatedAt = value.Time
			}
		case job.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				j.Title = value.String
			}
		case job.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				j.Slug = value.String
			}
		case job.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				j.Status = new(string)
				*j.Status = value.String
			}
		case job.FieldApplyTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field apply_to", values[i])
			} else if value.Valid {
				j.ApplyTo = value.String
			}
		case job.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				j.Description = value.String
			}
		case job.FieldCompanyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field company_id", values[i])
			} else if value.Valid {
				j.CompanyID = int(value.Int64)
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Job.
// This includes values selected through modifiers, order, etc.
func (j *Job) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return NewJobClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("created_at=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(j.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(j.Title)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(j.Slug)
	builder.WriteString(", ")
	if v := j.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("apply_to=")
	builder.WriteString(j.ApplyTo)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(j.Description)
	builder.WriteString(", ")
	builder.WriteString("company_id=")
	builder.WriteString(fmt.Sprintf("%v", j.CompanyID))
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job
