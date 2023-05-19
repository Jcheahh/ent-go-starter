// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/shippingaddress"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShippingAddressQuery is the builder for querying ShippingAddress entities.
type ShippingAddressQuery struct {
	config
	ctx        *QueryContext
	order      []shippingaddress.OrderOption
	inters     []Interceptor
	predicates []predicate.ShippingAddress
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*ShippingAddress) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShippingAddressQuery builder.
func (saq *ShippingAddressQuery) Where(ps ...predicate.ShippingAddress) *ShippingAddressQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit the number of records to be returned by this query.
func (saq *ShippingAddressQuery) Limit(limit int) *ShippingAddressQuery {
	saq.ctx.Limit = &limit
	return saq
}

// Offset to start from.
func (saq *ShippingAddressQuery) Offset(offset int) *ShippingAddressQuery {
	saq.ctx.Offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *ShippingAddressQuery) Unique(unique bool) *ShippingAddressQuery {
	saq.ctx.Unique = &unique
	return saq
}

// Order specifies how the records should be ordered.
func (saq *ShippingAddressQuery) Order(o ...shippingaddress.OrderOption) *ShippingAddressQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// First returns the first ShippingAddress entity from the query.
// Returns a *NotFoundError when no ShippingAddress was found.
func (saq *ShippingAddressQuery) First(ctx context.Context) (*ShippingAddress, error) {
	nodes, err := saq.Limit(1).All(setContextOp(ctx, saq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shippingaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *ShippingAddressQuery) FirstX(ctx context.Context) *ShippingAddress {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShippingAddress ID from the query.
// Returns a *NotFoundError when no ShippingAddress ID was found.
func (saq *ShippingAddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(1).IDs(setContextOp(ctx, saq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shippingaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *ShippingAddressQuery) FirstIDX(ctx context.Context) int {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShippingAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShippingAddress entity is found.
// Returns a *NotFoundError when no ShippingAddress entities are found.
func (saq *ShippingAddressQuery) Only(ctx context.Context) (*ShippingAddress, error) {
	nodes, err := saq.Limit(2).All(setContextOp(ctx, saq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shippingaddress.Label}
	default:
		return nil, &NotSingularError{shippingaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *ShippingAddressQuery) OnlyX(ctx context.Context) *ShippingAddress {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShippingAddress ID in the query.
// Returns a *NotSingularError when more than one ShippingAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *ShippingAddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(2).IDs(setContextOp(ctx, saq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shippingaddress.Label}
	default:
		err = &NotSingularError{shippingaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *ShippingAddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShippingAddresses.
func (saq *ShippingAddressQuery) All(ctx context.Context) ([]*ShippingAddress, error) {
	ctx = setContextOp(ctx, saq.ctx, "All")
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShippingAddress, *ShippingAddressQuery]()
	return withInterceptors[[]*ShippingAddress](ctx, saq, qr, saq.inters)
}

// AllX is like All, but panics if an error occurs.
func (saq *ShippingAddressQuery) AllX(ctx context.Context) []*ShippingAddress {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShippingAddress IDs.
func (saq *ShippingAddressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if saq.ctx.Unique == nil && saq.path != nil {
		saq.Unique(true)
	}
	ctx = setContextOp(ctx, saq.ctx, "IDs")
	if err = saq.Select(shippingaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *ShippingAddressQuery) IDsX(ctx context.Context) []int {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *ShippingAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, saq.ctx, "Count")
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, saq, querierCount[*ShippingAddressQuery](), saq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (saq *ShippingAddressQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *ShippingAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, saq.ctx, "Exist")
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *ShippingAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShippingAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *ShippingAddressQuery) Clone() *ShippingAddressQuery {
	if saq == nil {
		return nil
	}
	return &ShippingAddressQuery{
		config:     saq.config,
		ctx:        saq.ctx.Clone(),
		order:      append([]shippingaddress.OrderOption{}, saq.order...),
		inters:     append([]Interceptor{}, saq.inters...),
		predicates: append([]predicate.ShippingAddress{}, saq.predicates...),
		// clone intermediate query.
		sql:  saq.sql.Clone(),
		path: saq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ShippingAddress.Query().
//		GroupBy(shippingaddress.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (saq *ShippingAddressQuery) GroupBy(field string, fields ...string) *ShippingAddressGroupBy {
	saq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShippingAddressGroupBy{build: saq}
	grbuild.flds = &saq.ctx.Fields
	grbuild.label = shippingaddress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ShippingAddress.Query().
//		Select(shippingaddress.FieldName).
//		Scan(ctx, &v)
func (saq *ShippingAddressQuery) Select(fields ...string) *ShippingAddressSelect {
	saq.ctx.Fields = append(saq.ctx.Fields, fields...)
	sbuild := &ShippingAddressSelect{ShippingAddressQuery: saq}
	sbuild.label = shippingaddress.Label
	sbuild.flds, sbuild.scan = &saq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShippingAddressSelect configured with the given aggregations.
func (saq *ShippingAddressQuery) Aggregate(fns ...AggregateFunc) *ShippingAddressSelect {
	return saq.Select().Aggregate(fns...)
}

func (saq *ShippingAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range saq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, saq); err != nil {
				return err
			}
		}
	}
	for _, f := range saq.ctx.Fields {
		if !shippingaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *ShippingAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShippingAddress, error) {
	var (
		nodes   = []*ShippingAddress{}
		withFKs = saq.withFKs
		_spec   = saq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shippingaddress.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShippingAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShippingAddress{config: saq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range saq.loadTotal {
		if err := saq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (saq *ShippingAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	_spec.Node.Columns = saq.ctx.Fields
	if len(saq.ctx.Fields) > 0 {
		_spec.Unique = saq.ctx.Unique != nil && *saq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *ShippingAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shippingaddress.Table, shippingaddress.Columns, sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt))
	_spec.From = saq.sql
	if unique := saq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if saq.path != nil {
		_spec.Unique = true
	}
	if fields := saq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shippingaddress.FieldID)
		for i := range fields {
			if fields[i] != shippingaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *ShippingAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(shippingaddress.Table)
	columns := saq.ctx.Fields
	if len(columns) == 0 {
		columns = shippingaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.ctx.Unique != nil && *saq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShippingAddressGroupBy is the group-by builder for ShippingAddress entities.
type ShippingAddressGroupBy struct {
	selector
	build *ShippingAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *ShippingAddressGroupBy) Aggregate(fns ...AggregateFunc) *ShippingAddressGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the selector query and scans the result into the given value.
func (sagb *ShippingAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sagb.build.ctx, "GroupBy")
	if err := sagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShippingAddressQuery, *ShippingAddressGroupBy](ctx, sagb.build, sagb, sagb.build.inters, v)
}

func (sagb *ShippingAddressGroupBy) sqlScan(ctx context.Context, root *ShippingAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sagb.flds)+len(sagb.fns))
		for _, f := range *sagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShippingAddressSelect is the builder for selecting fields of ShippingAddress entities.
type ShippingAddressSelect struct {
	*ShippingAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sas *ShippingAddressSelect) Aggregate(fns ...AggregateFunc) *ShippingAddressSelect {
	sas.fns = append(sas.fns, fns...)
	return sas
}

// Scan applies the selector query and scans the result into the given value.
func (sas *ShippingAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sas.ctx, "Select")
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShippingAddressQuery, *ShippingAddressSelect](ctx, sas.ShippingAddressQuery, sas, sas.inters, v)
}

func (sas *ShippingAddressSelect) sqlScan(ctx context.Context, root *ShippingAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sas.fns))
	for _, fn := range sas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
