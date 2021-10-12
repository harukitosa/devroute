// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devroute/ent/companyuser"
	"devroute/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyUserDelete is the builder for deleting a CompanyUser entity.
type CompanyUserDelete struct {
	config
	hooks    []Hook
	mutation *CompanyUserMutation
}

// Where appends a list predicates to the CompanyUserDelete builder.
func (cud *CompanyUserDelete) Where(ps ...predicate.CompanyUser) *CompanyUserDelete {
	cud.mutation.Where(ps...)
	return cud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cud *CompanyUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cud.hooks) == 0 {
		affected, err = cud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cud.mutation = mutation
			affected, err = cud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cud.hooks) - 1; i >= 0; i-- {
			if cud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cud *CompanyUserDelete) ExecX(ctx context.Context) int {
	n, err := cud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cud *CompanyUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: companyuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: companyuser.FieldID,
			},
		},
	}
	if ps := cud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cud.driver, _spec)
}

// CompanyUserDeleteOne is the builder for deleting a single CompanyUser entity.
type CompanyUserDeleteOne struct {
	cud *CompanyUserDelete
}

// Exec executes the deletion query.
func (cudo *CompanyUserDeleteOne) Exec(ctx context.Context) error {
	n, err := cudo.cud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{companyuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cudo *CompanyUserDeleteOne) ExecX(ctx context.Context) {
	cudo.cud.ExecX(ctx)
}