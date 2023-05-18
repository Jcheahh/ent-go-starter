// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/linkvisit"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkVisitQuery is the builder for querying LinkVisit entities.
type LinkVisitQuery struct {
	config
	ctx        *QueryContext
	order      []linkvisit.Order
	inters     []Interceptor
	predicates []predicate.LinkVisit
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*LinkVisit) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinkVisitQuery builder.
func (lvq *LinkVisitQuery) Where(ps ...predicate.LinkVisit) *LinkVisitQuery {
	lvq.predicates = append(lvq.predicates, ps...)
	return lvq
}

// Limit the number of records to be returned by this query.
func (lvq *LinkVisitQuery) Limit(limit int) *LinkVisitQuery {
	lvq.ctx.Limit = &limit
	return lvq
}

// Offset to start from.
func (lvq *LinkVisitQuery) Offset(offset int) *LinkVisitQuery {
	lvq.ctx.Offset = &offset
	return lvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lvq *LinkVisitQuery) Unique(unique bool) *LinkVisitQuery {
	lvq.ctx.Unique = &unique
	return lvq
}

// Order specifies how the records should be ordered.
func (lvq *LinkVisitQuery) Order(o ...linkvisit.Order) *LinkVisitQuery {
	lvq.order = append(lvq.order, o...)
	return lvq
}

