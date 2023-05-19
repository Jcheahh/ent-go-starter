// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/productattribute"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductAttributeQuery is the builder for querying ProductAttribute entities.
type ProductAttributeQuery struct {
	config
	ctx        *QueryContext
	order      []productattribute.OrderOption
	inters     []Interceptor
	predicates []predicate.ProductAttribute
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*ProductAttribute) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductAttributeQuery builder.
func (paq *ProductAttributeQuery) Where(ps ...predicate.ProductAttribute) *ProductAttributeQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit the number of records to be returned by this query.
func (paq *ProductAttributeQuery) Limit(limit int) *ProductAttributeQuery {
	paq.ctx.Limit = &limit
	return paq
}

// Offset to start from.
func (paq *ProductAttributeQuery) Offset(offset int) *ProductAttributeQuery {
	paq.ctx.Offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *ProductAttributeQuery) Unique(unique bool) *ProductAttributeQuery {
	paq.ctx.Unique = &unique
	return paq
}

// Order specifies how the records should be ordered.
func (paq *ProductAttributeQuery) Order(o ...productattribute.OrderOption) *ProductAttributeQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// First returns the first ProductAttribute entity from the query.
// Returns a *NotFoundError when no ProductAttribute was found.
func (paq *ProductAttributeQuery) First(ctx context.Context) (*ProductAttribute, error) {
	nodes, err := paq.Limit(1).All(setContextOp(ctx, paq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productattribute.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *ProductAttributeQuery) FirstX(ctx context.Context) *ProductAttribute {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductAttribute ID from the query.
// Returns a *NotFoundError when no ProductAttribute ID was found.
func (paq *ProductAttributeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(1).IDs(setContextOp(ctx, paq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productattribute.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *ProductAttributeQuery) FirstIDX(ctx context.Context) int {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductAttribute entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductAttribute entity is found.
// Returns a *NotFoundError when no ProductAttribute entities are found.
func (paq *ProductAttributeQuery) Only(ctx context.Context) (*ProductAttribute, error) {
	nodes, err := paq.Limit(2).All(setContextOp(ctx, paq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productattribute.Label}
	default:
		return nil, &NotSingularError{productattribute.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *ProductAttributeQuery) OnlyX(ctx context.Context) *ProductAttribute {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductAttribute ID in the query.
// Returns a *NotSingularError when more than one ProductAttribute ID is found.
// Returns a *NotFoundError when no entities are found.
func (paq *ProductAttributeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(2).IDs(setContextOp(ctx, paq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productattribute.Label}
	default:
		err = &NotSingularError{productattribute.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *ProductAttributeQuery) OnlyIDX(ctx context.Context) int {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductAttributes.
func (paq *ProductAttributeQuery) All(ctx context.Context) ([]*ProductAttribute, error) {
	ctx = setContextOp(ctx, paq.ctx, "All")
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductAttribute, *ProductAttributeQuery]()
	return withInterceptors[[]*ProductAttribute](ctx, paq, qr, paq.inters)
}

// AllX is like All, but panics if an error occurs.
func (paq *ProductAttributeQuery) AllX(ctx context.Context) []*ProductAttribute {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductAttribute IDs.
func (paq *ProductAttributeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if paq.ctx.Unique == nil && paq.path != nil {
		paq.Unique(true)
	}
	ctx = setContextOp(ctx, paq.ctx, "IDs")
	if err = paq.Select(productattribute.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *ProductAttributeQuery) IDsX(ctx context.Context) []int {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *ProductAttributeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, paq.ctx, "Count")
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, paq, querierCount[*ProductAttributeQuery](), paq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (paq *ProductAttributeQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *ProductAttributeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, paq.ctx, "Exist")
	switch _, err := paq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *ProductAttributeQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductAttributeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *ProductAttributeQuery) Clone() *ProductAttributeQuery {
	if paq == nil {
		return nil
	}
	return &ProductAttributeQuery{
		config:     paq.config,
		ctx:        paq.ctx.Clone(),
		order:      append([]productattribute.OrderOption{}, paq.order...),
		inters:     append([]Interceptor{}, paq.inters...),
		predicates: append([]predicate.ProductAttribute{}, paq.predicates...),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name int `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductAttribute.Query().
//		GroupBy(productattribute.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (paq *ProductAttributeQuery) GroupBy(field string, fields ...string) *ProductAttributeGroupBy {
	paq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductAttributeGroupBy{build: paq}
	grbuild.flds = &paq.ctx.Fields
	grbuild.label = productattribute.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name int `json:"name,omitempty"`
//	}
//
//	client.ProductAttribute.Query().
//		Select(productattribute.FieldName).
//		Scan(ctx, &v)
func (paq *ProductAttributeQuery) Select(fields ...string) *ProductAttributeSelect {
	paq.ctx.Fields = append(paq.ctx.Fields, fields...)
	sbuild := &ProductAttributeSelect{ProductAttributeQuery: paq}
	sbuild.label = productattribute.Label
	sbuild.flds, sbuild.scan = &paq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductAttributeSelect configured with the given aggregations.
func (paq *ProductAttributeQuery) Aggregate(fns ...AggregateFunc) *ProductAttributeSelect {
	return paq.Select().Aggregate(fns...)
}

func (paq *ProductAttributeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range paq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, paq); err != nil {
				return err
			}
		}
	}
	for _, f := range paq.ctx.Fields {
		if !productattribute.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *ProductAttributeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductAttribute, error) {
	var (
		nodes   = []*ProductAttribute{}
		withFKs = paq.withFKs
		_spec   = paq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, productattribute.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductAttribute).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductAttribute{config: paq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range paq.loadTotal {
		if err := paq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (paq *ProductAttributeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	_spec.Node.Columns = paq.ctx.Fields
	if len(paq.ctx.Fields) > 0 {
		_spec.Unique = paq.ctx.Unique != nil && *paq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *ProductAttributeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productattribute.Table, productattribute.Columns, sqlgraph.NewFieldSpec(productattribute.FieldID, field.TypeInt))
	_spec.From = paq.sql
	if unique := paq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if paq.path != nil {
		_spec.Unique = true
	}
	if fields := paq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productattribute.FieldID)
		for i := range fields {
			if fields[i] != productattribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *ProductAttributeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(productattribute.Table)
	columns := paq.ctx.Fields
	if len(columns) == 0 {
		columns = productattribute.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if paq.ctx.Unique != nil && *paq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductAttributeGroupBy is the group-by builder for ProductAttribute entities.
type ProductAttributeGroupBy struct {
	selector
	build *ProductAttributeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *ProductAttributeGroupBy) Aggregate(fns ...AggregateFunc) *ProductAttributeGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the selector query and scans the result into the given value.
func (pagb *ProductAttributeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pagb.build.ctx, "GroupBy")
	if err := pagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductAttributeQuery, *ProductAttributeGroupBy](ctx, pagb.build, pagb, pagb.build.inters, v)
}

func (pagb *ProductAttributeGroupBy) sqlScan(ctx context.Context, root *ProductAttributeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pagb.flds)+len(pagb.fns))
		for _, f := range *pagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductAttributeSelect is the builder for selecting fields of ProductAttribute entities.
type ProductAttributeSelect struct {
	*ProductAttributeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pas *ProductAttributeSelect) Aggregate(fns ...AggregateFunc) *ProductAttributeSelect {
	pas.fns = append(pas.fns, fns...)
	return pas
}

// Scan applies the selector query and scans the result into the given value.
func (pas *ProductAttributeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pas.ctx, "Select")
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductAttributeQuery, *ProductAttributeSelect](ctx, pas.ProductAttributeQuery, pas, pas.inters, v)
}

func (pas *ProductAttributeSelect) sqlScan(ctx context.Context, root *ProductAttributeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pas.fns))
	for _, fn := range pas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}