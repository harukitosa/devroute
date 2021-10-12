// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devroute/ent/developer"
	"devroute/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeveloperUpdate is the builder for updating Developer entities.
type DeveloperUpdate struct {
	config
	hooks    []Hook
	mutation *DeveloperMutation
}

// Where appends a list predicates to the DeveloperUpdate builder.
func (du *DeveloperUpdate) Where(ps ...predicate.Developer) *DeveloperUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetLastName sets the "last_name" field.
func (du *DeveloperUpdate) SetLastName(s string) *DeveloperUpdate {
	du.mutation.SetLastName(s)
	return du
}

// SetFirstName sets the "first_name" field.
func (du *DeveloperUpdate) SetFirstName(s string) *DeveloperUpdate {
	du.mutation.SetFirstName(s)
	return du
}

// SetLastNameFurigana sets the "last_name_furigana" field.
func (du *DeveloperUpdate) SetLastNameFurigana(s string) *DeveloperUpdate {
	du.mutation.SetLastNameFurigana(s)
	return du
}

// SetFirstNameFurigana sets the "first_name_furigana" field.
func (du *DeveloperUpdate) SetFirstNameFurigana(s string) *DeveloperUpdate {
	du.mutation.SetFirstNameFurigana(s)
	return du
}

// SetProfileName sets the "profile_name" field.
func (du *DeveloperUpdate) SetProfileName(s string) *DeveloperUpdate {
	du.mutation.SetProfileName(s)
	return du
}

// SetIconURL sets the "icon_url" field.
func (du *DeveloperUpdate) SetIconURL(s string) *DeveloperUpdate {
	du.mutation.SetIconURL(s)
	return du
}

// SetGender sets the "gender" field.
func (du *DeveloperUpdate) SetGender(d developer.Gender) *DeveloperUpdate {
	du.mutation.SetGender(d)
	return du
}

// Mutation returns the DeveloperMutation object of the builder.
func (du *DeveloperUpdate) Mutation() *DeveloperMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeveloperUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeveloperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeveloperUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeveloperUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeveloperUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeveloperUpdate) check() error {
	if v, ok := du.mutation.LastName(); ok {
		if err := developer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	if v, ok := du.mutation.FirstName(); ok {
		if err := developer.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := du.mutation.LastNameFurigana(); ok {
		if err := developer.LastNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "last_name_furigana", err: fmt.Errorf("ent: validator failed for field \"last_name_furigana\": %w", err)}
		}
	}
	if v, ok := du.mutation.FirstNameFurigana(); ok {
		if err := developer.FirstNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "first_name_furigana", err: fmt.Errorf("ent: validator failed for field \"first_name_furigana\": %w", err)}
		}
	}
	if v, ok := du.mutation.ProfileName(); ok {
		if err := developer.ProfileNameValidator(v); err != nil {
			return &ValidationError{Name: "profile_name", err: fmt.Errorf("ent: validator failed for field \"profile_name\": %w", err)}
		}
	}
	if v, ok := du.mutation.IconURL(); ok {
		if err := developer.IconURLValidator(v); err != nil {
			return &ValidationError{Name: "icon_url", err: fmt.Errorf("ent: validator failed for field \"icon_url\": %w", err)}
		}
	}
	if v, ok := du.mutation.Gender(); ok {
		if err := developer.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (du *DeveloperUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   developer.Table,
			Columns: developer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: developer.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastName,
		})
	}
	if value, ok := du.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstName,
		})
	}
	if value, ok := du.mutation.LastNameFurigana(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastNameFurigana,
		})
	}
	if value, ok := du.mutation.FirstNameFurigana(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstNameFurigana,
		})
	}
	if value, ok := du.mutation.ProfileName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldProfileName,
		})
	}
	if value, ok := du.mutation.IconURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldIconURL,
		})
	}
	if value, ok := du.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: developer.FieldGender,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{developer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeveloperUpdateOne is the builder for updating a single Developer entity.
type DeveloperUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeveloperMutation
}

