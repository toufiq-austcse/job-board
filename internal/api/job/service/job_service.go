package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/thoas/go-funk"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	taxonomyEnam "github.com/toufiq-austcse/go-api-boilerplate/enums/taxonomy"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
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
	companyService     *service.CompanyService
}

func NewJobService(jobRepository *repository.JobRepository, taxonomyRepository *repository.TaxonomyRepository, companyService *service.CompanyService) *JobService {
	return &JobService{repository: jobRepository, taxonomyRepository: taxonomyRepository, companyService: companyService}
}

func (service JobService) Create(data req.CreateJobReqModel, company *ent.Company, ctx context.Context) (*res.CreateJobRes, error) {
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

	taxonomyTypes := []string{
		taxonomyEnam.CATEGORY,
		taxonomyEnam.REGION,
		taxonomyEnam.JOB_TYPE,
		taxonomyEnam.SKILLS,
		taxonomyEnam.SALARY_RANGE,
	}

	taxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(jobTaxonomyIds, taxonomyTypes, ctx)

	jobRes := &res.CreateJobRes{
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

func (service JobService) ListJobs(company *ent.Company, page int, limit int, status string, ctx *gin.Context) ([]*res.JobInListJobRes, *api_response.PaginationResponse, error) {
	var result = []*res.JobInListJobRes{}
	var jobList []*ent.Job
	var companies []*ent.Company
	var total int
	var err error

	if company == nil {
		jobList, total, err = service.repository.ListJobs(0, page, limit, status, ctx)

	} else {
		jobList, total, err = service.repository.ListJobs(company.ID, page, limit, status, ctx)
	}

	if err != nil {
		return result, nil, err
	}

	if len(jobList) == 0 {
		return result, nil, nil
	}

	jobIds := funk.Map(jobList, func(job *ent.Job) int {
		return job.ID
	}).([]int)

	if company == nil {
		companyIds := funk.Map(jobList, func(job *ent.Job) int {
			return job.CompanyID
		}).([]int)
		companies, err = service.companyService.ListCompanyByIds(companyIds, ctx)
		if err != nil {
			return result, nil, err
		}

	}

	allJobTaxonomies, err := service.repository.GetTaxonomies(jobIds, ctx)
	if err != nil {
		return result, nil, err
	}

	taxonomyIds := funk.Map(allJobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) int {
		return jobTaxonomy.TaxonomyID
	}).([]int)

	taxonomyTypes := []string{taxonomyEnam.JOB_TYPE}
	allTaxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(taxonomyIds, taxonomyTypes, ctx)

	if err != nil {
		return result, nil, err
	}

	for _, job := range jobList {

		jobTaxonomies := funk.Filter(allJobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) bool {
			return jobTaxonomy.JobID == job.ID
		})

		var taxonomies []res.JobTaxonomy

		for _, jobTaxonomy := range jobTaxonomies.([]*ent.JobTaxonomy) {
			taxonomy := funk.Find(allTaxonomies, func(taxonomy *ent.Taxonomy) bool {
				return taxonomy.ID == jobTaxonomy.TaxonomyID
			})
			if taxonomy != nil {
				taxonomy := taxonomy.(*ent.Taxonomy)
				taxonomies = append(taxonomies, res.JobTaxonomy{
					ID:    taxonomy.ID,
					Title: taxonomy.Title,
					Type:  taxonomy.Type,
					Slug:  taxonomy.Slug,
				})
			}

		}
		jobResponse := &res.JobInListJobRes{
			ID:         job.ID,
			Title:      job.Title,
			Slug:       job.Slug,
			Status:     job.Status,
			Taxonomies: taxonomies,
			Company:    res.JobCompany{},
			CreatedAt:  job.CreatedAt,
			UpdatedAt:  job.UpdatedAt,
		}
		if company != nil {
			jobResponse.Company = res.JobCompany{
				Name:     company.Name,
				Location: company.Location,
				LogoUrl:  company.LogoURL,
			}
		} else {
			jobCompany := funk.Find(companies, func(company *ent.Company) bool {
				return company.ID == job.CompanyID
			})
			if jobCompany != nil {
				jobCompany := jobCompany.(*ent.Company)
				jobResponse.Company = res.JobCompany{
					Name:     jobCompany.Name,
					Location: jobCompany.Location,
					LogoUrl:  jobCompany.LogoURL,
				}
			}

		}
		result = append(result, jobResponse)

	}
	if page == 0 && limit == 0 {
		return result, nil, nil
	}
	paginationData := utils.GetPaginationData(total, page, limit)
	return result, paginationData, nil
}

func (service JobService) GetTaxonomiesByJobId(jobId int, ctx context.Context) ([]*ent.Taxonomy, error) {
	jobTaxonomies, err := service.repository.GetJobTaxonomoyByJobId(jobId, ctx)
	if err != nil {
		return nil, err
	}
	taxonomyIds := funk.Map(jobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) int {
		return jobTaxonomy.ID
	}).([]int)

	taxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(taxonomyIds, []string{
		taxonomyEnam.JOB_TYPE,
		taxonomyEnam.SKILLS,
		taxonomyEnam.REGION,
		taxonomyEnam.CATEGORY,
		taxonomyEnam.SALARY_RANGE,
	}, ctx)

	if err != nil {
		return nil, err
	}
	return taxonomies, nil

}

func (service JobService) GetJobDetails(param req.JobDetailsReqParam, ctx context.Context) (*res.JobDetailsRes, error) {
	job, err := service.repository.FindJobBySlug(param.Slug, ctx)
	if err != nil {
		return nil, err
	}

	company, err := service.companyService.FindCompanyById(job.CompanyID, ctx)
	if err != nil {
		return nil, err
	}

	taxonomies, err := service.GetTaxonomiesByJobId(job.ID, ctx)
	if err != nil {
		return nil, err
	}

	jobTaxonomyRes := funk.Map(taxonomies, func(taxonomy *ent.Taxonomy) res.JobTaxonomy {
		return res.JobTaxonomy{
			ID:    taxonomy.ID,
			Title: taxonomy.Title,
			Slug:  taxonomy.Slug,
			Type:  taxonomy.Type,
		}
	}).([]res.JobTaxonomy)

	return &res.JobDetailsRes{
		ID:          job.ID,
		Title:       job.Title,
		Description: job.Description,
		ApplyTo:     job.ApplyTo,
		Slug:        job.Slug,
		Status:      job.Status,
		Company: res.JobCompanyInJobDetails{
			Name:       company.Name,
			Location:   company.Location,
			LogoUrl:    company.LogoURL,
			WebsiteUrl: company.WebsiteURL,
		},
		Taxonomies: jobTaxonomyRes,
		CreatedAt:  job.CreatedAt,
		UpdatedAt:  job.UpdatedAt,
	}, nil

}
