// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/linkvisit"
	"entdemo/ent/predicate"
	"entdemo/ent/review"
	"entdemo/ent/transaction"
	"entdemo/ent/user"
	"entdemo/ent/userbuyer"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBuyerUpdate is the builder for updating UserBuyer entities.
type UserBuyerUpdate struct {
	config
	hooks    []Hook
	mutation *UserBuyerMutation
}

// Where appends a list predicates to the UserBuyerUpdate builder.
func (ubu *UserBuyerUpdate) Where(ps ...predicate.UserBuyer) *UserBuyerUpdate {
	ubu.mutation.Where(ps...)
	return ubu
}

// SetPlaceholder sets the "placeholder" field.
func (ubu *UserBuyerUpdate) SetPlaceholder(i int) *UserBuyerUpdate {
	ubu.mutation.ResetPlaceholder()
	ubu.mutation.SetPlaceholder(i)
	return ubu
}

// SetNillablePlaceholder sets the "placeholder" field if the given value is not nil.
func (ubu *UserBuyerUpdate) SetNillablePlaceholder(i *int) *UserBuyerUpdate {
	if i != nil {
		ubu.SetPlaceholder(*i)
	}
	return ubu
}

// AddPlaceholder adds i to the "placeholder" field.
func (ubu *UserBuyerUpdate) AddPlaceholder(i int) *UserBuyerUpdate {
	ubu.mutation.AddPlaceholder(i)
	return ubu
}

// ClearPlaceholder clears the value of the "placeholder" field.
func (ubu *UserBuyerUpdate) ClearPlaceholder() *UserBuyerUpdate {
	ubu.mutation.ClearPlaceholder()
	return ubu
}

// AddUserProfileIDs adds the "userProfile" edge to the User entity by IDs.
func (ubu *UserBuyerUpdate) AddUserProfileIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.AddUserProfileIDs(ids...)
	return ubu
}

// AddUserProfile adds the "userProfile" edges to the User entity.
func (ubu *UserBuyerUpdate) AddUserProfile(u ...*User) *UserBuyerUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ubu.AddUserProfileIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ubu *UserBuyerUpdate) AddReviewIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.AddReviewIDs(ids...)
	return ubu
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ubu *UserBuyerUpdate) AddReviews(r ...*Review) *UserBuyerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ubu.AddReviewIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ubu *UserBuyerUpdate) AddTransactionIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.AddTransactionIDs(ids...)
	return ubu
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ubu *UserBuyerUpdate) AddTransactions(t ...*Transaction) *UserBuyerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ubu.AddTransactionIDs(ids...)
}

// AddLinksClickedIDs adds the "linksClicked" edge to the LinkVisit entity by IDs.
func (ubu *UserBuyerUpdate) AddLinksClickedIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.AddLinksClickedIDs(ids...)
	return ubu
}

// AddLinksClicked adds the "linksClicked" edges to the LinkVisit entity.
func (ubu *UserBuyerUpdate) AddLinksClicked(l ...*LinkVisit) *UserBuyerUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ubu.AddLinksClickedIDs(ids...)
}

// Mutation returns the UserBuyerMutation object of the builder.
func (ubu *UserBuyerUpdate) Mutation() *UserBuyerMutation {
	return ubu.mutation
}

// ClearUserProfile clears all "userProfile" edges to the User entity.
func (ubu *UserBuyerUpdate) ClearUserProfile() *UserBuyerUpdate {
	ubu.mutation.ClearUserProfile()
	return ubu
}

// RemoveUserProfileIDs removes the "userProfile" edge to User entities by IDs.
func (ubu *UserBuyerUpdate) RemoveUserProfileIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.RemoveUserProfileIDs(ids...)
	return ubu
}

// RemoveUserProfile removes "userProfile" edges to User entities.
func (ubu *UserBuyerUpdate) RemoveUserProfile(u ...*User) *UserBuyerUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ubu.RemoveUserProfileIDs(ids...)
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ubu *UserBuyerUpdate) ClearReviews() *UserBuyerUpdate {
	ubu.mutation.ClearReviews()
	return ubu
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ubu *UserBuyerUpdate) RemoveReviewIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.RemoveReviewIDs(ids...)
	return ubu
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ubu *UserBuyerUpdate) RemoveReviews(r ...*Review) *UserBuyerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ubu.RemoveReviewIDs(ids...)
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ubu *UserBuyerUpdate) ClearTransactions() *UserBuyerUpdate {
	ubu.mutation.ClearTransactions()
	return ubu
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ubu *UserBuyerUpdate) RemoveTransactionIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.RemoveTransactionIDs(ids...)
	return ubu
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ubu *UserBuyerUpdate) RemoveTransactions(t ...*Transaction) *UserBuyerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ubu.RemoveTransactionIDs(ids...)
}

// ClearLinksClicked clears all "linksClicked" edges to the LinkVisit entity.
func (ubu *UserBuyerUpdate) ClearLinksClicked() *UserBuyerUpdate {
	ubu.mutation.ClearLinksClicked()
	return ubu
}

