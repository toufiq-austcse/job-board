package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("slug"),
		field.String("apply_to"),
		field.String("description"),
		field.Int("company_id"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
