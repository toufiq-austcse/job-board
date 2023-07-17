// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/taxonomy"
)

// TaxonomyUpdate is the builder for updating Taxonomy entities.
type TaxonomyUpdate struct {
	config
	hooks    []Hook
	mutation *TaxonomyMutation
}

// Where appends a list predicates to the TaxonomyUpdate builder.
func (tu *TaxonomyUpdate) Where(ps ...predicate.Taxonomy) *TaxonomyUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetParentID sets the "parent_id" field.
func (tu *TaxonomyUpdate) SetParentID(s string) *TaxonomyUpdate {
	tu.mutation.SetParentID(s)
	return tu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (tu *TaxonomyUpdate) SetNillableParentID(s *string) *TaxonomyUpdate {
	if s != nil {
		tu.SetParentID(*s)
	}
	return tu
}

// ClearParentID clears the value of the "parent_id" field.
func (tu *TaxonomyUpdate) ClearParentID() *TaxonomyUpdate {
	tu.mutation.ClearParentID()
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaxonomyUpdate) SetTitle(s string) *TaxonomyUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TaxonomyUpdate) SetNillableTitle(s *string) *TaxonomyUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// ClearTitle clears the value of the "title" field.
func (tu *TaxonomyUpdate) ClearTitle() *TaxonomyUpdate {
	tu.mutation.ClearTitle()
	return tu
}

// SetSlug sets the "slug" field.
func (tu *TaxonomyUpdate) SetSlug(s string) *TaxonomyUpdate {
	tu.mutation.SetSlug(s)
	return tu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tu *TaxonomyUpdate) SetNillableSlug(s *string) *TaxonomyUpdate {
	if s != nil {
		tu.SetSlug(*s)
	}
	return tu
}

// ClearSlug clears the value of the "slug" field.
func (tu *TaxonomyUpdate) ClearSlug() *TaxonomyUpdate {
	tu.mutation.ClearSlug()
	return tu
}

// SetType sets the "type" field.
func (tu *TaxonomyUpdate) SetType(s string) *TaxonomyUpdate {
	tu.mutation.SetType(s)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TaxonomyUpdate) SetNillableType(s *string) *TaxonomyUpdate {
	if s != nil {
		tu.SetType(*s)
	}
	return tu
}

// ClearType clears the value of the "type" field.
func (tu *TaxonomyUpdate) ClearType() *TaxonomyUpdate {
	tu.mutation.ClearType()
	return tu
}