// RemoveLinksClickedIDs removes the "linksClicked" edge to LinkVisit entities by IDs.
func (ubu *UserBuyerUpdate) RemoveLinksClickedIDs(ids ...int) *UserBuyerUpdate {
	ubu.mutation.RemoveLinksClickedIDs(ids...)
	return ubu
}

// RemoveLinksClicked removes "linksClicked" edges to LinkVisit entities.
func (ubu *UserBuyerUpdate) RemoveLinksClicked(l ...*LinkVisit) *UserBuyerUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ubu.RemoveLinksClickedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ubu *UserBuyerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserBuyerMutation](ctx, ubu.sqlSave, ubu.mutation, ubu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ubu *UserBuyerUpdate) SaveX(ctx context.Context) int {
	affected, err := ubu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ubu *UserBuyerUpdate) Exec(ctx context.Context) error {
	_, err := ubu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubu *UserBuyerUpdate) ExecX(ctx context.Context) {
	if err := ubu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ubu *UserBuyerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userbuyer.Table, userbuyer.Columns, sqlgraph.NewFieldSpec(userbuyer.FieldID, field.TypeInt))
	if ps := ubu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ubu.mutation.Placeholder(); ok {
		_spec.SetField(userbuyer.FieldPlaceholder, field.TypeInt, value)
	}
	if value, ok := ubu.mutation.AddedPlaceholder(); ok {
		_spec.AddField(userbuyer.FieldPlaceholder, field.TypeInt, value)
	}
	if ubu.mutation.PlaceholderCleared() {
		_spec.ClearField(userbuyer.FieldPlaceholder, field.TypeInt)
	}
	if ubu.mutation.UserProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.RemovedUserProfileIDs(); len(nodes) > 0 && !ubu.mutation.UserProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.UserProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ubu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ubu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubu.mutation.LinksClickedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.RemovedLinksClickedIDs(); len(nodes) > 0 && !ubu.mutation.LinksClickedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubu.mutation.LinksClickedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ubu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userbuyer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ubu.mutation.done = true
	return n, nil
}

// UserBuyerUpdateOne is the builder for updating a single UserBuyer entity.
type UserBuyerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserBuyerMutation
}

// SetPlaceholder sets the "placeholder" field.
func (ubuo *UserBuyerUpdateOne) SetPlaceholder(i int) *UserBuyerUpdateOne {
	ubuo.mutation.ResetPlaceholder()
	ubuo.mutation.SetPlaceholder(i)
	return ubuo
}

// SetNillablePlaceholder sets the "placeholder" field if the given value is not nil.
func (ubuo *UserBuyerUpdateOne) SetNillablePlaceholder(i *int) *UserBuyerUpdateOne {
	if i != nil {
		ubuo.SetPlaceholder(*i)
	}
	return ubuo
}

// AddPlaceholder adds i to the "placeholder" field.
func (ubuo *UserBuyerUpdateOne) AddPlaceholder(i int) *UserBuyerUpdateOne {
	ubuo.mutation.AddPlaceholder(i)
	return ubuo
}

// ClearPlaceholder clears the value of the "placeholder" field.
func (ubuo *UserBuyerUpdateOne) ClearPlaceholder() *UserBuyerUpdateOne {
	ubuo.mutation.ClearPlaceholder()
	return ubuo
}

// AddUserProfileIDs adds the "userProfile" edge to the User entity by IDs.
func (ubuo *UserBuyerUpdateOne) AddUserProfileIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.AddUserProfileIDs(ids...)
	return ubuo
}

// AddUserProfile adds the "userProfile" edges to the User entity.
func (ubuo *UserBuyerUpdateOne) AddUserProfile(u ...*User) *UserBuyerUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ubuo.AddUserProfileIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ubuo *UserBuyerUpdateOne) AddReviewIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.AddReviewIDs(ids...)
	return ubuo
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ubuo *UserBuyerUpdateOne) AddReviews(r ...*Review) *UserBuyerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ubuo.AddReviewIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ubuo *UserBuyerUpdateOne) AddTransactionIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.AddTransactionIDs(ids...)
	return ubuo
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ubuo *UserBuyerUpdateOne) AddTransactions(t ...*Transaction) *UserBuyerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ubuo.AddTransactionIDs(ids...)
}

// AddLinksClickedIDs adds the "linksClicked" edge to the LinkVisit entity by IDs.
func (ubuo *UserBuyerUpdateOne) AddLinksClickedIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.AddLinksClickedIDs(ids...)
	return ubuo
}

// AddLinksClicked adds the "linksClicked" edges to the LinkVisit entity.
func (ubuo *UserBuyerUpdateOne) AddLinksClicked(l ...*LinkVisit) *UserBuyerUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ubuo.AddLinksClickedIDs(ids...)
}

// Mutation returns the UserBuyerMutation object of the builder.
func (ubuo *UserBuyerUpdateOne) Mutation() *UserBuyerMutation {
	return ubuo.mutation
}

// ClearUserProfile clears all "userProfile" edges to the User entity.
func (ubuo *UserBuyerUpdateOne) ClearUserProfile() *UserBuyerUpdateOne {
	ubuo.mutation.ClearUserProfile()
	return ubuo
}

