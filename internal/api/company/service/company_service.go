package service

import (
	"context"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/model"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/repository"
)

type CompanyRepository interface {
	CreateCompany(createCompanyModel *model.CreateCompany, context context.Context) (*ent.Company, error)
	FindCompanyByEmail(email string, context context.Context) (*ent.Company, error)
	FindCompanyById(id int, ctx context.Context) (*ent.Company, error)
	ListCompanyByIds(ids []int, ctx context.Context) ([]*ent.Company, error)
}

type CompanyService struct {
	companyRepo CompanyRepository
}

func NewCompanyService(repository *repository.Repository) *CompanyService {
	return &CompanyService{companyRepo: repository}
}

func (service *CompanyService) CreateCompany(name, email, password string, ctx context.Context) (*ent.Company, error) {
	companyModel := &model.CreateCompany{
		Name:     name,
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
