package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	jobEnums "github.com/toufiq-austcse/go-api-boilerplate/enums/job"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/apimodels/res"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	JobService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"net/http"
)

type CompanyController struct {
	companyService *companyService.CompanyService
	jobService     *JobService.JobService
}

func NewCompanyController(service *companyService.CompanyService, jobService *JobService.JobService) *CompanyController {
	return &CompanyController{companyService: service, jobService: jobService}
}

// GetCompany hosts godoc
// @Summary  Company Details
// @Tags     Company
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/companies/{slug} [get]
// @Param    slug   path      string  true  "Company Slug"
// @Success  200      {object}  api_response.Response{data=res.CompanyDetailsRes}
func (companyController *CompanyController) GetCompany(context *gin.Context) {
	param := req.CompanyDetailsReqParam{}
	if err := param.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	company, err := companyController.companyService.GetCompanyDetailsBySlug(param.Slug, context)
	if err != nil {
		var errRes api_response.Response
		if err.Error() == "ent: company not found" {
			errRes = api_response.BuildErrorResponse(http.StatusNotFound, "Not Found", err.Error(), nil)
		} else {
			errRes = api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		}
		context.JSON(errRes.Code, errRes)
		return
	}

	fmt.Println("companyDetails ", company)

	companyDetailsRes := api_response.BuildResponse(http.StatusOK, "Company Details", res.CompanyDetailsRes{
		Name:               company.Name,
		Location:           company.Location,
		LogoURL:            company.LogoURL,
		WebsiteURL:         company.WebsiteURL,
		Email:              company.Email,
		Size:               company.Size,
		Industry:           company.Industry,
		Established:        company.Established,
		Description:        company.Description,
		CultureDescription: company.CultureDescription,
		HiringDescription:  company.HiringDescription,
		Slug:               company.Slug,
	})
	context.JSON(companyDetailsRes.Code, companyDetailsRes)
}

// UpdateCompany hosts godoc
// @Summary  Update Company Details
// @Tags     Company
// @Security Authorization
// @name Authorization
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/companies/{slug} [patch]
// @Param    slug   path      string  true  "Company Slug"
// @Param    request  body      req.UpdateCompanyReqModel  true  "Signup Req Body"
// @Success  200      {object}  api_response.Response{data=res.CompanyDetailsRes}
func (companyController *CompanyController) UpdateCompany(context *gin.Context) {
	param := req.CompanyDetailsReqParam{}
	if err := param.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	body := req.UpdateCompanyReqModel{}
	if err := body.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}
	company, _ := context.Get("company")
	entCompany := company.(*ent.Company)
	companyDetails, err := companyController.companyService.UpdateCompany(param.Slug, entCompany, body, context)
	if err != nil {
		var errRes api_response.Response
		if err.Error() == "ent: company not found" {
			errRes = api_response.BuildErrorResponse(http.StatusNotFound, "Not Found", err.Error(), nil)
		} else if err.Error() == "unauthorized" {
			errRes = api_response.BuildErrorResponse(http.StatusUnauthorized, "Unauthorized", err.Error(), nil)
		} else {
			errRes = api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		}
		context.JSON(errRes.Code, errRes)
		return
	}

	companyDetailsRes := api_response.BuildResponse(http.StatusOK, "Company Details", companyDetails)
	context.JSON(companyDetailsRes.Code, companyDetailsRes)

}

// ListJobsByCompany hosts godoc
// @Summary List Jobs By Company
// @Param    slug   path      string  true  "Company Slug"
// @Tags     Company
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/companies/{slug}/jobs [get]
// @Success  201      {object}  api_response.ResponseWithPagination{data=[]res.JobInListJobRes}
func (companyController *CompanyController) ListJobsByCompany(context *gin.Context) {
	param := req.CompanyJobsReqParam{}
	if err := param.Validate(context); err != nil {
		errRes := api_response.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	company, err := companyController.companyService.GetCompanyDetailsBySlug(param.Slug, context)
	if err != nil {
		var errRes api_response.Response
		if err.Error() == "ent: company not found" {
			errRes = api_response.BuildErrorResponse(http.StatusNotFound, "Not Found", err.Error(), nil)
		} else {
			errRes = api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", err.Error(), nil)
		}
		context.JSON(errRes.Code, errRes)
		return
	}

	jobList, pagination, listJobsErr := companyController.jobService.ListJobsHandler(company, -1, -1, jobEnums.ACTIVE, "", context)
	if listJobsErr != nil {
		errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Server Error", listJobsErr.Error(), nil)
		context.JSON(errRes.Code, errRes)
		return
	}

	ListJobsRes := api_response.BuildResponseWithPagination(http.StatusOK, "Job List", jobList, pagination)
	context.JSON(ListJobsRes.Code, ListJobsRes)
}
