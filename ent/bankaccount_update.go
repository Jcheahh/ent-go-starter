// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/bankaccount"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BankAccountUpdate is the builder for updating BankAccount entities.
type BankAccountUpdate struct {
	config
	hooks    []Hook
	mutation *BankAccountMutation
}

// Where appends a list predicates to the BankAccountUpdate builder.
func (bau *BankAccountUpdate) Where(ps ...predicate.BankAccount) *BankAccountUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetXid sets the "xid" field.
func (bau *BankAccountUpdate) SetXid(i int) *BankAccountUpdate {
	bau.mutation.ResetXid()
	bau.mutation.SetXid(i)
	return bau
}

// AddXid adds i to the "xid" field.
func (bau *BankAccountUpdate) AddXid(i int) *BankAccountUpdate {
	bau.mutation.AddXid(i)
	return bau
}

// Mutation returns the BankAccountMutation object of the builder.
func (bau *BankAccountUpdate) Mutation() *BankAccountMutation {
	return bau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BankAccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BankAccountMutation](ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BankAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BankAccountUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BankAccountUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bau *BankAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bankaccount.Table, bankaccount.Columns, sqlgraph.NewFieldSpec(bankaccount.FieldID, field.TypeInt))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.Xid(); ok {
		_spec.SetField(bankaccount.FieldXid, field.TypeInt, value)
	}
	if value, ok := bau.mutation.AddedXid(); ok {
		_spec.AddField(bankaccount.FieldXid, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bankaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BankAccountUpdateOne is the builder for updating a single BankAccount entity.
type BankAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BankAccountMutation
}

// SetXid sets the "xid" field.
func (bauo *BankAccountUpdateOne) SetXid(i int) *BankAccountUpdateOne {
	bauo.mutation.ResetXid()
	bauo.mutation.SetXid(i)
	return bauo
}

// AddXid adds i to the "xid" field.
func (bauo *BankAccountUpdateOne) AddXid(i int) *BankAccountUpdateOne {
	bauo.mutation.AddXid(i)
	return bauo
}

// Mutation returns the BankAccountMutation object of the builder.
func (bauo *BankAccountUpdateOne) Mutation() *BankAccountMutation {
	return bauo.mutation
}

// Where appends a list predicates to the BankAccountUpdate builder.
func (bauo *BankAccountUpdateOne) Where(ps ...predicate.BankAccount) *BankAccountUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BankAccountUpdateOne) Select(field string, fields ...string) *BankAccountUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BankAccount entity.
func (bauo *BankAccountUpdateOne) Save(ctx context.Context) (*BankAccount, error) {
	return withHooks[*BankAccount, BankAccountMutation](ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BankAccountUpdateOne) SaveX(ctx context.Context) *BankAccount {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BankAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BankAccountUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bauo *BankAccountUpdateOne) sqlSave(ctx context.Context) (_node *BankAccount, err error) {
	_spec := sqlgraph.NewUpdateSpec(bankaccount.Table, bankaccount.Columns, sqlgraph.NewFieldSpec(bankaccount.FieldID, field.TypeInt))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BankAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bankaccount.FieldID)
		for _, f := range fields {
			if !bankaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bankaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.Xid(); ok {
		_spec.SetField(bankaccount.FieldXid, field.TypeInt, value)
	}
	if value, ok := bauo.mutation.AddedXid(); ok {
		_spec.AddField(bankaccount.FieldXid, field.TypeInt, value)
	}
	_node = &BankAccount{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bankaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}
