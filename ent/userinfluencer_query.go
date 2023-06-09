// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/predicate"
	"entdemo/ent/product"
	"entdemo/ent/referrallink"
	"entdemo/ent/review"
	"entdemo/ent/tag"
	"entdemo/ent/user"
	"entdemo/ent/userinfluencer"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfluencerQuery is the builder for querying UserInfluencer entities.
type UserInfluencerQuery struct {
	config
	ctx                    *QueryContext
	order                  []userinfluencer.OrderOption
	inters                 []Interceptor
	predicates             []predicate.UserInfluencer
	withUserProfile        *UserQuery
	withReferralLinks      *ReferralLinkQuery
	withReviews            *ReviewQuery
	withProducts           *ProductQuery
	withTags               *TagQuery
	withFKs                bool
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*UserInfluencer) error
	withNamedUserProfile   map[string]*UserQuery
	withNamedReferralLinks map[string]*ReferralLinkQuery
	withNamedReviews       map[string]*ReviewQuery
	withNamedProducts      map[string]*ProductQuery
	withNamedTags          map[string]*TagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserInfluencerQuery builder.
func (uiq *UserInfluencerQuery) Where(ps ...predicate.UserInfluencer) *UserInfluencerQuery {
	uiq.predicates = append(uiq.predicates, ps...)
	return uiq
}

// Limit the number of records to be returned by this query.
func (uiq *UserInfluencerQuery) Limit(limit int) *UserInfluencerQuery {
	uiq.ctx.Limit = &limit
	return uiq
}

// Offset to start from.
func (uiq *UserInfluencerQuery) Offset(offset int) *UserInfluencerQuery {
	uiq.ctx.Offset = &offset
	return uiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uiq *UserInfluencerQuery) Unique(unique bool) *UserInfluencerQuery {
	uiq.ctx.Unique = &unique
	return uiq
}

// Order specifies how the records should be ordered.
func (uiq *UserInfluencerQuery) Order(o ...userinfluencer.OrderOption) *UserInfluencerQuery {
	uiq.order = append(uiq.order, o...)
	return uiq
}

