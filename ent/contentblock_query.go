// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/contentblock"
	"entdemo/ent/image"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContentBlockQuery is the builder for querying ContentBlock entities.
type ContentBlockQuery struct {
	config
	ctx            *QueryContext
	order          []contentblock.OrderOption
	inters         []Interceptor
	predicates     []predicate.ContentBlock
	withImage      *ImageQuery
	withFKs        bool
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*ContentBlock) error
	withNamedImage map[string]*ImageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContentBlockQuery builder.
func (cbq *ContentBlockQuery) Where(ps ...predicate.ContentBlock) *ContentBlockQuery {
	cbq.predicates = append(cbq.predicates, ps...)
	return cbq
}

// Limit the number of records to be returned by this query.
func (cbq *ContentBlockQuery) Limit(limit int) *ContentBlockQuery {
	cbq.ctx.Limit = &limit
	return cbq
}

// Offset to start from.
func (cbq *ContentBlockQuery) Offset(offset int) *ContentBlockQuery {
	cbq.ctx.Offset = &offset
	return cbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cbq *ContentBlockQuery) Unique(unique bool) *ContentBlockQuery {
	cbq.ctx.Unique = &unique
	return cbq
}

// Order specifies how the records should be ordered.
func (cbq *ContentBlockQuery) Order(o ...contentblock.OrderOption) *ContentBlockQuery {
	cbq.order = append(cbq.order, o...)
	return cbq
}

