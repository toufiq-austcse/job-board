// Code generated by ent, DO NOT EDIT.

package jobtaxonomy

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the jobtaxonomy type in the database.
	Label = "job_taxonomy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldJobID holds the string denoting the job_id field in the database.
	FieldJobID = "job_id"
	// FieldTaxonomyID holds the string denoting the taxonomy_id field in the database.
	FieldTaxonomyID = "taxonomy_id"
	// Table holds the table name of the jobtaxonomy in the database.
	Table = "job_taxonomies"
)

// Columns holds all SQL columns for jobtaxonomy fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldJobID,
	FieldTaxonomyID,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the JobTaxonomy queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByJobID orders the results by the job_id field.
func ByJobID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJobID, opts...).ToFunc()
}

// ByTaxonomyID orders the results by the taxonomy_id field.
func ByTaxonomyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxonomyID, opts...).ToFunc()
}
