// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/product"
	"entdemo/ent/viewanalytics"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ViewAnalyticsCreate is the builder for creating a ViewAnalytics entity.
type ViewAnalyticsCreate struct {
	config
	mutation *ViewAnalyticsMutation
	hooks    []Hook
}

// SetViews sets the "views" field.
func (vac *ViewAnalyticsCreate) SetViews(i int) *ViewAnalyticsCreate {
	vac.mutation.SetViews(i)
	return vac
}

// SetScrolls sets the "scrolls" field.
func (vac *ViewAnalyticsCreate) SetScrolls(i int) *ViewAnalyticsCreate {
	vac.mutation.SetScrolls(i)
	return vac
}

// SetExits sets the "exits" field.
func (vac *ViewAnalyticsCreate) SetExits(i int) *ViewAnalyticsCreate {
	vac.mutation.SetExits(i)
	return vac
}

// SetDateCreated sets the "dateCreated" field.
func (vac *ViewAnalyticsCreate) SetDateCreated(s string) *ViewAnalyticsCreate {
	vac.mutation.SetDateCreated(s)
	return vac
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (vac *ViewAnalyticsCreate) AddProductIDs(ids ...int) *ViewAnalyticsCreate {
	vac.mutation.AddProductIDs(ids...)
	return vac
}

// AddProduct adds the "product" edges to the Product entity.
func (vac *ViewAnalyticsCreate) AddProduct(p ...*Product) *ViewAnalyticsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vac.AddProductIDs(ids...)
}

// Mutation returns the ViewAnalyticsMutation object of the builder.
func (vac *ViewAnalyticsCreate) Mutation() *ViewAnalyticsMutation {
	return vac.mutation
}

// Save creates the ViewAnalytics in the database.
func (vac *ViewAnalyticsCreate) Save(ctx context.Context) (*ViewAnalytics, error) {
	return withHooks[*ViewAnalytics, ViewAnalyticsMutation](ctx, vac.sqlSave, vac.mutation, vac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vac *ViewAnalyticsCreate) SaveX(ctx context.Context) *ViewAnalytics {
	v, err := vac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vac *ViewAnalyticsCreate) Exec(ctx context.Context) error {
	_, err := vac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vac *ViewAnalyticsCreate) ExecX(ctx context.Context) {
	if err := vac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vac *ViewAnalyticsCreate) check() error {
	if _, ok := vac.mutation.Views(); !ok {
		return &ValidationError{Name: "views", err: errors.New(`ent: missing required field "ViewAnalytics.views"`)}
	}
	if _, ok := vac.mutation.Scrolls(); !ok {
		return &ValidationError{Name: "scrolls", err: errors.New(`ent: missing required field "ViewAnalytics.scrolls"`)}
	}
	if _, ok := vac.mutation.Exits(); !ok {
		return &ValidationError{Name: "exits", err: errors.New(`ent: missing required field "ViewAnalytics.exits"`)}
	}
	if _, ok := vac.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "dateCreated", err: errors.New(`ent: missing required field "ViewAnalytics.dateCreated"`)}
	}
	return nil
}

func (vac *ViewAnalyticsCreate) sqlSave(ctx context.Context) (*ViewAnalytics, error) {
	if err := vac.check(); err != nil {
		return nil, err
	}
	_node, _spec := vac.createSpec()
	if err := sqlgraph.CreateNode(ctx, vac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vac.mutation.id = &_node.ID
	vac.mutation.done = true
	return _node, nil
}

func (vac *ViewAnalyticsCreate) createSpec() (*ViewAnalytics, *sqlgraph.CreateSpec) {
	var (
		_node = &ViewAnalytics{config: vac.config}
		_spec = sqlgraph.NewCreateSpec(viewanalytics.Table, sqlgraph.NewFieldSpec(viewanalytics.FieldID, field.TypeInt))
	)
	if value, ok := vac.mutation.Views(); ok {
		_spec.SetField(viewanalytics.FieldViews, field.TypeInt, value)
		_node.Views = value
	}
	if value, ok := vac.mutation.Scrolls(); ok {
		_spec.SetField(viewanalytics.FieldScrolls, field.TypeInt, value)
		_node.Scrolls = value
	}
	if value, ok := vac.mutation.Exits(); ok {
		_spec.SetField(viewanalytics.FieldExits, field.TypeInt, value)
		_node.Exits = value
	}
	if value, ok := vac.mutation.DateCreated(); ok {
		_spec.SetField(viewanalytics.FieldDateCreated, field.TypeString, value)
		_node.DateCreated = value
	}
	if nodes := vac.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   viewanalytics.ProductTable,
			Columns: []string{viewanalytics.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ViewAnalyticsCreateBulk is the builder for creating many ViewAnalytics entities in bulk.
type ViewAnalyticsCreateBulk struct {
	config
	builders []*ViewAnalyticsCreate
}

// Save creates the ViewAnalytics entities in the database.
func (vacb *ViewAnalyticsCreateBulk) Save(ctx context.Context) ([]*ViewAnalytics, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vacb.builders))
	nodes := make([]*ViewAnalytics, len(vacb.builders))
	mutators := make([]Mutator, len(vacb.builders))
	for i := range vacb.builders {
		func(i int, root context.Context) {
			builder := vacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ViewAnalyticsMutation)
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
					_, err = mutators[i+1].Mutate(root, vacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vacb *ViewAnalyticsCreateBulk) SaveX(ctx context.Context) []*ViewAnalytics {
	v, err := vacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vacb *ViewAnalyticsCreateBulk) Exec(ctx context.Context) error {
	_, err := vacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vacb *ViewAnalyticsCreateBulk) ExecX(ctx context.Context) {
	if err := vacb.Exec(ctx); err != nil {
		panic(err)
	}
}
