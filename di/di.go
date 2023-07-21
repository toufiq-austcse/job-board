package di

import (
	_ "github.com/lib/pq" // <------------ here
	authController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/controller"
	authService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/service"
	companyRepository "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/repository"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	jobController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/controller"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/repository"
	taxononomyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/job/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/orm/ent"
	"go.uber.org/dig"
)

func NewDiContainer() (*dig.Container, error) {
	c := dig.New()
	providers := []interface {
	}{
		ent.New,
		companyRepository.NewRepository,
		repository.NewTaxonomyRepository,
		authService.NewAuthService,
		companyService.NewCompanyService,
		taxononomyService.NewTaxonomyService,
		authController.NewAuthController,
		jobController.NewJobController,
		jobController.NewTaxonomyController,
	}
	for _, provider := range providers {
		if err := c.Provide(provider); err != nil {
			return nil, err
		}
	}
	return c, nil
}
