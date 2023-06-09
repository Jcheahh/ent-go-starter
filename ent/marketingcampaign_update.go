// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/marketingcampaign"
	"entdemo/ent/predicate"
	"entdemo/ent/product"
	"entdemo/ent/rewardtype"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MarketingCampaignUpdate is the builder for updating MarketingCampaign entities.
type MarketingCampaignUpdate struct {
	config
	hooks    []Hook
	mutation *MarketingCampaignMutation
}

// Where appends a list predicates to the MarketingCampaignUpdate builder.
func (mcu *MarketingCampaignUpdate) Where(ps ...predicate.MarketingCampaign) *MarketingCampaignUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetName sets the "name" field.
func (mcu *MarketingCampaignUpdate) SetName(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetName(s)
	return mcu
}

// SetDescription sets the "description" field.
func (mcu *MarketingCampaignUpdate) SetDescription(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetDescription(s)
	return mcu
}

// SetConsumerPurchaseValue sets the "consumerPurchaseValue" field.
func (mcu *MarketingCampaignUpdate) SetConsumerPurchaseValue(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetConsumerPurchaseValue(s)
	return mcu
}

// SetCustomerApplicationLogic sets the "customerApplicationLogic" field.
func (mcu *MarketingCampaignUpdate) SetCustomerApplicationLogic(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetCustomerApplicationLogic(s)
	return mcu
}

// SetInitialisationLogic sets the "initialisationLogic" field.
func (mcu *MarketingCampaignUpdate) SetInitialisationLogic(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetInitialisationLogic(s)
	return mcu
}

// SetStartDate sets the "startDate" field.
func (mcu *MarketingCampaignUpdate) SetStartDate(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetStartDate(s)
	return mcu
}

// SetEndDate sets the "endDate" field.
func (mcu *MarketingCampaignUpdate) SetEndDate(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetEndDate(s)
	return mcu
}

// SetDateCreated sets the "dateCreated" field.
func (mcu *MarketingCampaignUpdate) SetDateCreated(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetDateCreated(s)
	return mcu
}

// SetDateUpdated sets the "dateUpdated" field.
func (mcu *MarketingCampaignUpdate) SetDateUpdated(s string) *MarketingCampaignUpdate {
	mcu.mutation.SetDateUpdated(s)
	return mcu
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (mcu *MarketingCampaignUpdate) AddProductIDs(ids ...int) *MarketingCampaignUpdate {
	mcu.mutation.AddProductIDs(ids...)
	return mcu
}

// AddProduct adds the "product" edges to the Product entity.
func (mcu *MarketingCampaignUpdate) AddProduct(p ...*Product) *MarketingCampaignUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mcu.AddProductIDs(ids...)
}

// AddConsumerRewardIDs adds the "consumerReward" edge to the RewardType entity by IDs.
func (mcu *MarketingCampaignUpdate) AddConsumerRewardIDs(ids ...int) *MarketingCampaignUpdate {
	mcu.mutation.AddConsumerRewardIDs(ids...)
	return mcu
}

// AddConsumerReward adds the "consumerReward" edges to the RewardType entity.
func (mcu *MarketingCampaignUpdate) AddConsumerReward(r ...*RewardType) *MarketingCampaignUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mcu.AddConsumerRewardIDs(ids...)
}

