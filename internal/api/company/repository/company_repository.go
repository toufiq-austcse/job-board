package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/company"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/company/model"
)

type Repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (repository *Repository) CreateCompany(createCompanyModel *model.CreateCompany, context context.Context) (*ent.Company, error) {
	newCompany, err := repository.client.Company.Create().
		SetName(createCompanyModel.Name).
		SetEmail(createCompanyModel.Email).
		SetPassword(createCompanyModel.Password).
		Save(context)
	if err != nil {
		return nil, err
	}
	return newCompany, nil
}

func (repository *Repository) FindCompanyByEmail(email string, context context.Context) (*ent.Company, error) {
	aCompany, err := repository.client.Company.Query().Where(company.Email(email)).Only(context)
	if err != nil {
		return nil, err
	}
	return aCompany, nil
}

func (repository *Repository) FindCompanyById(id int, ctx context.Context) (*ent.Company, error) {
	aCompany, err := repository.client.Company.Query().Where(company.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return aCompany, nil
}

func (repository *Repository) ListCompanyByIds(ids []int, ctx context.Context) ([]*ent.Company, error) {
	return repository.client.Company.Query().Where(func(selector *sql.Selector) {
		selector.Where(sql.InInts(company.FieldID, ids...))
	}).All(ctx)
}
