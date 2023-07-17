// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/taxonomy"
)

// TaxonomyCreate is the builder for creating a Taxonomy entity.
type TaxonomyCreate struct {
	config
	mutation *TaxonomyMutation
	hooks    []Hook
}

// SetParentID sets the "parent_id" field.
func (tc *TaxonomyCreate) SetParentID(s string) *TaxonomyCreate {
	tc.mutation.SetParentID(s)
	return tc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (tc *TaxonomyCreate) SetNillableParentID(s *string) *TaxonomyCreate {
	if s != nil {
		tc.SetParentID(*s)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *TaxonomyCreate) SetTitle(s string) *TaxonomyCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tc *TaxonomyCreate) SetNillableTitle(s *string) *TaxonomyCreate {
	if s != nil {
		tc.SetTitle(*s)
	}
	return tc
}

// SetSlug sets the "slug" field.
func (tc *TaxonomyCreate) SetSlug(s string) *TaxonomyCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tc *TaxonomyCreate) SetNillableSlug(s *string) *TaxonomyCreate {
	if s != nil {
		tc.SetSlug(*s)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TaxonomyCreate) SetType(s string) *TaxonomyCreate {
	tc.mutation.SetType(s)
	return tc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tc *TaxonomyCreate) SetNillableType(s *string) *TaxonomyCreate {
	if s != nil {
		tc.SetType(*s)
	}
	return tc
}

// Mutation returns the TaxonomyMutation object of the builder.
func (tc *TaxonomyCreate) Mutation() *TaxonomyMutation {
	return tc.mutation
}

// Save creates the Taxonomy in the database.
func (tc *TaxonomyCreate) Save(ctx context.Context) (*Taxonomy, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaxonomyCreate) SaveX(ctx context.Context) *Taxonomy {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaxonomyCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaxonomyCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaxonomyCreate) check() error {
	return nil
}

func (tc *TaxonomyCreate) sqlSave(ctx context.Context) (*Taxonomy, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaxonomyCreate) createSpec() (*Taxonomy, *sqlgraph.CreateSpec) {
	var (
		_node = &Taxonomy{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(taxonomy.Table, sqlgraph.NewFieldSpec(taxonomy.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.ParentID(); ok {
		_spec.SetField(taxonomy.FieldParentID, field.TypeString, value)
		_node.ParentID = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(taxonomy.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(taxonomy.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(taxonomy.FieldType, field.TypeString, value)
		_node.Type = value
	}
	return _node, _spec
}

// TaxonomyCreateBulk is the builder for creating many Taxonomy entities in bulk.
type TaxonomyCreateBulk struct {
	config
	builders []*TaxonomyCreate
}

// Save creates the Taxonomy entities in the database.
func (tcb *TaxonomyCreateBulk) Save(ctx context.Context) ([]*Taxonomy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Taxonomy, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaxonomyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaxonomyCreateBulk) SaveX(ctx context.Context) []*Taxonomy {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaxonomyCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaxonomyCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
