// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/contentblock"
	"entdemo/ent/image"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContentBlockCreate is the builder for creating a ContentBlock entity.
type ContentBlockCreate struct {
	config
	mutation *ContentBlockMutation
	hooks    []Hook
}

// SetPrimaryMessage sets the "primaryMessage" field.
func (cbc *ContentBlockCreate) SetPrimaryMessage(s string) *ContentBlockCreate {
	cbc.mutation.SetPrimaryMessage(s)
	return cbc
}

// SetSecondaryMessage sets the "secondaryMessage" field.
func (cbc *ContentBlockCreate) SetSecondaryMessage(s string) *ContentBlockCreate {
	cbc.mutation.SetSecondaryMessage(s)
	return cbc
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (cbc *ContentBlockCreate) AddImageIDs(ids ...int) *ContentBlockCreate {
	cbc.mutation.AddImageIDs(ids...)
	return cbc
}

// AddImage adds the "image" edges to the Image entity.
func (cbc *ContentBlockCreate) AddImage(i ...*Image) *ContentBlockCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cbc.AddImageIDs(ids...)
}

// Mutation returns the ContentBlockMutation object of the builder.
func (cbc *ContentBlockCreate) Mutation() *ContentBlockMutation {
	return cbc.mutation
}

// Save creates the ContentBlock in the database.
func (cbc *ContentBlockCreate) Save(ctx context.Context) (*ContentBlock, error) {
	return withHooks[*ContentBlock, ContentBlockMutation](ctx, cbc.sqlSave, cbc.mutation, cbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cbc *ContentBlockCreate) SaveX(ctx context.Context) *ContentBlock {
	v, err := cbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbc *ContentBlockCreate) Exec(ctx context.Context) error {
	_, err := cbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbc *ContentBlockCreate) ExecX(ctx context.Context) {
	if err := cbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbc *ContentBlockCreate) check() error {
	if _, ok := cbc.mutation.PrimaryMessage(); !ok {
		return &ValidationError{Name: "primaryMessage", err: errors.New(`ent: missing required field "ContentBlock.primaryMessage"`)}
	}
	if _, ok := cbc.mutation.SecondaryMessage(); !ok {
		return &ValidationError{Name: "secondaryMessage", err: errors.New(`ent: missing required field "ContentBlock.secondaryMessage"`)}
	}
	return nil
}

func (cbc *ContentBlockCreate) sqlSave(ctx context.Context) (*ContentBlock, error) {
	if err := cbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cbc.mutation.id = &_node.ID
	cbc.mutation.done = true
	return _node, nil
}

func (cbc *ContentBlockCreate) createSpec() (*ContentBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &ContentBlock{config: cbc.config}
		_spec = sqlgraph.NewCreateSpec(contentblock.Table, sqlgraph.NewFieldSpec(contentblock.FieldID, field.TypeInt))
	)
	if value, ok := cbc.mutation.PrimaryMessage(); ok {
		_spec.SetField(contentblock.FieldPrimaryMessage, field.TypeString, value)
		_node.PrimaryMessage = value
	}
	if value, ok := cbc.mutation.SecondaryMessage(); ok {
		_spec.SetField(contentblock.FieldSecondaryMessage, field.TypeString, value)
		_node.SecondaryMessage = value
	}
	if nodes := cbc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contentblock.ImageTable,
			Columns: []string{contentblock.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContentBlockCreateBulk is the builder for creating many ContentBlock entities in bulk.
type ContentBlockCreateBulk struct {
	config
	builders []*ContentBlockCreate
}

// Save creates the ContentBlock entities in the database.
func (cbcb *ContentBlockCreateBulk) Save(ctx context.Context) ([]*ContentBlock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cbcb.builders))
	nodes := make([]*ContentBlock, len(cbcb.builders))
	mutators := make([]Mutator, len(cbcb.builders))
	for i := range cbcb.builders {
		func(i int, root context.Context) {
			builder := cbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContentBlockMutation)
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
					_, err = mutators[i+1].Mutate(root, cbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cbcb *ContentBlockCreateBulk) SaveX(ctx context.Context) []*ContentBlock {
	v, err := cbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbcb *ContentBlockCreateBulk) Exec(ctx context.Context) error {
	_, err := cbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbcb *ContentBlockCreateBulk) ExecX(ctx context.Context) {
	if err := cbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
