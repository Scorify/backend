// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/inject"
	"github.com/scorify/backend/pkg/ent/injectsubmission"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/structs"
)

// InjectUpdate is the builder for updating Inject entities.
type InjectUpdate struct {
	config
	hooks    []Hook
	mutation *InjectMutation
}

// Where appends a list predicates to the InjectUpdate builder.
func (iu *InjectUpdate) Where(ps ...predicate.Inject) *InjectUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *InjectUpdate) SetUpdateTime(t time.Time) *InjectUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetTitle sets the "title" field.
func (iu *InjectUpdate) SetTitle(s string) *InjectUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (iu *InjectUpdate) SetNillableTitle(s *string) *InjectUpdate {
	if s != nil {
		iu.SetTitle(*s)
	}
	return iu
}

// SetStartTime sets the "start_time" field.
func (iu *InjectUpdate) SetStartTime(t time.Time) *InjectUpdate {
	iu.mutation.SetStartTime(t)
	return iu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (iu *InjectUpdate) SetNillableStartTime(t *time.Time) *InjectUpdate {
	if t != nil {
		iu.SetStartTime(*t)
	}
	return iu
}

// SetEndTime sets the "end_time" field.
func (iu *InjectUpdate) SetEndTime(t time.Time) *InjectUpdate {
	iu.mutation.SetEndTime(t)
	return iu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (iu *InjectUpdate) SetNillableEndTime(t *time.Time) *InjectUpdate {
	if t != nil {
		iu.SetEndTime(*t)
	}
	return iu
}

// SetFiles sets the "files" field.
func (iu *InjectUpdate) SetFiles(s []structs.File) *InjectUpdate {
	iu.mutation.SetFiles(s)
	return iu
}

// AppendFiles appends s to the "files" field.
func (iu *InjectUpdate) AppendFiles(s []structs.File) *InjectUpdate {
	iu.mutation.AppendFiles(s)
	return iu
}

// AddSubmissionIDs adds the "submissions" edge to the InjectSubmission entity by IDs.
func (iu *InjectUpdate) AddSubmissionIDs(ids ...uuid.UUID) *InjectUpdate {
	iu.mutation.AddSubmissionIDs(ids...)
	return iu
}

// AddSubmissions adds the "submissions" edges to the InjectSubmission entity.
func (iu *InjectUpdate) AddSubmissions(i ...*InjectSubmission) *InjectUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.AddSubmissionIDs(ids...)
}

// Mutation returns the InjectMutation object of the builder.
func (iu *InjectUpdate) Mutation() *InjectMutation {
	return iu.mutation
}

// ClearSubmissions clears all "submissions" edges to the InjectSubmission entity.
func (iu *InjectUpdate) ClearSubmissions() *InjectUpdate {
	iu.mutation.ClearSubmissions()
	return iu
}

// RemoveSubmissionIDs removes the "submissions" edge to InjectSubmission entities by IDs.
func (iu *InjectUpdate) RemoveSubmissionIDs(ids ...uuid.UUID) *InjectUpdate {
	iu.mutation.RemoveSubmissionIDs(ids...)
	return iu
}

// RemoveSubmissions removes "submissions" edges to InjectSubmission entities.
func (iu *InjectUpdate) RemoveSubmissions(i ...*InjectSubmission) *InjectUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.RemoveSubmissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InjectUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InjectUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InjectUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InjectUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InjectUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := inject.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InjectUpdate) check() error {
	if v, ok := iu.mutation.Title(); ok {
		if err := inject.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Inject.title": %w`, err)}
		}
	}
	return nil
}