// QueryImage chains the current query on the "image" edge.
func (cbq *ContentBlockQuery) QueryImage() *ImageQuery {
	query := (&ImageClient{config: cbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contentblock.Table, contentblock.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contentblock.ImageTable, contentblock.ImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(cbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ContentBlock entity from the query.
// Returns a *NotFoundError when no ContentBlock was found.
func (cbq *ContentBlockQuery) First(ctx context.Context) (*ContentBlock, error) {
	nodes, err := cbq.Limit(1).All(setContextOp(ctx, cbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contentblock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cbq *ContentBlockQuery) FirstX(ctx context.Context) *ContentBlock {
	node, err := cbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContentBlock ID from the query.
// Returns a *NotFoundError when no ContentBlock ID was found.
func (cbq *ContentBlockQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cbq.Limit(1).IDs(setContextOp(ctx, cbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contentblock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cbq *ContentBlockQuery) FirstIDX(ctx context.Context) int {
	id, err := cbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContentBlock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContentBlock entity is found.
// Returns a *NotFoundError when no ContentBlock entities are found.
func (cbq *ContentBlockQuery) Only(ctx context.Context) (*ContentBlock, error) {
	nodes, err := cbq.Limit(2).All(setContextOp(ctx, cbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contentblock.Label}
	default:
		return nil, &NotSingularError{contentblock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cbq *ContentBlockQuery) OnlyX(ctx context.Context) *ContentBlock {
	node, err := cbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContentBlock ID in the query.
// Returns a *NotSingularError when more than one ContentBlock ID is found.
// Returns a *NotFoundError when no entities are found.
func (cbq *ContentBlockQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cbq.Limit(2).IDs(setContextOp(ctx, cbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contentblock.Label}
	default:
		err = &NotSingularError{contentblock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cbq *ContentBlockQuery) OnlyIDX(ctx context.Context) int {
	id, err := cbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContentBlocks.
func (cbq *ContentBlockQuery) All(ctx context.Context) ([]*ContentBlock, error) {
	ctx = setContextOp(ctx, cbq.ctx, "All")
	if err := cbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContentBlock, *ContentBlockQuery]()
	return withInterceptors[[]*ContentBlock](ctx, cbq, qr, cbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cbq *ContentBlockQuery) AllX(ctx context.Context) []*ContentBlock {
	nodes, err := cbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContentBlock IDs.
func (cbq *ContentBlockQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cbq.ctx.Unique == nil && cbq.path != nil {
		cbq.Unique(true)
	}
	ctx = setContextOp(ctx, cbq.ctx, "IDs")
	if err = cbq.Select(contentblock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cbq *ContentBlockQuery) IDsX(ctx context.Context) []int {
	ids, err := cbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cbq *ContentBlockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cbq.ctx, "Count")
	if err := cbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cbq, querierCount[*ContentBlockQuery](), cbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cbq *ContentBlockQuery) CountX(ctx context.Context) int {
	count, err := cbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cbq *ContentBlockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cbq.ctx, "Exist")
	switch _, err := cbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cbq *ContentBlockQuery) ExistX(ctx context.Context) bool {
	exist, err := cbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContentBlockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cbq *ContentBlockQuery) Clone() *ContentBlockQuery {
	if cbq == nil {
		return nil
	}
	return &ContentBlockQuery{
		config:     cbq.config,
		ctx:        cbq.ctx.Clone(),
		order:      append([]contentblock.OrderOption{}, cbq.order...),
		inters:     append([]Interceptor{}, cbq.inters...),
		predicates: append([]predicate.ContentBlock{}, cbq.predicates...),
		withImage:  cbq.withImage.Clone(),
		// clone intermediate query.
		sql:  cbq.sql.Clone(),
		path: cbq.path,
	}
}

// WithImage tells the query-builder to eager-load the nodes that are connected to
// the "image" edge. The optional arguments are used to configure the query builder of the edge.
func (cbq *ContentBlockQuery) WithImage(opts ...func(*ImageQuery)) *ContentBlockQuery {
	query := (&ImageClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cbq.withImage = query
	return cbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PrimaryMessage string `json:"primaryMessage,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ContentBlock.Query().
//		GroupBy(contentblock.FieldPrimaryMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cbq *ContentBlockQuery) GroupBy(field string, fields ...string) *ContentBlockGroupBy {
	cbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContentBlockGroupBy{build: cbq}
	grbuild.flds = &cbq.ctx.Fields
	grbuild.label = contentblock.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PrimaryMessage string `json:"primaryMessage,omitempty"`
//	}
//
//	client.ContentBlock.Query().
//		Select(contentblock.FieldPrimaryMessage).
//		Scan(ctx, &v)
func (cbq *ContentBlockQuery) Select(fields ...string) *ContentBlockSelect {
	cbq.ctx.Fields = append(cbq.ctx.Fields, fields...)
	sbuild := &ContentBlockSelect{ContentBlockQuery: cbq}
	sbuild.label = contentblock.Label
	sbuild.flds, sbuild.scan = &cbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContentBlockSelect configured with the given aggregations.
func (cbq *ContentBlockQuery) Aggregate(fns ...AggregateFunc) *ContentBlockSelect {
	return cbq.Select().Aggregate(fns...)
}

func (cbq *ContentBlockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cbq); err != nil {
				return err
			}
		}
	}
	for _, f := range cbq.ctx.Fields {
		if !contentblock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cbq.path != nil {
		prev, err := cbq.path(ctx)
		if err != nil {
			return err
		}
		cbq.sql = prev
	}
	return nil
}

func (cbq *ContentBlockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContentBlock, error) {
	var (
		nodes       = []*ContentBlock{}
		withFKs     = cbq.withFKs
		_spec       = cbq.querySpec()
		loadedTypes = [1]bool{
			cbq.withImage != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, contentblock.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContentBlock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContentBlock{config: cbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cbq.withImage; query != nil {
		if err := cbq.loadImage(ctx, query, nodes,
			func(n *ContentBlock) { n.Edges.Image = []*Image{} },
			func(n *ContentBlock, e *Image) { n.Edges.Image = append(n.Edges.Image, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cbq.withNamedImage {
		if err := cbq.loadImage(ctx, query, nodes,
			func(n *ContentBlock) { n.appendNamedImage(name) },
			func(n *ContentBlock, e *Image) { n.appendNamedImage(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cbq.loadTotal {
		if err := cbq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cbq *ContentBlockQuery) loadImage(ctx context.Context, query *ImageQuery, nodes []*ContentBlock, init func(*ContentBlock), assign func(*ContentBlock, *Image)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ContentBlock)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Image(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(contentblock.ImageColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.content_block_image
		if fk == nil {
			return fmt.Errorf(`foreign-key "content_block_image" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "content_block_image" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cbq *ContentBlockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cbq.querySpec()
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	_spec.Node.Columns = cbq.ctx.Fields
	if len(cbq.ctx.Fields) > 0 {
		_spec.Unique = cbq.ctx.Unique != nil && *cbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cbq.driver, _spec)
}

func (cbq *ContentBlockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contentblock.Table, contentblock.Columns, sqlgraph.NewFieldSpec(contentblock.FieldID, field.TypeInt))
	_spec.From = cbq.sql
	if unique := cbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cbq.path != nil {
		_spec.Unique = true
	}
	if fields := cbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contentblock.FieldID)
		for i := range fields {
			if fields[i] != contentblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cbq *ContentBlockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cbq.driver.Dialect())
	t1 := builder.Table(contentblock.Table)
	columns := cbq.ctx.Fields
	if len(columns) == 0 {
		columns = contentblock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cbq.sql != nil {
		selector = cbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cbq.ctx.Unique != nil && *cbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cbq.predicates {
		p(selector)
	}
	for _, p := range cbq.order {
		p(selector)
	}
	if offset := cbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedImage tells the query-builder to eager-load the nodes that are connected to the "image"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cbq *ContentBlockQuery) WithNamedImage(name string, opts ...func(*ImageQuery)) *ContentBlockQuery {
	query := (&ImageClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cbq.withNamedImage == nil {
		cbq.withNamedImage = make(map[string]*ImageQuery)
	}
	cbq.withNamedImage[name] = query
	return cbq
}

// ContentBlockGroupBy is the group-by builder for ContentBlock entities.
type ContentBlockGroupBy struct {
	selector
	build *ContentBlockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cbgb *ContentBlockGroupBy) Aggregate(fns ...AggregateFunc) *ContentBlockGroupBy {
	cbgb.fns = append(cbgb.fns, fns...)
	return cbgb
}

// Scan applies the selector query and scans the result into the given value.
func (cbgb *ContentBlockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbgb.build.ctx, "GroupBy")
	if err := cbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContentBlockQuery, *ContentBlockGroupBy](ctx, cbgb.build, cbgb, cbgb.build.inters, v)
}

func (cbgb *ContentBlockGroupBy) sqlScan(ctx context.Context, root *ContentBlockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cbgb.fns))
	for _, fn := range cbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cbgb.flds)+len(cbgb.fns))
		for _, f := range *cbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContentBlockSelect is the builder for selecting fields of ContentBlock entities.
type ContentBlockSelect struct {
	*ContentBlockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cbs *ContentBlockSelect) Aggregate(fns ...AggregateFunc) *ContentBlockSelect {
	cbs.fns = append(cbs.fns, fns...)
	return cbs
}

// Scan applies the selector query and scans the result into the given value.
func (cbs *ContentBlockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbs.ctx, "Select")
	if err := cbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContentBlockQuery, *ContentBlockSelect](ctx, cbs.ContentBlockQuery, cbs, cbs.inters, v)
}

func (cbs *ContentBlockSelect) sqlScan(ctx context.Context, root *ContentBlockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cbs.fns))
	for _, fn := range cbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
