// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/groupbuy"
	"entdemo/ent/predicate"
	"entdemo/ent/product"
	"entdemo/ent/transaction"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupBuyQuery is the builder for querying GroupBuy entities.
type GroupBuyQuery struct {
	config
	ctx                  *QueryContext
	order                []groupbuy.Order
	inters               []Interceptor
	predicates           []predicate.GroupBuy
	withProduct          *ProductQuery
	withTransaction      *TransactionQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*GroupBuy) error
	withNamedProduct     map[string]*ProductQuery
	withNamedTransaction map[string]*TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupBuyQuery builder.
func (gbq *GroupBuyQuery) Where(ps ...predicate.GroupBuy) *GroupBuyQuery {
	gbq.predicates = append(gbq.predicates, ps...)
	return gbq
}

// Limit the number of records to be returned by this query.
func (gbq *GroupBuyQuery) Limit(limit int) *GroupBuyQuery {
	gbq.ctx.Limit = &limit
	return gbq
}

// Offset to start from.
func (gbq *GroupBuyQuery) Offset(offset int) *GroupBuyQuery {
	gbq.ctx.Offset = &offset
	return gbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gbq *GroupBuyQuery) Unique(unique bool) *GroupBuyQuery {
	gbq.ctx.Unique = &unique
	return gbq
}

// Order specifies how the records should be ordered.
func (gbq *GroupBuyQuery) Order(o ...groupbuy.Order) *GroupBuyQuery {
	gbq.order = append(gbq.order, o...)
	return gbq
}

