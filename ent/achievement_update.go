// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devroute/ent/achievement"
	"devroute/ent/predicate"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AchievementUpdate is the builder for updating Achievement entities.
type AchievementUpdate struct {
	config
	hooks    []Hook
	mutation *AchievementMutation
}

// Where appends a list predicates to the AchievementUpdate builder.
func (au *AchievementUpdate) Where(ps ...predicate.Achievement) *AchievementUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetTitle sets the "title" field.
func (au *AchievementUpdate) SetTitle(s string) *AchievementUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetStartTime sets the "start_time" field.
func (au *AchievementUpdate) SetStartTime(t time.Time) *AchievementUpdate {
	au.mutation.SetStartTime(t)
	return au
}

// SetEndTime sets the "end_time" field.
func (au *AchievementUpdate) SetEndTime(t time.Time) *AchievementUpdate {
	au.mutation.SetEndTime(t)
	return au
}

// SetContent sets the "content" field.
func (au *AchievementUpdate) SetContent(s string) *AchievementUpdate {
	au.mutation.SetContent(s)
	return au
}

// Mutation returns the AchievementMutation object of the builder.
func (au *AchievementUpdate) Mutation() *AchievementMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AchievementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AchievementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AchievementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AchievementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AchievementUpdate) check() error {
	if v, ok := au.mutation.Title(); ok {
		if err := achievement.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (au *AchievementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   achievement.Table,
			Columns: achievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: achievement.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldTitle,
		})
	}
	if value, ok := au.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: achievement.FieldStartTime,
		})
	}
	if value, ok := au.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: achievement.FieldEndTime,
		})
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldContent,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AchievementUpdateOne is the builder for updating a single Achievement entity.
type AchievementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AchievementMutation
}

// SetTitle sets the "title" field.
func (auo *AchievementUpdateOne) SetTitle(s string) *AchievementUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetStartTime sets the "start_time" field.
func (auo *AchievementUpdateOne) SetStartTime(t time.Time) *AchievementUpdateOne {
	auo.mutation.SetStartTime(t)
	return auo
}

// SetEndTime sets the "end_time" field.
func (auo *AchievementUpdateOne) SetEndTime(t time.Time) *AchievementUpdateOne {
	auo.mutation.SetEndTime(t)
	return auo
}

// SetContent sets the "content" field.
func (auo *AchievementUpdateOne) SetContent(s string) *AchievementUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// Mutation returns the AchievementMutation object of the builder.
func (auo *AchievementUpdateOne) Mutation() *AchievementMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AchievementUpdateOne) Select(field string, fields ...string) *AchievementUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Achievement entity.
func (auo *AchievementUpdateOne) Save(ctx context.Context) (*Achievement, error) {
	var (
		err  error
		node *Achievement
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AchievementUpdateOne) SaveX(ctx context.Context) *Achievement {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AchievementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AchievementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AchievementUpdateOne) check() error {
	if v, ok := auo.mutation.Title(); ok {
		if err := achievement.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (auo *AchievementUpdateOne) sqlSave(ctx context.Context) (_node *Achievement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   achievement.Table,
			Columns: achievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: achievement.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Achievement.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievement.FieldID)
		for _, f := range fields {
			if !achievement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != achievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldTitle,
		})
	}
	if value, ok := auo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: achievement.FieldStartTime,
		})
	}
	if value, ok := auo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: achievement.FieldEndTime,
		})
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldContent,
		})
	}
	_node = &Achievement{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
