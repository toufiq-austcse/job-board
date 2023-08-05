package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/res"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"github.com/toufiq-austcse/go-api-boilerplate/utils"
	"strconv"
)

type JobService struct {
	repository         *repository.JobRepository
	taxonomyRepository *repository.TaxonomyRepository
}

func NewJobService(jobRepository *repository.JobRepository, taxonomyRepository *repository.TaxonomyRepository) *JobService {
	return &JobService{repository: jobRepository, taxonomyRepository: taxonomyRepository}
}

func (service JobService) Create(data req.CreateJobReqModel, company *ent.Company, ctx context.Context) (*res.JobDetailsRes, error) {
	jobSlug := slug.MakeLang(data.Title, "en")
	currentAvailableJobsCount, err := service.repository.GetJobCount(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("currentAvailableJobsCount ", currentAvailableJobsCount)
	if currentAvailableJobsCount > 0 {
		jobSlug = jobSlug + "-" + strconv.Itoa(currentAvailableJobsCount)
	}

	createdJob, err := service.repository.Create(data.Title, jobSlug, data.ApplyTo, data.Description, company.ID, ctx)
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

	jobRes := &res.JobDetailsRes{
		ID:          createdJob.ID,
		Title:       createdJob.Title,
		Slug:        createdJob.Slug,
		Status:      createdJob.Status,
		ApplyTo:     createdJob.ApplyTo,
		Description: createdJob.Description,
		Taxonomies:  nil,
		CreatedAt:   createdJob.CreatedAt,
		UpdatedAt:   createdJob.UpdatedAt,
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

func (service JobService) ListJobs(company *ent.Company, page int, limit int, ctx *gin.Context) ([]*res.JobInListJobRes, *api_response.PaginationResponse, error) {
	var result []*res.JobInListJobRes
	jobList, total, err := service.repository.ListJobs(company.ID, page, limit, ctx)
	if err != nil {
		return result, nil, err
	}
	for _, job := range jobList {
		result = append(result, &res.JobInListJobRes{
			ID:        job.ID,
			Title:     job.Title,
			Slug:      job.Slug,
			Status:    job.Status,
			CreatedAt: job.CreatedAt,
			UpdatedAt: job.UpdatedAt,
		})
	}
	if page == 0 && limit == 0 {
		return result, nil, nil
	}
	paginationData := utils.GetPaginationData(total, page, limit)
	return result, paginationData, nil
}