// Mutation returns the MarketingCampaignMutation object of the builder.
func (mcu *MarketingCampaignUpdate) Mutation() *MarketingCampaignMutation {
	return mcu.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (mcu *MarketingCampaignUpdate) ClearProduct() *MarketingCampaignUpdate {
	mcu.mutation.ClearProduct()
	return mcu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (mcu *MarketingCampaignUpdate) RemoveProductIDs(ids ...int) *MarketingCampaignUpdate {
	mcu.mutation.RemoveProductIDs(ids...)
	return mcu
}

// RemoveProduct removes "product" edges to Product entities.
func (mcu *MarketingCampaignUpdate) RemoveProduct(p ...*Product) *MarketingCampaignUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mcu.RemoveProductIDs(ids...)
}

// ClearConsumerReward clears all "consumerReward" edges to the RewardType entity.
func (mcu *MarketingCampaignUpdate) ClearConsumerReward() *MarketingCampaignUpdate {
	mcu.mutation.ClearConsumerReward()
	return mcu
}

// RemoveConsumerRewardIDs removes the "consumerReward" edge to RewardType entities by IDs.
func (mcu *MarketingCampaignUpdate) RemoveConsumerRewardIDs(ids ...int) *MarketingCampaignUpdate {
	mcu.mutation.RemoveConsumerRewardIDs(ids...)
	return mcu
}

// RemoveConsumerReward removes "consumerReward" edges to RewardType entities.
func (mcu *MarketingCampaignUpdate) RemoveConsumerReward(r ...*RewardType) *MarketingCampaignUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mcu.RemoveConsumerRewardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MarketingCampaignUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MarketingCampaignMutation](ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MarketingCampaignUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MarketingCampaignUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MarketingCampaignUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcu *MarketingCampaignUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(marketingcampaign.Table, marketingcampaign.Columns, sqlgraph.NewFieldSpec(marketingcampaign.FieldID, field.TypeInt))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.Name(); ok {
		_spec.SetField(marketingcampaign.FieldName, field.TypeString, value)
	}
	if value, ok := mcu.mutation.Description(); ok {
		_spec.SetField(marketingcampaign.FieldDescription, field.TypeString, value)
	}
	if value, ok := mcu.mutation.ConsumerPurchaseValue(); ok {
		_spec.SetField(marketingcampaign.FieldConsumerPurchaseValue, field.TypeString, value)
	}
	if value, ok := mcu.mutation.CustomerApplicationLogic(); ok {
		_spec.SetField(marketingcampaign.FieldCustomerApplicationLogic, field.TypeString, value)
	}
	if value, ok := mcu.mutation.InitialisationLogic(); ok {
		_spec.SetField(marketingcampaign.FieldInitialisationLogic, field.TypeString, value)
	}
	if value, ok := mcu.mutation.StartDate(); ok {
		_spec.SetField(marketingcampaign.FieldStartDate, field.TypeString, value)
	}
	if value, ok := mcu.mutation.EndDate(); ok {
		_spec.SetField(marketingcampaign.FieldEndDate, field.TypeString, value)
	}
	if value, ok := mcu.mutation.DateCreated(); ok {
		_spec.SetField(marketingcampaign.FieldDateCreated, field.TypeString, value)
	}
	if value, ok := mcu.mutation.DateUpdated(); ok {
		_spec.SetField(marketingcampaign.FieldDateUpdated, field.TypeString, value)
	}
	if mcu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedProductIDs(); len(nodes) > 0 && !mcu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.ConsumerRewardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedConsumerRewardIDs(); len(nodes) > 0 && !mcu.mutation.ConsumerRewardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.ConsumerRewardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{marketingcampaign.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MarketingCampaignUpdateOne is the builder for updating a single MarketingCampaign entity.
type MarketingCampaignUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MarketingCampaignMutation
}

// SetName sets the "name" field.
func (mcuo *MarketingCampaignUpdateOne) SetName(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetName(s)
	return mcuo
}

// SetDescription sets the "description" field.
func (mcuo *MarketingCampaignUpdateOne) SetDescription(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetDescription(s)
	return mcuo
}

// SetConsumerPurchaseValue sets the "consumerPurchaseValue" field.
func (mcuo *MarketingCampaignUpdateOne) SetConsumerPurchaseValue(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetConsumerPurchaseValue(s)
	return mcuo
}

// SetCustomerApplicationLogic sets the "customerApplicationLogic" field.
func (mcuo *MarketingCampaignUpdateOne) SetCustomerApplicationLogic(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetCustomerApplicationLogic(s)
	return mcuo
}

// SetInitialisationLogic sets the "initialisationLogic" field.
func (mcuo *MarketingCampaignUpdateOne) SetInitialisationLogic(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetInitialisationLogic(s)
	return mcuo
}

// SetStartDate sets the "startDate" field.
func (mcuo *MarketingCampaignUpdateOne) SetStartDate(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetStartDate(s)
	return mcuo
}

// SetEndDate sets the "endDate" field.
func (mcuo *MarketingCampaignUpdateOne) SetEndDate(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetEndDate(s)
	return mcuo
}

// SetDateCreated sets the "dateCreated" field.
func (mcuo *MarketingCampaignUpdateOne) SetDateCreated(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetDateCreated(s)
	return mcuo
}

// SetDateUpdated sets the "dateUpdated" field.
func (mcuo *MarketingCampaignUpdateOne) SetDateUpdated(s string) *MarketingCampaignUpdateOne {
	mcuo.mutation.SetDateUpdated(s)
	return mcuo
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (mcuo *MarketingCampaignUpdateOne) AddProductIDs(ids ...int) *MarketingCampaignUpdateOne {
	mcuo.mutation.AddProductIDs(ids...)
	return mcuo
}

// AddProduct adds the "product" edges to the Product entity.
func (mcuo *MarketingCampaignUpdateOne) AddProduct(p ...*Product) *MarketingCampaignUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mcuo.AddProductIDs(ids...)
}

// AddConsumerRewardIDs adds the "consumerReward" edge to the RewardType entity by IDs.
func (mcuo *MarketingCampaignUpdateOne) AddConsumerRewardIDs(ids ...int) *MarketingCampaignUpdateOne {
	mcuo.mutation.AddConsumerRewardIDs(ids...)
	return mcuo
}

// AddConsumerReward adds the "consumerReward" edges to the RewardType entity.
func (mcuo *MarketingCampaignUpdateOne) AddConsumerReward(r ...*RewardType) *MarketingCampaignUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mcuo.AddConsumerRewardIDs(ids...)
}