// QueryUserProfile chains the current query on the "userProfile" edge.
func (uiq *UserInfluencerQuery) QueryUserProfile() *UserQuery {
	query := (&UserClient{config: uiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userinfluencer.Table, userinfluencer.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userinfluencer.UserProfileTable, userinfluencer.UserProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(uiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReferralLinks chains the current query on the "referralLinks" edge.
func (uiq *UserInfluencerQuery) QueryReferralLinks() *ReferralLinkQuery {
	query := (&ReferralLinkClient{config: uiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userinfluencer.Table, userinfluencer.FieldID, selector),
			sqlgraph.To(referrallink.Table, referrallink.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userinfluencer.ReferralLinksTable, userinfluencer.ReferralLinksColumn),
		)
		fromU = sqlgraph.SetNeighbors(uiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReviews chains the current query on the "reviews" edge.
func (uiq *UserInfluencerQuery) QueryReviews() *ReviewQuery {
	query := (&ReviewClient{config: uiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userinfluencer.Table, userinfluencer.FieldID, selector),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userinfluencer.ReviewsTable, userinfluencer.ReviewsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProducts chains the current query on the "products" edge.
func (uiq *UserInfluencerQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: uiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userinfluencer.Table, userinfluencer.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userinfluencer.ProductsTable, userinfluencer.ProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (uiq *UserInfluencerQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: uiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userinfluencer.Table, userinfluencer.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userinfluencer.TagsTable, userinfluencer.TagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserInfluencer entity from the query.
// Returns a *NotFoundError when no UserInfluencer was found.
func (uiq *UserInfluencerQuery) First(ctx context.Context) (*UserInfluencer, error) {
	nodes, err := uiq.Limit(1).All(setContextOp(ctx, uiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userinfluencer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uiq *UserInfluencerQuery) FirstX(ctx context.Context) *UserInfluencer {
	node, err := uiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserInfluencer ID from the query.
// Returns a *NotFoundError when no UserInfluencer ID was found.
func (uiq *UserInfluencerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uiq.Limit(1).IDs(setContextOp(ctx, uiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userinfluencer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uiq *UserInfluencerQuery) FirstIDX(ctx context.Context) int {
	id, err := uiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserInfluencer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserInfluencer entity is found.
// Returns a *NotFoundError when no UserInfluencer entities are found.
func (uiq *UserInfluencerQuery) Only(ctx context.Context) (*UserInfluencer, error) {
	nodes, err := uiq.Limit(2).All(setContextOp(ctx, uiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userinfluencer.Label}
	default:
		return nil, &NotSingularError{userinfluencer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uiq *UserInfluencerQuery) OnlyX(ctx context.Context) *UserInfluencer {
	node, err := uiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserInfluencer ID in the query.
// Returns a *NotSingularError when more than one UserInfluencer ID is found.
// Returns a *NotFoundError when no entities are found.
func (uiq *UserInfluencerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uiq.Limit(2).IDs(setContextOp(ctx, uiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userinfluencer.Label}
	default:
		err = &NotSingularError{userinfluencer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uiq *UserInfluencerQuery) OnlyIDX(ctx context.Context) int {
	id, err := uiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserInfluencers.
func (uiq *UserInfluencerQuery) All(ctx context.Context) ([]*UserInfluencer, error) {
	ctx = setContextOp(ctx, uiq.ctx, "All")
	if err := uiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserInfluencer, *UserInfluencerQuery]()
	return withInterceptors[[]*UserInfluencer](ctx, uiq, qr, uiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uiq *UserInfluencerQuery) AllX(ctx context.Context) []*UserInfluencer {
	nodes, err := uiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserInfluencer IDs.
func (uiq *UserInfluencerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uiq.ctx.Unique == nil && uiq.path != nil {
		uiq.Unique(true)
	}
	ctx = setContextOp(ctx, uiq.ctx, "IDs")
	if err = uiq.Select(userinfluencer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uiq *UserInfluencerQuery) IDsX(ctx context.Context) []int {
	ids, err := uiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uiq *UserInfluencerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uiq.ctx, "Count")
	if err := uiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uiq, querierCount[*UserInfluencerQuery](), uiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uiq *UserInfluencerQuery) CountX(ctx context.Context) int {
	count, err := uiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uiq *UserInfluencerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uiq.ctx, "Exist")
	switch _, err := uiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uiq *UserInfluencerQuery) ExistX(ctx context.Context) bool {
	exist, err := uiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserInfluencerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uiq *UserInfluencerQuery) Clone() *UserInfluencerQuery {
	if uiq == nil {
		return nil
	}
	return &UserInfluencerQuery{
		config:            uiq.config,
		ctx:               uiq.ctx.Clone(),
		order:             append([]userinfluencer.OrderOption{}, uiq.order...),
		inters:            append([]Interceptor{}, uiq.inters...),
		predicates:        append([]predicate.UserInfluencer{}, uiq.predicates...),
		withUserProfile:   uiq.withUserProfile.Clone(),
		withReferralLinks: uiq.withReferralLinks.Clone(),
		withReviews:       uiq.withReviews.Clone(),
		withProducts:      uiq.withProducts.Clone(),
		withTags:          uiq.withTags.Clone(),
		// clone intermediate query.
		sql:  uiq.sql.Clone(),
		path: uiq.path,
	}
}

// WithUserProfile tells the query-builder to eager-load the nodes that are connected to
// the "userProfile" edge. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithUserProfile(opts ...func(*UserQuery)) *UserInfluencerQuery {
	query := (&UserClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uiq.withUserProfile = query
	return uiq
}

// WithReferralLinks tells the query-builder to eager-load the nodes that are connected to
// the "referralLinks" edge. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithReferralLinks(opts ...func(*ReferralLinkQuery)) *UserInfluencerQuery {
	query := (&ReferralLinkClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uiq.withReferralLinks = query
	return uiq
}

// WithReviews tells the query-builder to eager-load the nodes that are connected to
// the "reviews" edge. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithReviews(opts ...func(*ReviewQuery)) *UserInfluencerQuery {
	query := (&ReviewClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uiq.withReviews = query
	return uiq
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithProducts(opts ...func(*ProductQuery)) *UserInfluencerQuery {
	query := (&ProductClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uiq.withProducts = query
	return uiq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithTags(opts ...func(*TagQuery)) *UserInfluencerQuery {
	query := (&TagClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uiq.withTags = query
	return uiq
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
//	client.UserInfluencer.Query().
//		GroupBy(userinfluencer.FieldPlaceholder).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uiq *UserInfluencerQuery) GroupBy(field string, fields ...string) *UserInfluencerGroupBy {
	uiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserInfluencerGroupBy{build: uiq}
	grbuild.flds = &uiq.ctx.Fields
	grbuild.label = userinfluencer.Label
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
//	client.UserInfluencer.Query().
//		Select(userinfluencer.FieldPlaceholder).
//		Scan(ctx, &v)
func (uiq *UserInfluencerQuery) Select(fields ...string) *UserInfluencerSelect {
	uiq.ctx.Fields = append(uiq.ctx.Fields, fields...)
	sbuild := &UserInfluencerSelect{UserInfluencerQuery: uiq}
	sbuild.label = userinfluencer.Label
	sbuild.flds, sbuild.scan = &uiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserInfluencerSelect configured with the given aggregations.
func (uiq *UserInfluencerQuery) Aggregate(fns ...AggregateFunc) *UserInfluencerSelect {
	return uiq.Select().Aggregate(fns...)
}

func (uiq *UserInfluencerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uiq); err != nil {
				return err
			}
		}
	}
	for _, f := range uiq.ctx.Fields {
		if !userinfluencer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uiq.path != nil {
		prev, err := uiq.path(ctx)
		if err != nil {
			return err
		}
		uiq.sql = prev
	}
	return nil
}

func (uiq *UserInfluencerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserInfluencer, error) {
	var (
		nodes       = []*UserInfluencer{}
		withFKs     = uiq.withFKs
		_spec       = uiq.querySpec()
		loadedTypes = [5]bool{
			uiq.withUserProfile != nil,
			uiq.withReferralLinks != nil,
			uiq.withReviews != nil,
			uiq.withProducts != nil,
			uiq.withTags != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userinfluencer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserInfluencer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserInfluencer{config: uiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uiq.modifiers) > 0 {
		_spec.Modifiers = uiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uiq.withUserProfile; query != nil {
		if err := uiq.loadUserProfile(ctx, query, nodes,
			func(n *UserInfluencer) { n.Edges.UserProfile = []*User{} },
			func(n *UserInfluencer, e *User) { n.Edges.UserProfile = append(n.Edges.UserProfile, e) }); err != nil {
			return nil, err
		}
	}
	if query := uiq.withReferralLinks; query != nil {
		if err := uiq.loadReferralLinks(ctx, query, nodes,
			func(n *UserInfluencer) { n.Edges.ReferralLinks = []*ReferralLink{} },
			func(n *UserInfluencer, e *ReferralLink) { n.Edges.ReferralLinks = append(n.Edges.ReferralLinks, e) }); err != nil {
			return nil, err
		}
	}
	if query := uiq.withReviews; query != nil {
		if err := uiq.loadReviews(ctx, query, nodes,
			func(n *UserInfluencer) { n.Edges.Reviews = []*Review{} },
			func(n *UserInfluencer, e *Review) { n.Edges.Reviews = append(n.Edges.Reviews, e) }); err != nil {
			return nil, err
		}
	}
	if query := uiq.withProducts; query != nil {
		if err := uiq.loadProducts(ctx, query, nodes,
			func(n *UserInfluencer) { n.Edges.Products = []*Product{} },
			func(n *UserInfluencer, e *Product) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	if query := uiq.withTags; query != nil {
		if err := uiq.loadTags(ctx, query, nodes,
			func(n *UserInfluencer) { n.Edges.Tags = []*Tag{} },
			func(n *UserInfluencer, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uiq.withNamedUserProfile {
		if err := uiq.loadUserProfile(ctx, query, nodes,
			func(n *UserInfluencer) { n.appendNamedUserProfile(name) },
			func(n *UserInfluencer, e *User) { n.appendNamedUserProfile(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uiq.withNamedReferralLinks {
		if err := uiq.loadReferralLinks(ctx, query, nodes,
			func(n *UserInfluencer) { n.appendNamedReferralLinks(name) },
			func(n *UserInfluencer, e *ReferralLink) { n.appendNamedReferralLinks(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uiq.withNamedReviews {
		if err := uiq.loadReviews(ctx, query, nodes,
			func(n *UserInfluencer) { n.appendNamedReviews(name) },
			func(n *UserInfluencer, e *Review) { n.appendNamedReviews(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uiq.withNamedProducts {
		if err := uiq.loadProducts(ctx, query, nodes,
			func(n *UserInfluencer) { n.appendNamedProducts(name) },
			func(n *UserInfluencer, e *Product) { n.appendNamedProducts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uiq.withNamedTags {
		if err := uiq.loadTags(ctx, query, nodes,
			func(n *UserInfluencer) { n.appendNamedTags(name) },
			func(n *UserInfluencer, e *Tag) { n.appendNamedTags(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range uiq.loadTotal {
		if err := uiq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uiq *UserInfluencerQuery) loadUserProfile(ctx context.Context, query *UserQuery, nodes []*UserInfluencer, init func(*UserInfluencer), assign func(*UserInfluencer, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserInfluencer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userinfluencer.UserProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_influencer_user_profile
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_influencer_user_profile" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_influencer_user_profile" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uiq *UserInfluencerQuery) loadReferralLinks(ctx context.Context, query *ReferralLinkQuery, nodes []*UserInfluencer, init func(*UserInfluencer), assign func(*UserInfluencer, *ReferralLink)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserInfluencer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ReferralLink(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userinfluencer.ReferralLinksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_influencer_referral_links
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_influencer_referral_links" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_influencer_referral_links" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uiq *UserInfluencerQuery) loadReviews(ctx context.Context, query *ReviewQuery, nodes []*UserInfluencer, init func(*UserInfluencer), assign func(*UserInfluencer, *Review)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserInfluencer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Review(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userinfluencer.ReviewsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_influencer_reviews
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_influencer_reviews" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_influencer_reviews" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uiq *UserInfluencerQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*UserInfluencer, init func(*UserInfluencer), assign func(*UserInfluencer, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserInfluencer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userinfluencer.ProductsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_influencer_products
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_influencer_products" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_influencer_products" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uiq *UserInfluencerQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*UserInfluencer, init func(*UserInfluencer), assign func(*UserInfluencer, *Tag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserInfluencer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userinfluencer.TagsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_influencer_tags
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_influencer_tags" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_influencer_tags" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uiq *UserInfluencerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uiq.querySpec()
	if len(uiq.modifiers) > 0 {
		_spec.Modifiers = uiq.modifiers
	}
	_spec.Node.Columns = uiq.ctx.Fields
	if len(uiq.ctx.Fields) > 0 {
		_spec.Unique = uiq.ctx.Unique != nil && *uiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uiq.driver, _spec)
}

func (uiq *UserInfluencerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userinfluencer.Table, userinfluencer.Columns, sqlgraph.NewFieldSpec(userinfluencer.FieldID, field.TypeInt))
	_spec.From = uiq.sql
	if unique := uiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uiq.path != nil {
		_spec.Unique = true
	}
	if fields := uiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userinfluencer.FieldID)
		for i := range fields {
			if fields[i] != userinfluencer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uiq *UserInfluencerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uiq.driver.Dialect())
	t1 := builder.Table(userinfluencer.Table)
	columns := uiq.ctx.Fields
	if len(columns) == 0 {
		columns = userinfluencer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uiq.sql != nil {
		selector = uiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uiq.ctx.Unique != nil && *uiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uiq.predicates {
		p(selector)
	}
	for _, p := range uiq.order {
		p(selector)
	}
	if offset := uiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedUserProfile tells the query-builder to eager-load the nodes that are connected to the "userProfile"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithNamedUserProfile(name string, opts ...func(*UserQuery)) *UserInfluencerQuery {
	query := (&UserClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uiq.withNamedUserProfile == nil {
		uiq.withNamedUserProfile = make(map[string]*UserQuery)
	}
	uiq.withNamedUserProfile[name] = query
	return uiq
}

// WithNamedReferralLinks tells the query-builder to eager-load the nodes that are connected to the "referralLinks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithNamedReferralLinks(name string, opts ...func(*ReferralLinkQuery)) *UserInfluencerQuery {
	query := (&ReferralLinkClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uiq.withNamedReferralLinks == nil {
		uiq.withNamedReferralLinks = make(map[string]*ReferralLinkQuery)
	}
	uiq.withNamedReferralLinks[name] = query
	return uiq
}

// WithNamedReviews tells the query-builder to eager-load the nodes that are connected to the "reviews"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithNamedReviews(name string, opts ...func(*ReviewQuery)) *UserInfluencerQuery {
	query := (&ReviewClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uiq.withNamedReviews == nil {
		uiq.withNamedReviews = make(map[string]*ReviewQuery)
	}
	uiq.withNamedReviews[name] = query
	return uiq
}

// WithNamedProducts tells the query-builder to eager-load the nodes that are connected to the "products"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithNamedProducts(name string, opts ...func(*ProductQuery)) *UserInfluencerQuery {
	query := (&ProductClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uiq.withNamedProducts == nil {
		uiq.withNamedProducts = make(map[string]*ProductQuery)
	}
	uiq.withNamedProducts[name] = query
	return uiq
}

// WithNamedTags tells the query-builder to eager-load the nodes that are connected to the "tags"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uiq *UserInfluencerQuery) WithNamedTags(name string, opts ...func(*TagQuery)) *UserInfluencerQuery {
	query := (&TagClient{config: uiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uiq.withNamedTags == nil {
		uiq.withNamedTags = make(map[string]*TagQuery)
	}
	uiq.withNamedTags[name] = query
	return uiq
}

// UserInfluencerGroupBy is the group-by builder for UserInfluencer entities.
type UserInfluencerGroupBy struct {
	selector
	build *UserInfluencerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uigb *UserInfluencerGroupBy) Aggregate(fns ...AggregateFunc) *UserInfluencerGroupBy {
	uigb.fns = append(uigb.fns, fns...)
	return uigb
}

// Scan applies the selector query and scans the result into the given value.
func (uigb *UserInfluencerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uigb.build.ctx, "GroupBy")
	if err := uigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserInfluencerQuery, *UserInfluencerGroupBy](ctx, uigb.build, uigb, uigb.build.inters, v)
}

func (uigb *UserInfluencerGroupBy) sqlScan(ctx context.Context, root *UserInfluencerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uigb.fns))
	for _, fn := range uigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uigb.flds)+len(uigb.fns))
		for _, f := range *uigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserInfluencerSelect is the builder for selecting fields of UserInfluencer entities.
type UserInfluencerSelect struct {
	*UserInfluencerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uis *UserInfluencerSelect) Aggregate(fns ...AggregateFunc) *UserInfluencerSelect {
	uis.fns = append(uis.fns, fns...)
	return uis
}

// Scan applies the selector query and scans the result into the given value.
func (uis *UserInfluencerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uis.ctx, "Select")
	if err := uis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserInfluencerQuery, *UserInfluencerSelect](ctx, uis.UserInfluencerQuery, uis, uis.inters, v)
}

func (uis *UserInfluencerSelect) sqlScan(ctx context.Context, root *UserInfluencerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uis.fns))
	for _, fn := range uis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
