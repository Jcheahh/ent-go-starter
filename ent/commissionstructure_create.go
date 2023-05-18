// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/commissionstructure"
	"entdemo/ent/userseller"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionStructureCreate is the builder for creating a CommissionStructure entity.
type CommissionStructureCreate struct {
	config
	mutation *CommissionStructureMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (csc *CommissionStructureCreate) SetName(s string) *CommissionStructureCreate {
	csc.mutation.SetName(s)
	return csc
}

// SetDescription sets the "description" field.
func (csc *CommissionStructureCreate) SetDescription(s string) *CommissionStructureCreate {
	csc.mutation.SetDescription(s)
	return csc
}

// SetCommissionValue sets the "commissionValue" field.
func (csc *CommissionStructureCreate) SetCommissionValue(s string) *CommissionStructureCreate {
	csc.mutation.SetCommissionValue(s)
	return csc
}

// SetCommissionPercentage sets the "commissionPercentage" field.
func (csc *CommissionStructureCreate) SetCommissionPercentage(s string) *CommissionStructureCreate {
	csc.mutation.SetCommissionPercentage(s)
	return csc
}

// AddProductSellerIDs adds the "productSeller" edge to the UserSeller entity by IDs.
func (csc *CommissionStructureCreate) AddProductSellerIDs(ids ...int) *CommissionStructureCreate {
	csc.mutation.AddProductSellerIDs(ids...)
	return csc
}

// AddProductSeller adds the "productSeller" edges to the UserSeller entity.
func (csc *CommissionStructureCreate) AddProductSeller(u ...*UserSeller) *CommissionStructureCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csc.AddProductSellerIDs(ids...)
}

// Mutation returns the CommissionStructureMutation object of the builder.
func (csc *CommissionStructureCreate) Mutation() *CommissionStructureMutation {
	return csc.mutation
}

// Save creates the CommissionStructure in the database.
func (csc *CommissionStructureCreate) Save(ctx context.Context) (*CommissionStructure, error) {
	return withHooks[*CommissionStructure, CommissionStructureMutation](ctx, csc.sqlSave, csc.mutation, csc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (csc *CommissionStructureCreate) SaveX(ctx context.Context) *CommissionStructure {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csc *CommissionStructureCreate) Exec(ctx context.Context) error {
	_, err := csc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csc *CommissionStructureCreate) ExecX(ctx context.Context) {
	if err := csc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csc *CommissionStructureCreate) check() error {
	if _, ok := csc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CommissionStructure.name"`)}
	}
	if _, ok := csc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "CommissionStructure.description"`)}
	}
	if _, ok := csc.mutation.CommissionValue(); !ok {
		return &ValidationError{Name: "commissionValue", err: errors.New(`ent: missing required field "CommissionStructure.commissionValue"`)}
	}
	if _, ok := csc.mutation.CommissionPercentage(); !ok {
		return &ValidationError{Name: "commissionPercentage", err: errors.New(`ent: missing required field "CommissionStructure.commissionPercentage"`)}
	}
	return nil
}

func (csc *CommissionStructureCreate) sqlSave(ctx context.Context) (*CommissionStructure, error) {
	if err := csc.check(); err != nil {
		return nil, err
	}
	_node, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	csc.mutation.id = &_node.ID
	csc.mutation.done = true
	return _node, nil
}

func (csc *CommissionStructureCreate) createSpec() (*CommissionStructure, *sqlgraph.CreateSpec) {
	var (
		_node = &CommissionStructure{config: csc.config}
		_spec = sqlgraph.NewCreateSpec(commissionstructure.Table, sqlgraph.NewFieldSpec(commissionstructure.FieldID, field.TypeInt))
	)
	if value, ok := csc.mutation.Name(); ok {
		_spec.SetField(commissionstructure.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := csc.mutation.Description(); ok {
		_spec.SetField(commissionstructure.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := csc.mutation.CommissionValue(); ok {
		_spec.SetField(commissionstructure.FieldCommissionValue, field.TypeString, value)
		_node.CommissionValue = value
	}
	if value, ok := csc.mutation.CommissionPercentage(); ok {
		_spec.SetField(commissionstructure.FieldCommissionPercentage, field.TypeString, value)
		_node.CommissionPercentage = value
	}
	if nodes := csc.mutation.ProductSellerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructure.ProductSellerTable,
			Columns: []string{commissionstructure.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommissionStructureCreateBulk is the builder for creating many CommissionStructure entities in bulk.
type CommissionStructureCreateBulk struct {
	config
	builders []*CommissionStructureCreate
}

// Save creates the CommissionStructure entities in the database.
func (cscb *CommissionStructureCreateBulk) Save(ctx context.Context) ([]*CommissionStructure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cscb.builders))
	nodes := make([]*CommissionStructure, len(cscb.builders))
	mutators := make([]Mutator, len(cscb.builders))
	for i := range cscb.builders {
		func(i int, root context.Context) {
			builder := cscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommissionStructureMutation)
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
					_, err = mutators[i+1].Mutate(root, cscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cscb *CommissionStructureCreateBulk) SaveX(ctx context.Context) []*CommissionStructure {
	v, err := cscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cscb *CommissionStructureCreateBulk) Exec(ctx context.Context) error {
	_, err := cscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscb *CommissionStructureCreateBulk) ExecX(ctx context.Context) {
	if err := cscb.Exec(ctx); err != nil {
		panic(err)
	}
}
