// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/linkvisit"
	"entdemo/ent/review"
	"entdemo/ent/transaction"
	"entdemo/ent/user"
	"entdemo/ent/userbuyer"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBuyerCreate is the builder for creating a UserBuyer entity.
type UserBuyerCreate struct {
	config
	mutation *UserBuyerMutation
	hooks    []Hook
}

// SetPlaceholder sets the "placeholder" field.
func (ubc *UserBuyerCreate) SetPlaceholder(i int) *UserBuyerCreate {
	ubc.mutation.SetPlaceholder(i)
	return ubc
}

// SetNillablePlaceholder sets the "placeholder" field if the given value is not nil.
func (ubc *UserBuyerCreate) SetNillablePlaceholder(i *int) *UserBuyerCreate {
	if i != nil {
		ubc.SetPlaceholder(*i)
	}
	return ubc
}

// AddUserProfileIDs adds the "userProfile" edge to the User entity by IDs.
func (ubc *UserBuyerCreate) AddUserProfileIDs(ids ...int) *UserBuyerCreate {
	ubc.mutation.AddUserProfileIDs(ids...)
	return ubc
}

// AddUserProfile adds the "userProfile" edges to the User entity.
func (ubc *UserBuyerCreate) AddUserProfile(u ...*User) *UserBuyerCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ubc.AddUserProfileIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ubc *UserBuyerCreate) AddReviewIDs(ids ...int) *UserBuyerCreate {
	ubc.mutation.AddReviewIDs(ids...)
	return ubc
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ubc *UserBuyerCreate) AddReviews(r ...*Review) *UserBuyerCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ubc.AddReviewIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ubc *UserBuyerCreate) AddTransactionIDs(ids ...int) *UserBuyerCreate {
	ubc.mutation.AddTransactionIDs(ids...)
	return ubc
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ubc *UserBuyerCreate) AddTransactions(t ...*Transaction) *UserBuyerCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ubc.AddTransactionIDs(ids...)
}

// AddLinksClickedIDs adds the "linksClicked" edge to the LinkVisit entity by IDs.
func (ubc *UserBuyerCreate) AddLinksClickedIDs(ids ...int) *UserBuyerCreate {
	ubc.mutation.AddLinksClickedIDs(ids...)
	return ubc
}

// AddLinksClicked adds the "linksClicked" edges to the LinkVisit entity.
func (ubc *UserBuyerCreate) AddLinksClicked(l ...*LinkVisit) *UserBuyerCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ubc.AddLinksClickedIDs(ids...)
}

// Mutation returns the UserBuyerMutation object of the builder.
func (ubc *UserBuyerCreate) Mutation() *UserBuyerMutation {
	return ubc.mutation
}

// Save creates the UserBuyer in the database.
func (ubc *UserBuyerCreate) Save(ctx context.Context) (*UserBuyer, error) {
	return withHooks[*UserBuyer, UserBuyerMutation](ctx, ubc.sqlSave, ubc.mutation, ubc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ubc *UserBuyerCreate) SaveX(ctx context.Context) *UserBuyer {
	v, err := ubc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubc *UserBuyerCreate) Exec(ctx context.Context) error {
	_, err := ubc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubc *UserBuyerCreate) ExecX(ctx context.Context) {
	if err := ubc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubc *UserBuyerCreate) check() error {
	return nil
}

func (ubc *UserBuyerCreate) sqlSave(ctx context.Context) (*UserBuyer, error) {
	if err := ubc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ubc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ubc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ubc.mutation.id = &_node.ID
	ubc.mutation.done = true
	return _node, nil
}

func (ubc *UserBuyerCreate) createSpec() (*UserBuyer, *sqlgraph.CreateSpec) {
	var (
		_node = &UserBuyer{config: ubc.config}
		_spec = sqlgraph.NewCreateSpec(userbuyer.Table, sqlgraph.NewFieldSpec(userbuyer.FieldID, field.TypeInt))
	)
	if value, ok := ubc.mutation.Placeholder(); ok {
		_spec.SetField(userbuyer.FieldPlaceholder, field.TypeInt, value)
		_node.Placeholder = value
	}
	if nodes := ubc.mutation.UserProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ubc.mutation.ReviewsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ubc.mutation.TransactionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ubc.mutation.LinksClickedIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserBuyerCreateBulk is the builder for creating many UserBuyer entities in bulk.
type UserBuyerCreateBulk struct {
	config
	builders []*UserBuyerCreate
}

// Save creates the UserBuyer entities in the database.
func (ubcb *UserBuyerCreateBulk) Save(ctx context.Context) ([]*UserBuyer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ubcb.builders))
	nodes := make([]*UserBuyer, len(ubcb.builders))
	mutators := make([]Mutator, len(ubcb.builders))
	for i := range ubcb.builders {
		func(i int, root context.Context) {
			builder := ubcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserBuyerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ubcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ubcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ubcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ubcb *UserBuyerCreateBulk) SaveX(ctx context.Context) []*UserBuyer {
	v, err := ubcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubcb *UserBuyerCreateBulk) Exec(ctx context.Context) error {
	_, err := ubcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubcb *UserBuyerCreateBulk) ExecX(ctx context.Context) {
	if err := ubcb.Exec(ctx); err != nil {
		panic(err)
	}
}