// SetLastName sets the "last_name" field.
func (duo *DeveloperUpdateOne) SetLastName(s string) *DeveloperUpdateOne {
	duo.mutation.SetLastName(s)
	return duo
}

// SetFirstName sets the "first_name" field.
func (duo *DeveloperUpdateOne) SetFirstName(s string) *DeveloperUpdateOne {
	duo.mutation.SetFirstName(s)
	return duo
}

// SetLastNameFurigana sets the "last_name_furigana" field.
func (duo *DeveloperUpdateOne) SetLastNameFurigana(s string) *DeveloperUpdateOne {
	duo.mutation.SetLastNameFurigana(s)
	return duo
}

// SetFirstNameFurigana sets the "first_name_furigana" field.
func (duo *DeveloperUpdateOne) SetFirstNameFurigana(s string) *DeveloperUpdateOne {
	duo.mutation.SetFirstNameFurigana(s)
	return duo
}

// SetProfileName sets the "profile_name" field.
func (duo *DeveloperUpdateOne) SetProfileName(s string) *DeveloperUpdateOne {
	duo.mutation.SetProfileName(s)
	return duo
}

// SetIconURL sets the "icon_url" field.
func (duo *DeveloperUpdateOne) SetIconURL(s string) *DeveloperUpdateOne {
	duo.mutation.SetIconURL(s)
	return duo
}

// SetGender sets the "gender" field.
func (duo *DeveloperUpdateOne) SetGender(d developer.Gender) *DeveloperUpdateOne {
	duo.mutation.SetGender(d)
	return duo
}

// Mutation returns the DeveloperMutation object of the builder.
func (duo *DeveloperUpdateOne) Mutation() *DeveloperMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeveloperUpdateOne) Select(field string, fields ...string) *DeveloperUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Developer entity.
func (duo *DeveloperUpdateOne) Save(ctx context.Context) (*Developer, error) {
	var (
		err  error
		node *Developer
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeveloperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeveloperUpdateOne) SaveX(ctx context.Context) *Developer {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeveloperUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeveloperUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeveloperUpdateOne) check() error {
	if v, ok := duo.mutation.LastName(); ok {
		if err := developer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.FirstName(); ok {
		if err := developer.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.LastNameFurigana(); ok {
		if err := developer.LastNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "last_name_furigana", err: fmt.Errorf("ent: validator failed for field \"last_name_furigana\": %w", err)}
		}
	}
	if v, ok := duo.mutation.FirstNameFurigana(); ok {
		if err := developer.FirstNameFuriganaValidator(v); err != nil {
			return &ValidationError{Name: "first_name_furigana", err: fmt.Errorf("ent: validator failed for field \"first_name_furigana\": %w", err)}
		}
	}
	if v, ok := duo.mutation.ProfileName(); ok {
		if err := developer.ProfileNameValidator(v); err != nil {
			return &ValidationError{Name: "profile_name", err: fmt.Errorf("ent: validator failed for field \"profile_name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.IconURL(); ok {
		if err := developer.IconURLValidator(v); err != nil {
			return &ValidationError{Name: "icon_url", err: fmt.Errorf("ent: validator failed for field \"icon_url\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Gender(); ok {
		if err := developer.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (duo *DeveloperUpdateOne) sqlSave(ctx context.Context) (_node *Developer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   developer.Table,
			Columns: developer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: developer.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Developer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, developer.FieldID)
		for _, f := range fields {
			if !developer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != developer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastName,
		})
	}
	if value, ok := duo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstName,
		})
	}
	if value, ok := duo.mutation.LastNameFurigana(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldLastNameFurigana,
		})
	}
	if value, ok := duo.mutation.FirstNameFurigana(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldFirstNameFurigana,
		})
	}
	if value, ok := duo.mutation.ProfileName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldProfileName,
		})
	}
	if value, ok := duo.mutation.IconURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: developer.FieldIconURL,
		})
	}
	if value, ok := duo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: developer.FieldGender,
		})
	}
	_node = &Developer{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{developer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
