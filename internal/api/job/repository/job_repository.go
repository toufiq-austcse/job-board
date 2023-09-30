package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/job"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/jobtaxonomy"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/schema"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/taxonomy"
	"time"
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
func (repository JobRepository) FindJobBySlug(slug string, ctx context.Context) (*ent.Job, error) {
	return repository.client.Job.Query().Where(job.Slug(slug)).Only(ctx)
}
func (repository JobRepository) FindJobById(id int, ctx context.Context) (*ent.Job, error) {
	return repository.client.Job.Query().Where(job.ID(id)).Only(ctx)
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

func (repository JobRepository) GetJobTaxonomyByJobId(jobId int, ctx context.Context) ([]*ent.JobTaxonomy, error) {
	return repository.client.JobTaxonomy.Query().Where(jobtaxonomy.JobID(jobId)).All(ctx)
}

func (repository JobRepository) ListJobs(companyId int, page int, limit int, status string, ctx *gin.Context) ([]*ent.Job, int, error) {
	var predicates []predicate.Job
	if companyId != 0 {
		predicates = append(predicates, job.CompanyID(companyId))
	}
	if status != "" {
		predicates = append(predicates, job.Status(status))
	}

	var jobs []*ent.Job
	query := repository.client.Job.Query().Where(predicates...).Order(ent.Desc(job.FieldCreatedAt))

	if limit == -1 {
		if allJobs, err := query.All(ctx); err != nil {
			return nil, 0, err
		} else {
			jobs = allJobs
		}
	} else {
		if allJobs, err := query.Limit(limit).Offset((page - 1) * limit).All(ctx); err != nil {
			return nil, 0, err
		} else {
			jobs = allJobs
		}
	}

	if len(jobs) == 0 {
		return []*ent.Job{}, 0, nil
	}
	count, err := repository.client.Job.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return jobs, count, nil
}

func (repository JobRepository) GetJobsByTaxonomyId(taxonomyId int, companyId int, page int, limit int, status string, ctx context.Context) ([]*ent.Job, int, error) {
	var jobs []*ent.Job
	predicates := []*sql.Predicate{
		sql.EQ(jobtaxonomy.FieldTaxonomyID, taxonomyId),
	}

	if companyId != 0 {
		predicates = append(predicates, sql.EQ(job.FieldCompanyID, companyId))
	}
	if status != "" {
		predicates = append(predicates, sql.EQ(job.FieldStatus, status))
	}
	query := repository.client.JobTaxonomy.Query().Where(func(selector *sql.Selector) {
		jobTableView := sql.Table(job.Table)
		selector.Where(sql.And(predicates...))
		selector.LeftJoin(jobTableView).On(selector.C(jobtaxonomy.FieldJobID), jobTableView.C(job.FieldID)).
			Select(jobTableView.C(job.FieldID), jobTableView.C(job.FieldTitle), jobTableView.C(job.FieldSlug),
				jobTableView.C(job.FieldStatus), jobTableView.C(job.FieldCompanyID), jobTableView.C(job.FieldApplyTo),
				jobTableView.C(job.FieldDescription), jobTableView.C(job.FieldCreatedAt), jobTableView.C(job.FieldUpdatedAt))
	})

	if page == 0 {
		page = 1
	}
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if limit == -1 {
		if err := query.Select().Scan(ctx, &jobs); err != nil {
			return nil, 0, err
		}
	} else {
		if err := query.Limit(limit).Offset((page-1)*limit).Select().Scan(ctx, &jobs); err != nil {
			return nil, 0, err
		}
	}

	return jobs, count, nil
}
func (repository JobRepository) GetJobTaxonomiesByJobId(jobId int, ctx context.Context) ([]schema.JobTaxonomyDetails, error) {
	var data []struct {
		Id         int       `json:"id"`
		TaxonomyId int       `json:"taxonomy_id"`
		ParentId   int       `json:"parent_id"`
		Title      string    `json:"title"`
		Slug       string    `json:"slug"`
		Type       string    `json:"type"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
	err := repository.client.JobTaxonomy.Query().Where(func(selector *sql.Selector) {
		taxonomyTableView := sql.Table(taxonomy.Table)
		jobTaxonomyTableView := sql.Table(jobtaxonomy.Table)
		selector.Where(sql.EQ(jobtaxonomy.FieldJobID, jobId))
		selector.LeftJoin(taxonomyTableView).On(selector.C(jobtaxonomy.FieldTaxonomyID), taxonomyTableView.C(taxonomy.FieldID)).
			Select(jobTaxonomyTableView.C(jobtaxonomy.FieldID), jobtaxonomy.FieldTaxonomyID,
				taxonomyTableView.C(taxonomy.FieldParentID), taxonomyTableView.C(taxonomy.FieldTitle),
				taxonomyTableView.C(taxonomy.FieldSlug), taxonomyTableView.C(taxonomy.FieldType),
				taxonomyTableView.C(taxonomy.FieldCreatedAt), taxonomyTableView.C(taxonomy.FieldUpdatedAt))
	}).Select().Scan(ctx, &data)
	if err != nil {
		return nil, err
	}

	var res []schema.JobTaxonomyDetails
	for _, currentData := range data {
		res = append(res, schema.JobTaxonomyDetails{
			Id: currentData.Id,
			Job: ent.Job{
				ID: jobId,
			},
			Taxonomies: ent.Taxonomy{
				ID:        currentData.TaxonomyId,
				ParentID:  "",
				Title:     currentData.Title,
				Slug:      currentData.Slug,
				Type:      currentData.Type,
				CreatedAt: currentData.CreatedAt,
				UpdatedAt: currentData.UpdatedAt,
			},
		})
	}
	return res, nil
}
