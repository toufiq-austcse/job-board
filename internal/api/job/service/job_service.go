package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/schema"
	taxonomyEnam "github.com/toufiq-austcse/go-api-boilerplate/enums/taxonomy"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/res"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"github.com/toufiq-austcse/go-api-boilerplate/utils"
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
	currentAvailableJobsCount, err := service.repository.GetJobCount(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("currentAvailableJobsCount ", currentAvailableJobsCount)

	jobSlug := utils.GetSlug(data.Title, currentAvailableJobsCount)

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

func (service JobService) ListJobsByQuery(taxonomySlug string, company *ent.Company, page int, limit int, status string, ctx *gin.Context) ([]*ent.Job, int, error) {
	companyId := 0

	if company != nil {
		companyId = company.ID
	}

	if taxonomySlug != "" {
		taxonomy, err := service.taxonomyRepository.GetTaxonomyBySlug(taxonomySlug, ctx)
		if err != nil {
			return []*ent.Job{}, 0, err
		}
		return service.repository.GetJobsByTaxonomyId(taxonomy.ID, companyId, page, limit, status, ctx)

	} else {
		return service.repository.ListJobs(companyId, page, limit, status, ctx)
	}

}

func (service JobService) ListJobsHandler(company *ent.Company, page int, limit int, status string, taxonomySlug string, ctx *gin.Context) ([]*res.JobInListJobRes, *api_response.PaginationResponse, error) {

	var companies []*ent.Company

	jobList, total, err := service.ListJobsByQuery(taxonomySlug, company, page, limit, status, ctx)

	if err != nil {
		return []*res.JobInListJobRes{}, nil, err
	}

	if len(jobList) == 0 {
		return []*res.JobInListJobRes{}, nil, nil
	}

	jobIds := funk.Map(jobList, func(job *ent.Job) int {
		return job.ID
	}).([]int)

	if company == nil {
		companyIds := funk.Map(jobList, func(job *ent.Job) int {
			return job.CompanyID
		}).([]int)
		if companyList, err := service.companyService.ListCompanyByIds(companyIds, ctx); err != nil {
			return []*res.JobInListJobRes{}, nil, err
		} else {
			companies = companyList
		}

	}

	allJobTaxonomies, err := service.repository.GetTaxonomies(jobIds, ctx)
	if err != nil {
		return []*res.JobInListJobRes{}, nil, err
	}

	taxonomyIds := funk.Map(allJobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) int {
		return jobTaxonomy.TaxonomyID
	}).([]int)

	taxonomyTypes := []string{taxonomyEnam.JOB_TYPE, taxonomyEnam.REGION}
	allTaxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(taxonomyIds, taxonomyTypes, ctx)

	if err != nil {
		return []*res.JobInListJobRes{}, nil, err
	}

	var result []*res.JobInListJobRes

	for _, job := range jobList {

		jobTaxonomies := funk.Filter(allJobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) bool {
			return jobTaxonomy.JobID == job.ID
		})

		taxonomies := []res.JobTaxonomy{}

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
	jobTaxonomies, err := service.repository.GetJobTaxonomyByJobId(jobId, ctx)
	if err != nil {
		return nil, err
	}
	taxonomyIds := funk.Map(jobTaxonomies, func(jobTaxonomy *ent.JobTaxonomy) int {
		return jobTaxonomy.TaxonomyID
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

func (service JobService) GetIsMine(jobCompany *ent.Company, loginCompany *ent.Company) bool {
	if loginCompany == nil {
		return false
	}
	return jobCompany.ID == loginCompany.ID
}

func (service JobService) GetJobDetails(company *ent.Company, param req.JobDetailsReqParam, ctx context.Context) (*res.JobDetailsRes, error) {
	job, err := service.repository.FindJobBySlug(param.Slug, ctx)
	if err != nil {
		return nil, err
	}

	jobCompany, err := service.companyService.FindCompanyById(job.CompanyID, ctx)
	if err != nil {
		return nil, err
	}

	taxonomies, err := service.GetTaxonomiesByJobId(job.ID, ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("taxonomies ", taxonomies)

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
		IsMine:      service.GetIsMine(jobCompany, company),
		Company: res.JobCompanyInJobDetails{
			Name:       jobCompany.Name,
			Location:   jobCompany.Location,
			Slug:       jobCompany.Slug,
			LogoUrl:    jobCompany.LogoURL,
			WebsiteUrl: jobCompany.WebsiteURL,
		},
		Taxonomies: jobTaxonomyRes,
		CreatedAt:  job.CreatedAt,
		UpdatedAt:  job.UpdatedAt,
	}, nil

}

func (service JobService) Update(param req.UpdateJobReqParam, body req.UpdateJobReqModel, company *ent.Company, ctx *gin.Context) ([]schema.JobTaxonomyDetails, error) {

	fmt.Println("body ", body.ApplyTo)

	job, err := service.repository.FindJobById(param.Id, ctx)
	if err != nil {
		return nil, err
	}
	if job.CompanyID != company.ID {
		return nil, errors.New("unauthorized")
	}

	updateJobQuery := job.Update()

	if body.ApplyTo != "" {
		updateJobQuery.SetApplyTo(body.ApplyTo)
	}
	if body.Title != "" {
		updateJobQuery.SetTitle(body.Title)
		//currentAvailableJobsCount, err := service.repository.GetJobCount(ctx)
		//if err != nil {
		//	return nil, err
		//}
		//
		//updateJobQuery.SetSlug(utils.GetSlug(body.Title, currentAvailableJobsCount))
	}
	if body.Description != "" {
		updateJobQuery.SetDescription(body.Description)
	}
	if body.Status != "" {
		updateJobQuery.SetStatus(body.Status)
	}
	_, updateJobErr := updateJobQuery.Save(ctx)
	if updateJobErr != nil {
		return nil, err
	}
	if len(body.Taxonomies) > 0 {
		jobTaxonomies, err := service.repository.GetJobTaxonomiesByJobId(job.ID, ctx)
		if err != nil {
			return nil, err
		}

		err = service.UpdateCategory(param.Id, body.Taxonomies, jobTaxonomies, ctx)
		if err != nil {
			return nil, err
		}

		err = service.UpdateSalaryRange(param.Id, body.Taxonomies, jobTaxonomies, ctx)
		if err != nil {
			return nil, err
		}

		err = service.UpdateJobType(param.Id, body.Taxonomies, jobTaxonomies, ctx)
		if err != nil {
			return nil, err
		}

		err = service.UpdateJobRegion(param.Id, body.Taxonomies, jobTaxonomies, ctx)
		if err != nil {
			return nil, err
		}

		err = service.updateJobSkills(param.Id, body.Taxonomies, jobTaxonomies, ctx)
		if err != nil {
			return nil, err
		}

	}
	return nil, err

}

func (service JobService) UpdateCategory(jobId int, taxonomies []req.UpdateJobTaxonomyModel, currentJobTaxonomies []schema.JobTaxonomyDetails, ctx context.Context) error {
	categoryToUpdate := funk.Find(taxonomies, func(taxonomy req.UpdateJobTaxonomyModel) bool {
		return taxonomy.Type == taxonomyEnam.CATEGORY
	})

	if categoryToUpdate == nil {
		return nil
	}
	categoryToUpdateTyped := categoryToUpdate.(req.UpdateJobTaxonomyModel)

	currentCategoryJobTaxonomy := funk.Find(currentJobTaxonomies, func(jobTaxonomy schema.JobTaxonomyDetails) bool {
		return jobTaxonomy.Taxonomy.Type == taxonomyEnam.CATEGORY
	})

	if currentCategoryJobTaxonomy == nil {
		_, err := service.repository.CreateJobTaxonomy(jobId, []int{categoryToUpdateTyped.Id}, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	currentCategoryJobTaxonomyTyped := currentCategoryJobTaxonomy.(schema.JobTaxonomyDetails)
	_, err := service.repository.UpdateJobTaxonomyById(currentCategoryJobTaxonomyTyped.Id, categoryToUpdateTyped.Id, ctx)

	if err != nil {
		return err
	}
	return nil

}

func (service JobService) UpdateSalaryRange(jobId int, taxonomies []req.UpdateJobTaxonomyModel, currentJobTaxonomies []schema.JobTaxonomyDetails, ctx context.Context) error {
	salaryRangeToUpdate := funk.Find(taxonomies, func(taxonomy req.UpdateJobTaxonomyModel) bool {
		return taxonomy.Type == taxonomyEnam.SALARY_RANGE
	})

	if salaryRangeToUpdate == nil {
		return nil
	}
	categoryToUpdateTyped := salaryRangeToUpdate.(req.UpdateJobTaxonomyModel)

	currentSalaryRangeJobTaxonomy := funk.Find(currentJobTaxonomies, func(jobTaxonomy schema.JobTaxonomyDetails) bool {
		return jobTaxonomy.Taxonomy.Type == taxonomyEnam.SALARY_RANGE
	})

	if currentSalaryRangeJobTaxonomy == nil {
		_, err := service.repository.CreateJobTaxonomy(jobId, []int{categoryToUpdateTyped.Id}, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	currentSalaryRangeJobTaxonomyTyped := currentSalaryRangeJobTaxonomy.(schema.JobTaxonomyDetails)
	_, err := service.repository.UpdateJobTaxonomyById(currentSalaryRangeJobTaxonomyTyped.Id, categoryToUpdateTyped.Id, ctx)

	if err != nil {
		return err
	}
	return nil
}
func (service JobService) UpdateJobType(jobId int, taxonomies []req.UpdateJobTaxonomyModel, currentJobTaxonomies []schema.JobTaxonomyDetails, ctx context.Context) error {
	jobTypeToUpdate := funk.Find(taxonomies, func(taxonomy req.UpdateJobTaxonomyModel) bool {
		return taxonomy.Type == taxonomyEnam.JOB_TYPE
	})

	if jobTypeToUpdate == nil {
		return nil
	}
	jobTypeToUpdateTyped := jobTypeToUpdate.(req.UpdateJobTaxonomyModel)

	currentJobTypeJobTaxonomy := funk.Find(currentJobTaxonomies, func(jobTaxonomy schema.JobTaxonomyDetails) bool {
		return jobTaxonomy.Taxonomy.Type == taxonomyEnam.JOB_TYPE
	})

	if currentJobTypeJobTaxonomy == nil {
		_, err := service.repository.CreateJobTaxonomy(jobId, []int{jobTypeToUpdateTyped.Id}, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	currentJobTypeJobTaxonomyTyped := currentJobTypeJobTaxonomy.(schema.JobTaxonomyDetails)
	_, err := service.repository.UpdateJobTaxonomyById(currentJobTypeJobTaxonomyTyped.Id, jobTypeToUpdateTyped.Id, ctx)

	if err != nil {
		return err
	}
	return nil
}

func (service JobService) UpdateJobRegion(jobId int, taxonomies []req.UpdateJobTaxonomyModel, currentJobTaxonomies []schema.JobTaxonomyDetails, ctx context.Context) error {
	jobRegionsToUpdate := funk.Filter(taxonomies, func(taxonomy req.UpdateJobTaxonomyModel) bool {
		return taxonomy.Type == taxonomyEnam.REGION
	}).([]req.UpdateJobTaxonomyModel)

	if len(jobRegionsToUpdate) == 0 {
		return nil
	}

	jobRegionIdsToUpdate := funk.Map(jobRegionsToUpdate, func(taxonomy req.UpdateJobTaxonomyModel) int {
		return taxonomy.Id
	}).([]int)
	regionTaxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(jobRegionIdsToUpdate, []string{taxonomyEnam.REGION}, ctx)
	if err != nil {
		return err
	}
	if len(regionTaxonomies) != len(jobRegionsToUpdate) {
		return errors.New("invalid Region")
	}

	//jobRegionsToUpdateTyped := jobRegionsToUpdate.([]req.UpdateJobTaxonomyModel)

	currentJobRegionsJobTaxonomy := funk.Filter(currentJobTaxonomies, func(jobTaxonomy schema.JobTaxonomyDetails) bool {
		return jobTaxonomy.Taxonomy.Type == taxonomyEnam.REGION
	})

	if currentJobRegionsJobTaxonomy == nil {
		var taxonomyIds []int

		for _, region := range jobRegionsToUpdate {
			taxonomyIds = append(taxonomyIds, region.Id)
		}
		_, err := service.repository.CreateJobTaxonomy(jobId, taxonomyIds, ctx)
		if err != nil {
			return err
		}
		return nil
	}
	currentJobRegionsJobTaxonomyTyped := currentJobRegionsJobTaxonomy.([]schema.JobTaxonomyDetails)

	jobRegionIdsToDelete := service.GetJobRegionsToDelete(currentJobRegionsJobTaxonomyTyped, jobRegionsToUpdate)
	_, deleteErr := service.repository.DeleteJobTaxonomies(jobId, jobRegionIdsToDelete, ctx)
	if deleteErr != nil {
		return deleteErr
	}

	jobRegionIdsToCreate := service.GetJobRegionsToCreate(currentJobRegionsJobTaxonomyTyped, jobRegionsToUpdate)
	_, createErr := service.repository.CreateJobTaxonomy(jobId, jobRegionIdsToCreate, ctx)
	if createErr != nil {
		return createErr
	}

	return nil
}

func (service JobService) GetJobRegionsToDelete(currentJobRegions []schema.JobTaxonomyDetails, jobRegionsToUpdate []req.UpdateJobTaxonomyModel) []int {
	var jobRegionsIdsToDelete []int
	for _, currentRegion := range currentJobRegions {
		jobRegionToDelete := funk.Find(jobRegionsToUpdate, func(jobRegion req.UpdateJobTaxonomyModel) bool {
			return jobRegion.Id == currentRegion.Taxonomy.ID
		})
		if jobRegionToDelete == nil {
			jobRegionsIdsToDelete = append(jobRegionsIdsToDelete, currentRegion.Taxonomy.ID)
		}

	}

	return jobRegionsIdsToDelete

}

func (service JobService) GetJobRegionsToCreate(currentJobRegions []schema.JobTaxonomyDetails, jobRegionsToUpdate []req.UpdateJobTaxonomyModel) []int {
	var jobRegionsIdsToCreate []int

	for _, jobRegionToUpdate := range jobRegionsToUpdate {
		existingJobRegion := funk.Find(currentJobRegions, func(currentJobRegion schema.JobTaxonomyDetails) bool {
			return currentJobRegion.Taxonomy.ID == jobRegionToUpdate.Id
		})
		if existingJobRegion == nil {
			jobRegionsIdsToCreate = append(jobRegionsIdsToCreate, jobRegionToUpdate.Id)
		}
	}

	return jobRegionsIdsToCreate

}

func (service JobService) GetJobSkillsToDelete(currentJobSkills []schema.JobTaxonomyDetails, jobSkillsToUpdate []req.UpdateJobTaxonomyModel) []int {
	var jobSkillsIdsToDelete []int

	for _, currentJobSkill := range currentJobSkills {
		jobSkillInUpdateReq := funk.Find(jobSkillsToUpdate, func(jobSkill req.UpdateJobTaxonomyModel) bool {
			return jobSkill.Id == currentJobSkill.Taxonomy.ID
		})
		if jobSkillInUpdateReq == nil {
			jobSkillsIdsToDelete = append(jobSkillsIdsToDelete, currentJobSkill.Taxonomy.ID)
		}

	}
	return jobSkillsIdsToDelete

}

func (service JobService) GetJobSkillsToCreate(currentJobSkills []schema.JobTaxonomyDetails, jobSkillsToUpdate []req.UpdateJobTaxonomyModel) []int {
	var jobRegionsIdsToCreate []int

	for _, jobSkillToUpdate := range jobSkillsToUpdate {
		existingJobSkill := funk.Find(currentJobSkills, func(currentJobRegion schema.JobTaxonomyDetails) bool {
			return currentJobRegion.Taxonomy.ID == jobSkillToUpdate.Id
		})
		if existingJobSkill == nil {
			jobRegionsIdsToCreate = append(jobRegionsIdsToCreate, jobSkillToUpdate.Id)
		}
	}

	return jobRegionsIdsToCreate

}

func (service JobService) updateJobSkills(jobId int, taxonomiesToUpdate []req.UpdateJobTaxonomyModel, taxonomies []schema.JobTaxonomyDetails, ctx *gin.Context) error {

	skillsToUpdate := funk.Filter(taxonomies, func(taxonomy req.UpdateJobTaxonomyModel) bool {
		return taxonomy.Type == taxonomyEnam.REGION
	}).([]req.UpdateJobTaxonomyModel)

	if len(skillsToUpdate) == 0 {
		return nil
	}

	skillIdsToUpdate := funk.Map(skillsToUpdate, func(taxonomy req.UpdateJobTaxonomyModel) int {
		return taxonomy.Id
	}).([]int)

	skillTaxonomies, err := service.taxonomyRepository.GetTaxonomyByIds(skillIdsToUpdate, []string{taxonomyEnam.SKILLS}, ctx)
	if err != nil {
		return err
	}
	if len(skillTaxonomies) != len(taxonomiesToUpdate) {
		return errors.New("invalid skills")
	}

	//jobRegionsToUpdateTyped := jobRegionsToUpdate.([]req.UpdateJobTaxonomyModel)

	currentJobSKillsTaxonomy := funk.Filter(taxonomies, func(jobTaxonomy schema.JobTaxonomyDetails) bool {
		return jobTaxonomy.Taxonomy.Type == taxonomyEnam.SKILLS
	})

	if currentJobSKillsTaxonomy == nil {
		var taxonomyIds []int

		for _, skill := range skillsToUpdate {
			taxonomyIds = append(taxonomyIds, skill.Id)
		}
		_, err := service.repository.CreateJobTaxonomy(jobId, taxonomyIds, ctx)
		if err != nil {
			return err
		}
		return nil
	}
	currentJobSkillsTaxonomyTyped := currentJobSKillsTaxonomy.([]schema.JobTaxonomyDetails)

	jobSkillIdsToDelete := service.GetJobSkillsToDelete(currentJobSkillsTaxonomyTyped, skillsToUpdate)
	_, deleteErr := service.repository.DeleteJobTaxonomies(jobId, jobSkillIdsToDelete, ctx)
	if deleteErr != nil {
		return deleteErr
	}

	jobSKillIdsToCreate := service.GetJobSkillsToCreate(currentJobSkillsTaxonomyTyped, skillsToUpdate)
	_, createErr := service.repository.CreateJobTaxonomy(jobId, jobSKillIdsToCreate, ctx)
	if createErr != nil {
		return createErr
	}

	return nil

}

func (service JobService) OnUpdateOne(ctx context.Context, mutation ent.Mutation) {
	updatedFields := mutation.Fields()
	fmt.Println(updatedFields)
}
