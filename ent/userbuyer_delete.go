// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/userbuyer"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBuyerDelete is the builder for deleting a UserBuyer entity.
type UserBuyerDelete struct {
	config
	hooks    []Hook
	mutation *UserBuyerMutation
}

// Where appends a list predicates to the UserBuyerDelete builder.
func (ubd *UserBuyerDelete) Where(ps ...predicate.UserBuyer) *UserBuyerDelete {
	ubd.mutation.Where(ps...)
	return ubd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ubd *UserBuyerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserBuyerMutation](ctx, ubd.sqlExec, ubd.mutation, ubd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ubd *UserBuyerDelete) ExecX(ctx context.Context) int {
	n, err := ubd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ubd *UserBuyerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userbuyer.Table, sqlgraph.NewFieldSpec(userbuyer.FieldID, field.TypeInt))
	if ps := ubd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ubd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ubd.mutation.done = true
	return affected, err
}

// UserBuyerDeleteOne is the builder for deleting a single UserBuyer entity.
type UserBuyerDeleteOne struct {
	ubd *UserBuyerDelete
}

// Where appends a list predicates to the UserBuyerDelete builder.
func (ubdo *UserBuyerDeleteOne) Where(ps ...predicate.UserBuyer) *UserBuyerDeleteOne {
	ubdo.ubd.mutation.Where(ps...)
	return ubdo
}

// Exec executes the deletion query.
func (ubdo *UserBuyerDeleteOne) Exec(ctx context.Context) error {
	n, err := ubdo.ubd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userbuyer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ubdo *UserBuyerDeleteOne) ExecX(ctx context.Context) {
	if err := ubdo.Exec(ctx); err != nil {
		panic(err)
	}
}
