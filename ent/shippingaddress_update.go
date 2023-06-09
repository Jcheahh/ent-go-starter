// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/shippingaddress"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShippingAddressUpdate is the builder for updating ShippingAddress entities.
type ShippingAddressUpdate struct {
	config
	hooks    []Hook
	mutation *ShippingAddressMutation
}

// Where appends a list predicates to the ShippingAddressUpdate builder.
func (sau *ShippingAddressUpdate) Where(ps ...predicate.ShippingAddress) *ShippingAddressUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetName sets the "name" field.
func (sau *ShippingAddressUpdate) SetName(s string) *ShippingAddressUpdate {
	sau.mutation.SetName(s)
	return sau
}

// SetAddress sets the "address" field.
func (sau *ShippingAddressUpdate) SetAddress(s string) *ShippingAddressUpdate {
	sau.mutation.SetAddress(s)
	return sau
}

// SetCity sets the "city" field.
func (sau *ShippingAddressUpdate) SetCity(s string) *ShippingAddressUpdate {
	sau.mutation.SetCity(s)
	return sau
}

// SetState sets the "state" field.
func (sau *ShippingAddressUpdate) SetState(s string) *ShippingAddressUpdate {
	sau.mutation.SetState(s)
	return sau
}

// SetZip sets the "zip" field.
func (sau *ShippingAddressUpdate) SetZip(s string) *ShippingAddressUpdate {
	sau.mutation.SetZip(s)
	return sau
}

// SetCountry sets the "country" field.
func (sau *ShippingAddressUpdate) SetCountry(s string) *ShippingAddressUpdate {
	sau.mutation.SetCountry(s)
	return sau
}

// SetDateUpdated sets the "dateUpdated" field.
func (sau *ShippingAddressUpdate) SetDateUpdated(s string) *ShippingAddressUpdate {
	sau.mutation.SetDateUpdated(s)
	return sau
}

// SetNillableDateUpdated sets the "dateUpdated" field if the given value is not nil.
func (sau *ShippingAddressUpdate) SetNillableDateUpdated(s *string) *ShippingAddressUpdate {
	if s != nil {
		sau.SetDateUpdated(*s)
	}
	return sau
}

// Mutation returns the ShippingAddressMutation object of the builder.
func (sau *ShippingAddressUpdate) Mutation() *ShippingAddressMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *ShippingAddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ShippingAddressMutation](ctx, sau.sqlSave, sau.mutation, sau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sau *ShippingAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *ShippingAddressUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *ShippingAddressUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sau *ShippingAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(shippingaddress.Table, shippingaddress.Columns, sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt))
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.Name(); ok {
		_spec.SetField(shippingaddress.FieldName, field.TypeString, value)
	}
	if value, ok := sau.mutation.Address(); ok {
		_spec.SetField(shippingaddress.FieldAddress, field.TypeString, value)
	}
	if value, ok := sau.mutation.City(); ok {
		_spec.SetField(shippingaddress.FieldCity, field.TypeString, value)
	}
	if value, ok := sau.mutation.State(); ok {
		_spec.SetField(shippingaddress.FieldState, field.TypeString, value)
	}
	if value, ok := sau.mutation.Zip(); ok {
		_spec.SetField(shippingaddress.FieldZip, field.TypeString, value)
	}
	if value, ok := sau.mutation.Country(); ok {
		_spec.SetField(shippingaddress.FieldCountry, field.TypeString, value)
	}
	if value, ok := sau.mutation.DateUpdated(); ok {
		_spec.SetField(shippingaddress.FieldDateUpdated, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shippingaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sau.mutation.done = true
	return n, nil
}

// ShippingAddressUpdateOne is the builder for updating a single ShippingAddress entity.
type ShippingAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShippingAddressMutation
}

// SetName sets the "name" field.
func (sauo *ShippingAddressUpdateOne) SetName(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetName(s)
	return sauo
}

// SetAddress sets the "address" field.
func (sauo *ShippingAddressUpdateOne) SetAddress(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetAddress(s)
	return sauo
}

// SetCity sets the "city" field.
func (sauo *ShippingAddressUpdateOne) SetCity(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetCity(s)
	return sauo
}

// SetState sets the "state" field.
func (sauo *ShippingAddressUpdateOne) SetState(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetState(s)
	return sauo
}

// SetZip sets the "zip" field.
func (sauo *ShippingAddressUpdateOne) SetZip(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetZip(s)
	return sauo
}

// SetCountry sets the "country" field.
func (sauo *ShippingAddressUpdateOne) SetCountry(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetCountry(s)
	return sauo
}

// SetDateUpdated sets the "dateUpdated" field.
func (sauo *ShippingAddressUpdateOne) SetDateUpdated(s string) *ShippingAddressUpdateOne {
	sauo.mutation.SetDateUpdated(s)
	return sauo
}

// SetNillableDateUpdated sets the "dateUpdated" field if the given value is not nil.
func (sauo *ShippingAddressUpdateOne) SetNillableDateUpdated(s *string) *ShippingAddressUpdateOne {
	if s != nil {
		sauo.SetDateUpdated(*s)
	}
	return sauo
}

// Mutation returns the ShippingAddressMutation object of the builder.
func (sauo *ShippingAddressUpdateOne) Mutation() *ShippingAddressMutation {
	return sauo.mutation
}

// Where appends a list predicates to the ShippingAddressUpdate builder.
func (sauo *ShippingAddressUpdateOne) Where(ps ...predicate.ShippingAddress) *ShippingAddressUpdateOne {
	sauo.mutation.Where(ps...)
	return sauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *ShippingAddressUpdateOne) Select(field string, fields ...string) *ShippingAddressUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated ShippingAddress entity.
func (sauo *ShippingAddressUpdateOne) Save(ctx context.Context) (*ShippingAddress, error) {
	return withHooks[*ShippingAddress, ShippingAddressMutation](ctx, sauo.sqlSave, sauo.mutation, sauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *ShippingAddressUpdateOne) SaveX(ctx context.Context) *ShippingAddress {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *ShippingAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *ShippingAddressUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sauo *ShippingAddressUpdateOne) sqlSave(ctx context.Context) (_node *ShippingAddress, err error) {
	_spec := sqlgraph.NewUpdateSpec(shippingaddress.Table, shippingaddress.Columns, sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt))
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShippingAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shippingaddress.FieldID)
		for _, f := range fields {
			if !shippingaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shippingaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.Name(); ok {
		_spec.SetField(shippingaddress.FieldName, field.TypeString, value)
	}
	if value, ok := sauo.mutation.Address(); ok {
		_spec.SetField(shippingaddress.FieldAddress, field.TypeString, value)
	}
	if value, ok := sauo.mutation.City(); ok {
		_spec.SetField(shippingaddress.FieldCity, field.TypeString, value)
	}
	if value, ok := sauo.mutation.State(); ok {
		_spec.SetField(shippingaddress.FieldState, field.TypeString, value)
	}
	if value, ok := sauo.mutation.Zip(); ok {
		_spec.SetField(shippingaddress.FieldZip, field.TypeString, value)
	}
	if value, ok := sauo.mutation.Country(); ok {
		_spec.SetField(shippingaddress.FieldCountry, field.TypeString, value)
	}
	if value, ok := sauo.mutation.DateUpdated(); ok {
		_spec.SetField(shippingaddress.FieldDateUpdated, field.TypeString, value)
	}
	_node = &ShippingAddress{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shippingaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sauo.mutation.done = true
	return _node, nil
}
