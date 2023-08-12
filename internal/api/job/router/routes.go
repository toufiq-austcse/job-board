package router

import (
	"github.com/gin-gonic/gin"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	jobController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/controller"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/middleware"
)

func SetupJobRoutes(group *gin.RouterGroup, controller *jobController.JobController, companyService *companyService.CompanyService) {
	group.Use(middleware.AuthMiddleware(companyService, true))
	group.GET("", controller.ListJobs)
	group.GET(":slug", controller.GetJobBySlug)
	group.Use(middleware.AuthMiddleware(companyService, false))
	group.POST("", controller.Create)

}

func SetupTaxonomyRoutes(group *gin.RouterGroup, controller *jobController.TaxonomyController) {
	group.GET("", controller.ListTaxonomies)

}
