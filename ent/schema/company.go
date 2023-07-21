package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"errors"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/company"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("location").Optional(),
		field.String("logo_url").Optional(),
		field.String("website_url").Optional(),
		field.String("email").Unique(),
		field.String("password"),
		field.String("size").Optional(),
		field.String("industry").Optional(),
		field.String("established").Optional(),
		field.String("description").Optional(),
		field.String("culture_description").Optional(),
		field.String("hiring_description").Optional(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
func (Company) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
					oldPass, ok := mutation.Field(company.FieldPassword)
					if !ok {
						return nil, errors.New("password field is required")
					}
					hashedPass, err := hashAndSaltPassword([]byte(oldPass.(string)))
					if err != nil {
						return nil, err
					}
					setFieldErr := mutation.SetField(company.FieldPassword, hashedPass)
					if setFieldErr != nil {
						return nil, setFieldErr
					}
					return next.Mutate(ctx, mutation)
				})
			}, ent.OpCreate),
	}
}

func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}

func hashAndSaltPassword(password []byte) (hashPassword string, err error) {
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}
