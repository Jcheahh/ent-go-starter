// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/commissionstructure"
	"entdemo/ent/predicate"
	"entdemo/ent/userseller"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionStructureUpdate is the builder for updating CommissionStructure entities.
type CommissionStructureUpdate struct {
	config
	hooks    []Hook
	mutation *CommissionStructureMutation
}

// Where appends a list predicates to the CommissionStructureUpdate builder.
func (csu *CommissionStructureUpdate) Where(ps ...predicate.CommissionStructure) *CommissionStructureUpdate {
	csu.mutation.Where(ps...)
	return csu
}

// SetName sets the "name" field.
func (csu *CommissionStructureUpdate) SetName(s string) *CommissionStructureUpdate {
	csu.mutation.SetName(s)
	return csu
}

// SetDescription sets the "description" field.
func (csu *CommissionStructureUpdate) SetDescription(s string) *CommissionStructureUpdate {
	csu.mutation.SetDescription(s)
	return csu
}

// SetCommissionValue sets the "commissionValue" field.
func (csu *CommissionStructureUpdate) SetCommissionValue(s string) *CommissionStructureUpdate {
	csu.mutation.SetCommissionValue(s)
	return csu
}

// SetCommissionPercentage sets the "commissionPercentage" field.
func (csu *CommissionStructureUpdate) SetCommissionPercentage(s string) *CommissionStructureUpdate {
	csu.mutation.SetCommissionPercentage(s)
	return csu
}

// AddProductSellerIDs adds the "productSeller" edge to the UserSeller entity by IDs.
func (csu *CommissionStructureUpdate) AddProductSellerIDs(ids ...int) *CommissionStructureUpdate {
	csu.mutation.AddProductSellerIDs(ids...)
	return csu
}

// AddProductSeller adds the "productSeller" edges to the UserSeller entity.
func (csu *CommissionStructureUpdate) AddProductSeller(u ...*UserSeller) *CommissionStructureUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csu.AddProductSellerIDs(ids...)
}

// Mutation returns the CommissionStructureMutation object of the builder.
func (csu *CommissionStructureUpdate) Mutation() *CommissionStructureMutation {
	return csu.mutation
}

// ClearProductSeller clears all "productSeller" edges to the UserSeller entity.
func (csu *CommissionStructureUpdate) ClearProductSeller() *CommissionStructureUpdate {
	csu.mutation.ClearProductSeller()
	return csu
}

// RemoveProductSellerIDs removes the "productSeller" edge to UserSeller entities by IDs.
func (csu *CommissionStructureUpdate) RemoveProductSellerIDs(ids ...int) *CommissionStructureUpdate {
	csu.mutation.RemoveProductSellerIDs(ids...)
	return csu
}

// RemoveProductSeller removes "productSeller" edges to UserSeller entities.
func (csu *CommissionStructureUpdate) RemoveProductSeller(u ...*UserSeller) *CommissionStructureUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csu.RemoveProductSellerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *CommissionStructureUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CommissionStructureMutation](ctx, csu.sqlSave, csu.mutation, csu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csu *CommissionStructureUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *CommissionStructureUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *CommissionStructureUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csu *CommissionStructureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(commissionstructure.Table, commissionstructure.Columns, sqlgraph.NewFieldSpec(commissionstructure.FieldID, field.TypeInt))
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.Name(); ok {
		_spec.SetField(commissionstructure.FieldName, field.TypeString, value)
	}
	if value, ok := csu.mutation.Description(); ok {
		_spec.SetField(commissionstructure.FieldDescription, field.TypeString, value)
	}
	if value, ok := csu.mutation.CommissionValue(); ok {
		_spec.SetField(commissionstructure.FieldCommissionValue, field.TypeString, value)
	}
	if value, ok := csu.mutation.CommissionPercentage(); ok {
		_spec.SetField(commissionstructure.FieldCommissionPercentage, field.TypeString, value)
	}
	if csu.mutation.ProductSellerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.RemovedProductSellerIDs(); len(nodes) > 0 && !csu.mutation.ProductSellerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.ProductSellerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionstructure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	csu.mutation.done = true
	return n, nil
}

// CommissionStructureUpdateOne is the builder for updating a single CommissionStructure entity.
type CommissionStructureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommissionStructureMutation
}

// SetName sets the "name" field.
func (csuo *CommissionStructureUpdateOne) SetName(s string) *CommissionStructureUpdateOne {
	csuo.mutation.SetName(s)
	return csuo
}