// RemoveUserProfileIDs removes the "userProfile" edge to User entities by IDs.
func (ubuo *UserBuyerUpdateOne) RemoveUserProfileIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.RemoveUserProfileIDs(ids...)
	return ubuo
}

// RemoveUserProfile removes "userProfile" edges to User entities.
func (ubuo *UserBuyerUpdateOne) RemoveUserProfile(u ...*User) *UserBuyerUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ubuo.RemoveUserProfileIDs(ids...)
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ubuo *UserBuyerUpdateOne) ClearReviews() *UserBuyerUpdateOne {
	ubuo.mutation.ClearReviews()
	return ubuo
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ubuo *UserBuyerUpdateOne) RemoveReviewIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.RemoveReviewIDs(ids...)
	return ubuo
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ubuo *UserBuyerUpdateOne) RemoveReviews(r ...*Review) *UserBuyerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ubuo.RemoveReviewIDs(ids...)
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ubuo *UserBuyerUpdateOne) ClearTransactions() *UserBuyerUpdateOne {
	ubuo.mutation.ClearTransactions()
	return ubuo
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ubuo *UserBuyerUpdateOne) RemoveTransactionIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.RemoveTransactionIDs(ids...)
	return ubuo
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ubuo *UserBuyerUpdateOne) RemoveTransactions(t ...*Transaction) *UserBuyerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ubuo.RemoveTransactionIDs(ids...)
}

// ClearLinksClicked clears all "linksClicked" edges to the LinkVisit entity.
func (ubuo *UserBuyerUpdateOne) ClearLinksClicked() *UserBuyerUpdateOne {
	ubuo.mutation.ClearLinksClicked()
	return ubuo
}

// RemoveLinksClickedIDs removes the "linksClicked" edge to LinkVisit entities by IDs.
func (ubuo *UserBuyerUpdateOne) RemoveLinksClickedIDs(ids ...int) *UserBuyerUpdateOne {
	ubuo.mutation.RemoveLinksClickedIDs(ids...)
	return ubuo
}

// RemoveLinksClicked removes "linksClicked" edges to LinkVisit entities.
func (ubuo *UserBuyerUpdateOne) RemoveLinksClicked(l ...*LinkVisit) *UserBuyerUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ubuo.RemoveLinksClickedIDs(ids...)
}

// Where appends a list predicates to the UserBuyerUpdate builder.
func (ubuo *UserBuyerUpdateOne) Where(ps ...predicate.UserBuyer) *UserBuyerUpdateOne {
	ubuo.mutation.Where(ps...)
	return ubuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ubuo *UserBuyerUpdateOne) Select(field string, fields ...string) *UserBuyerUpdateOne {
	ubuo.fields = append([]string{field}, fields...)
	return ubuo
}

// Save executes the query and returns the updated UserBuyer entity.
func (ubuo *UserBuyerUpdateOne) Save(ctx context.Context) (*UserBuyer, error) {
	return withHooks[*UserBuyer, UserBuyerMutation](ctx, ubuo.sqlSave, ubuo.mutation, ubuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ubuo *UserBuyerUpdateOne) SaveX(ctx context.Context) *UserBuyer {
	node, err := ubuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ubuo *UserBuyerUpdateOne) Exec(ctx context.Context) error {
	_, err := ubuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubuo *UserBuyerUpdateOne) ExecX(ctx context.Context) {
	if err := ubuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ubuo *UserBuyerUpdateOne) sqlSave(ctx context.Context) (_node *UserBuyer, err error) {
	_spec := sqlgraph.NewUpdateSpec(userbuyer.Table, userbuyer.Columns, sqlgraph.NewFieldSpec(userbuyer.FieldID, field.TypeInt))
	id, ok := ubuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserBuyer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ubuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userbuyer.FieldID)
		for _, f := range fields {
			if !userbuyer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userbuyer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ubuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ubuo.mutation.Placeholder(); ok {
		_spec.SetField(userbuyer.FieldPlaceholder, field.TypeInt, value)
	}
	if value, ok := ubuo.mutation.AddedPlaceholder(); ok {
		_spec.AddField(userbuyer.FieldPlaceholder, field.TypeInt, value)
	}
	if ubuo.mutation.PlaceholderCleared() {
		_spec.ClearField(userbuyer.FieldPlaceholder, field.TypeInt)
	}
	if ubuo.mutation.UserProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.RemovedUserProfileIDs(); len(nodes) > 0 && !ubuo.mutation.UserProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.UserProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.UserProfileTable,
			Columns: []string{userbuyer.UserProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ubuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.ReviewsTable,
			Columns: []string{userbuyer.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ubuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.TransactionsTable,
			Columns: []string{userbuyer.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubuo.mutation.LinksClickedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.RemovedLinksClickedIDs(); len(nodes) > 0 && !ubuo.mutation.LinksClickedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubuo.mutation.LinksClickedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userbuyer.LinksClickedTable,
			Columns: []string{userbuyer.LinksClickedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linkvisit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserBuyer{config: ubuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ubuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userbuyer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ubuo.mutation.done = true
	return _node, nil
}
