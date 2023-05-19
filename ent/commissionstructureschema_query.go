// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/commissionstructureschema"
	"entdemo/ent/predicate"
	"entdemo/ent/userseller"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionStructureSchemaQuery is the builder for querying CommissionStructureSchema entities.
type CommissionStructureSchemaQuery struct {
	config
	ctx                    *QueryContext
	order                  []commissionstructureschema.OrderOption
	inters                 []Interceptor
	predicates             []predicate.CommissionStructureSchema
	withProductSeller      *UserSellerQuery
	withFKs                bool
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*CommissionStructureSchema) error
	withNamedProductSeller map[string]*UserSellerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommissionStructureSchemaQuery builder.
func (cssq *CommissionStructureSchemaQuery) Where(ps ...predicate.CommissionStructureSchema) *CommissionStructureSchemaQuery {
	cssq.predicates = append(cssq.predicates, ps...)
	return cssq
}

// Limit the number of records to be returned by this query.
func (cssq *CommissionStructureSchemaQuery) Limit(limit int) *CommissionStructureSchemaQuery {
	cssq.ctx.Limit = &limit
	return cssq
}

// Offset to start from.
func (cssq *CommissionStructureSchemaQuery) Offset(offset int) *CommissionStructureSchemaQuery {
	cssq.ctx.Offset = &offset
	return cssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cssq *CommissionStructureSchemaQuery) Unique(unique bool) *CommissionStructureSchemaQuery {
	cssq.ctx.Unique = &unique
	return cssq
}

// Order specifies how the records should be ordered.
func (cssq *CommissionStructureSchemaQuery) Order(o ...commissionstructureschema.OrderOption) *CommissionStructureSchemaQuery {
	cssq.order = append(cssq.order, o...)
	return cssq
}

