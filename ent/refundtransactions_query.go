// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/predicate"
	"entdemo/ent/refundtransactions"
	"entdemo/ent/transaction"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefundTransactionsQuery is the builder for querying RefundTransactions entities.
type RefundTransactionsQuery struct {
	config
	ctx                  *QueryContext
	order                []refundtransactions.OrderOption
	inters               []Interceptor
	predicates           []predicate.RefundTransactions
	withTransaction      *TransactionQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*RefundTransactions) error
	withNamedTransaction map[string]*TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RefundTransactionsQuery builder.
func (rtq *RefundTransactionsQuery) Where(ps ...predicate.RefundTransactions) *RefundTransactionsQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit the number of records to be returned by this query.
func (rtq *RefundTransactionsQuery) Limit(limit int) *RefundTransactionsQuery {
	rtq.ctx.Limit = &limit
	return rtq
}

// Offset to start from.
func (rtq *RefundTransactionsQuery) Offset(offset int) *RefundTransactionsQuery {
	rtq.ctx.Offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *RefundTransactionsQuery) Unique(unique bool) *RefundTransactionsQuery {
	rtq.ctx.Unique = &unique
	return rtq
}

// Order specifies how the records should be ordered.
func (rtq *RefundTransactionsQuery) Order(o ...refundtransactions.OrderOption) *RefundTransactionsQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryTransaction chains the current query on the "transaction" edge.
func (rtq *RefundTransactionsQuery) QueryTransaction() *TransactionQuery {
	query := (&TransactionClient{config: rtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(refundtransactions.Table, refundtransactions.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, refundtransactions.TransactionTable, refundtransactions.TransactionColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RefundTransactions entity from the query.
// Returns a *NotFoundError when no RefundTransactions was found.
func (rtq *RefundTransactionsQuery) First(ctx context.Context) (*RefundTransactions, error) {
	nodes, err := rtq.Limit(1).All(setContextOp(ctx, rtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{refundtransactions.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) FirstX(ctx context.Context) *RefundTransactions {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RefundTransactions ID from the query.
// Returns a *NotFoundError when no RefundTransactions ID was found.
func (rtq *RefundTransactionsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(setContextOp(ctx, rtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{refundtransactions.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) FirstIDX(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RefundTransactions entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RefundTransactions entity is found.
// Returns a *NotFoundError when no RefundTransactions entities are found.
func (rtq *RefundTransactionsQuery) Only(ctx context.Context) (*RefundTransactions, error) {
	nodes, err := rtq.Limit(2).All(setContextOp(ctx, rtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{refundtransactions.Label}
	default:
		return nil, &NotSingularError{refundtransactions.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) OnlyX(ctx context.Context) *RefundTransactions {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RefundTransactions ID in the query.
// Returns a *NotSingularError when more than one RefundTransactions ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *RefundTransactionsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(setContextOp(ctx, rtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{refundtransactions.Label}
	default:
		err = &NotSingularError{refundtransactions.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RefundTransactionsSlice.
func (rtq *RefundTransactionsQuery) All(ctx context.Context) ([]*RefundTransactions, error) {
	ctx = setContextOp(ctx, rtq.ctx, "All")
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RefundTransactions, *RefundTransactionsQuery]()
	return withInterceptors[[]*RefundTransactions](ctx, rtq, qr, rtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) AllX(ctx context.Context) []*RefundTransactions {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RefundTransactions IDs.
func (rtq *RefundTransactionsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rtq.ctx.Unique == nil && rtq.path != nil {
		rtq.Unique(true)
	}
	ctx = setContextOp(ctx, rtq.ctx, "IDs")
	if err = rtq.Select(refundtransactions.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *RefundTransactionsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Count")
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rtq, querierCount[*RefundTransactionsQuery](), rtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *RefundTransactionsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Exist")
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *RefundTransactionsQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RefundTransactionsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *RefundTransactionsQuery) Clone() *RefundTransactionsQuery {
	if rtq == nil {
		return nil
	}
	return &RefundTransactionsQuery{
		config:          rtq.config,
		ctx:             rtq.ctx.Clone(),
		order:           append([]refundtransactions.OrderOption{}, rtq.order...),
		inters:          append([]Interceptor{}, rtq.inters...),
		predicates:      append([]predicate.RefundTransactions{}, rtq.predicates...),
		withTransaction: rtq.withTransaction.Clone(),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

// WithTransaction tells the query-builder to eager-load the nodes that are connected to
// the "transaction" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *RefundTransactionsQuery) WithTransaction(opts ...func(*TransactionQuery)) *RefundTransactionsQuery {
	query := (&TransactionClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rtq.withTransaction = query
	return rtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RefundAmount string `json:"refundAmount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RefundTransactions.Query().
//		GroupBy(refundtransactions.FieldRefundAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rtq *RefundTransactionsQuery) GroupBy(field string, fields ...string) *RefundTransactionsGroupBy {
	rtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RefundTransactionsGroupBy{build: rtq}
	grbuild.flds = &rtq.ctx.Fields
	grbuild.label = refundtransactions.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RefundAmount string `json:"refundAmount,omitempty"`
//	}
//
//	client.RefundTransactions.Query().
//		Select(refundtransactions.FieldRefundAmount).
//		Scan(ctx, &v)
func (rtq *RefundTransactionsQuery) Select(fields ...string) *RefundTransactionsSelect {
	rtq.ctx.Fields = append(rtq.ctx.Fields, fields...)
	sbuild := &RefundTransactionsSelect{RefundTransactionsQuery: rtq}
	sbuild.label = refundtransactions.Label
	sbuild.flds, sbuild.scan = &rtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RefundTransactionsSelect configured with the given aggregations.
func (rtq *RefundTransactionsQuery) Aggregate(fns ...AggregateFunc) *RefundTransactionsSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *RefundTransactionsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rtq.ctx.Fields {
		if !refundtransactions.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *RefundTransactionsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RefundTransactions, error) {
	var (
		nodes       = []*RefundTransactions{}
		_spec       = rtq.querySpec()
		loadedTypes = [1]bool{
			rtq.withTransaction != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RefundTransactions).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RefundTransactions{config: rtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rtq.modifiers) > 0 {
		_spec.Modifiers = rtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rtq.withTransaction; query != nil {
		if err := rtq.loadTransaction(ctx, query, nodes,
			func(n *RefundTransactions) { n.Edges.Transaction = []*Transaction{} },
			func(n *RefundTransactions, e *Transaction) { n.Edges.Transaction = append(n.Edges.Transaction, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rtq.withNamedTransaction {
		if err := rtq.loadTransaction(ctx, query, nodes,
			func(n *RefundTransactions) { n.appendNamedTransaction(name) },
			func(n *RefundTransactions, e *Transaction) { n.appendNamedTransaction(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rtq.loadTotal {
		if err := rtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rtq *RefundTransactionsQuery) loadTransaction(ctx context.Context, query *TransactionQuery, nodes []*RefundTransactions, init func(*RefundTransactions), assign func(*RefundTransactions, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*RefundTransactions)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(refundtransactions.TransactionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.refund_transactions_transaction
		if fk == nil {
			return fmt.Errorf(`foreign-key "refund_transactions_transaction" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "refund_transactions_transaction" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rtq *RefundTransactionsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	if len(rtq.modifiers) > 0 {
		_spec.Modifiers = rtq.modifiers
	}
	_spec.Node.Columns = rtq.ctx.Fields
	if len(rtq.ctx.Fields) > 0 {
		_spec.Unique = rtq.ctx.Unique != nil && *rtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *RefundTransactionsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(refundtransactions.Table, refundtransactions.Columns, sqlgraph.NewFieldSpec(refundtransactions.FieldID, field.TypeInt))
	_spec.From = rtq.sql
	if unique := rtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rtq.path != nil {
		_spec.Unique = true
	}
	if fields := rtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refundtransactions.FieldID)
		for i := range fields {
			if fields[i] != refundtransactions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *RefundTransactionsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(refundtransactions.Table)
	columns := rtq.ctx.Fields
	if len(columns) == 0 {
		columns = refundtransactions.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.ctx.Unique != nil && *rtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTransaction tells the query-builder to eager-load the nodes that are connected to the "transaction"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rtq *RefundTransactionsQuery) WithNamedTransaction(name string, opts ...func(*TransactionQuery)) *RefundTransactionsQuery {
	query := (&TransactionClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rtq.withNamedTransaction == nil {
		rtq.withNamedTransaction = make(map[string]*TransactionQuery)
	}
	rtq.withNamedTransaction[name] = query
	return rtq
}

// RefundTransactionsGroupBy is the group-by builder for RefundTransactions entities.
type RefundTransactionsGroupBy struct {
	selector
	build *RefundTransactionsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *RefundTransactionsGroupBy) Aggregate(fns ...AggregateFunc) *RefundTransactionsGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rtgb *RefundTransactionsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rtgb.build.ctx, "GroupBy")
	if err := rtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RefundTransactionsQuery, *RefundTransactionsGroupBy](ctx, rtgb.build, rtgb, rtgb.build.inters, v)
}

func (rtgb *RefundTransactionsGroupBy) sqlScan(ctx context.Context, root *RefundTransactionsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rtgb.flds)+len(rtgb.fns))
		for _, f := range *rtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RefundTransactionsSelect is the builder for selecting fields of RefundTransactions entities.
type RefundTransactionsSelect struct {
	*RefundTransactionsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *RefundTransactionsSelect) Aggregate(fns ...AggregateFunc) *RefundTransactionsSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *RefundTransactionsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rts.ctx, "Select")
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RefundTransactionsQuery, *RefundTransactionsSelect](ctx, rts.RefundTransactionsQuery, rts, rts.inters, v)
}

func (rts *RefundTransactionsSelect) sqlScan(ctx context.Context, root *RefundTransactionsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
