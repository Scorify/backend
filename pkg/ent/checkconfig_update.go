// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/status"
)

// CheckConfigUpdate is the builder for updating CheckConfig entities.
type CheckConfigUpdate struct {
	config
	hooks    []Hook
	mutation *CheckConfigMutation
}

// Where appends a list predicates to the CheckConfigUpdate builder.
func (ccu *CheckConfigUpdate) Where(ps ...predicate.CheckConfig) *CheckConfigUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetUpdateTime sets the "update_time" field.
func (ccu *CheckConfigUpdate) SetUpdateTime(t time.Time) *CheckConfigUpdate {
	ccu.mutation.SetUpdateTime(t)
	return ccu
}

// SetConfig sets the "config" field.
func (ccu *CheckConfigUpdate) SetConfig(m map[string]interface{}) *CheckConfigUpdate {
	ccu.mutation.SetConfig(m)
	return ccu
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (ccu *CheckConfigUpdate) AddStatusIDs(ids ...uuid.UUID) *CheckConfigUpdate {
	ccu.mutation.AddStatusIDs(ids...)
	return ccu
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (ccu *CheckConfigUpdate) AddStatuses(s ...*Status) *CheckConfigUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ccu.AddStatusIDs(ids...)
}

// Mutation returns the CheckConfigMutation object of the builder.
func (ccu *CheckConfigUpdate) Mutation() *CheckConfigMutation {
	return ccu.mutation
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (ccu *CheckConfigUpdate) ClearStatuses() *CheckConfigUpdate {
	ccu.mutation.ClearStatuses()
	return ccu
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (ccu *CheckConfigUpdate) RemoveStatusIDs(ids ...uuid.UUID) *CheckConfigUpdate {
	ccu.mutation.RemoveStatusIDs(ids...)
	return ccu
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (ccu *CheckConfigUpdate) RemoveStatuses(s ...*Status) *CheckConfigUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ccu.RemoveStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CheckConfigUpdate) Save(ctx context.Context) (int, error) {
	ccu.defaults()
	return withHooks(ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CheckConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CheckConfigUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CheckConfigUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *CheckConfigUpdate) defaults() {
	if _, ok := ccu.mutation.UpdateTime(); !ok {
		v := checkconfig.UpdateDefaultUpdateTime()
		ccu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *CheckConfigUpdate) check() error {
	if _, ok := ccu.mutation.CheckID(); ccu.mutation.CheckCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CheckConfig.check"`)
	}
	if _, ok := ccu.mutation.UserID(); ccu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CheckConfig.user"`)
	}
	return nil
}

func (ccu *CheckConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ccu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(checkconfig.Table, checkconfig.Columns, sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID))
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.UpdateTime(); ok {
		_spec.SetField(checkconfig.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ccu.mutation.Config(); ok {
		_spec.SetField(checkconfig.FieldConfig, field.TypeJSON, value)
	}
	if ccu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !ccu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// CheckConfigUpdateOne is the builder for updating a single CheckConfig entity.
type CheckConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckConfigMutation
}

// SetUpdateTime sets the "update_time" field.
func (ccuo *CheckConfigUpdateOne) SetUpdateTime(t time.Time) *CheckConfigUpdateOne {
	ccuo.mutation.SetUpdateTime(t)
	return ccuo
}

// SetConfig sets the "config" field.
func (ccuo *CheckConfigUpdateOne) SetConfig(m map[string]interface{}) *CheckConfigUpdateOne {
	ccuo.mutation.SetConfig(m)
	return ccuo
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (ccuo *CheckConfigUpdateOne) AddStatusIDs(ids ...uuid.UUID) *CheckConfigUpdateOne {
	ccuo.mutation.AddStatusIDs(ids...)
	return ccuo
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (ccuo *CheckConfigUpdateOne) AddStatuses(s ...*Status) *CheckConfigUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ccuo.AddStatusIDs(ids...)
}

// Mutation returns the CheckConfigMutation object of the builder.
func (ccuo *CheckConfigUpdateOne) Mutation() *CheckConfigMutation {
	return ccuo.mutation
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (ccuo *CheckConfigUpdateOne) ClearStatuses() *CheckConfigUpdateOne {
	ccuo.mutation.ClearStatuses()
	return ccuo
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (ccuo *CheckConfigUpdateOne) RemoveStatusIDs(ids ...uuid.UUID) *CheckConfigUpdateOne {
	ccuo.mutation.RemoveStatusIDs(ids...)
	return ccuo
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (ccuo *CheckConfigUpdateOne) RemoveStatuses(s ...*Status) *CheckConfigUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ccuo.RemoveStatusIDs(ids...)
}

// Where appends a list predicates to the CheckConfigUpdate builder.
func (ccuo *CheckConfigUpdateOne) Where(ps ...predicate.CheckConfig) *CheckConfigUpdateOne {
	ccuo.mutation.Where(ps...)
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CheckConfigUpdateOne) Select(field string, fields ...string) *CheckConfigUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CheckConfig entity.
func (ccuo *CheckConfigUpdateOne) Save(ctx context.Context) (*CheckConfig, error) {
	ccuo.defaults()
	return withHooks(ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CheckConfigUpdateOne) SaveX(ctx context.Context) *CheckConfig {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CheckConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CheckConfigUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *CheckConfigUpdateOne) defaults() {
	if _, ok := ccuo.mutation.UpdateTime(); !ok {
		v := checkconfig.UpdateDefaultUpdateTime()
		ccuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *CheckConfigUpdateOne) check() error {
	if _, ok := ccuo.mutation.CheckID(); ccuo.mutation.CheckCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CheckConfig.check"`)
	}
	if _, ok := ccuo.mutation.UserID(); ccuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CheckConfig.user"`)
	}
	return nil
}

func (ccuo *CheckConfigUpdateOne) sqlSave(ctx context.Context) (_node *CheckConfig, err error) {
	if err := ccuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(checkconfig.Table, checkconfig.Columns, sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID))
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CheckConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checkconfig.FieldID)
		for _, f := range fields {
			if !checkconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != checkconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.UpdateTime(); ok {
		_spec.SetField(checkconfig.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ccuo.mutation.Config(); ok {
		_spec.SetField(checkconfig.FieldConfig, field.TypeJSON, value)
	}
	if ccuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !ccuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checkconfig.StatusesTable,
			Columns: []string{checkconfig.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CheckConfig{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}
