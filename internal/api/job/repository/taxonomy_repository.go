package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/taxonomy"
)

type TaxonomyRepository struct {
	client *ent.Client
}

func NewTaxonomyRepository(client *ent.Client) *TaxonomyRepository {
	return &TaxonomyRepository{
		client: client,
	}
}

func (repository TaxonomyRepository) ListTaxonomies(taxonomyType string, ctx context.Context) []*ent.Taxonomy {
	if taxonomyType == "" {
		return repository.client.Taxonomy.Query().AllX(ctx)
	}
	return repository.client.Taxonomy.Query().Where(taxonomy.Type(taxonomyType)).AllX(ctx)
}

func (repository TaxonomyRepository) GetTaxonomyByIds(ids []int, fieldTypes []string, ctx context.Context) ([]*ent.Taxonomy, error) {
	fieldTypeInterface := make([]interface{}, len(fieldTypes))
	for i, fieldType := range fieldTypes {
		fieldTypeInterface[i] = fieldType
	}
	return repository.client.Taxonomy.Query().Where(
		taxonomy.And(
			func(s *sql.Selector) {
				s.Where(sql.InInts(taxonomy.FieldID, ids...))
			},
			func(s *sql.Selector) {
				s.Where(sql.In(taxonomy.FieldType, fieldTypeInterface...))
			},
		),
	).All(ctx)
}

func (repository TaxonomyRepository) GetTaxonomyBySlug(slug string, ctx context.Context) (*ent.Taxonomy, error) {
	return repository.client.Taxonomy.Query().Where(taxonomy.Slug(slug)).Only(ctx)

}
