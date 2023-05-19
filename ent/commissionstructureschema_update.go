// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/commissionstructureschema"
	"entdemo/ent/predicate"
	"entdemo/ent/userseller"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionStructureSchemaUpdate is the builder for updating CommissionStructureSchema entities.
type CommissionStructureSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *CommissionStructureSchemaMutation
}

// Where appends a list predicates to the CommissionStructureSchemaUpdate builder.
func (cssu *CommissionStructureSchemaUpdate) Where(ps ...predicate.CommissionStructureSchema) *CommissionStructureSchemaUpdate {
	cssu.mutation.Where(ps...)
	return cssu
}

// SetName sets the "name" field.
func (cssu *CommissionStructureSchemaUpdate) SetName(s string) *CommissionStructureSchemaUpdate {
	cssu.mutation.SetName(s)
	return cssu
}

// SetDescription sets the "description" field.
func (cssu *CommissionStructureSchemaUpdate) SetDescription(s string) *CommissionStructureSchemaUpdate {
	cssu.mutation.SetDescription(s)
	return cssu
}

// SetCommissionValue sets the "commissionValue" field.
func (cssu *CommissionStructureSchemaUpdate) SetCommissionValue(s string) *CommissionStructureSchemaUpdate {
	cssu.mutation.SetCommissionValue(s)
	return cssu
}

// SetCommissionPercentage sets the "commissionPercentage" field.
func (cssu *CommissionStructureSchemaUpdate) SetCommissionPercentage(s string) *CommissionStructureSchemaUpdate {
	cssu.mutation.SetCommissionPercentage(s)
	return cssu
}

// AddProductSellerIDs adds the "productSeller" edge to the UserSeller entity by IDs.
func (cssu *CommissionStructureSchemaUpdate) AddProductSellerIDs(ids ...int) *CommissionStructureSchemaUpdate {
	cssu.mutation.AddProductSellerIDs(ids...)
	return cssu
}

// AddProductSeller adds the "productSeller" edges to the UserSeller entity.
func (cssu *CommissionStructureSchemaUpdate) AddProductSeller(u ...*UserSeller) *CommissionStructureSchemaUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cssu.AddProductSellerIDs(ids...)
}

// Mutation returns the CommissionStructureSchemaMutation object of the builder.
func (cssu *CommissionStructureSchemaUpdate) Mutation() *CommissionStructureSchemaMutation {
	return cssu.mutation
}

// ClearProductSeller clears all "productSeller" edges to the UserSeller entity.
func (cssu *CommissionStructureSchemaUpdate) ClearProductSeller() *CommissionStructureSchemaUpdate {
	cssu.mutation.ClearProductSeller()
	return cssu
}

// RemoveProductSellerIDs removes the "productSeller" edge to UserSeller entities by IDs.
func (cssu *CommissionStructureSchemaUpdate) RemoveProductSellerIDs(ids ...int) *CommissionStructureSchemaUpdate {
	cssu.mutation.RemoveProductSellerIDs(ids...)
	return cssu
}

// RemoveProductSeller removes "productSeller" edges to UserSeller entities.
func (cssu *CommissionStructureSchemaUpdate) RemoveProductSeller(u ...*UserSeller) *CommissionStructureSchemaUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cssu.RemoveProductSellerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cssu *CommissionStructureSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CommissionStructureSchemaMutation](ctx, cssu.sqlSave, cssu.mutation, cssu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cssu *CommissionStructureSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := cssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cssu *CommissionStructureSchemaUpdate) Exec(ctx context.Context) error {
	_, err := cssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cssu *CommissionStructureSchemaUpdate) ExecX(ctx context.Context) {
	if err := cssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cssu *CommissionStructureSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(commissionstructureschema.Table, commissionstructureschema.Columns, sqlgraph.NewFieldSpec(commissionstructureschema.FieldID, field.TypeInt))
	if ps := cssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cssu.mutation.Name(); ok {
		_spec.SetField(commissionstructureschema.FieldName, field.TypeString, value)
	}
	if value, ok := cssu.mutation.Description(); ok {
		_spec.SetField(commissionstructureschema.FieldDescription, field.TypeString, value)
	}
	if value, ok := cssu.mutation.CommissionValue(); ok {
		_spec.SetField(commissionstructureschema.FieldCommissionValue, field.TypeString, value)
	}
	if value, ok := cssu.mutation.CommissionPercentage(); ok {
		_spec.SetField(commissionstructureschema.FieldCommissionPercentage, field.TypeString, value)
	}
	if cssu.mutation.ProductSellerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cssu.mutation.RemovedProductSellerIDs(); len(nodes) > 0 && !cssu.mutation.ProductSellerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cssu.mutation.ProductSellerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionstructureschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cssu.mutation.done = true
	return n, nil
}

// CommissionStructureSchemaUpdateOne is the builder for updating a single CommissionStructureSchema entity.
type CommissionStructureSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommissionStructureSchemaMutation
}

