// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/predicate"
	"api/ent/segment"
	"api/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SegmentUpdate is the builder for updating Segment entities.
type SegmentUpdate struct {
	config
	hooks    []Hook
	mutation *SegmentMutation
}

// Where appends a list predicates to the SegmentUpdate builder.
func (su *SegmentUpdate) Where(ps ...predicate.Segment) *SegmentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SegmentUpdate) SetName(s string) *SegmentUpdate {
	su.mutation.SetName(s)
	return su
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (su *SegmentUpdate) AddUserIDs(ids ...int) *SegmentUpdate {
	su.mutation.AddUserIDs(ids...)
	return su
}

// AddUsers adds the "users" edges to the User entity.
func (su *SegmentUpdate) AddUsers(u ...*User) *SegmentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserIDs(ids...)
}

// Mutation returns the SegmentMutation object of the builder.
func (su *SegmentUpdate) Mutation() *SegmentMutation {
	return su.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (su *SegmentUpdate) ClearUsers() *SegmentUpdate {
	su.mutation.ClearUsers()
	return su
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (su *SegmentUpdate) RemoveUserIDs(ids ...int) *SegmentUpdate {
	su.mutation.RemoveUserIDs(ids...)
	return su
}

// RemoveUsers removes "users" edges to User entities.
func (su *SegmentUpdate) RemoveUsers(u ...*User) *SegmentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SegmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SegmentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SegmentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SegmentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SegmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(segment.Table, segment.Columns, sqlgraph.NewFieldSpec(segment.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(segment.FieldName, field.TypeString, value)
	}
	if su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUsersIDs(); len(nodes) > 0 && !su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{segment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SegmentUpdateOne is the builder for updating a single Segment entity.
type SegmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SegmentMutation
}

// SetName sets the "name" field.
func (suo *SegmentUpdateOne) SetName(s string) *SegmentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (suo *SegmentUpdateOne) AddUserIDs(ids ...int) *SegmentUpdateOne {
	suo.mutation.AddUserIDs(ids...)
	return suo
}

// AddUsers adds the "users" edges to the User entity.
func (suo *SegmentUpdateOne) AddUsers(u ...*User) *SegmentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserIDs(ids...)
}

// Mutation returns the SegmentMutation object of the builder.
func (suo *SegmentUpdateOne) Mutation() *SegmentMutation {
	return suo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (suo *SegmentUpdateOne) ClearUsers() *SegmentUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (suo *SegmentUpdateOne) RemoveUserIDs(ids ...int) *SegmentUpdateOne {
	suo.mutation.RemoveUserIDs(ids...)
	return suo
}

// RemoveUsers removes "users" edges to User entities.
func (suo *SegmentUpdateOne) RemoveUsers(u ...*User) *SegmentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the SegmentUpdate builder.
func (suo *SegmentUpdateOne) Where(ps ...predicate.Segment) *SegmentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SegmentUpdateOne) Select(field string, fields ...string) *SegmentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Segment entity.
func (suo *SegmentUpdateOne) Save(ctx context.Context) (*Segment, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SegmentUpdateOne) SaveX(ctx context.Context) *Segment {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SegmentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SegmentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SegmentUpdateOne) sqlSave(ctx context.Context) (_node *Segment, err error) {
	_spec := sqlgraph.NewUpdateSpec(segment.Table, segment.Columns, sqlgraph.NewFieldSpec(segment.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Segment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, segment.FieldID)
		for _, f := range fields {
			if !segment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != segment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(segment.FieldName, field.TypeString, value)
	}
	if suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   segment.UsersTable,
			Columns: segment.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Segment{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{segment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}