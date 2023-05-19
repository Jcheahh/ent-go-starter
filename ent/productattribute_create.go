// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/productattribute"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductAttributeCreate is the builder for creating a ProductAttribute entity.
type ProductAttributeCreate struct {
	config
	mutation *ProductAttributeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pac *ProductAttributeCreate) SetName(i int) *ProductAttributeCreate {
	pac.mutation.SetName(i)
	return pac
}

// SetDescription sets the "description" field.
func (pac *ProductAttributeCreate) SetDescription(i int) *ProductAttributeCreate {
	pac.mutation.SetDescription(i)
	return pac
}

// SetValue sets the "value" field.
func (pac *ProductAttributeCreate) SetValue(i int) *ProductAttributeCreate {
	pac.mutation.SetValue(i)
	return pac
}

// Mutation returns the ProductAttributeMutation object of the builder.
func (pac *ProductAttributeCreate) Mutation() *ProductAttributeMutation {
	return pac.mutation
}

// Save creates the ProductAttribute in the database.
func (pac *ProductAttributeCreate) Save(ctx context.Context) (*ProductAttribute, error) {
	return withHooks[*ProductAttribute, ProductAttributeMutation](ctx, pac.sqlSave, pac.mutation, pac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pac *ProductAttributeCreate) SaveX(ctx context.Context) *ProductAttribute {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *ProductAttributeCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *ProductAttributeCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *ProductAttributeCreate) check() error {
	if _, ok := pac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProductAttribute.name"`)}
	}
	if _, ok := pac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ProductAttribute.description"`)}
	}
	if _, ok := pac.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "ProductAttribute.value"`)}
	}
	return nil
}

func (pac *ProductAttributeCreate) sqlSave(ctx context.Context) (*ProductAttribute, error) {
	if err := pac.check(); err != nil {
		return nil, err
	}
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pac.mutation.id = &_node.ID
	pac.mutation.done = true
	return _node, nil
}

func (pac *ProductAttributeCreate) createSpec() (*ProductAttribute, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductAttribute{config: pac.config}
		_spec = sqlgraph.NewCreateSpec(productattribute.Table, sqlgraph.NewFieldSpec(productattribute.FieldID, field.TypeInt))
	)
	if value, ok := pac.mutation.Name(); ok {
		_spec.SetField(productattribute.FieldName, field.TypeInt, value)
		_node.Name = value
	}
	if value, ok := pac.mutation.Description(); ok {
		_spec.SetField(productattribute.FieldDescription, field.TypeInt, value)
		_node.Description = value
	}
	if value, ok := pac.mutation.Value(); ok {
		_spec.SetField(productattribute.FieldValue, field.TypeInt, value)
		_node.Value = value
	}
	return _node, _spec
}

// ProductAttributeCreateBulk is the builder for creating many ProductAttribute entities in bulk.
type ProductAttributeCreateBulk struct {
	config
	builders []*ProductAttributeCreate
}

// Save creates the ProductAttribute entities in the database.
func (pacb *ProductAttributeCreateBulk) Save(ctx context.Context) ([]*ProductAttribute, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*ProductAttribute, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductAttributeMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *ProductAttributeCreateBulk) SaveX(ctx context.Context) []*ProductAttribute {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *ProductAttributeCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *ProductAttributeCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}