// QueryProductSeller chains the current query on the "productSeller" edge.
func (cssq *CommissionStructureSchemaQuery) QueryProductSeller() *UserSellerQuery {
	query := (&UserSellerClient{config: cssq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cssq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cssq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commissionstructureschema.Table, commissionstructureschema.FieldID, selector),
			sqlgraph.To(userseller.Table, userseller.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, commissionstructureschema.ProductSellerTable, commissionstructureschema.ProductSellerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cssq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommissionStructureSchema entity from the query.
// Returns a *NotFoundError when no CommissionStructureSchema was found.
func (cssq *CommissionStructureSchemaQuery) First(ctx context.Context) (*CommissionStructureSchema, error) {
	nodes, err := cssq.Limit(1).All(setContextOp(ctx, cssq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commissionstructureschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) FirstX(ctx context.Context) *CommissionStructureSchema {
	node, err := cssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommissionStructureSchema ID from the query.
// Returns a *NotFoundError when no CommissionStructureSchema ID was found.
func (cssq *CommissionStructureSchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cssq.Limit(1).IDs(setContextOp(ctx, cssq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commissionstructureschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := cssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommissionStructureSchema entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommissionStructureSchema entity is found.
// Returns a *NotFoundError when no CommissionStructureSchema entities are found.
func (cssq *CommissionStructureSchemaQuery) Only(ctx context.Context) (*CommissionStructureSchema, error) {
	nodes, err := cssq.Limit(2).All(setContextOp(ctx, cssq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commissionstructureschema.Label}
	default:
		return nil, &NotSingularError{commissionstructureschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) OnlyX(ctx context.Context) *CommissionStructureSchema {
	node, err := cssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommissionStructureSchema ID in the query.
// Returns a *NotSingularError when more than one CommissionStructureSchema ID is found.
// Returns a *NotFoundError when no entities are found.
func (cssq *CommissionStructureSchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cssq.Limit(2).IDs(setContextOp(ctx, cssq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commissionstructureschema.Label}
	default:
		err = &NotSingularError{commissionstructureschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := cssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommissionStructureSchemas.
func (cssq *CommissionStructureSchemaQuery) All(ctx context.Context) ([]*CommissionStructureSchema, error) {
	ctx = setContextOp(ctx, cssq.ctx, "All")
	if err := cssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommissionStructureSchema, *CommissionStructureSchemaQuery]()
	return withInterceptors[[]*CommissionStructureSchema](ctx, cssq, qr, cssq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) AllX(ctx context.Context) []*CommissionStructureSchema {
	nodes, err := cssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommissionStructureSchema IDs.
func (cssq *CommissionStructureSchemaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cssq.ctx.Unique == nil && cssq.path != nil {
		cssq.Unique(true)
	}
	ctx = setContextOp(ctx, cssq.ctx, "IDs")
	if err = cssq.Select(commissionstructureschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := cssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cssq *CommissionStructureSchemaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cssq.ctx, "Count")
	if err := cssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cssq, querierCount[*CommissionStructureSchemaQuery](), cssq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) CountX(ctx context.Context) int {
	count, err := cssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cssq *CommissionStructureSchemaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cssq.ctx, "Exist")
	switch _, err := cssq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cssq *CommissionStructureSchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := cssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommissionStructureSchemaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cssq *CommissionStructureSchemaQuery) Clone() *CommissionStructureSchemaQuery {
	if cssq == nil {
		return nil
	}
	return &CommissionStructureSchemaQuery{
		config:            cssq.config,
		ctx:               cssq.ctx.Clone(),
		order:             append([]commissionstructureschema.OrderOption{}, cssq.order...),
		inters:            append([]Interceptor{}, cssq.inters...),
		predicates:        append([]predicate.CommissionStructureSchema{}, cssq.predicates...),
		withProductSeller: cssq.withProductSeller.Clone(),
		// clone intermediate query.
		sql:  cssq.sql.Clone(),
		path: cssq.path,
	}
}

// WithProductSeller tells the query-builder to eager-load the nodes that are connected to
// the "productSeller" edge. The optional arguments are used to configure the query builder of the edge.
func (cssq *CommissionStructureSchemaQuery) WithProductSeller(opts ...func(*UserSellerQuery)) *CommissionStructureSchemaQuery {
	query := (&UserSellerClient{config: cssq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cssq.withProductSeller = query
	return cssq
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
//	client.CommissionStructureSchema.Query().
//		GroupBy(commissionstructureschema.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cssq *CommissionStructureSchemaQuery) GroupBy(field string, fields ...string) *CommissionStructureSchemaGroupBy {
	cssq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommissionStructureSchemaGroupBy{build: cssq}
	grbuild.flds = &cssq.ctx.Fields
	grbuild.label = commissionstructureschema.Label
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
//	client.CommissionStructureSchema.Query().
//		Select(commissionstructureschema.FieldName).
//		Scan(ctx, &v)
func (cssq *CommissionStructureSchemaQuery) Select(fields ...string) *CommissionStructureSchemaSelect {
	cssq.ctx.Fields = append(cssq.ctx.Fields, fields...)
	sbuild := &CommissionStructureSchemaSelect{CommissionStructureSchemaQuery: cssq}
	sbuild.label = commissionstructureschema.Label
	sbuild.flds, sbuild.scan = &cssq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommissionStructureSchemaSelect configured with the given aggregations.
func (cssq *CommissionStructureSchemaQuery) Aggregate(fns ...AggregateFunc) *CommissionStructureSchemaSelect {
	return cssq.Select().Aggregate(fns...)
}

func (cssq *CommissionStructureSchemaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cssq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cssq); err != nil {
				return err
			}
		}
	}
	for _, f := range cssq.ctx.Fields {
		if !commissionstructureschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cssq.path != nil {
		prev, err := cssq.path(ctx)
		if err != nil {
			return err
		}
		cssq.sql = prev
	}
	return nil
}

func (cssq *CommissionStructureSchemaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommissionStructureSchema, error) {
	var (
		nodes       = []*CommissionStructureSchema{}
		withFKs     = cssq.withFKs
		_spec       = cssq.querySpec()
		loadedTypes = [1]bool{
			cssq.withProductSeller != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, commissionstructureschema.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommissionStructureSchema).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommissionStructureSchema{config: cssq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cssq.modifiers) > 0 {
		_spec.Modifiers = cssq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cssq.withProductSeller; query != nil {
		if err := cssq.loadProductSeller(ctx, query, nodes,
			func(n *CommissionStructureSchema) { n.Edges.ProductSeller = []*UserSeller{} },
			func(n *CommissionStructureSchema, e *UserSeller) {
				n.Edges.ProductSeller = append(n.Edges.ProductSeller, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range cssq.withNamedProductSeller {
		if err := cssq.loadProductSeller(ctx, query, nodes,
			func(n *CommissionStructureSchema) { n.appendNamedProductSeller(name) },
			func(n *CommissionStructureSchema, e *UserSeller) { n.appendNamedProductSeller(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cssq.loadTotal {
		if err := cssq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cssq *CommissionStructureSchemaQuery) loadProductSeller(ctx context.Context, query *UserSellerQuery, nodes []*CommissionStructureSchema, init func(*CommissionStructureSchema), assign func(*CommissionStructureSchema, *UserSeller)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CommissionStructureSchema)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserSeller(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(commissionstructureschema.ProductSellerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.commission_structure_schema_product_seller
		if fk == nil {
			return fmt.Errorf(`foreign-key "commission_structure_schema_product_seller" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "commission_structure_schema_product_seller" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cssq *CommissionStructureSchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cssq.querySpec()
	if len(cssq.modifiers) > 0 {
		_spec.Modifiers = cssq.modifiers
	}
	_spec.Node.Columns = cssq.ctx.Fields
	if len(cssq.ctx.Fields) > 0 {
		_spec.Unique = cssq.ctx.Unique != nil && *cssq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cssq.driver, _spec)
}

func (cssq *CommissionStructureSchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commissionstructureschema.Table, commissionstructureschema.Columns, sqlgraph.NewFieldSpec(commissionstructureschema.FieldID, field.TypeInt))
	_spec.From = cssq.sql
	if unique := cssq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cssq.path != nil {
		_spec.Unique = true
	}
	if fields := cssq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissionstructureschema.FieldID)
		for i := range fields {
			if fields[i] != commissionstructureschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cssq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cssq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cssq *CommissionStructureSchemaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cssq.driver.Dialect())
	t1 := builder.Table(commissionstructureschema.Table)
	columns := cssq.ctx.Fields
	if len(columns) == 0 {
		columns = commissionstructureschema.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cssq.sql != nil {
		selector = cssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cssq.ctx.Unique != nil && *cssq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cssq.predicates {
		p(selector)
	}
	for _, p := range cssq.order {
		p(selector)
	}
	if offset := cssq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cssq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProductSeller tells the query-builder to eager-load the nodes that are connected to the "productSeller"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cssq *CommissionStructureSchemaQuery) WithNamedProductSeller(name string, opts ...func(*UserSellerQuery)) *CommissionStructureSchemaQuery {
	query := (&UserSellerClient{config: cssq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cssq.withNamedProductSeller == nil {
		cssq.withNamedProductSeller = make(map[string]*UserSellerQuery)
	}
	cssq.withNamedProductSeller[name] = query
	return cssq
}

// CommissionStructureSchemaGroupBy is the group-by builder for CommissionStructureSchema entities.
type CommissionStructureSchemaGroupBy struct {
	selector
	build *CommissionStructureSchemaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cssgb *CommissionStructureSchemaGroupBy) Aggregate(fns ...AggregateFunc) *CommissionStructureSchemaGroupBy {
	cssgb.fns = append(cssgb.fns, fns...)
	return cssgb
}

// Scan applies the selector query and scans the result into the given value.
func (cssgb *CommissionStructureSchemaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cssgb.build.ctx, "GroupBy")
	if err := cssgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommissionStructureSchemaQuery, *CommissionStructureSchemaGroupBy](ctx, cssgb.build, cssgb, cssgb.build.inters, v)
}

func (cssgb *CommissionStructureSchemaGroupBy) sqlScan(ctx context.Context, root *CommissionStructureSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cssgb.fns))
	for _, fn := range cssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cssgb.flds)+len(cssgb.fns))
		for _, f := range *cssgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cssgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cssgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommissionStructureSchemaSelect is the builder for selecting fields of CommissionStructureSchema entities.
type CommissionStructureSchemaSelect struct {
	*CommissionStructureSchemaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (csss *CommissionStructureSchemaSelect) Aggregate(fns ...AggregateFunc) *CommissionStructureSchemaSelect {
	csss.fns = append(csss.fns, fns...)
	return csss
}

// Scan applies the selector query and scans the result into the given value.
func (csss *CommissionStructureSchemaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csss.ctx, "Select")
	if err := csss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommissionStructureSchemaQuery, *CommissionStructureSchemaSelect](ctx, csss.CommissionStructureSchemaQuery, csss, csss.inters, v)
}

func (csss *CommissionStructureSchemaSelect) sqlScan(ctx context.Context, root *CommissionStructureSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(csss.fns))
	for _, fn := range csss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*csss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
