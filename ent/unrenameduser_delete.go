// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/unrenameduser"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnrenamedUserDelete is the builder for deleting a UnrenamedUser entity.
type UnrenamedUserDelete struct {
	config
	hooks    []Hook
	mutation *UnrenamedUserMutation
}

// Where appends a list predicates to the UnrenamedUserDelete builder.
func (uud *UnrenamedUserDelete) Where(ps ...predicate.UnrenamedUser) *UnrenamedUserDelete {
	uud.mutation.Where(ps...)
	return uud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uud *UnrenamedUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uud.sqlExec, uud.mutation, uud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uud *UnrenamedUserDelete) ExecX(ctx context.Context) int {
	n, err := uud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uud *UnrenamedUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(unrenameduser.Table, sqlgraph.NewFieldSpec(unrenameduser.FieldID, field.TypeInt))
	if ps := uud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uud.mutation.done = true
	return affected, err
}

// UnrenamedUserDeleteOne is the builder for deleting a single UnrenamedUser entity.
type UnrenamedUserDeleteOne struct {
	uud *UnrenamedUserDelete
}

// Where appends a list predicates to the UnrenamedUserDelete builder.
func (uudo *UnrenamedUserDeleteOne) Where(ps ...predicate.UnrenamedUser) *UnrenamedUserDeleteOne {
	uudo.uud.mutation.Where(ps...)
	return uudo
}

// Exec executes the deletion query.
func (uudo *UnrenamedUserDeleteOne) Exec(ctx context.Context) error {
	n, err := uudo.uud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{unrenameduser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uudo *UnrenamedUserDeleteOne) ExecX(ctx context.Context) {
	if err := uudo.Exec(ctx); err != nil {
		panic(err)
	}
}
