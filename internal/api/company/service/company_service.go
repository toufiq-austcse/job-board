package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/apimodels/req"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/apimodels/res"
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

func (service *CompanyService) GetCompanyDetailsBySlug(slug string, ctx *gin.Context) (res.CompanyDetailsRes, error) {
	company, err := service.companyRepo.GetCompanyBySlug(slug, ctx)
	if err != nil {
		return res.CompanyDetailsRes{}, err
	}
	return res.CompanyDetailsRes{
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
	}, nil
}

func (service *CompanyService) UpdateCompany(slug string, authCompany *ent.Company, body req.UpdateCompanyReqModel, ctx *gin.Context) (res.CompanyDetailsRes, error) {
	company, err := service.companyRepo.GetCompanyBySlug(slug, ctx)
	if err != nil {
		return res.CompanyDetailsRes{}, err
	}
	if company.Email != authCompany.Email {
		return res.CompanyDetailsRes{}, errors.New("unauthorized")
	}
	updateCompanyQuery := company.Update()

	if body.Name != "" {
		updateCompanyQuery.SetName(body.Name)
	}
	if body.Location != "" {
		updateCompanyQuery.SetLocation(body.Location)
	}
	if body.LogoURL != "" {
		updateCompanyQuery.SetLogoURL(body.LogoURL)
	}
	if body.WebsiteURL != "" {
		updateCompanyQuery.SetWebsiteURL(body.WebsiteURL)
	}
	if body.Size != "" {
		updateCompanyQuery.SetSize(body.Size)
	}
	if body.Industry != "" {
		updateCompanyQuery.SetIndustry(body.Industry)
	}
	if body.Established != "" {
		updateCompanyQuery.SetEstablished(body.Established)
	}
	if body.Description != "" {
		updateCompanyQuery.SetDescription(body.Description)
	}
	if body.CultureDescription != "" {
		updateCompanyQuery.SetCultureDescription(body.CultureDescription)
	}
	if body.HiringDescription != "" {
		updateCompanyQuery.SetHiringDescription(body.HiringDescription)
	}

	updatedCompany, err := updateCompanyQuery.Save(ctx)
	if err != nil {
		return res.CompanyDetailsRes{}, err
	}

	return res.CompanyDetailsRes{
		Name:               updatedCompany.Name,
		Location:           updatedCompany.Location,
		LogoURL:            updatedCompany.LogoURL,
		WebsiteURL:         updatedCompany.WebsiteURL,
		Email:              updatedCompany.Email,
		Size:               updatedCompany.Size,
		Industry:           updatedCompany.Industry,
		Established:        updatedCompany.Established,
		Description:        updatedCompany.Description,
		CultureDescription: updatedCompany.CultureDescription,
		HiringDescription:  updatedCompany.HiringDescription,
		Slug:               updatedCompany.Slug,
	}, nil
}
