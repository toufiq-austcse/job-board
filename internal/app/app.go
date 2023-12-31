package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toufiq-austcse/go-api-boilerplate/config"
	"github.com/toufiq-austcse/go-api-boilerplate/di"
	"github.com/toufiq-austcse/go-api-boilerplate/docs"
	authController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/controller"
	authRouter "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/router"
	companyController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/controller"
	companyRouter "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/router"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	indexRouter "github.com/toufiq-austcse/go-api-boilerplate/internal/api/index/router"
	jobController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/controller"
	jobRouter "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/router"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/handlers"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/server"
	"time"
)

func Run(configPath string) error {
	config.Init(configPath)
	apiServer := server.NewServer()
	apiServer.GinEngine.Use(gin.CustomRecovery(exception.GlobalErrorHandler))
	enableCors(apiServer.GinEngine)

	setupSwagger(apiServer)
	container, err := di.NewDiContainer()
	if err != nil {
		return err
	}
	err = container.Invoke(func(
		authController *authController.AuthController,
		jobController *jobController.JobController,
		taxonomyController *jobController.TaxonomyController,
		companyService *companyService.CompanyService,
		companyController *companyController.CompanyController,
	) {
		indexRouterGroup := apiServer.GinEngine.Group("")
		indexRouter.Setup(indexRouterGroup)

		authV1RouterGroup := apiServer.GinEngine.Group("api/v1/auth")
		authRouter.Setup(authV1RouterGroup, authController, companyService)

		jobsV1RouterGroup := apiServer.GinEngine.Group("api/v1/jobs")
		jobRouter.SetupJobRoutes(jobsV1RouterGroup, jobController, companyService)

		taxonomyV1RouterGroup := apiServer.GinEngine.Group("api/v1/taxonomies")
		jobRouter.SetupTaxonomyRoutes(taxonomyV1RouterGroup, taxonomyController)

		companyV1RouterGroup := apiServer.GinEngine.Group("api/v1/companies")
		companyRouter.SetupCompanyRoutes(companyV1RouterGroup, companyController, companyService)

	})
	if err != nil {
		return err
	}

	err = apiServer.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupSwagger(apiServer *server.Server) {
	docs.SwaggerInfo.Title = config.AppConfig.APP_NAME + " API DOC"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = config.AppConfig.APP_URL
	apiServer.GinEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func enableCors(engine *gin.Engine) {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With", "Referer", "guest", "publicKey", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	engine.Use(cors.New(corsConfig))
}
