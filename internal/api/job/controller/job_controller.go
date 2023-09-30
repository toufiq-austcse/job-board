package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

type JobController struct {
	service *service.JobService
}

func NewJobController(jobService *service.JobService) *JobController {
	return &JobController{
		service: jobService,
	}
}

// Create hosts godoc
// @Summary  Create New Job
// @Param    request  body      req.CreateJobReqModel  true  "Signup Req Body"
// @Security Authorization
// @name Authorization
// @Tags     Jobs
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/jobs [post]
// @Success  201      {object}  api_response.Response{data=res.JobDetailsRes}
func (controller *JobController) Create(context *gin.Context) {
	company, _ := context.Get("company")
	entCompany := company.(*ent.Company)

	body := req.CreateJobReqModel{}
	if err := body.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	createdJob, err := controller.service.Create(body, entCompany, context)

	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	res := api_response.BuildResponse(http.StatusCreated, "Created", createdJob)
	context.JSON(res.Code, res)
}

// Update hosts godoc
// @Summary  Update New Job
// @Param    request  body      req.UpdateJobReqModel  true  "Update Req Body"
// @Security Authorization
// @name Authorization
// @Param    id   path      int  true  "Job id"
// @Tags     Jobs
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/jobs/{id} [patch]
// @Success  201      {object}  api_response.Response{data=res.JobDetailsRes}
func (controller *JobController) Update(context *gin.Context) {
	company, _ := context.Get("company")
	entCompany := company.(*ent.Company)

	param := req.UpdateJobReqParam{}
	if err := param.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	body := req.UpdateJobReqModel{}
	if err := body.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	updatedJob, err := controller.service.Update(param, body, entCompany, context)

	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	fmt.Println("updatedJob ", updatedJob)

	res := api_response.BuildResponse(http.StatusOK, "Updated", updatedJob)
	context.JSON(res.Code, res)
}

// ListJobs hosts godoc
// @Summary List Jobs
// @Security Authorization
// @name Authorization
// @Param    request  query      req.JobListQuery  true  "List Job Query"
// @Tags     Jobs
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/jobs [get]
// @Success  201      {object}  api_response.ResponseWithPagination{data=[]res.JobInListJobRes}
func (controller *JobController) ListJobs(context *gin.Context) {

	query := req.JobListQuery{}
	if err := query.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	fmt.Println("query ", query)

	var entCompany *ent.Company

	company, isCompanyExist := context.Get("company")
	if isCompanyExist {
		entCompany = company.(*ent.Company)
	}

	jobList, pagination, err := controller.service.ListJobsHandler(entCompany, query.Page, query.Limit, query.Status, query.TaxonomySlug, context)
	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	res := api_response.BuildResponseWithPagination(http.StatusOK, "Job List", jobList, pagination)
	context.JSON(res.Code, res)

}

// GetJobBySlug hosts godoc
// @Summary Get Job Details
// @Security Authorization
// @name Authorization
// @Param    slug   path      string  true  "Job Slug"
// @Tags     Jobs
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/jobs/{slug} [get]
// @Success  201      {object}  api_response.Response{data=res.JobDetailsRes}
func (controller *JobController) GetJobBySlug(context *gin.Context) {
	param := req.JobDetailsReqParam{}

	if err := param.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	var entCompany *ent.Company

	company, isCompanyExist := context.Get("company")
	if isCompanyExist {
		entCompany = company.(*ent.Company)
	}

	jobDetails, err := controller.service.GetJobDetails(entCompany, param, context)

	if err != nil {
		var errRes api_response.Response
		if err.Error() == "ent: job not found" {
			errRes = api_response.BuildErrorResponse(http.StatusNotFound, "Not Found", err.Error(), nil)
		} else {
			errRes = api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		}
		context.JSON(errRes.Code, errRes)
		return
	}

	res := api_response.BuildResponse(http.StatusOK, "Job Details", jobDetails)
	context.JSON(res.Code, res)

}
