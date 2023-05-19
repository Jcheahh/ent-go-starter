// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/paymentmethod"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentMethodCreate is the builder for creating a PaymentMethod entity.
type PaymentMethodCreate struct {
	config
	mutation *PaymentMethodMutation
	hooks    []Hook
}

// SetXid sets the "xid" field.
func (pmc *PaymentMethodCreate) SetXid(i int) *PaymentMethodCreate {
	pmc.mutation.SetXid(i)
	return pmc
}

// Mutation returns the PaymentMethodMutation object of the builder.
func (pmc *PaymentMethodCreate) Mutation() *PaymentMethodMutation {
	return pmc.mutation
}

// Save creates the PaymentMethod in the database.
func (pmc *PaymentMethodCreate) Save(ctx context.Context) (*PaymentMethod, error) {
	return withHooks[*PaymentMethod, PaymentMethodMutation](ctx, pmc.sqlSave, pmc.mutation, pmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pmc *PaymentMethodCreate) SaveX(ctx context.Context) *PaymentMethod {
	v, err := pmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmc *PaymentMethodCreate) Exec(ctx context.Context) error {
	_, err := pmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmc *PaymentMethodCreate) ExecX(ctx context.Context) {
	if err := pmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmc *PaymentMethodCreate) check() error {
	if _, ok := pmc.mutation.Xid(); !ok {
		return &ValidationError{Name: "xid", err: errors.New(`ent: missing required field "PaymentMethod.xid"`)}
	}
	return nil
}

func (pmc *PaymentMethodCreate) sqlSave(ctx context.Context) (*PaymentMethod, error) {
	if err := pmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pmc.mutation.id = &_node.ID
	pmc.mutation.done = true
	return _node, nil
}

func (pmc *PaymentMethodCreate) createSpec() (*PaymentMethod, *sqlgraph.CreateSpec) {
	var (
		_node = &PaymentMethod{config: pmc.config}
		_spec = sqlgraph.NewCreateSpec(paymentmethod.Table, sqlgraph.NewFieldSpec(paymentmethod.FieldID, field.TypeInt))
	)
	if value, ok := pmc.mutation.Xid(); ok {
		_spec.SetField(paymentmethod.FieldXid, field.TypeInt, value)
		_node.Xid = value
	}
	return _node, _spec
}

// PaymentMethodCreateBulk is the builder for creating many PaymentMethod entities in bulk.
type PaymentMethodCreateBulk struct {
	config
	builders []*PaymentMethodCreate
}

// Save creates the PaymentMethod entities in the database.
func (pmcb *PaymentMethodCreateBulk) Save(ctx context.Context) ([]*PaymentMethod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pmcb.builders))
	nodes := make([]*PaymentMethod, len(pmcb.builders))
	mutators := make([]Mutator, len(pmcb.builders))
	for i := range pmcb.builders {
		func(i int, root context.Context) {
			builder := pmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMethodMutation)
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
					_, err = mutators[i+1].Mutate(root, pmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmcb *PaymentMethodCreateBulk) SaveX(ctx context.Context) []*PaymentMethod {
	v, err := pmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmcb *PaymentMethodCreateBulk) Exec(ctx context.Context) error {
	_, err := pmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmcb *PaymentMethodCreateBulk) ExecX(ctx context.Context) {
	if err := pmcb.Exec(ctx); err != nil {
		panic(err)
	}
}