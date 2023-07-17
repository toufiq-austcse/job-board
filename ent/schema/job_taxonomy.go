package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// JobTaxonomy holds the schema definition for the JobTaxonomy entity.
type JobTaxonomy struct {
	ent.Schema
}

// Fields of the JobTaxonomy.
func (JobTaxonomy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("job_id"),
		field.Int("taxonomy_id"),
	}
}

// Edges of the JobTaxonomy.
func (JobTaxonomy) Edges() []ent.Edge {
	return nil
}