// SetDescription sets the "description" field.
func (csuo *CommissionStructureUpdateOne) SetDescription(s string) *CommissionStructureUpdateOne {
	csuo.mutation.SetDescription(s)
	return csuo
}

// SetCommissionValue sets the "commissionValue" field.
func (csuo *CommissionStructureUpdateOne) SetCommissionValue(s string) *CommissionStructureUpdateOne {
	csuo.mutation.SetCommissionValue(s)
	return csuo
}

// SetCommissionPercentage sets the "commissionPercentage" field.
func (csuo *CommissionStructureUpdateOne) SetCommissionPercentage(s string) *CommissionStructureUpdateOne {
	csuo.mutation.SetCommissionPercentage(s)
	return csuo
}

// AddProductSellerIDs adds the "productSeller" edge to the UserSeller entity by IDs.
func (csuo *CommissionStructureUpdateOne) AddProductSellerIDs(ids ...int) *CommissionStructureUpdateOne {
	csuo.mutation.AddProductSellerIDs(ids...)
	return csuo
}

// AddProductSeller adds the "productSeller" edges to the UserSeller entity.
func (csuo *CommissionStructureUpdateOne) AddProductSeller(u ...*UserSeller) *CommissionStructureUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csuo.AddProductSellerIDs(ids...)
}

// Mutation returns the CommissionStructureMutation object of the builder.
func (csuo *CommissionStructureUpdateOne) Mutation() *CommissionStructureMutation {
	return csuo.mutation
}

// ClearProductSeller clears all "productSeller" edges to the UserSeller entity.
func (csuo *CommissionStructureUpdateOne) ClearProductSeller() *CommissionStructureUpdateOne {
	csuo.mutation.ClearProductSeller()
	return csuo
}

// RemoveProductSellerIDs removes the "productSeller" edge to UserSeller entities by IDs.
func (csuo *CommissionStructureUpdateOne) RemoveProductSellerIDs(ids ...int) *CommissionStructureUpdateOne {
	csuo.mutation.RemoveProductSellerIDs(ids...)
	return csuo
}

// RemoveProductSeller removes "productSeller" edges to UserSeller entities.
func (csuo *CommissionStructureUpdateOne) RemoveProductSeller(u ...*UserSeller) *CommissionStructureUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csuo.RemoveProductSellerIDs(ids...)
}

// Where appends a list predicates to the CommissionStructureUpdate builder.
func (csuo *CommissionStructureUpdateOne) Where(ps ...predicate.CommissionStructure) *CommissionStructureUpdateOne {
	csuo.mutation.Where(ps...)
	return csuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csuo *CommissionStructureUpdateOne) Select(field string, fields ...string) *CommissionStructureUpdateOne {
	csuo.fields = append([]string{field}, fields...)
	return csuo
}

// Save executes the query and returns the updated CommissionStructure entity.
func (csuo *CommissionStructureUpdateOne) Save(ctx context.Context) (*CommissionStructure, error) {
	return withHooks[*CommissionStructure, CommissionStructureMutation](ctx, csuo.sqlSave, csuo.mutation, csuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *CommissionStructureUpdateOne) SaveX(ctx context.Context) *CommissionStructure {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *CommissionStructureUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *CommissionStructureUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csuo *CommissionStructureUpdateOne) sqlSave(ctx context.Context) (_node *CommissionStructure, err error) {
	_spec := sqlgraph.NewUpdateSpec(commissionstructure.Table, commissionstructure.Columns, sqlgraph.NewFieldSpec(commissionstructure.FieldID, field.TypeInt))
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CommissionStructure.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissionstructure.FieldID)
		for _, f := range fields {
			if !commissionstructure.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commissionstructure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.Name(); ok {
		_spec.SetField(commissionstructure.FieldName, field.TypeString, value)
	}
	if value, ok := csuo.mutation.Description(); ok {
		_spec.SetField(commissionstructure.FieldDescription, field.TypeString, value)
	}
	if value, ok := csuo.mutation.CommissionValue(); ok {
		_spec.SetField(commissionstructure.FieldCommissionValue, field.TypeString, value)
	}
	if value, ok := csuo.mutation.CommissionPercentage(); ok {
		_spec.SetField(commissionstructure.FieldCommissionPercentage, field.TypeString, value)
	}
	if csuo.mutation.ProductSellerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.RemovedProductSellerIDs(); len(nodes) > 0 && !csuo.mutation.ProductSellerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.ProductSellerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CommissionStructure{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionstructure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	csuo.mutation.done = true
	return _node, nil
}