// QueryProduct chains the current query on the "product" edge.
func (gbq *GroupBuyQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: gbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupbuy.Table, groupbuy.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, groupbuy.ProductTable, groupbuy.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(gbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransaction chains the current query on the "transaction" edge.
func (gbq *GroupBuyQuery) QueryTransaction() *TransactionQuery {
	query := (&TransactionClient{config: gbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupbuy.Table, groupbuy.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, groupbuy.TransactionTable, groupbuy.TransactionColumn),
		)
		fromU = sqlgraph.SetNeighbors(gbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupBuy entity from the query.
// Returns a *NotFoundError when no GroupBuy was found.
func (gbq *GroupBuyQuery) First(ctx context.Context) (*GroupBuy, error) {
	nodes, err := gbq.Limit(1).All(setContextOp(ctx, gbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupbuy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gbq *GroupBuyQuery) FirstX(ctx context.Context) *GroupBuy {
	node, err := gbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupBuy ID from the query.
// Returns a *NotFoundError when no GroupBuy ID was found.
func (gbq *GroupBuyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gbq.Limit(1).IDs(setContextOp(ctx, gbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupbuy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gbq *GroupBuyQuery) FirstIDX(ctx context.Context) int {
	id, err := gbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupBuy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupBuy entity is found.
// Returns a *NotFoundError when no GroupBuy entities are found.
func (gbq *GroupBuyQuery) Only(ctx context.Context) (*GroupBuy, error) {
	nodes, err := gbq.Limit(2).All(setContextOp(ctx, gbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupbuy.Label}
	default:
		return nil, &NotSingularError{groupbuy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gbq *GroupBuyQuery) OnlyX(ctx context.Context) *GroupBuy {
	node, err := gbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupBuy ID in the query.
// Returns a *NotSingularError when more than one GroupBuy ID is found.
// Returns a *NotFoundError when no entities are found.
func (gbq *GroupBuyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gbq.Limit(2).IDs(setContextOp(ctx, gbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupbuy.Label}
	default:
		err = &NotSingularError{groupbuy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gbq *GroupBuyQuery) OnlyIDX(ctx context.Context) int {
	id, err := gbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupBuys.
func (gbq *GroupBuyQuery) All(ctx context.Context) ([]*GroupBuy, error) {
	ctx = setContextOp(ctx, gbq.ctx, "All")
	if err := gbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupBuy, *GroupBuyQuery]()
	return withInterceptors[[]*GroupBuy](ctx, gbq, qr, gbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gbq *GroupBuyQuery) AllX(ctx context.Context) []*GroupBuy {
	nodes, err := gbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupBuy IDs.
func (gbq *GroupBuyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gbq.ctx.Unique == nil && gbq.path != nil {
		gbq.Unique(true)
	}
	ctx = setContextOp(ctx, gbq.ctx, "IDs")
	if err = gbq.Select(groupbuy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gbq *GroupBuyQuery) IDsX(ctx context.Context) []int {
	ids, err := gbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gbq *GroupBuyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gbq.ctx, "Count")
	if err := gbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gbq, querierCount[*GroupBuyQuery](), gbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gbq *GroupBuyQuery) CountX(ctx context.Context) int {
	count, err := gbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gbq *GroupBuyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gbq.ctx, "Exist")
	switch _, err := gbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gbq *GroupBuyQuery) ExistX(ctx context.Context) bool {
	exist, err := gbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupBuyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gbq *GroupBuyQuery) Clone() *GroupBuyQuery {
	if gbq == nil {
		return nil
	}
	return &GroupBuyQuery{
		config:          gbq.config,
		ctx:             gbq.ctx.Clone(),
		order:           append([]groupbuy.Order{}, gbq.order...),
		inters:          append([]Interceptor{}, gbq.inters...),
		predicates:      append([]predicate.GroupBuy{}, gbq.predicates...),
		withProduct:     gbq.withProduct.Clone(),
		withTransaction: gbq.withTransaction.Clone(),
		// clone intermediate query.
		sql:  gbq.sql.Clone(),
		path: gbq.path,
	}
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (gbq *GroupBuyQuery) WithProduct(opts ...func(*ProductQuery)) *GroupBuyQuery {
	query := (&ProductClient{config: gbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gbq.withProduct = query
	return gbq
}

// WithTransaction tells the query-builder to eager-load the nodes that are connected to
// the "transaction" edge. The optional arguments are used to configure the query builder of the edge.
func (gbq *GroupBuyQuery) WithTransaction(opts ...func(*TransactionQuery)) *GroupBuyQuery {
	query := (&TransactionClient{config: gbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gbq.withTransaction = query
	return gbq
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
//	client.GroupBuy.Query().
//		GroupBy(groupbuy.FieldDateCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gbq *GroupBuyQuery) GroupBy(field string, fields ...string) *GroupBuyGroupBy {
	gbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupBuyGroupBy{build: gbq}
	grbuild.flds = &gbq.ctx.Fields
	grbuild.label = groupbuy.Label
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
//	client.GroupBuy.Query().
//		Select(groupbuy.FieldDateCreated).
//		Scan(ctx, &v)
func (gbq *GroupBuyQuery) Select(fields ...string) *GroupBuySelect {
	gbq.ctx.Fields = append(gbq.ctx.Fields, fields...)
	sbuild := &GroupBuySelect{GroupBuyQuery: gbq}
	sbuild.label = groupbuy.Label
	sbuild.flds, sbuild.scan = &gbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupBuySelect configured with the given aggregations.
func (gbq *GroupBuyQuery) Aggregate(fns ...AggregateFunc) *GroupBuySelect {
	return gbq.Select().Aggregate(fns...)
}

func (gbq *GroupBuyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gbq); err != nil {
				return err
			}
		}
	}
	for _, f := range gbq.ctx.Fields {
		if !groupbuy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gbq.path != nil {
		prev, err := gbq.path(ctx)
		if err != nil {
			return err
		}
		gbq.sql = prev
	}
	return nil
}

func (gbq *GroupBuyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupBuy, error) {
	var (
		nodes       = []*GroupBuy{}
		withFKs     = gbq.withFKs
		_spec       = gbq.querySpec()
		loadedTypes = [2]bool{
			gbq.withProduct != nil,
			gbq.withTransaction != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupbuy.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupBuy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupBuy{config: gbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gbq.modifiers) > 0 {
		_spec.Modifiers = gbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gbq.withProduct; query != nil {
		if err := gbq.loadProduct(ctx, query, nodes,
			func(n *GroupBuy) { n.Edges.Product = []*Product{} },
			func(n *GroupBuy, e *Product) { n.Edges.Product = append(n.Edges.Product, e) }); err != nil {
			return nil, err
		}
	}
	if query := gbq.withTransaction; query != nil {
		if err := gbq.loadTransaction(ctx, query, nodes,
			func(n *GroupBuy) { n.Edges.Transaction = []*Transaction{} },
			func(n *GroupBuy, e *Transaction) { n.Edges.Transaction = append(n.Edges.Transaction, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range gbq.withNamedProduct {
		if err := gbq.loadProduct(ctx, query, nodes,
			func(n *GroupBuy) { n.appendNamedProduct(name) },
			func(n *GroupBuy, e *Product) { n.appendNamedProduct(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range gbq.withNamedTransaction {
		if err := gbq.loadTransaction(ctx, query, nodes,
			func(n *GroupBuy) { n.appendNamedTransaction(name) },
			func(n *GroupBuy, e *Transaction) { n.appendNamedTransaction(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range gbq.loadTotal {
		if err := gbq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gbq *GroupBuyQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*GroupBuy, init func(*GroupBuy), assign func(*GroupBuy, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GroupBuy)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(groupbuy.ProductColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.group_buy_product
		if fk == nil {
			return fmt.Errorf(`foreign-key "group_buy_product" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_buy_product" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gbq *GroupBuyQuery) loadTransaction(ctx context.Context, query *TransactionQuery, nodes []*GroupBuy, init func(*GroupBuy), assign func(*GroupBuy, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GroupBuy)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(groupbuy.TransactionColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.group_buy_transaction
		if fk == nil {
			return fmt.Errorf(`foreign-key "group_buy_transaction" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_buy_transaction" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (gbq *GroupBuyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gbq.querySpec()
	if len(gbq.modifiers) > 0 {
		_spec.Modifiers = gbq.modifiers
	}
	_spec.Node.Columns = gbq.ctx.Fields
	if len(gbq.ctx.Fields) > 0 {
		_spec.Unique = gbq.ctx.Unique != nil && *gbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gbq.driver, _spec)
}

func (gbq *GroupBuyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupbuy.Table, groupbuy.Columns, sqlgraph.NewFieldSpec(groupbuy.FieldID, field.TypeInt))
	_spec.From = gbq.sql
	if unique := gbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gbq.path != nil {
		_spec.Unique = true
	}
	if fields := gbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupbuy.FieldID)
		for i := range fields {
			if fields[i] != groupbuy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gbq *GroupBuyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gbq.driver.Dialect())
	t1 := builder.Table(groupbuy.Table)
	columns := gbq.ctx.Fields
	if len(columns) == 0 {
		columns = groupbuy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gbq.sql != nil {
		selector = gbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gbq.ctx.Unique != nil && *gbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gbq.predicates {
		p(selector)
	}
	for _, p := range gbq.order {
		p(selector)
	}
	if offset := gbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProduct tells the query-builder to eager-load the nodes that are connected to the "product"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (gbq *GroupBuyQuery) WithNamedProduct(name string, opts ...func(*ProductQuery)) *GroupBuyQuery {
	query := (&ProductClient{config: gbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if gbq.withNamedProduct == nil {
		gbq.withNamedProduct = make(map[string]*ProductQuery)
	}
	gbq.withNamedProduct[name] = query
	return gbq
}

// WithNamedTransaction tells the query-builder to eager-load the nodes that are connected to the "transaction"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (gbq *GroupBuyQuery) WithNamedTransaction(name string, opts ...func(*TransactionQuery)) *GroupBuyQuery {
	query := (&TransactionClient{config: gbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if gbq.withNamedTransaction == nil {
		gbq.withNamedTransaction = make(map[string]*TransactionQuery)
	}
	gbq.withNamedTransaction[name] = query
	return gbq
}

// GroupBuyGroupBy is the group-by builder for GroupBuy entities.
type GroupBuyGroupBy struct {
	selector
	build *GroupBuyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gbgb *GroupBuyGroupBy) Aggregate(fns ...AggregateFunc) *GroupBuyGroupBy {
	gbgb.fns = append(gbgb.fns, fns...)
	return gbgb
}

// Scan applies the selector query and scans the result into the given value.
func (gbgb *GroupBuyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gbgb.build.ctx, "GroupBy")
	if err := gbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupBuyQuery, *GroupBuyGroupBy](ctx, gbgb.build, gbgb, gbgb.build.inters, v)
}

func (gbgb *GroupBuyGroupBy) sqlScan(ctx context.Context, root *GroupBuyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gbgb.fns))
	for _, fn := range gbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gbgb.flds)+len(gbgb.fns))
		for _, f := range *gbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupBuySelect is the builder for selecting fields of GroupBuy entities.
type GroupBuySelect struct {
	*GroupBuyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gbs *GroupBuySelect) Aggregate(fns ...AggregateFunc) *GroupBuySelect {
	gbs.fns = append(gbs.fns, fns...)
	return gbs
}

// Scan applies the selector query and scans the result into the given value.
func (gbs *GroupBuySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gbs.ctx, "Select")
	if err := gbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupBuyQuery, *GroupBuySelect](ctx, gbs.GroupBuyQuery, gbs, gbs.inters, v)
}

func (gbs *GroupBuySelect) sqlScan(ctx context.Context, root *GroupBuyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gbs.fns))
	for _, fn := range gbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
