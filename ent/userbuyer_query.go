// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/linkvisit"
	"entdemo/ent/predicate"
	"entdemo/ent/review"
	"entdemo/ent/transaction"
	"entdemo/ent/user"
	"entdemo/ent/userbuyer"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBuyerQuery is the builder for querying UserBuyer entities.
type UserBuyerQuery struct {
	config
	ctx                   *QueryContext
	order                 []userbuyer.OrderOption
	inters                []Interceptor
	predicates            []predicate.UserBuyer
	withUserProfile       *UserQuery
	withReviews           *ReviewQuery
	withTransactions      *TransactionQuery
	withLinksClicked      *LinkVisitQuery
	withFKs               bool
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*UserBuyer) error
	withNamedUserProfile  map[string]*UserQuery
	withNamedReviews      map[string]*ReviewQuery
	withNamedTransactions map[string]*TransactionQuery
	withNamedLinksClicked map[string]*LinkVisitQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserBuyerQuery builder.
func (ubq *UserBuyerQuery) Where(ps ...predicate.UserBuyer) *UserBuyerQuery {
	ubq.predicates = append(ubq.predicates, ps...)
	return ubq
}

// Limit the number of records to be returned by this query.
func (ubq *UserBuyerQuery) Limit(limit int) *UserBuyerQuery {
	ubq.ctx.Limit = &limit
	return ubq
}

// Offset to start from.
func (ubq *UserBuyerQuery) Offset(offset int) *UserBuyerQuery {
	ubq.ctx.Offset = &offset
	return ubq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ubq *UserBuyerQuery) Unique(unique bool) *UserBuyerQuery {
	ubq.ctx.Unique = &unique
	return ubq
}

// Order specifies how the records should be ordered.
func (ubq *UserBuyerQuery) Order(o ...userbuyer.OrderOption) *UserBuyerQuery {
	ubq.order = append(ubq.order, o...)
	return ubq
}

