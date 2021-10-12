// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devroute/ent/developer"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeveloperCreate is the builder for creating a Developer entity.
type DeveloperCreate struct {
	config
	mutation *DeveloperMutation
	hooks    []Hook
}

// SetLastName sets the "last_name" field.
func (dc *DeveloperCreate) SetLastName(s string) *DeveloperCreate {
	dc.mutation.SetLastName(s)
	return dc
}

// SetFirstName sets the "first_name" field.
func (dc *DeveloperCreate) SetFirstName(s string) *DeveloperCreate {
	dc.mutation.SetFirstName(s)
	return dc
}

// SetLastNameFurigana sets the "last_name_furigana" field.
func (dc *DeveloperCreate) SetLastNameFurigana(s string) *DeveloperCreate {
	dc.mutation.SetLastNameFurigana(s)
	return dc
}

// SetFirstNameFurigana sets the "first_name_furigana" field.
func (dc *DeveloperCreate) SetFirstNameFurigana(s string) *DeveloperCreate {
	dc.mutation.SetFirstNameFurigana(s)
	return dc
}

// SetProfileName sets the "profile_name" field.
func (dc *DeveloperCreate) SetProfileName(s string) *DeveloperCreate {
	dc.mutation.SetProfileName(s)
	return dc
}

// SetIconURL sets the "icon_url" field.
func (dc *DeveloperCreate) SetIconURL(s string) *DeveloperCreate {
	dc.mutation.SetIconURL(s)
	return dc
}

// SetGender sets the "gender" field.
func (dc *DeveloperCreate) SetGender(d developer.Gender) *DeveloperCreate {
	dc.mutation.SetGender(d)
	return dc
}

// Mutation returns the DeveloperMutation object of the builder.
func (dc *DeveloperCreate) Mutation() *DeveloperMutation {
	return dc.mutation
}

// Save creates the Developer in the database.
func (dc *DeveloperCreate) Save(ctx context.Context) (*Developer, error) {
	var (
		err  error
		node *Developer
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeveloperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeveloperCreate) SaveX(ctx context.Context) *Developer {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeveloperCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeveloperCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeveloperCreate) check() error {
	if _, ok := dc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "last_name"`)}
	}
	if v, ok := dc.mutation.LastName(); ok {
		if err := developer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "last_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "first_name"`)}
	}
	if v, ok := dc.mutation.FirstName(); ok {
		if err := developer.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "first_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.LastNameFurigana(); !ok {
		return &ValidationError{Name: "last_name_furigana", err: errors.New(`ent: missing required field "last_name_furigana"`)}
	}
	if v, ok := dc.mutation.LastNameFurigana(); ok {
		if err := developer.LastNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "last_name_furigana", err: fmt.Errorf(`ent: validator failed for field "last_name_furigana": %w`, err)}
		}
	}
	if _, ok := dc.mutation.FirstNameFurigana(); !ok {
		return &ValidationError{Name: "first_name_furigana", err: errors.New(`ent: missing required field "first_name_furigana"`)}
	}
	if v, ok := dc.mutation.FirstNameFurigana(); ok {
		if err := developer.FirstNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "first_name_furigana", err: fmt.Errorf(`ent: validator failed for field "first_name_furigana": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ProfileName(); !ok {
		return &ValidationError{Name: "profile_name", err: errors.New(`ent: missing required field "profile_name"`)}
	}
	if v, ok := dc.mutation.ProfileName(); ok {
		if err := developer.ProfileNameValidator(v); err != nil {
			return &ValidationError{Name: "profile_name", err: fmt.Errorf(`ent: validator failed for field "profile_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.IconURL(); !ok {
		return &ValidationError{Name: "icon_url", err: errors.New(`ent: missing required field "icon_url"`)}
	}
	if v, ok := dc.mutation.IconURL(); ok {
		if err := developer.IconURLValidator(v); err != nil {
			return &ValidationError{Name: "icon_url", err: fmt.Errorf(`ent: validator failed for field "icon_url": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "gender"`)}
	}
	if v, ok := dc.mutation.Gender(); ok {
		if err := developer.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "gender": %w`, err)}
		}
	}
	return nil
}

func (dc *DeveloperCreate) sqlSave(ctx context.Context) (*Developer, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DeveloperCreate) createSpec() (*Developer, *sqlgraph.CreateSpec) {
	var (
		_node = &Developer{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: developer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: developer.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := dc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := dc.mutation.LastNameFurigana(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastNameFurigana,
		})
		_node.LastNameFurigana = value
	}
	if value, ok := dc.mutation.FirstNameFurigana(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstNameFurigana,
		})
		_node.FirstNameFurigana = value
	}
	if value, ok := dc.mutation.ProfileName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldProfileName,
		})
		_node.ProfileName = value
	}
	if value, ok := dc.mutation.IconURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldIconURL,
		})
		_node.IconURL = value
	}
	if value, ok := dc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: developer.FieldGender,
		})
		_node.Gender = value
	}
	return _node, _spec
}

// DeveloperCreateBulk is the builder for creating many Developer entities in bulk.
type DeveloperCreateBulk struct {
	config
	builders []*DeveloperCreate
}

// Save creates the Developer entities in the database.
func (dcb *DeveloperCreateBulk) Save(ctx context.Context) ([]*Developer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Developer, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeveloperMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeveloperCreateBulk) SaveX(ctx context.Context) []*Developer {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeveloperCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeveloperCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}