// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devroute/ent/recruitment"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecruitmentCreate is the builder for creating a Recruitment entity.
type RecruitmentCreate struct {
	config
	mutation *RecruitmentMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (rc *RecruitmentCreate) SetTitle(s string) *RecruitmentCreate {
	rc.mutation.SetTitle(s)
	return rc
}

// SetContent sets the "content" field.
func (rc *RecruitmentCreate) SetContent(s string) *RecruitmentCreate {
	rc.mutation.SetContent(s)
	return rc
}

// Mutation returns the RecruitmentMutation object of the builder.
func (rc *RecruitmentCreate) Mutation() *RecruitmentMutation {
	return rc.mutation
}

// Save creates the Recruitment in the database.
func (rc *RecruitmentCreate) Save(ctx context.Context) (*Recruitment, error) {
	var (
		err  error
		node *Recruitment
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecruitmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecruitmentCreate) SaveX(ctx context.Context) *Recruitment {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecruitmentCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecruitmentCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecruitmentCreate) check() error {
	if _, ok := rc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := rc.mutation.Title(); ok {
		if err := recruitment.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if v, ok := rc.mutation.Content(); ok {
		if err := recruitment.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "content": %w`, err)}
		}
	}
	return nil
}

func (rc *RecruitmentCreate) sqlSave(ctx context.Context) (*Recruitment, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RecruitmentCreate) createSpec() (*Recruitment, *sqlgraph.CreateSpec) {
	var (
		_node = &Recruitment{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recruitment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recruitment.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recruitment.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recruitment.FieldContent,
		})
		_node.Content = value
	}
	return _node, _spec
}

// RecruitmentCreateBulk is the builder for creating many Recruitment entities in bulk.
type RecruitmentCreateBulk struct {
	config
	builders []*RecruitmentCreate
}

// Save creates the Recruitment entities in the database.
func (rcb *RecruitmentCreateBulk) Save(ctx context.Context) ([]*Recruitment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recruitment, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecruitmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecruitmentCreateBulk) SaveX(ctx context.Context) []*Recruitment {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecruitmentCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecruitmentCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
