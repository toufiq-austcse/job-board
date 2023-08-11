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
	company, _ := context.Get("company")
	entCompany := company.(*ent.Company)

	query := req.JobListQuery{}
	if err := query.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	fmt.Println("query ", query)
	jobList, pagination, err := controller.service.ListJobs(entCompany, query.Page, query.Limit, query.Status, context)
	if err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	res := api_response.BuildResponseWithPagination(http.StatusOK, "Job List", jobList, pagination)
	context.JSON(res.Code, res)

}
