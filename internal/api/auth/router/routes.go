package router

import (
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/controller"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/middleware"
)

func Setup(group *gin.RouterGroup, controller *controller.AuthController, companyService *companyService.CompanyService) {
	group.POST("signup", controller.SignUp)
	group.POST("login", controller.Login)
	group.Use(middleware.AuthMiddleware(companyService, false))
	group.GET("me", controller.Me)

}