// Mutation returns the MarketingCampaignMutation object of the builder.
func (mcuo *MarketingCampaignUpdateOne) Mutation() *MarketingCampaignMutation {
	return mcuo.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (mcuo *MarketingCampaignUpdateOne) ClearProduct() *MarketingCampaignUpdateOne {
	mcuo.mutation.ClearProduct()
	return mcuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (mcuo *MarketingCampaignUpdateOne) RemoveProductIDs(ids ...int) *MarketingCampaignUpdateOne {
	mcuo.mutation.RemoveProductIDs(ids...)
	return mcuo
}

// RemoveProduct removes "product" edges to Product entities.
func (mcuo *MarketingCampaignUpdateOne) RemoveProduct(p ...*Product) *MarketingCampaignUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mcuo.RemoveProductIDs(ids...)
}

// ClearConsumerReward clears all "consumerReward" edges to the RewardType entity.
func (mcuo *MarketingCampaignUpdateOne) ClearConsumerReward() *MarketingCampaignUpdateOne {
	mcuo.mutation.ClearConsumerReward()
	return mcuo
}

// RemoveConsumerRewardIDs removes the "consumerReward" edge to RewardType entities by IDs.
func (mcuo *MarketingCampaignUpdateOne) RemoveConsumerRewardIDs(ids ...int) *MarketingCampaignUpdateOne {
	mcuo.mutation.RemoveConsumerRewardIDs(ids...)
	return mcuo
}

// RemoveConsumerReward removes "consumerReward" edges to RewardType entities.
func (mcuo *MarketingCampaignUpdateOne) RemoveConsumerReward(r ...*RewardType) *MarketingCampaignUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mcuo.RemoveConsumerRewardIDs(ids...)
}

// Where appends a list predicates to the MarketingCampaignUpdate builder.
func (mcuo *MarketingCampaignUpdateOne) Where(ps ...predicate.MarketingCampaign) *MarketingCampaignUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MarketingCampaignUpdateOne) Select(field string, fields ...string) *MarketingCampaignUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MarketingCampaign entity.
func (mcuo *MarketingCampaignUpdateOne) Save(ctx context.Context) (*MarketingCampaign, error) {
	return withHooks[*MarketingCampaign, MarketingCampaignMutation](ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MarketingCampaignUpdateOne) SaveX(ctx context.Context) *MarketingCampaign {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MarketingCampaignUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MarketingCampaignUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcuo *MarketingCampaignUpdateOne) sqlSave(ctx context.Context) (_node *MarketingCampaign, err error) {
	_spec := sqlgraph.NewUpdateSpec(marketingcampaign.Table, marketingcampaign.Columns, sqlgraph.NewFieldSpec(marketingcampaign.FieldID, field.TypeInt))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MarketingCampaign.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, marketingcampaign.FieldID)
		for _, f := range fields {
			if !marketingcampaign.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != marketingcampaign.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.Name(); ok {
		_spec.SetField(marketingcampaign.FieldName, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.Description(); ok {
		_spec.SetField(marketingcampaign.FieldDescription, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.ConsumerPurchaseValue(); ok {
		_spec.SetField(marketingcampaign.FieldConsumerPurchaseValue, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.CustomerApplicationLogic(); ok {
		_spec.SetField(marketingcampaign.FieldCustomerApplicationLogic, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.InitialisationLogic(); ok {
		_spec.SetField(marketingcampaign.FieldInitialisationLogic, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.StartDate(); ok {
		_spec.SetField(marketingcampaign.FieldStartDate, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.EndDate(); ok {
		_spec.SetField(marketingcampaign.FieldEndDate, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.DateCreated(); ok {
		_spec.SetField(marketingcampaign.FieldDateCreated, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.DateUpdated(); ok {
		_spec.SetField(marketingcampaign.FieldDateUpdated, field.TypeString, value)
	}
	if mcuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !mcuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ProductTable,
			Columns: []string{marketingcampaign.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.ConsumerRewardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedConsumerRewardIDs(); len(nodes) > 0 && !mcuo.mutation.ConsumerRewardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.ConsumerRewardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   marketingcampaign.ConsumerRewardTable,
			Columns: []string{marketingcampaign.ConsumerRewardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rewardtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MarketingCampaign{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{marketingcampaign.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}
