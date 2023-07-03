package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/internal/api/todo/models"
)

type Repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}
func (repository *Repository) Create(model *models.CreateTodoModel, context context.Context) (*ent.Todo, error) {
	todo, err := repository.client.Todo.Create().
		SetID(uuid.New()).
		SetTitle(model.Title).
		SetStatus(model.Status).
		Save(context)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (repository *Repository) GetAll(ctx context.Context) ([]*ent.Todo, error) {
	all, err := repository.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (repository *Repository) FindOne(id uuid.UUID, ctx context.Context) (*ent.Todo, error) {
	todo, err := repository.client.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