// SetName sets the "name" field.
func (cssuo *CommissionStructureSchemaUpdateOne) SetName(s string) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.SetName(s)
	return cssuo
}

// SetDescription sets the "description" field.
func (cssuo *CommissionStructureSchemaUpdateOne) SetDescription(s string) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.SetDescription(s)
	return cssuo
}

// SetCommissionValue sets the "commissionValue" field.
func (cssuo *CommissionStructureSchemaUpdateOne) SetCommissionValue(s string) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.SetCommissionValue(s)
	return cssuo
}

// SetCommissionPercentage sets the "commissionPercentage" field.
func (cssuo *CommissionStructureSchemaUpdateOne) SetCommissionPercentage(s string) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.SetCommissionPercentage(s)
	return cssuo
}

// AddProductSellerIDs adds the "productSeller" edge to the UserSeller entity by IDs.
func (cssuo *CommissionStructureSchemaUpdateOne) AddProductSellerIDs(ids ...int) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.AddProductSellerIDs(ids...)
	return cssuo
}

// AddProductSeller adds the "productSeller" edges to the UserSeller entity.
func (cssuo *CommissionStructureSchemaUpdateOne) AddProductSeller(u ...*UserSeller) *CommissionStructureSchemaUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cssuo.AddProductSellerIDs(ids...)
}

// Mutation returns the CommissionStructureSchemaMutation object of the builder.
func (cssuo *CommissionStructureSchemaUpdateOne) Mutation() *CommissionStructureSchemaMutation {
	return cssuo.mutation
}

// ClearProductSeller clears all "productSeller" edges to the UserSeller entity.
func (cssuo *CommissionStructureSchemaUpdateOne) ClearProductSeller() *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.ClearProductSeller()
	return cssuo
}

// RemoveProductSellerIDs removes the "productSeller" edge to UserSeller entities by IDs.
func (cssuo *CommissionStructureSchemaUpdateOne) RemoveProductSellerIDs(ids ...int) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.RemoveProductSellerIDs(ids...)
	return cssuo
}

// RemoveProductSeller removes "productSeller" edges to UserSeller entities.
func (cssuo *CommissionStructureSchemaUpdateOne) RemoveProductSeller(u ...*UserSeller) *CommissionStructureSchemaUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cssuo.RemoveProductSellerIDs(ids...)
}

// Where appends a list predicates to the CommissionStructureSchemaUpdate builder.
func (cssuo *CommissionStructureSchemaUpdateOne) Where(ps ...predicate.CommissionStructureSchema) *CommissionStructureSchemaUpdateOne {
	cssuo.mutation.Where(ps...)
	return cssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cssuo *CommissionStructureSchemaUpdateOne) Select(field string, fields ...string) *CommissionStructureSchemaUpdateOne {
	cssuo.fields = append([]string{field}, fields...)
	return cssuo
}

// Save executes the query and returns the updated CommissionStructureSchema entity.
func (cssuo *CommissionStructureSchemaUpdateOne) Save(ctx context.Context) (*CommissionStructureSchema, error) {
	return withHooks[*CommissionStructureSchema, CommissionStructureSchemaMutation](ctx, cssuo.sqlSave, cssuo.mutation, cssuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cssuo *CommissionStructureSchemaUpdateOne) SaveX(ctx context.Context) *CommissionStructureSchema {
	node, err := cssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cssuo *CommissionStructureSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := cssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cssuo *CommissionStructureSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := cssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cssuo *CommissionStructureSchemaUpdateOne) sqlSave(ctx context.Context) (_node *CommissionStructureSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(commissionstructureschema.Table, commissionstructureschema.Columns, sqlgraph.NewFieldSpec(commissionstructureschema.FieldID, field.TypeInt))
	id, ok := cssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CommissionStructureSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissionstructureschema.FieldID)
		for _, f := range fields {
			if !commissionstructureschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commissionstructureschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cssuo.mutation.Name(); ok {
		_spec.SetField(commissionstructureschema.FieldName, field.TypeString, value)
	}
	if value, ok := cssuo.mutation.Description(); ok {
		_spec.SetField(commissionstructureschema.FieldDescription, field.TypeString, value)
	}
	if value, ok := cssuo.mutation.CommissionValue(); ok {
		_spec.SetField(commissionstructureschema.FieldCommissionValue, field.TypeString, value)
	}
	if value, ok := cssuo.mutation.CommissionPercentage(); ok {
		_spec.SetField(commissionstructureschema.FieldCommissionPercentage, field.TypeString, value)
	}
	if cssuo.mutation.ProductSellerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cssuo.mutation.RemovedProductSellerIDs(); len(nodes) > 0 && !cssuo.mutation.ProductSellerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cssuo.mutation.ProductSellerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionstructureschema.ProductSellerTable,
			Columns: []string{commissionstructureschema.ProductSellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userseller.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CommissionStructureSchema{config: cssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionstructureschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cssuo.mutation.done = true
	return _node, nil
}