// First returns the first LinkVisit entity from the query.
// Returns a *NotFoundError when no LinkVisit was found.
func (lvq *LinkVisitQuery) First(ctx context.Context) (*LinkVisit, error) {
	nodes, err := lvq.Limit(1).All(setContextOp(ctx, lvq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linkvisit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lvq *LinkVisitQuery) FirstX(ctx context.Context) *LinkVisit {
	node, err := lvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LinkVisit ID from the query.
// Returns a *NotFoundError when no LinkVisit ID was found.
func (lvq *LinkVisitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lvq.Limit(1).IDs(setContextOp(ctx, lvq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linkvisit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lvq *LinkVisitQuery) FirstIDX(ctx context.Context) int {
	id, err := lvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LinkVisit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LinkVisit entity is found.
// Returns a *NotFoundError when no LinkVisit entities are found.
func (lvq *LinkVisitQuery) Only(ctx context.Context) (*LinkVisit, error) {
	nodes, err := lvq.Limit(2).All(setContextOp(ctx, lvq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{linkvisit.Label}
	default:
		return nil, &NotSingularError{linkvisit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lvq *LinkVisitQuery) OnlyX(ctx context.Context) *LinkVisit {
	node, err := lvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LinkVisit ID in the query.
// Returns a *NotSingularError when more than one LinkVisit ID is found.
// Returns a *NotFoundError when no entities are found.
func (lvq *LinkVisitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lvq.Limit(2).IDs(setContextOp(ctx, lvq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{linkvisit.Label}
	default:
		err = &NotSingularError{linkvisit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lvq *LinkVisitQuery) OnlyIDX(ctx context.Context) int {
	id, err := lvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LinkVisits.
func (lvq *LinkVisitQuery) All(ctx context.Context) ([]*LinkVisit, error) {
	ctx = setContextOp(ctx, lvq.ctx, "All")
	if err := lvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LinkVisit, *LinkVisitQuery]()
	return withInterceptors[[]*LinkVisit](ctx, lvq, qr, lvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lvq *LinkVisitQuery) AllX(ctx context.Context) []*LinkVisit {
	nodes, err := lvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LinkVisit IDs.
func (lvq *LinkVisitQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lvq.ctx.Unique == nil && lvq.path != nil {
		lvq.Unique(true)
	}
	ctx = setContextOp(ctx, lvq.ctx, "IDs")
	if err = lvq.Select(linkvisit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lvq *LinkVisitQuery) IDsX(ctx context.Context) []int {
	ids, err := lvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lvq *LinkVisitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lvq.ctx, "Count")
	if err := lvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lvq, querierCount[*LinkVisitQuery](), lvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lvq *LinkVisitQuery) CountX(ctx context.Context) int {
	count, err := lvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lvq *LinkVisitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lvq.ctx, "Exist")
	switch _, err := lvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lvq *LinkVisitQuery) ExistX(ctx context.Context) bool {
	exist, err := lvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinkVisitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lvq *LinkVisitQuery) Clone() *LinkVisitQuery {
	if lvq == nil {
		return nil
	}
	return &LinkVisitQuery{
		config:     lvq.config,
		ctx:        lvq.ctx.Clone(),
		order:      append([]linkvisit.Order{}, lvq.order...),
		inters:     append([]Interceptor{}, lvq.inters...),
		predicates: append([]predicate.LinkVisit{}, lvq.predicates...),
		// clone intermediate query.
		sql:  lvq.sql.Clone(),
		path: lvq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DateCreated string `json:"dateCreated,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LinkVisit.Query().
//		GroupBy(linkvisit.FieldDateCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lvq *LinkVisitQuery) GroupBy(field string, fields ...string) *LinkVisitGroupBy {
	lvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LinkVisitGroupBy{build: lvq}
	grbuild.flds = &lvq.ctx.Fields
	grbuild.label = linkvisit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DateCreated string `json:"dateCreated,omitempty"`
//	}
//
//	client.LinkVisit.Query().
//		Select(linkvisit.FieldDateCreated).
//		Scan(ctx, &v)
func (lvq *LinkVisitQuery) Select(fields ...string) *LinkVisitSelect {
	lvq.ctx.Fields = append(lvq.ctx.Fields, fields...)
	sbuild := &LinkVisitSelect{LinkVisitQuery: lvq}
	sbuild.label = linkvisit.Label
	sbuild.flds, sbuild.scan = &lvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LinkVisitSelect configured with the given aggregations.
func (lvq *LinkVisitQuery) Aggregate(fns ...AggregateFunc) *LinkVisitSelect {
	return lvq.Select().Aggregate(fns...)
}

func (lvq *LinkVisitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lvq); err != nil {
				return err
			}
		}
	}
	for _, f := range lvq.ctx.Fields {
		if !linkvisit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lvq.path != nil {
		prev, err := lvq.path(ctx)
		if err != nil {
			return err
		}
		lvq.sql = prev
	}
	return nil
}

func (lvq *LinkVisitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LinkVisit, error) {
	var (
		nodes   = []*LinkVisit{}
		withFKs = lvq.withFKs
		_spec   = lvq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, linkvisit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LinkVisit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LinkVisit{config: lvq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(lvq.modifiers) > 0 {
		_spec.Modifiers = lvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range lvq.loadTotal {
		if err := lvq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lvq *LinkVisitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lvq.querySpec()
	if len(lvq.modifiers) > 0 {
		_spec.Modifiers = lvq.modifiers
	}
	_spec.Node.Columns = lvq.ctx.Fields
	if len(lvq.ctx.Fields) > 0 {
		_spec.Unique = lvq.ctx.Unique != nil && *lvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lvq.driver, _spec)
}

func (lvq *LinkVisitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(linkvisit.Table, linkvisit.Columns, sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt))
	_spec.From = lvq.sql
	if unique := lvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lvq.path != nil {
		_spec.Unique = true
	}
	if fields := lvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linkvisit.FieldID)
		for i := range fields {
			if fields[i] != linkvisit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lvq *LinkVisitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lvq.driver.Dialect())
	t1 := builder.Table(linkvisit.Table)
	columns := lvq.ctx.Fields
	if len(columns) == 0 {
		columns = linkvisit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lvq.sql != nil {
		selector = lvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lvq.ctx.Unique != nil && *lvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lvq.predicates {
		p(selector)
	}
	for _, p := range lvq.order {
		p(selector)
	}
	if offset := lvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LinkVisitGroupBy is the group-by builder for LinkVisit entities.
type LinkVisitGroupBy struct {
	selector
	build *LinkVisitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lvgb *LinkVisitGroupBy) Aggregate(fns ...AggregateFunc) *LinkVisitGroupBy {
	lvgb.fns = append(lvgb.fns, fns...)
	return lvgb
}

// Scan applies the selector query and scans the result into the given value.
func (lvgb *LinkVisitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lvgb.build.ctx, "GroupBy")
	if err := lvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LinkVisitQuery, *LinkVisitGroupBy](ctx, lvgb.build, lvgb, lvgb.build.inters, v)
}

func (lvgb *LinkVisitGroupBy) sqlScan(ctx context.Context, root *LinkVisitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lvgb.fns))
	for _, fn := range lvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lvgb.flds)+len(lvgb.fns))
		for _, f := range *lvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LinkVisitSelect is the builder for selecting fields of LinkVisit entities.
type LinkVisitSelect struct {
	*LinkVisitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lvs *LinkVisitSelect) Aggregate(fns ...AggregateFunc) *LinkVisitSelect {
	lvs.fns = append(lvs.fns, fns...)
	return lvs
}

// Scan applies the selector query and scans the result into the given value.
func (lvs *LinkVisitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lvs.ctx, "Select")
	if err := lvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LinkVisitQuery, *LinkVisitSelect](ctx, lvs.LinkVisitQuery, lvs, lvs.inters, v)
}

func (lvs *LinkVisitSelect) sqlScan(ctx context.Context, root *LinkVisitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lvs.fns))
	for _, fn := range lvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