// Mutation returns the TaxonomyMutation object of the builder.
func (tu *TaxonomyUpdate) Mutation() *TaxonomyMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaxonomyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaxonomyUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaxonomyUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaxonomyUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaxonomyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(taxonomy.Table, taxonomy.Columns, sqlgraph.NewFieldSpec(taxonomy.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.ParentID(); ok {
		_spec.SetField(taxonomy.FieldParentID, field.TypeString, value)
	}
	if tu.mutation.ParentIDCleared() {
		_spec.ClearField(taxonomy.FieldParentID, field.TypeString)
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(taxonomy.FieldTitle, field.TypeString, value)
	}
	if tu.mutation.TitleCleared() {
		_spec.ClearField(taxonomy.FieldTitle, field.TypeString)
	}
	if value, ok := tu.mutation.Slug(); ok {
		_spec.SetField(taxonomy.FieldSlug, field.TypeString, value)
	}
	if tu.mutation.SlugCleared() {
		_spec.ClearField(taxonomy.FieldSlug, field.TypeString)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(taxonomy.FieldType, field.TypeString, value)
	}
	if tu.mutation.TypeCleared() {
		_spec.ClearField(taxonomy.FieldType, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taxonomy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaxonomyUpdateOne is the builder for updating a single Taxonomy entity.
type TaxonomyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaxonomyMutation
}

// SetParentID sets the "parent_id" field.
func (tuo *TaxonomyUpdateOne) SetParentID(s string) *TaxonomyUpdateOne {
	tuo.mutation.SetParentID(s)
	return tuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (tuo *TaxonomyUpdateOne) SetNillableParentID(s *string) *TaxonomyUpdateOne {
	if s != nil {
		tuo.SetParentID(*s)
	}
	return tuo
}

// ClearParentID clears the value of the "parent_id" field.
func (tuo *TaxonomyUpdateOne) ClearParentID() *TaxonomyUpdateOne {
	tuo.mutation.ClearParentID()
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TaxonomyUpdateOne) SetTitle(s string) *TaxonomyUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TaxonomyUpdateOne) SetNillableTitle(s *string) *TaxonomyUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// ClearTitle clears the value of the "title" field.
func (tuo *TaxonomyUpdateOne) ClearTitle() *TaxonomyUpdateOne {
	tuo.mutation.ClearTitle()
	return tuo
}

// SetSlug sets the "slug" field.
func (tuo *TaxonomyUpdateOne) SetSlug(s string) *TaxonomyUpdateOne {
	tuo.mutation.SetSlug(s)
	return tuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tuo *TaxonomyUpdateOne) SetNillableSlug(s *string) *TaxonomyUpdateOne {
	if s != nil {
		tuo.SetSlug(*s)
	}
	return tuo
}

// ClearSlug clears the value of the "slug" field.
func (tuo *TaxonomyUpdateOne) ClearSlug() *TaxonomyUpdateOne {
	tuo.mutation.ClearSlug()
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaxonomyUpdateOne) SetType(s string) *TaxonomyUpdateOne {
	tuo.mutation.SetType(s)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TaxonomyUpdateOne) SetNillableType(s *string) *TaxonomyUpdateOne {
	if s != nil {
		tuo.SetType(*s)
	}
	return tuo
}

// ClearType clears the value of the "type" field.
func (tuo *TaxonomyUpdateOne) ClearType() *TaxonomyUpdateOne {
	tuo.mutation.ClearType()
	return tuo
}

// Mutation returns the TaxonomyMutation object of the builder.
func (tuo *TaxonomyUpdateOne) Mutation() *TaxonomyMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaxonomyUpdate builder.
func (tuo *TaxonomyUpdateOne) Where(ps ...predicate.Taxonomy) *TaxonomyUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaxonomyUpdateOne) Select(field string, fields ...string) *TaxonomyUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Taxonomy entity.
func (tuo *TaxonomyUpdateOne) Save(ctx context.Context) (*Taxonomy, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaxonomyUpdateOne) SaveX(ctx context.Context) *Taxonomy {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaxonomyUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaxonomyUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaxonomyUpdateOne) sqlSave(ctx context.Context) (_node *Taxonomy, err error) {
	_spec := sqlgraph.NewUpdateSpec(taxonomy.Table, taxonomy.Columns, sqlgraph.NewFieldSpec(taxonomy.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Taxonomy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taxonomy.FieldID)
		for _, f := range fields {
			if !taxonomy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taxonomy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.ParentID(); ok {
		_spec.SetField(taxonomy.FieldParentID, field.TypeString, value)
	}
	if tuo.mutation.ParentIDCleared() {
		_spec.ClearField(taxonomy.FieldParentID, field.TypeString)
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(taxonomy.FieldTitle, field.TypeString, value)
	}
	if tuo.mutation.TitleCleared() {
		_spec.ClearField(taxonomy.FieldTitle, field.TypeString)
	}
	if value, ok := tuo.mutation.Slug(); ok {
		_spec.SetField(taxonomy.FieldSlug, field.TypeString, value)
	}
	if tuo.mutation.SlugCleared() {
		_spec.ClearField(taxonomy.FieldSlug, field.TypeString)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(taxonomy.FieldType, field.TypeString, value)
	}
	if tuo.mutation.TypeCleared() {
		_spec.ClearField(taxonomy.FieldType, field.TypeString)
	}
	_node = &Taxonomy{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taxonomy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}