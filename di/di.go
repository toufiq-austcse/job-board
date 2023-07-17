package di

import (
	_ "github.com/lib/pq" // <------------ here
	authController "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/controller"
	authService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/auth/service"
	companyRepository "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/repository"
	companyService "github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/service"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/orm/ent"
	"go.uber.org/dig"
)

func NewDiContainer() (*dig.Container, error) {
	c := dig.New()
	providers := []interface {
	}{
		ent.New,
		companyRepository.NewRepository,
		companyService.NewCompanyService,
		authService.NewAuthService,
		authController.NewAuthController,
	}
	for _, provider := range providers {
		if err := c.Provide(provider); err != nil {
			return nil, err
		}
	}
	return c, nil
}
