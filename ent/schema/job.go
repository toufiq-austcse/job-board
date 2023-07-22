package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/hook"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/job"
	jobEnums "github.com/toufiq-austcse/go-api-boilerplate/enums/job"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("slug").Unique(),
		field.String("status").Nillable(),
		field.String("apply_to"),
		field.String("description"),
		field.Int("company_id"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}

func (Job) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
					fmt.Println("called hook")

					err := mutation.SetField(job.FieldStatus, jobEnums.ACTIVE)
					if err != nil {
						return nil, err
					}
					return next.Mutate(ctx, mutation)
				})
			}, ent.OpCreate),
	}
}

func (Job) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}
