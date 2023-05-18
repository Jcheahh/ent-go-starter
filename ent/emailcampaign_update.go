// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/emailcampaign"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmailCampaignUpdate is the builder for updating EmailCampaign entities.
type EmailCampaignUpdate struct {
	config
	hooks    []Hook
	mutation *EmailCampaignMutation
}

// Where appends a list predicates to the EmailCampaignUpdate builder.
func (ecu *EmailCampaignUpdate) Where(ps ...predicate.EmailCampaign) *EmailCampaignUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetXid sets the "xid" field.
func (ecu *EmailCampaignUpdate) SetXid(i int) *EmailCampaignUpdate {
	ecu.mutation.ResetXid()
	ecu.mutation.SetXid(i)
	return ecu
}

// AddXid adds i to the "xid" field.
func (ecu *EmailCampaignUpdate) AddXid(i int) *EmailCampaignUpdate {
	ecu.mutation.AddXid(i)
	return ecu
}

// Mutation returns the EmailCampaignMutation object of the builder.
func (ecu *EmailCampaignUpdate) Mutation() *EmailCampaignMutation {
	return ecu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EmailCampaignUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EmailCampaignMutation](ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EmailCampaignUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EmailCampaignUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EmailCampaignUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecu *EmailCampaignUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(emailcampaign.Table, emailcampaign.Columns, sqlgraph.NewFieldSpec(emailcampaign.FieldID, field.TypeInt))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Xid(); ok {
		_spec.SetField(emailcampaign.FieldXid, field.TypeInt, value)
	}
	if value, ok := ecu.mutation.AddedXid(); ok {
		_spec.AddField(emailcampaign.FieldXid, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailcampaign.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EmailCampaignUpdateOne is the builder for updating a single EmailCampaign entity.
type EmailCampaignUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailCampaignMutation
}

// SetXid sets the "xid" field.
func (ecuo *EmailCampaignUpdateOne) SetXid(i int) *EmailCampaignUpdateOne {
	ecuo.mutation.ResetXid()
	ecuo.mutation.SetXid(i)
	return ecuo
}

// AddXid adds i to the "xid" field.
func (ecuo *EmailCampaignUpdateOne) AddXid(i int) *EmailCampaignUpdateOne {
	ecuo.mutation.AddXid(i)
	return ecuo
}

// Mutation returns the EmailCampaignMutation object of the builder.
func (ecuo *EmailCampaignUpdateOne) Mutation() *EmailCampaignMutation {
	return ecuo.mutation
}

// Where appends a list predicates to the EmailCampaignUpdate builder.
func (ecuo *EmailCampaignUpdateOne) Where(ps ...predicate.EmailCampaign) *EmailCampaignUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EmailCampaignUpdateOne) Select(field string, fields ...string) *EmailCampaignUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EmailCampaign entity.
func (ecuo *EmailCampaignUpdateOne) Save(ctx context.Context) (*EmailCampaign, error) {
	return withHooks[*EmailCampaign, EmailCampaignMutation](ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EmailCampaignUpdateOne) SaveX(ctx context.Context) *EmailCampaign {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EmailCampaignUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EmailCampaignUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecuo *EmailCampaignUpdateOne) sqlSave(ctx context.Context) (_node *EmailCampaign, err error) {
	_spec := sqlgraph.NewUpdateSpec(emailcampaign.Table, emailcampaign.Columns, sqlgraph.NewFieldSpec(emailcampaign.FieldID, field.TypeInt))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailCampaign.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailcampaign.FieldID)
		for _, f := range fields {
			if !emailcampaign.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emailcampaign.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Xid(); ok {
		_spec.SetField(emailcampaign.FieldXid, field.TypeInt, value)
	}
	if value, ok := ecuo.mutation.AddedXid(); ok {
		_spec.AddField(emailcampaign.FieldXid, field.TypeInt, value)
	}
	_node = &EmailCampaign{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailcampaign.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
