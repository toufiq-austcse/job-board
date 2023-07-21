package repository

import (
	"context"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
)

type JobRepository struct {
	client *ent.Client
}

func NewJobRepository(client *ent.Client) *JobRepository {
	return &JobRepository{
		client: client,
	}
}

func (repository JobRepository) Create(title, slug, applyTo, description string, companyId int, ctx context.Context) (*ent.Job, error) {
	return repository.client.Job.Create().
		SetTitle(title).
		SetSlug(slug).
		SetApplyTo(applyTo).
		SetDescription(description).
		SetCompanyID(companyId).
		Save(ctx)

}

func (repository JobRepository) CreateJobTaxonomy(jobId int, taxonomyIds []int, ctx context.Context) ([]*ent.JobTaxonomy, error) {
	bulk := make([]*ent.JobTaxonomyCreate, len(taxonomyIds))

	for i, taxonomyId := range taxonomyIds {
		bulk[i] = repository.client.JobTaxonomy.Create().SetTaxonomyID(taxonomyId).SetJobID(jobId)
	}
	jobTaxonomies, err := repository.client.JobTaxonomy.CreateBulk(bulk...).Save(ctx)

	return jobTaxonomies, err
}