func (iu *InjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(inject.Table, inject.Columns, sqlgraph.NewFieldSpec(inject.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.SetField(inject.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.SetField(inject.FieldTitle, field.TypeString, value)
	}
	if value, ok := iu.mutation.StartTime(); ok {
		_spec.SetField(inject.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := iu.mutation.EndTime(); ok {
		_spec.SetField(inject.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Files(); ok {
		_spec.SetField(inject.FieldFiles, field.TypeJSON, value)
	}
	if value, ok := iu.mutation.AppendedFiles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, inject.FieldFiles, value)
		})
	}
	if iu.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedSubmissionsIDs(); len(nodes) > 0 && !iu.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.SubmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InjectUpdateOne is the builder for updating a single Inject entity.
type InjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InjectMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *InjectUpdateOne) SetUpdateTime(t time.Time) *InjectUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetTitle sets the "title" field.
func (iuo *InjectUpdateOne) SetTitle(s string) *InjectUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (iuo *InjectUpdateOne) SetNillableTitle(s *string) *InjectUpdateOne {
	if s != nil {
		iuo.SetTitle(*s)
	}
	return iuo
}

// SetStartTime sets the "start_time" field.
func (iuo *InjectUpdateOne) SetStartTime(t time.Time) *InjectUpdateOne {
	iuo.mutation.SetStartTime(t)
	return iuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (iuo *InjectUpdateOne) SetNillableStartTime(t *time.Time) *InjectUpdateOne {
	if t != nil {
		iuo.SetStartTime(*t)
	}
	return iuo
}

// SetEndTime sets the "end_time" field.
func (iuo *InjectUpdateOne) SetEndTime(t time.Time) *InjectUpdateOne {
	iuo.mutation.SetEndTime(t)
	return iuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (iuo *InjectUpdateOne) SetNillableEndTime(t *time.Time) *InjectUpdateOne {
	if t != nil {
		iuo.SetEndTime(*t)
	}
	return iuo
}

// SetFiles sets the "files" field.
func (iuo *InjectUpdateOne) SetFiles(s []structs.File) *InjectUpdateOne {
	iuo.mutation.SetFiles(s)
	return iuo
}

// AppendFiles appends s to the "files" field.
func (iuo *InjectUpdateOne) AppendFiles(s []structs.File) *InjectUpdateOne {
	iuo.mutation.AppendFiles(s)
	return iuo
}

// AddSubmissionIDs adds the "submissions" edge to the InjectSubmission entity by IDs.
func (iuo *InjectUpdateOne) AddSubmissionIDs(ids ...uuid.UUID) *InjectUpdateOne {
	iuo.mutation.AddSubmissionIDs(ids...)
	return iuo
}

// AddSubmissions adds the "submissions" edges to the InjectSubmission entity.
func (iuo *InjectUpdateOne) AddSubmissions(i ...*InjectSubmission) *InjectUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.AddSubmissionIDs(ids...)
}

// Mutation returns the InjectMutation object of the builder.
func (iuo *InjectUpdateOne) Mutation() *InjectMutation {
	return iuo.mutation
}

// ClearSubmissions clears all "submissions" edges to the InjectSubmission entity.
func (iuo *InjectUpdateOne) ClearSubmissions() *InjectUpdateOne {
	iuo.mutation.ClearSubmissions()
	return iuo
}

// RemoveSubmissionIDs removes the "submissions" edge to InjectSubmission entities by IDs.
func (iuo *InjectUpdateOne) RemoveSubmissionIDs(ids ...uuid.UUID) *InjectUpdateOne {
	iuo.mutation.RemoveSubmissionIDs(ids...)
	return iuo
}

// RemoveSubmissions removes "submissions" edges to InjectSubmission entities.
func (iuo *InjectUpdateOne) RemoveSubmissions(i ...*InjectSubmission) *InjectUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.RemoveSubmissionIDs(ids...)
}

// Where appends a list predicates to the InjectUpdate builder.
func (iuo *InjectUpdateOne) Where(ps ...predicate.Inject) *InjectUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InjectUpdateOne) Select(field string, fields ...string) *InjectUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Inject entity.
func (iuo *InjectUpdateOne) Save(ctx context.Context) (*Inject, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InjectUpdateOne) SaveX(ctx context.Context) *Inject {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InjectUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InjectUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InjectUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := inject.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InjectUpdateOne) check() error {
	if v, ok := iuo.mutation.Title(); ok {
		if err := inject.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Inject.title": %w`, err)}
		}
	}
	return nil
}

func (iuo *InjectUpdateOne) sqlSave(ctx context.Context) (_node *Inject, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(inject.Table, inject.Columns, sqlgraph.NewFieldSpec(inject.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Inject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inject.FieldID)
		for _, f := range fields {
			if !inject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != inject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.SetField(inject.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.SetField(inject.FieldTitle, field.TypeString, value)
	}
	if value, ok := iuo.mutation.StartTime(); ok {
		_spec.SetField(inject.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.EndTime(); ok {
		_spec.SetField(inject.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Files(); ok {
		_spec.SetField(inject.FieldFiles, field.TypeJSON, value)
	}
	if value, ok := iuo.mutation.AppendedFiles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, inject.FieldFiles, value)
		})
	}
	if iuo.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedSubmissionsIDs(); len(nodes) > 0 && !iuo.mutation.SubmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.SubmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inject.SubmissionsTable,
			Columns: []string{inject.SubmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Inject{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
