// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/structs"
)

// CheckUpdate is the builder for updating Check entities.
type CheckUpdate struct {
	config
	hooks    []Hook
	mutation *CheckMutation
}

// Where appends a list predicates to the CheckUpdate builder.
func (cu *CheckUpdate) Where(ps ...predicate.Check) *CheckUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CheckUpdate) SetName(s string) *CheckUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableName(s *string) *CheckUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetSource sets the "source" field.
func (cu *CheckUpdate) SetSource(s string) *CheckUpdate {
	cu.mutation.SetSource(s)
	return cu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableSource(s *string) *CheckUpdate {
	if s != nil {
		cu.SetSource(*s)
	}
	return cu
}

// SetWeight sets the "weight" field.
func (cu *CheckUpdate) SetWeight(i int) *CheckUpdate {
	cu.mutation.ResetWeight()
	cu.mutation.SetWeight(i)
	return cu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableWeight(i *int) *CheckUpdate {
	if i != nil {
		cu.SetWeight(*i)
	}
	return cu
}

// AddWeight adds i to the "weight" field.
func (cu *CheckUpdate) AddWeight(i int) *CheckUpdate {
	cu.mutation.AddWeight(i)
	return cu
}

// SetDefaultConfig sets the "default_config" field.
func (cu *CheckUpdate) SetDefaultConfig(sc structs.CheckConfiguration) *CheckUpdate {
	cu.mutation.SetDefaultConfig(sc)
	return cu
}

// SetNillableDefaultConfig sets the "default_config" field if the given value is not nil.
func (cu *CheckUpdate) SetNillableDefaultConfig(sc *structs.CheckConfiguration) *CheckUpdate {
	if sc != nil {
		cu.SetDefaultConfig(*sc)
	}
	return cu
}

// AddConfigIDs adds the "configs" edge to the CheckConfig entity by IDs.
func (cu *CheckUpdate) AddConfigIDs(ids ...uuid.UUID) *CheckUpdate {
	cu.mutation.AddConfigIDs(ids...)
	return cu
}

// AddConfigs adds the "configs" edges to the CheckConfig entity.
func (cu *CheckUpdate) AddConfigs(c ...*CheckConfig) *CheckUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddConfigIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (cu *CheckUpdate) AddStatusIDs(ids ...uuid.UUID) *CheckUpdate {
	cu.mutation.AddStatusIDs(ids...)
	return cu
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (cu *CheckUpdate) AddStatuses(s ...*Status) *CheckUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStatusIDs(ids...)
}

// Mutation returns the CheckMutation object of the builder.
func (cu *CheckUpdate) Mutation() *CheckMutation {
	return cu.mutation
}

// ClearConfigs clears all "configs" edges to the CheckConfig entity.
func (cu *CheckUpdate) ClearConfigs() *CheckUpdate {
	cu.mutation.ClearConfigs()
	return cu
}

// RemoveConfigIDs removes the "configs" edge to CheckConfig entities by IDs.
func (cu *CheckUpdate) RemoveConfigIDs(ids ...uuid.UUID) *CheckUpdate {
	cu.mutation.RemoveConfigIDs(ids...)
	return cu
}

// RemoveConfigs removes "configs" edges to CheckConfig entities.
func (cu *CheckUpdate) RemoveConfigs(c ...*CheckConfig) *CheckUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveConfigIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (cu *CheckUpdate) ClearStatuses() *CheckUpdate {
	cu.mutation.ClearStatuses()
	return cu
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (cu *CheckUpdate) RemoveStatusIDs(ids ...uuid.UUID) *CheckUpdate {
	cu.mutation.RemoveStatusIDs(ids...)
	return cu
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (cu *CheckUpdate) RemoveStatuses(s ...*Status) *CheckUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CheckUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := check.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Check.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Source(); ok {
		if err := check.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "Check.source": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Weight(); ok {
		if err := check.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Check.weight": %w`, err)}
		}
	}
	return nil
}

func (cu *CheckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Source(); ok {
		_spec.SetField(check.FieldSource, field.TypeString, value)
	}
	if value, ok := cu.mutation.Weight(); ok {
		_spec.SetField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedWeight(); ok {
		_spec.AddField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cu.mutation.DefaultConfig(); ok {
		_spec.SetField(check.FieldDefaultConfig, field.TypeJSON, value)
	}
	if cu.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !cu.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !cu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
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
	if nodes := cu.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CheckUpdateOne is the builder for updating a single Check entity.
type CheckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckMutation
}

// SetName sets the "name" field.
func (cuo *CheckUpdateOne) SetName(s string) *CheckUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableName(s *string) *CheckUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetSource sets the "source" field.
func (cuo *CheckUpdateOne) SetSource(s string) *CheckUpdateOne {
	cuo.mutation.SetSource(s)
	return cuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableSource(s *string) *CheckUpdateOne {
	if s != nil {
		cuo.SetSource(*s)
	}
	return cuo
}

// SetWeight sets the "weight" field.
func (cuo *CheckUpdateOne) SetWeight(i int) *CheckUpdateOne {
	cuo.mutation.ResetWeight()
	cuo.mutation.SetWeight(i)
	return cuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableWeight(i *int) *CheckUpdateOne {
	if i != nil {
		cuo.SetWeight(*i)
	}
	return cuo
}

// AddWeight adds i to the "weight" field.
func (cuo *CheckUpdateOne) AddWeight(i int) *CheckUpdateOne {
	cuo.mutation.AddWeight(i)
	return cuo
}

// SetDefaultConfig sets the "default_config" field.
func (cuo *CheckUpdateOne) SetDefaultConfig(sc structs.CheckConfiguration) *CheckUpdateOne {
	cuo.mutation.SetDefaultConfig(sc)
	return cuo
}

// SetNillableDefaultConfig sets the "default_config" field if the given value is not nil.
func (cuo *CheckUpdateOne) SetNillableDefaultConfig(sc *structs.CheckConfiguration) *CheckUpdateOne {
	if sc != nil {
		cuo.SetDefaultConfig(*sc)
	}
	return cuo
}

// AddConfigIDs adds the "configs" edge to the CheckConfig entity by IDs.
func (cuo *CheckUpdateOne) AddConfigIDs(ids ...uuid.UUID) *CheckUpdateOne {
	cuo.mutation.AddConfigIDs(ids...)
	return cuo
}

// AddConfigs adds the "configs" edges to the CheckConfig entity.
func (cuo *CheckUpdateOne) AddConfigs(c ...*CheckConfig) *CheckUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddConfigIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (cuo *CheckUpdateOne) AddStatusIDs(ids ...uuid.UUID) *CheckUpdateOne {
	cuo.mutation.AddStatusIDs(ids...)
	return cuo
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (cuo *CheckUpdateOne) AddStatuses(s ...*Status) *CheckUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStatusIDs(ids...)
}

// Mutation returns the CheckMutation object of the builder.
func (cuo *CheckUpdateOne) Mutation() *CheckMutation {
	return cuo.mutation
}

// ClearConfigs clears all "configs" edges to the CheckConfig entity.
func (cuo *CheckUpdateOne) ClearConfigs() *CheckUpdateOne {
	cuo.mutation.ClearConfigs()
	return cuo
}

// RemoveConfigIDs removes the "configs" edge to CheckConfig entities by IDs.
func (cuo *CheckUpdateOne) RemoveConfigIDs(ids ...uuid.UUID) *CheckUpdateOne {
	cuo.mutation.RemoveConfigIDs(ids...)
	return cuo
}

// RemoveConfigs removes "configs" edges to CheckConfig entities.
func (cuo *CheckUpdateOne) RemoveConfigs(c ...*CheckConfig) *CheckUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveConfigIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (cuo *CheckUpdateOne) ClearStatuses() *CheckUpdateOne {
	cuo.mutation.ClearStatuses()
	return cuo
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (cuo *CheckUpdateOne) RemoveStatusIDs(ids ...uuid.UUID) *CheckUpdateOne {
	cuo.mutation.RemoveStatusIDs(ids...)
	return cuo
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (cuo *CheckUpdateOne) RemoveStatuses(s ...*Status) *CheckUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStatusIDs(ids...)
}

// Where appends a list predicates to the CheckUpdate builder.
func (cuo *CheckUpdateOne) Where(ps ...predicate.Check) *CheckUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckUpdateOne) Select(field string, fields ...string) *CheckUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Check entity.
func (cuo *CheckUpdateOne) Save(ctx context.Context) (*Check, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckUpdateOne) SaveX(ctx context.Context) *Check {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CheckUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := check.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Check.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Source(); ok {
		if err := check.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "Check.source": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Weight(); ok {
		if err := check.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Check.weight": %w`, err)}
		}
	}
	return nil
}

func (cuo *CheckUpdateOne) sqlSave(ctx context.Context) (_node *Check, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Check.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, check.FieldID)
		for _, f := range fields {
			if !check.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != check.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Source(); ok {
		_spec.SetField(check.FieldSource, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Weight(); ok {
		_spec.SetField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedWeight(); ok {
		_spec.AddField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.DefaultConfig(); ok {
		_spec.SetField(check.FieldDefaultConfig, field.TypeJSON, value)
	}
	if cuo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !cuo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.ConfigsTable,
			Columns: []string{check.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !cuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
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
	if nodes := cuo.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusesTable,
			Columns: []string{check.StatusesColumn},
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
	_node = &Check{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
