package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	ent2 "github.com/toufiq-austcse/go-api-boilerplate/ent"
)

// JobTaxonomy holds the schema definition for the JobTaxonomy entity.
type JobTaxonomy struct {
	ent.Schema
}
type JobTaxonomyDetails struct {
	Id         int           `json:"id"`
	Job        ent2.Job      `json:"job"`
	Taxonomies ent2.Taxonomy `json:"taxonomies"`
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
	return []ent.Edge{}
}

func (JobTaxonomy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}