// QueryUserProfile chains the current query on the "userProfile" edge.
func (ubq *UserBuyerQuery) QueryUserProfile() *UserQuery {
	query := (&UserClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbuyer.Table, userbuyer.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userbuyer.UserProfileTable, userbuyer.UserProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReviews chains the current query on the "reviews" edge.
func (ubq *UserBuyerQuery) QueryReviews() *ReviewQuery {
	query := (&ReviewClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbuyer.Table, userbuyer.FieldID, selector),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userbuyer.ReviewsTable, userbuyer.ReviewsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransactions chains the current query on the "transactions" edge.
func (ubq *UserBuyerQuery) QueryTransactions() *TransactionQuery {
	query := (&TransactionClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbuyer.Table, userbuyer.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userbuyer.TransactionsTable, userbuyer.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinksClicked chains the current query on the "linksClicked" edge.
func (ubq *UserBuyerQuery) QueryLinksClicked() *LinkVisitQuery {
	query := (&LinkVisitClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbuyer.Table, userbuyer.FieldID, selector),
			sqlgraph.To(linkvisit.Table, linkvisit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userbuyer.LinksClickedTable, userbuyer.LinksClickedColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserBuyer entity from the query.
// Returns a *NotFoundError when no UserBuyer was found.
func (ubq *UserBuyerQuery) First(ctx context.Context) (*UserBuyer, error) {
	nodes, err := ubq.Limit(1).All(setContextOp(ctx, ubq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userbuyer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ubq *UserBuyerQuery) FirstX(ctx context.Context) *UserBuyer {
	node, err := ubq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserBuyer ID from the query.
// Returns a *NotFoundError when no UserBuyer ID was found.
func (ubq *UserBuyerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ubq.Limit(1).IDs(setContextOp(ctx, ubq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userbuyer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ubq *UserBuyerQuery) FirstIDX(ctx context.Context) int {
	id, err := ubq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserBuyer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserBuyer entity is found.
// Returns a *NotFoundError when no UserBuyer entities are found.
func (ubq *UserBuyerQuery) Only(ctx context.Context) (*UserBuyer, error) {
	nodes, err := ubq.Limit(2).All(setContextOp(ctx, ubq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userbuyer.Label}
	default:
		return nil, &NotSingularError{userbuyer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ubq *UserBuyerQuery) OnlyX(ctx context.Context) *UserBuyer {
	node, err := ubq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserBuyer ID in the query.
// Returns a *NotSingularError when more than one UserBuyer ID is found.
// Returns a *NotFoundError when no entities are found.
func (ubq *UserBuyerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ubq.Limit(2).IDs(setContextOp(ctx, ubq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userbuyer.Label}
	default:
		err = &NotSingularError{userbuyer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ubq *UserBuyerQuery) OnlyIDX(ctx context.Context) int {
	id, err := ubq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserBuyers.
func (ubq *UserBuyerQuery) All(ctx context.Context) ([]*UserBuyer, error) {
	ctx = setContextOp(ctx, ubq.ctx, "All")
	if err := ubq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserBuyer, *UserBuyerQuery]()
	return withInterceptors[[]*UserBuyer](ctx, ubq, qr, ubq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ubq *UserBuyerQuery) AllX(ctx context.Context) []*UserBuyer {
	nodes, err := ubq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserBuyer IDs.
func (ubq *UserBuyerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ubq.ctx.Unique == nil && ubq.path != nil {
		ubq.Unique(true)
	}
	ctx = setContextOp(ctx, ubq.ctx, "IDs")
	if err = ubq.Select(userbuyer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ubq *UserBuyerQuery) IDsX(ctx context.Context) []int {
	ids, err := ubq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ubq *UserBuyerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ubq.ctx, "Count")
	if err := ubq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ubq, querierCount[*UserBuyerQuery](), ubq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ubq *UserBuyerQuery) CountX(ctx context.Context) int {
	count, err := ubq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ubq *UserBuyerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ubq.ctx, "Exist")
	switch _, err := ubq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ubq *UserBuyerQuery) ExistX(ctx context.Context) bool {
	exist, err := ubq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserBuyerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ubq *UserBuyerQuery) Clone() *UserBuyerQuery {
	if ubq == nil {
		return nil
	}
	return &UserBuyerQuery{
		config:           ubq.config,
		ctx:              ubq.ctx.Clone(),
		order:            append([]userbuyer.OrderOption{}, ubq.order...),
		inters:           append([]Interceptor{}, ubq.inters...),
		predicates:       append([]predicate.UserBuyer{}, ubq.predicates...),
		withUserProfile:  ubq.withUserProfile.Clone(),
		withReviews:      ubq.withReviews.Clone(),
		withTransactions: ubq.withTransactions.Clone(),
		withLinksClicked: ubq.withLinksClicked.Clone(),
		// clone intermediate query.
		sql:  ubq.sql.Clone(),
		path: ubq.path,
	}
}

// WithUserProfile tells the query-builder to eager-load the nodes that are connected to
// the "userProfile" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithUserProfile(opts ...func(*UserQuery)) *UserBuyerQuery {
	query := (&UserClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withUserProfile = query
	return ubq
}

// WithReviews tells the query-builder to eager-load the nodes that are connected to
// the "reviews" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithReviews(opts ...func(*ReviewQuery)) *UserBuyerQuery {
	query := (&ReviewClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withReviews = query
	return ubq
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithTransactions(opts ...func(*TransactionQuery)) *UserBuyerQuery {
	query := (&TransactionClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withTransactions = query
	return ubq
}

// WithLinksClicked tells the query-builder to eager-load the nodes that are connected to
// the "linksClicked" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithLinksClicked(opts ...func(*LinkVisitQuery)) *UserBuyerQuery {
	query := (&LinkVisitClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withLinksClicked = query
	return ubq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Placeholder int `json:"placeholder,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserBuyer.Query().
//		GroupBy(userbuyer.FieldPlaceholder).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ubq *UserBuyerQuery) GroupBy(field string, fields ...string) *UserBuyerGroupBy {
	ubq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserBuyerGroupBy{build: ubq}
	grbuild.flds = &ubq.ctx.Fields
	grbuild.label = userbuyer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Placeholder int `json:"placeholder,omitempty"`
//	}
//
//	client.UserBuyer.Query().
//		Select(userbuyer.FieldPlaceholder).
//		Scan(ctx, &v)
func (ubq *UserBuyerQuery) Select(fields ...string) *UserBuyerSelect {
	ubq.ctx.Fields = append(ubq.ctx.Fields, fields...)
	sbuild := &UserBuyerSelect{UserBuyerQuery: ubq}
	sbuild.label = userbuyer.Label
	sbuild.flds, sbuild.scan = &ubq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserBuyerSelect configured with the given aggregations.
func (ubq *UserBuyerQuery) Aggregate(fns ...AggregateFunc) *UserBuyerSelect {
	return ubq.Select().Aggregate(fns...)
}

func (ubq *UserBuyerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ubq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ubq); err != nil {
				return err
			}
		}
	}
	for _, f := range ubq.ctx.Fields {
		if !userbuyer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ubq.path != nil {
		prev, err := ubq.path(ctx)
		if err != nil {
			return err
		}
		ubq.sql = prev
	}
	return nil
}

func (ubq *UserBuyerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserBuyer, error) {
	var (
		nodes       = []*UserBuyer{}
		withFKs     = ubq.withFKs
		_spec       = ubq.querySpec()
		loadedTypes = [4]bool{
			ubq.withUserProfile != nil,
			ubq.withReviews != nil,
			ubq.withTransactions != nil,
			ubq.withLinksClicked != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userbuyer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserBuyer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserBuyer{config: ubq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ubq.modifiers) > 0 {
		_spec.Modifiers = ubq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ubq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ubq.withUserProfile; query != nil {
		if err := ubq.loadUserProfile(ctx, query, nodes,
			func(n *UserBuyer) { n.Edges.UserProfile = []*User{} },
			func(n *UserBuyer, e *User) { n.Edges.UserProfile = append(n.Edges.UserProfile, e) }); err != nil {
			return nil, err
		}
	}
	if query := ubq.withReviews; query != nil {
		if err := ubq.loadReviews(ctx, query, nodes,
			func(n *UserBuyer) { n.Edges.Reviews = []*Review{} },
			func(n *UserBuyer, e *Review) { n.Edges.Reviews = append(n.Edges.Reviews, e) }); err != nil {
			return nil, err
		}
	}
	if query := ubq.withTransactions; query != nil {
		if err := ubq.loadTransactions(ctx, query, nodes,
			func(n *UserBuyer) { n.Edges.Transactions = []*Transaction{} },
			func(n *UserBuyer, e *Transaction) { n.Edges.Transactions = append(n.Edges.Transactions, e) }); err != nil {
			return nil, err
		}
	}
	if query := ubq.withLinksClicked; query != nil {
		if err := ubq.loadLinksClicked(ctx, query, nodes,
			func(n *UserBuyer) { n.Edges.LinksClicked = []*LinkVisit{} },
			func(n *UserBuyer, e *LinkVisit) { n.Edges.LinksClicked = append(n.Edges.LinksClicked, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ubq.withNamedUserProfile {
		if err := ubq.loadUserProfile(ctx, query, nodes,
			func(n *UserBuyer) { n.appendNamedUserProfile(name) },
			func(n *UserBuyer, e *User) { n.appendNamedUserProfile(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ubq.withNamedReviews {
		if err := ubq.loadReviews(ctx, query, nodes,
			func(n *UserBuyer) { n.appendNamedReviews(name) },
			func(n *UserBuyer, e *Review) { n.appendNamedReviews(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ubq.withNamedTransactions {
		if err := ubq.loadTransactions(ctx, query, nodes,
			func(n *UserBuyer) { n.appendNamedTransactions(name) },
			func(n *UserBuyer, e *Transaction) { n.appendNamedTransactions(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ubq.withNamedLinksClicked {
		if err := ubq.loadLinksClicked(ctx, query, nodes,
			func(n *UserBuyer) { n.appendNamedLinksClicked(name) },
			func(n *UserBuyer, e *LinkVisit) { n.appendNamedLinksClicked(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ubq.loadTotal {
		if err := ubq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ubq *UserBuyerQuery) loadUserProfile(ctx context.Context, query *UserQuery, nodes []*UserBuyer, init func(*UserBuyer), assign func(*UserBuyer, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserBuyer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userbuyer.UserProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_buyer_user_profile
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_buyer_user_profile" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_buyer_user_profile" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ubq *UserBuyerQuery) loadReviews(ctx context.Context, query *ReviewQuery, nodes []*UserBuyer, init func(*UserBuyer), assign func(*UserBuyer, *Review)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserBuyer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Review(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userbuyer.ReviewsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_buyer_reviews
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_buyer_reviews" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_buyer_reviews" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ubq *UserBuyerQuery) loadTransactions(ctx context.Context, query *TransactionQuery, nodes []*UserBuyer, init func(*UserBuyer), assign func(*UserBuyer, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserBuyer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userbuyer.TransactionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_buyer_transactions
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_buyer_transactions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_buyer_transactions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ubq *UserBuyerQuery) loadLinksClicked(ctx context.Context, query *LinkVisitQuery, nodes []*UserBuyer, init func(*UserBuyer), assign func(*UserBuyer, *LinkVisit)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserBuyer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.LinkVisit(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userbuyer.LinksClickedColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_buyer_links_clicked
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_buyer_links_clicked" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_buyer_links_clicked" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ubq *UserBuyerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ubq.querySpec()
	if len(ubq.modifiers) > 0 {
		_spec.Modifiers = ubq.modifiers
	}
	_spec.Node.Columns = ubq.ctx.Fields
	if len(ubq.ctx.Fields) > 0 {
		_spec.Unique = ubq.ctx.Unique != nil && *ubq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ubq.driver, _spec)
}

func (ubq *UserBuyerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userbuyer.Table, userbuyer.Columns, sqlgraph.NewFieldSpec(userbuyer.FieldID, field.TypeInt))
	_spec.From = ubq.sql
	if unique := ubq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ubq.path != nil {
		_spec.Unique = true
	}
	if fields := ubq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userbuyer.FieldID)
		for i := range fields {
			if fields[i] != userbuyer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ubq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ubq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ubq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ubq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ubq *UserBuyerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ubq.driver.Dialect())
	t1 := builder.Table(userbuyer.Table)
	columns := ubq.ctx.Fields
	if len(columns) == 0 {
		columns = userbuyer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ubq.sql != nil {
		selector = ubq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ubq.ctx.Unique != nil && *ubq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ubq.predicates {
		p(selector)
	}
	for _, p := range ubq.order {
		p(selector)
	}
	if offset := ubq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ubq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedUserProfile tells the query-builder to eager-load the nodes that are connected to the "userProfile"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithNamedUserProfile(name string, opts ...func(*UserQuery)) *UserBuyerQuery {
	query := (&UserClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ubq.withNamedUserProfile == nil {
		ubq.withNamedUserProfile = make(map[string]*UserQuery)
	}
	ubq.withNamedUserProfile[name] = query
	return ubq
}

// WithNamedReviews tells the query-builder to eager-load the nodes that are connected to the "reviews"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithNamedReviews(name string, opts ...func(*ReviewQuery)) *UserBuyerQuery {
	query := (&ReviewClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ubq.withNamedReviews == nil {
		ubq.withNamedReviews = make(map[string]*ReviewQuery)
	}
	ubq.withNamedReviews[name] = query
	return ubq
}

// WithNamedTransactions tells the query-builder to eager-load the nodes that are connected to the "transactions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithNamedTransactions(name string, opts ...func(*TransactionQuery)) *UserBuyerQuery {
	query := (&TransactionClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ubq.withNamedTransactions == nil {
		ubq.withNamedTransactions = make(map[string]*TransactionQuery)
	}
	ubq.withNamedTransactions[name] = query
	return ubq
}

// WithNamedLinksClicked tells the query-builder to eager-load the nodes that are connected to the "linksClicked"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBuyerQuery) WithNamedLinksClicked(name string, opts ...func(*LinkVisitQuery)) *UserBuyerQuery {
	query := (&LinkVisitClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ubq.withNamedLinksClicked == nil {
		ubq.withNamedLinksClicked = make(map[string]*LinkVisitQuery)
	}
	ubq.withNamedLinksClicked[name] = query
	return ubq
}

// UserBuyerGroupBy is the group-by builder for UserBuyer entities.
type UserBuyerGroupBy struct {
	selector
	build *UserBuyerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ubgb *UserBuyerGroupBy) Aggregate(fns ...AggregateFunc) *UserBuyerGroupBy {
	ubgb.fns = append(ubgb.fns, fns...)
	return ubgb
}

// Scan applies the selector query and scans the result into the given value.
func (ubgb *UserBuyerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ubgb.build.ctx, "GroupBy")
	if err := ubgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBuyerQuery, *UserBuyerGroupBy](ctx, ubgb.build, ubgb, ubgb.build.inters, v)
}

func (ubgb *UserBuyerGroupBy) sqlScan(ctx context.Context, root *UserBuyerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ubgb.fns))
	for _, fn := range ubgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ubgb.flds)+len(ubgb.fns))
		for _, f := range *ubgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ubgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserBuyerSelect is the builder for selecting fields of UserBuyer entities.
type UserBuyerSelect struct {
	*UserBuyerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ubs *UserBuyerSelect) Aggregate(fns ...AggregateFunc) *UserBuyerSelect {
	ubs.fns = append(ubs.fns, fns...)
	return ubs
}

// Scan applies the selector query and scans the result into the given value.
func (ubs *UserBuyerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ubs.ctx, "Select")
	if err := ubs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBuyerQuery, *UserBuyerSelect](ctx, ubs.UserBuyerQuery, ubs, ubs.inters, v)
}

func (ubs *UserBuyerSelect) sqlScan(ctx context.Context, root *UserBuyerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ubs.fns))
	for _, fn := range ubs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ubs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
