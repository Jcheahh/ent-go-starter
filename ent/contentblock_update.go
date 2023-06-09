// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/contentblock"
	"entdemo/ent/image"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContentBlockUpdate is the builder for updating ContentBlock entities.
type ContentBlockUpdate struct {
	config
	hooks    []Hook
	mutation *ContentBlockMutation
}

// Where appends a list predicates to the ContentBlockUpdate builder.
func (cbu *ContentBlockUpdate) Where(ps ...predicate.ContentBlock) *ContentBlockUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetPrimaryMessage sets the "primaryMessage" field.
func (cbu *ContentBlockUpdate) SetPrimaryMessage(s string) *ContentBlockUpdate {
	cbu.mutation.SetPrimaryMessage(s)
	return cbu
}

// SetSecondaryMessage sets the "secondaryMessage" field.
func (cbu *ContentBlockUpdate) SetSecondaryMessage(s string) *ContentBlockUpdate {
	cbu.mutation.SetSecondaryMessage(s)
	return cbu
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (cbu *ContentBlockUpdate) AddImageIDs(ids ...int) *ContentBlockUpdate {
	cbu.mutation.AddImageIDs(ids...)
	return cbu
}

// AddImage adds the "image" edges to the Image entity.
func (cbu *ContentBlockUpdate) AddImage(i ...*Image) *ContentBlockUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cbu.AddImageIDs(ids...)
}

// Mutation returns the ContentBlockMutation object of the builder.
func (cbu *ContentBlockUpdate) Mutation() *ContentBlockMutation {
	return cbu.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (cbu *ContentBlockUpdate) ClearImage() *ContentBlockUpdate {
	cbu.mutation.ClearImage()
	return cbu
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (cbu *ContentBlockUpdate) RemoveImageIDs(ids ...int) *ContentBlockUpdate {
	cbu.mutation.RemoveImageIDs(ids...)
	return cbu
}

// RemoveImage removes "image" edges to Image entities.
func (cbu *ContentBlockUpdate) RemoveImage(i ...*Image) *ContentBlockUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cbu.RemoveImageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *ContentBlockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ContentBlockMutation](ctx, cbu.sqlSave, cbu.mutation, cbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *ContentBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *ContentBlockUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *ContentBlockUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cbu *ContentBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contentblock.Table, contentblock.Columns, sqlgraph.NewFieldSpec(contentblock.FieldID, field.TypeInt))
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.PrimaryMessage(); ok {
		_spec.SetField(contentblock.FieldPrimaryMessage, field.TypeString, value)
	}
	if value, ok := cbu.mutation.SecondaryMessage(); ok {
		_spec.SetField(contentblock.FieldSecondaryMessage, field.TypeString, value)
	}
	if cbu.mutation.ImageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.RemovedImageIDs(); len(nodes) > 0 && !cbu.mutation.ImageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.ImageIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contentblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cbu.mutation.done = true
	return n, nil
}

// ContentBlockUpdateOne is the builder for updating a single ContentBlock entity.
type ContentBlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContentBlockMutation
}

// SetPrimaryMessage sets the "primaryMessage" field.
func (cbuo *ContentBlockUpdateOne) SetPrimaryMessage(s string) *ContentBlockUpdateOne {
	cbuo.mutation.SetPrimaryMessage(s)
	return cbuo
}

// SetSecondaryMessage sets the "secondaryMessage" field.
func (cbuo *ContentBlockUpdateOne) SetSecondaryMessage(s string) *ContentBlockUpdateOne {
	cbuo.mutation.SetSecondaryMessage(s)
	return cbuo
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (cbuo *ContentBlockUpdateOne) AddImageIDs(ids ...int) *ContentBlockUpdateOne {
	cbuo.mutation.AddImageIDs(ids...)
	return cbuo
}

// AddImage adds the "image" edges to the Image entity.
func (cbuo *ContentBlockUpdateOne) AddImage(i ...*Image) *ContentBlockUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cbuo.AddImageIDs(ids...)
}

// Mutation returns the ContentBlockMutation object of the builder.
func (cbuo *ContentBlockUpdateOne) Mutation() *ContentBlockMutation {
	return cbuo.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (cbuo *ContentBlockUpdateOne) ClearImage() *ContentBlockUpdateOne {
	cbuo.mutation.ClearImage()
	return cbuo
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (cbuo *ContentBlockUpdateOne) RemoveImageIDs(ids ...int) *ContentBlockUpdateOne {
	cbuo.mutation.RemoveImageIDs(ids...)
	return cbuo
}

// RemoveImage removes "image" edges to Image entities.
func (cbuo *ContentBlockUpdateOne) RemoveImage(i ...*Image) *ContentBlockUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cbuo.RemoveImageIDs(ids...)
}

// Where appends a list predicates to the ContentBlockUpdate builder.
func (cbuo *ContentBlockUpdateOne) Where(ps ...predicate.ContentBlock) *ContentBlockUpdateOne {
	cbuo.mutation.Where(ps...)
	return cbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *ContentBlockUpdateOne) Select(field string, fields ...string) *ContentBlockUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated ContentBlock entity.
func (cbuo *ContentBlockUpdateOne) Save(ctx context.Context) (*ContentBlock, error) {
	return withHooks[*ContentBlock, ContentBlockMutation](ctx, cbuo.sqlSave, cbuo.mutation, cbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *ContentBlockUpdateOne) SaveX(ctx context.Context) *ContentBlock {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *ContentBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *ContentBlockUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cbuo *ContentBlockUpdateOne) sqlSave(ctx context.Context) (_node *ContentBlock, err error) {
	_spec := sqlgraph.NewUpdateSpec(contentblock.Table, contentblock.Columns, sqlgraph.NewFieldSpec(contentblock.FieldID, field.TypeInt))
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContentBlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contentblock.FieldID)
		for _, f := range fields {
			if !contentblock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contentblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.PrimaryMessage(); ok {
		_spec.SetField(contentblock.FieldPrimaryMessage, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.SecondaryMessage(); ok {
		_spec.SetField(contentblock.FieldSecondaryMessage, field.TypeString, value)
	}
	if cbuo.mutation.ImageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.RemovedImageIDs(); len(nodes) > 0 && !cbuo.mutation.ImageCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.ImageIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContentBlock{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contentblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cbuo.mutation.done = true
	return _node, nil
}
