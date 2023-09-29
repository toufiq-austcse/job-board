package router

import (
	"github.com/gin-gonic/gin"
	companyController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/controller"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/middleware"
)

func SetupCompanyRoutes(group *gin.RouterGroup, companyController *companyController.CompanyController, companyService *service.CompanyService) {
	group.GET(":slug", companyController.GetCompany)
	group.GET(":slug/jobs", companyController.ListJobsByCompany)

	group.Use(middleware.AuthMiddleware(companyService, false))
	group.PATCH(":slug", companyController.UpdateCompany)

}
