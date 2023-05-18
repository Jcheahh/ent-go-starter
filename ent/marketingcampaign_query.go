// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/marketingcampaign"
	"entdemo/ent/predicate"
	"entdemo/ent/product"
	"entdemo/ent/rewardtype"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MarketingCampaignQuery is the builder for querying MarketingCampaign entities.
type MarketingCampaignQuery struct {
	config
	ctx                     *QueryContext
	order                   []marketingcampaign.Order
	inters                  []Interceptor
	predicates              []predicate.MarketingCampaign
	withProduct             *ProductQuery
	withConsumerReward      *RewardTypeQuery
	withFKs                 bool
	modifiers               []func(*sql.Selector)
	loadTotal               []func(context.Context, []*MarketingCampaign) error
	withNamedProduct        map[string]*ProductQuery
	withNamedConsumerReward map[string]*RewardTypeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MarketingCampaignQuery builder.
func (mcq *MarketingCampaignQuery) Where(ps ...predicate.MarketingCampaign) *MarketingCampaignQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit the number of records to be returned by this query.
func (mcq *MarketingCampaignQuery) Limit(limit int) *MarketingCampaignQuery {
	mcq.ctx.Limit = &limit
	return mcq
}

// Offset to start from.
func (mcq *MarketingCampaignQuery) Offset(offset int) *MarketingCampaignQuery {
	mcq.ctx.Offset = &offset
	return mcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mcq *MarketingCampaignQuery) Unique(unique bool) *MarketingCampaignQuery {
	mcq.ctx.Unique = &unique
	return mcq
}

// Order specifies how the records should be ordered.
func (mcq *MarketingCampaignQuery) Order(o ...marketingcampaign.Order) *MarketingCampaignQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// QueryProduct chains the current query on the "product" edge.
func (mcq *MarketingCampaignQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(marketingcampaign.Table, marketingcampaign.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, marketingcampaign.ProductTable, marketingcampaign.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConsumerReward chains the current query on the "consumerReward" edge.
func (mcq *MarketingCampaignQuery) QueryConsumerReward() *RewardTypeQuery {
	query := (&RewardTypeClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(marketingcampaign.Table, marketingcampaign.FieldID, selector),
			sqlgraph.To(rewardtype.Table, rewardtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, marketingcampaign.ConsumerRewardTable, marketingcampaign.ConsumerRewardColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MarketingCampaign entity from the query.
// Returns a *NotFoundError when no MarketingCampaign was found.
func (mcq *MarketingCampaignQuery) First(ctx context.Context) (*MarketingCampaign, error) {
	nodes, err := mcq.Limit(1).All(setContextOp(ctx, mcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{marketingcampaign.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) FirstX(ctx context.Context) *MarketingCampaign {
	node, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MarketingCampaign ID from the query.
// Returns a *NotFoundError when no MarketingCampaign ID was found.
func (mcq *MarketingCampaignQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(1).IDs(setContextOp(ctx, mcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{marketingcampaign.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) FirstIDX(ctx context.Context) int {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MarketingCampaign entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MarketingCampaign entity is found.
// Returns a *NotFoundError when no MarketingCampaign entities are found.
func (mcq *MarketingCampaignQuery) Only(ctx context.Context) (*MarketingCampaign, error) {
	nodes, err := mcq.Limit(2).All(setContextOp(ctx, mcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{marketingcampaign.Label}
	default:
		return nil, &NotSingularError{marketingcampaign.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) OnlyX(ctx context.Context) *MarketingCampaign {
	node, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MarketingCampaign ID in the query.
// Returns a *NotSingularError when more than one MarketingCampaign ID is found.
// Returns a *NotFoundError when no entities are found.
func (mcq *MarketingCampaignQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(2).IDs(setContextOp(ctx, mcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{marketingcampaign.Label}
	default:
		err = &NotSingularError{marketingcampaign.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) OnlyIDX(ctx context.Context) int {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MarketingCampaigns.
func (mcq *MarketingCampaignQuery) All(ctx context.Context) ([]*MarketingCampaign, error) {
	ctx = setContextOp(ctx, mcq.ctx, "All")
	if err := mcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MarketingCampaign, *MarketingCampaignQuery]()
	return withInterceptors[[]*MarketingCampaign](ctx, mcq, qr, mcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) AllX(ctx context.Context) []*MarketingCampaign {
	nodes, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MarketingCampaign IDs.
func (mcq *MarketingCampaignQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mcq.ctx.Unique == nil && mcq.path != nil {
		mcq.Unique(true)
	}
	ctx = setContextOp(ctx, mcq.ctx, "IDs")
	if err = mcq.Select(marketingcampaign.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) IDsX(ctx context.Context) []int {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MarketingCampaignQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Count")
	if err := mcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mcq, querierCount[*MarketingCampaignQuery](), mcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MarketingCampaignQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Exist")
	switch _, err := mcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MarketingCampaignQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MarketingCampaignQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MarketingCampaignQuery) Clone() *MarketingCampaignQuery {
	if mcq == nil {
		return nil
	}
	return &MarketingCampaignQuery{
		config:             mcq.config,
		ctx:                mcq.ctx.Clone(),
		order:              append([]marketingcampaign.Order{}, mcq.order...),
		inters:             append([]Interceptor{}, mcq.inters...),
		predicates:         append([]predicate.MarketingCampaign{}, mcq.predicates...),
		withProduct:        mcq.withProduct.Clone(),
		withConsumerReward: mcq.withConsumerReward.Clone(),
		// clone intermediate query.
		sql:  mcq.sql.Clone(),
		path: mcq.path,
	}
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MarketingCampaignQuery) WithProduct(opts ...func(*ProductQuery)) *MarketingCampaignQuery {
	query := (&ProductClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withProduct = query
	return mcq
}

// WithConsumerReward tells the query-builder to eager-load the nodes that are connected to
// the "consumerReward" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MarketingCampaignQuery) WithConsumerReward(opts ...func(*RewardTypeQuery)) *MarketingCampaignQuery {
	query := (&RewardTypeClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withConsumerReward = query
	return mcq
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
//	client.MarketingCampaign.Query().
//		GroupBy(marketingcampaign.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mcq *MarketingCampaignQuery) GroupBy(field string, fields ...string) *MarketingCampaignGroupBy {
	mcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MarketingCampaignGroupBy{build: mcq}
	grbuild.flds = &mcq.ctx.Fields
	grbuild.label = marketingcampaign.Label
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
//	client.MarketingCampaign.Query().
//		Select(marketingcampaign.FieldName).
//		Scan(ctx, &v)
func (mcq *MarketingCampaignQuery) Select(fields ...string) *MarketingCampaignSelect {
	mcq.ctx.Fields = append(mcq.ctx.Fields, fields...)
	sbuild := &MarketingCampaignSelect{MarketingCampaignQuery: mcq}
	sbuild.label = marketingcampaign.Label
	sbuild.flds, sbuild.scan = &mcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MarketingCampaignSelect configured with the given aggregations.
func (mcq *MarketingCampaignQuery) Aggregate(fns ...AggregateFunc) *MarketingCampaignSelect {
	return mcq.Select().Aggregate(fns...)
}

func (mcq *MarketingCampaignQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mcq); err != nil {
				return err
			}
		}
	}
	for _, f := range mcq.ctx.Fields {
		if !marketingcampaign.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mcq.path != nil {
		prev, err := mcq.path(ctx)
		if err != nil {
			return err
		}
		mcq.sql = prev
	}
	return nil
}

func (mcq *MarketingCampaignQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MarketingCampaign, error) {
	var (
		nodes       = []*MarketingCampaign{}
		withFKs     = mcq.withFKs
		_spec       = mcq.querySpec()
		loadedTypes = [2]bool{
			mcq.withProduct != nil,
			mcq.withConsumerReward != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, marketingcampaign.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MarketingCampaign).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MarketingCampaign{config: mcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mcq.modifiers) > 0 {
		_spec.Modifiers = mcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mcq.withProduct; query != nil {
		if err := mcq.loadProduct(ctx, query, nodes,
			func(n *MarketingCampaign) { n.Edges.Product = []*Product{} },
			func(n *MarketingCampaign, e *Product) { n.Edges.Product = append(n.Edges.Product, e) }); err != nil {
			return nil, err
		}
	}
	if query := mcq.withConsumerReward; query != nil {
		if err := mcq.loadConsumerReward(ctx, query, nodes,
			func(n *MarketingCampaign) { n.Edges.ConsumerReward = []*RewardType{} },
			func(n *MarketingCampaign, e *RewardType) { n.Edges.ConsumerReward = append(n.Edges.ConsumerReward, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mcq.withNamedProduct {
		if err := mcq.loadProduct(ctx, query, nodes,
			func(n *MarketingCampaign) { n.appendNamedProduct(name) },
			func(n *MarketingCampaign, e *Product) { n.appendNamedProduct(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mcq.withNamedConsumerReward {
		if err := mcq.loadConsumerReward(ctx, query, nodes,
			func(n *MarketingCampaign) { n.appendNamedConsumerReward(name) },
			func(n *MarketingCampaign, e *RewardType) { n.appendNamedConsumerReward(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range mcq.loadTotal {
		if err := mcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mcq *MarketingCampaignQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*MarketingCampaign, init func(*MarketingCampaign), assign func(*MarketingCampaign, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MarketingCampaign)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(marketingcampaign.ProductColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.marketing_campaign_product
		if fk == nil {
			return fmt.Errorf(`foreign-key "marketing_campaign_product" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "marketing_campaign_product" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mcq *MarketingCampaignQuery) loadConsumerReward(ctx context.Context, query *RewardTypeQuery, nodes []*MarketingCampaign, init func(*MarketingCampaign), assign func(*MarketingCampaign, *RewardType)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MarketingCampaign)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RewardType(func(s *sql.Selector) {
		s.Where(sql.InValues(marketingcampaign.ConsumerRewardColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.marketing_campaign_consumer_reward
		if fk == nil {
			return fmt.Errorf(`foreign-key "marketing_campaign_consumer_reward" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "marketing_campaign_consumer_reward" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mcq *MarketingCampaignQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcq.querySpec()
	if len(mcq.modifiers) > 0 {
		_spec.Modifiers = mcq.modifiers
	}
	_spec.Node.Columns = mcq.ctx.Fields
	if len(mcq.ctx.Fields) > 0 {
		_spec.Unique = mcq.ctx.Unique != nil && *mcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mcq.driver, _spec)
}

func (mcq *MarketingCampaignQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(marketingcampaign.Table, marketingcampaign.Columns, sqlgraph.NewFieldSpec(marketingcampaign.FieldID, field.TypeInt))
	_spec.From = mcq.sql
	if unique := mcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mcq.path != nil {
		_spec.Unique = true
	}
	if fields := mcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, marketingcampaign.FieldID)
		for i := range fields {
			if fields[i] != marketingcampaign.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcq *MarketingCampaignQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(marketingcampaign.Table)
	columns := mcq.ctx.Fields
	if len(columns) == 0 {
		columns = marketingcampaign.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mcq.ctx.Unique != nil && *mcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProduct tells the query-builder to eager-load the nodes that are connected to the "product"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mcq *MarketingCampaignQuery) WithNamedProduct(name string, opts ...func(*ProductQuery)) *MarketingCampaignQuery {
	query := (&ProductClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mcq.withNamedProduct == nil {
		mcq.withNamedProduct = make(map[string]*ProductQuery)
	}
	mcq.withNamedProduct[name] = query
	return mcq
}

// WithNamedConsumerReward tells the query-builder to eager-load the nodes that are connected to the "consumerReward"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mcq *MarketingCampaignQuery) WithNamedConsumerReward(name string, opts ...func(*RewardTypeQuery)) *MarketingCampaignQuery {
	query := (&RewardTypeClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mcq.withNamedConsumerReward == nil {
		mcq.withNamedConsumerReward = make(map[string]*RewardTypeQuery)
	}
	mcq.withNamedConsumerReward[name] = query
	return mcq
}

// MarketingCampaignGroupBy is the group-by builder for MarketingCampaign entities.
type MarketingCampaignGroupBy struct {
	selector
	build *MarketingCampaignQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MarketingCampaignGroupBy) Aggregate(fns ...AggregateFunc) *MarketingCampaignGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the selector query and scans the result into the given value.
func (mcgb *MarketingCampaignGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcgb.build.ctx, "GroupBy")
	if err := mcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MarketingCampaignQuery, *MarketingCampaignGroupBy](ctx, mcgb.build, mcgb, mcgb.build.inters, v)
}

func (mcgb *MarketingCampaignGroupBy) sqlScan(ctx context.Context, root *MarketingCampaignQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mcgb.fns))
	for _, fn := range mcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mcgb.flds)+len(mcgb.fns))
		for _, f := range *mcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MarketingCampaignSelect is the builder for selecting fields of MarketingCampaign entities.
type MarketingCampaignSelect struct {
	*MarketingCampaignQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mcs *MarketingCampaignSelect) Aggregate(fns ...AggregateFunc) *MarketingCampaignSelect {
	mcs.fns = append(mcs.fns, fns...)
	return mcs
}

// Scan applies the selector query and scans the result into the given value.
func (mcs *MarketingCampaignSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcs.ctx, "Select")
	if err := mcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MarketingCampaignQuery, *MarketingCampaignSelect](ctx, mcs.MarketingCampaignQuery, mcs, mcs.inters, v)
}

func (mcs *MarketingCampaignSelect) sqlScan(ctx context.Context, root *MarketingCampaignQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mcs.fns))
	for _, fn := range mcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
