package service

import (
	"context"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/res"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
)

type JobService struct {
	repository         *repository.JobRepository
	taxonomyRepository *repository.TaxonomyRepository
}

func NewJobService(jobRepository *repository.JobRepository, taxonomyRepository *repository.TaxonomyRepository) *JobService {
	return &JobService{repository: jobRepository, taxonomyRepository: taxonomyRepository}
}

func (service JobService) Create(data req.CreateJobReqModel, company *ent.Company, ctx context.Context) (*res.JobRes, error) {
	createdJob, err := service.repository.Create(data.Title, "", data.ApplyTo, data.Description, company.ID, ctx)
	if err != nil {
		return nil, err
	}

	jobTaxonomies, err := service.repository.CreateJobTaxonomy(createdJob.ID, data.Taxonomies, ctx)
	if err != nil {
		return nil, err
	}

	jobTaxonomyIds := make([]int, len(jobTaxonomies))

	for i, jobTaxonomy := range jobTaxonomies {
		jobTaxonomyIds[i] = jobTaxonomy.TaxonomyID
	}

	taxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(jobTaxonomyIds, ctx)

	jobRes := &res.JobRes{
		ID:          createdJob.ID,
		Title:       createdJob.Title,
		ApplyTo:     createdJob.ApplyTo,
		Description: createdJob.Description,
		Taxonomies:  nil,
	}

	for _, taxonomy := range taxonomies {
		jobRes.Taxonomies = append(jobRes.Taxonomies, res.JobTaxonomy{
			ID:    taxonomy.ID,
			Title: taxonomy.Title,
			Type:  taxonomy.Type,
			Slug:  taxonomy.Slug,
		})
	}
	return jobRes, nil

}
