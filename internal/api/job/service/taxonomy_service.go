package service

import (
	"context"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/res"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
)

type TaxonomyService struct {
	taxonomyRepo *repository.TaxonomyRepository
}

func NewTaxonomyService(repository *repository.TaxonomyRepository) *TaxonomyService {
	return &TaxonomyService{taxonomyRepo: repository}
}

func (service *TaxonomyService) ListTaxonomy(taxonomyType string, ctx context.Context) []res.TaxonomyInListRes {
	var taxonomies []res.TaxonomyInListRes
	dbTaxonomies := service.taxonomyRepo.ListTaxonomies(taxonomyType, ctx)

	for _, taxonomy := range dbTaxonomies {
		taxonomies = append(taxonomies, res.TaxonomyInListRes{
			ID:    taxonomy.ID,
			Title: taxonomy.Title,
			Slug:  taxonomy.Slug,
			Type:  taxonomy.Type,
		})
	}
	return taxonomies

}
