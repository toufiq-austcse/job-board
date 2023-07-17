// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/jobtaxonomy"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
)

// JobTaxonomyDelete is the builder for deleting a JobTaxonomy entity.
type JobTaxonomyDelete struct {
	config
	hooks    []Hook
	mutation *JobTaxonomyMutation
}

// Where appends a list predicates to the JobTaxonomyDelete builder.
func (jtd *JobTaxonomyDelete) Where(ps ...predicate.JobTaxonomy) *JobTaxonomyDelete {
	jtd.mutation.Where(ps...)
	return jtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jtd *JobTaxonomyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jtd.sqlExec, jtd.mutation, jtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jtd *JobTaxonomyDelete) ExecX(ctx context.Context) int {
	n, err := jtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jtd *JobTaxonomyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(jobtaxonomy.Table, sqlgraph.NewFieldSpec(jobtaxonomy.FieldID, field.TypeInt))
	if ps := jtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jtd.mutation.done = true
	return affected, err
}

// JobTaxonomyDeleteOne is the builder for deleting a single JobTaxonomy entity.
type JobTaxonomyDeleteOne struct {
	jtd *JobTaxonomyDelete
}

// Where appends a list predicates to the JobTaxonomyDelete builder.
func (jtdo *JobTaxonomyDeleteOne) Where(ps ...predicate.JobTaxonomy) *JobTaxonomyDeleteOne {
	jtdo.jtd.mutation.Where(ps...)
	return jtdo
}

// Exec executes the deletion query.
func (jtdo *JobTaxonomyDeleteOne) Exec(ctx context.Context) error {
	n, err := jtdo.jtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jobtaxonomy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jtdo *JobTaxonomyDeleteOne) ExecX(ctx context.Context) {
	if err := jtdo.Exec(ctx); err != nil {
		panic(err)
	}
}