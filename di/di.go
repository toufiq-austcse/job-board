package di

import (
	_ "github.com/lib/pq" // <------------ here
	authController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/controller"
	authService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/service"
	companyController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/controller"
	companyRepository "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/repository"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	jobController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/controller"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
	jobServices "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/orm/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/providers/influxdb"
	"go.uber.org/dig"
)

func NewDiContainer() (*dig.Container, error) {
	c := dig.New()
	providers := []interface {
	}{
		ent.New,
		influxdb.OpenInfluxDbConnection,
		companyRepository.NewRepository,
		repository.NewTaxonomyRepository,
		repository.NewJobRepository,
		authService.NewAuthService,
		companyService.NewCompanyService,
		jobServices.NewJobService,
		jobServices.NewTaxonomyService,
		authController.NewAuthController,
		jobController.NewJobController,
		jobController.NewTaxonomyController,
		companyController.NewCompanyController,
	}
	for _, provider := range providers {
		if err := c.Provide(provider); err != nil {
			return nil, err
		}
	}
	return c, nil
}
