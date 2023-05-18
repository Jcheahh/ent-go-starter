// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/shop"
	"entdemo/ent/user"
	"entdemo/ent/userseller"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSellerCreate is the builder for creating a UserSeller entity.
type UserSellerCreate struct {
	config
	mutation *UserSellerMutation
	hooks    []Hook
}

// SetBrandName sets the "brandName" field.
func (usc *UserSellerCreate) SetBrandName(s string) *UserSellerCreate {
	usc.mutation.SetBrandName(s)
	return usc
}

// AddUserProfileIDs adds the "userProfile" edge to the User entity by IDs.
func (usc *UserSellerCreate) AddUserProfileIDs(ids ...int) *UserSellerCreate {
	usc.mutation.AddUserProfileIDs(ids...)
	return usc
}

// AddUserProfile adds the "userProfile" edges to the User entity.
func (usc *UserSellerCreate) AddUserProfile(u ...*User) *UserSellerCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return usc.AddUserProfileIDs(ids...)
}

// AddShopIDs adds the "shops" edge to the Shop entity by IDs.
func (usc *UserSellerCreate) AddShopIDs(ids ...int) *UserSellerCreate {
	usc.mutation.AddShopIDs(ids...)
	return usc
}

// AddShops adds the "shops" edges to the Shop entity.
func (usc *UserSellerCreate) AddShops(s ...*Shop) *UserSellerCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return usc.AddShopIDs(ids...)
}

// Mutation returns the UserSellerMutation object of the builder.
func (usc *UserSellerCreate) Mutation() *UserSellerMutation {
	return usc.mutation
}

// Save creates the UserSeller in the database.
func (usc *UserSellerCreate) Save(ctx context.Context) (*UserSeller, error) {
	return withHooks[*UserSeller, UserSellerMutation](ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSellerCreate) SaveX(ctx context.Context) *UserSeller {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSellerCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSellerCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSellerCreate) check() error {
	if _, ok := usc.mutation.BrandName(); !ok {
		return &ValidationError{Name: "brandName", err: errors.New(`ent: missing required field "UserSeller.brandName"`)}
	}
	return nil
}

func (usc *UserSellerCreate) sqlSave(ctx context.Context) (*UserSeller, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSellerCreate) createSpec() (*UserSeller, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSeller{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(userseller.Table, sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.BrandName(); ok {
		_spec.SetField(userseller.FieldBrandName, field.TypeString, value)
		_node.BrandName = value
	}
	if nodes := usc.mutation.UserProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userseller.UserProfileTable,
			Columns: []string{userseller.UserProfileColumn},
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
	if nodes := usc.mutation.ShopsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userseller.ShopsTable,
			Columns: []string{userseller.ShopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shop.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSellerCreateBulk is the builder for creating many UserSeller entities in bulk.
type UserSellerCreateBulk struct {
	config
	builders []*UserSellerCreate
}

// Save creates the UserSeller entities in the database.
func (uscb *UserSellerCreateBulk) Save(ctx context.Context) ([]*UserSeller, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSeller, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSellerMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSellerCreateBulk) SaveX(ctx context.Context) []*UserSeller {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSellerCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSellerCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
