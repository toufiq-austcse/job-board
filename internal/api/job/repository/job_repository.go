package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/job"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/jobtaxonomy"
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

func (repository JobRepository) FindJobCountByTitle(title string, ctx context.Context) (int, error) {
	return repository.client.Job.Query().Where(job.Title(title)).Count(ctx)
}
func (repository JobRepository) FindJobCountBySlug(slug string, ctx context.Context) (int, error) {
	return repository.client.Job.Query().Where(job.Slug(slug)).Count(ctx)
}

func (repository JobRepository) GetJobCount(ctx context.Context) (int, error) {
	return repository.client.Job.Query().Count(ctx)
}

func (repository JobRepository) CreateJobTaxonomy(jobId int, taxonomyIds []int, ctx context.Context) ([]*ent.JobTaxonomy, error) {
	bulk := make([]*ent.JobTaxonomyCreate, len(taxonomyIds))

	for i, taxonomyId := range taxonomyIds {
		bulk[i] = repository.client.JobTaxonomy.Create().SetTaxonomyID(taxonomyId).SetJobID(jobId)
	}
	jobTaxonomies, err := repository.client.JobTaxonomy.CreateBulk(bulk...).Save(ctx)

	return jobTaxonomies, err
}

func (repository JobRepository) GetTaxonomies(jobIds []int, ctx context.Context) ([]*ent.JobTaxonomy, error) {
	return repository.client.JobTaxonomy.Query().Where(func(selector *sql.Selector) {
		selector.Where(sql.InInts(jobtaxonomy.FieldJobID, jobIds...))
	}).All(ctx)
}

func (repository JobRepository) ListJobs(companyId int, page int, limit int, ctx *gin.Context) ([]*ent.Job, int, error) {
	jobs, err := repository.client.Job.Query().Where(job.CompanyID(companyId)).Limit(limit).Offset((page - 1) * limit).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := repository.client.Job.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return jobs, count, nil
}
