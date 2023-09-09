package service

import (
	"context"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/model"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/repository"
	"github.com/toufiq-austcse/go-api-boilerplate/utils"
)

type CompanyService struct {
	companyRepo *repository.CompanyRepository
}

func NewCompanyService(repository *repository.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepo: repository}
}

func (service *CompanyService) CreateCompany(name, email, password string, ctx context.Context) (*ent.Company, error) {
	companyCount, err := service.companyRepo.GetCompanyCount(ctx)
	if err != nil {
		return nil, err
	}
	companySlug := utils.GetSlug(name, companyCount)
	companyModel := &model.CreateCompany{
		Name:     name,
		Slug:     companySlug,
		Password: password,
		Email:    email,
	}
	return service.companyRepo.CreateCompany(companyModel, ctx)
}

func (service *CompanyService) FindCompanyByEmail(email string, ctx context.Context) (*ent.Company, error) {
	return service.companyRepo.FindCompanyByEmail(email, ctx)
}

func (service *CompanyService) FindCompanyById(id int, ctx context.Context) (*ent.Company, error) {
	return service.companyRepo.FindCompanyById(id, ctx)
}

func (service *CompanyService) ListCompanyByIds(ids []int, ctx context.Context) ([]*ent.Company, error) {
	return service.companyRepo.ListCompanyByIds(ids, ctx)

}
