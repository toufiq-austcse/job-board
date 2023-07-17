package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Taxonomy struct {
	ent.Schema
}

// Fields of the Todo.
func (Taxonomy) Fields() []ent.Field {
	return []ent.Field{
		field.String("parent_id").Optional(),
		field.String("title").Optional(),
		field.String("slug").Optional(),
		field.String("type").Optional(),
	}
}

// Edges of the Todo.
func (Taxonomy) Edges() []ent.Edge {
	return nil
}
