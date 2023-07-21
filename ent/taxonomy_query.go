// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/taxonomy"
)

// TaxonomyQuery is the builder for querying Taxonomy entities.
type TaxonomyQuery struct {
	config
	ctx        *QueryContext
	order      []taxonomy.OrderOption
	inters     []Interceptor
	predicates []predicate.Taxonomy
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaxonomyQuery builder.
func (tq *TaxonomyQuery) Where(ps ...predicate.Taxonomy) *TaxonomyQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TaxonomyQuery) Limit(limit int) *TaxonomyQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TaxonomyQuery) Offset(offset int) *TaxonomyQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TaxonomyQuery) Unique(unique bool) *TaxonomyQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TaxonomyQuery) Order(o ...taxonomy.OrderOption) *TaxonomyQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Taxonomy entity from the query.
// Returns a *NotFoundError when no Taxonomy was found.
func (tq *TaxonomyQuery) First(ctx context.Context) (*Taxonomy, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taxonomy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TaxonomyQuery) FirstX(ctx context.Context) *Taxonomy {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Taxonomy ID from the query.
// Returns a *NotFoundError when no Taxonomy ID was found.
func (tq *TaxonomyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taxonomy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TaxonomyQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Taxonomy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Taxonomy entity is found.
// Returns a *NotFoundError when no Taxonomy entities are found.
func (tq *TaxonomyQuery) Only(ctx context.Context) (*Taxonomy, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taxonomy.Label}
	default:
		return nil, &NotSingularError{taxonomy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TaxonomyQuery) OnlyX(ctx context.Context) *Taxonomy {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Taxonomy ID in the query.
// Returns a *NotSingularError when more than one Taxonomy ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TaxonomyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taxonomy.Label}
	default:
		err = &NotSingularError{taxonomy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TaxonomyQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Taxonomies.
func (tq *TaxonomyQuery) All(ctx context.Context) ([]*Taxonomy, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Taxonomy, *TaxonomyQuery]()
	return withInterceptors[[]*Taxonomy](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TaxonomyQuery) AllX(ctx context.Context) []*Taxonomy {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Taxonomy IDs.
func (tq *TaxonomyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(taxonomy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TaxonomyQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TaxonomyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TaxonomyQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TaxonomyQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TaxonomyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TaxonomyQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaxonomyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TaxonomyQuery) Clone() *TaxonomyQuery {
	if tq == nil {
		return nil
	}
	return &TaxonomyQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]taxonomy.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Taxonomy{}, tq.predicates...),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Taxonomy.Query().
//		GroupBy(taxonomy.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TaxonomyQuery) GroupBy(field string, fields ...string) *TaxonomyGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaxonomyGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = taxonomy.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Taxonomy.Query().
//		Select(taxonomy.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TaxonomyQuery) Select(fields ...string) *TaxonomySelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TaxonomySelect{TaxonomyQuery: tq}
	sbuild.label = taxonomy.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaxonomySelect configured with the given aggregations.
func (tq *TaxonomyQuery) Aggregate(fns ...AggregateFunc) *TaxonomySelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TaxonomyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !taxonomy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TaxonomyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Taxonomy, error) {
	var (
		nodes = []*Taxonomy{}
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Taxonomy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Taxonomy{config: tq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tq *TaxonomyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TaxonomyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(taxonomy.Table, taxonomy.Columns, sqlgraph.NewFieldSpec(taxonomy.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taxonomy.FieldID)
		for i := range fields {
			if fields[i] != taxonomy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TaxonomyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(taxonomy.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = taxonomy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaxonomyGroupBy is the group-by builder for Taxonomy entities.
type TaxonomyGroupBy struct {
	selector
	build *TaxonomyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TaxonomyGroupBy) Aggregate(fns ...AggregateFunc) *TaxonomyGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TaxonomyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaxonomyQuery, *TaxonomyGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TaxonomyGroupBy) sqlScan(ctx context.Context, root *TaxonomyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TaxonomySelect is the builder for selecting fields of Taxonomy entities.
type TaxonomySelect struct {
	*TaxonomyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TaxonomySelect) Aggregate(fns ...AggregateFunc) *TaxonomySelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TaxonomySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaxonomyQuery, *TaxonomySelect](ctx, ts.TaxonomyQuery, ts, ts.inters, v)
}

func (ts *TaxonomySelect) sqlScan(ctx context.Context, root *TaxonomyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
