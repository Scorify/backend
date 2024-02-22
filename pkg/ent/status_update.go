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
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
)

// StatusUpdate is the builder for updating Status entities.
type StatusUpdate struct {
	config
	hooks    []Hook
	mutation *StatusMutation
}

// Where appends a list predicates to the StatusUpdate builder.
func (su *StatusUpdate) Where(ps ...predicate.Status) *StatusUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *StatusUpdate) SetUpdateTime(t time.Time) *StatusUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetError sets the "error" field.
func (su *StatusUpdate) SetError(s string) *StatusUpdate {
	su.mutation.SetError(s)
	return su
}

// SetNillableError sets the "error" field if the given value is not nil.
func (su *StatusUpdate) SetNillableError(s *string) *StatusUpdate {
	if s != nil {
		su.SetError(*s)
	}
	return su
}

// ClearError clears the value of the "error" field.
func (su *StatusUpdate) ClearError() *StatusUpdate {
	su.mutation.ClearError()
	return su
}

// SetStatus sets the "status" field.
func (su *StatusUpdate) SetStatus(s status.Status) *StatusUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *StatusUpdate) SetNillableStatus(s *status.Status) *StatusUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// SetWeight sets the "weight" field.
func (su *StatusUpdate) SetWeight(i int) *StatusUpdate {
	su.mutation.ResetWeight()
	su.mutation.SetWeight(i)
	return su
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (su *StatusUpdate) SetNillableWeight(i *int) *StatusUpdate {
	if i != nil {
		su.SetWeight(*i)
	}
	return su
}

// AddWeight adds i to the "weight" field.
func (su *StatusUpdate) AddWeight(i int) *StatusUpdate {
	su.mutation.AddWeight(i)
	return su
}

// SetCheckID sets the "check" edge to the Check entity by ID.
func (su *StatusUpdate) SetCheckID(id uuid.UUID) *StatusUpdate {
	su.mutation.SetCheckID(id)
	return su
}

// SetCheck sets the "check" edge to the Check entity.
func (su *StatusUpdate) SetCheck(c *Check) *StatusUpdate {
	return su.SetCheckID(c.ID)
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (su *StatusUpdate) SetRoundID(id uuid.UUID) *StatusUpdate {
	su.mutation.SetRoundID(id)
	return su
}

// SetRound sets the "round" edge to the Round entity.
func (su *StatusUpdate) SetRound(r *Round) *StatusUpdate {
	return su.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *StatusUpdate) SetUserID(id uuid.UUID) *StatusUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *StatusUpdate) SetUser(u *User) *StatusUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the StatusMutation object of the builder.
func (su *StatusUpdate) Mutation() *StatusMutation {
	return su.mutation
}

// ClearCheck clears the "check" edge to the Check entity.
func (su *StatusUpdate) ClearCheck() *StatusUpdate {
	su.mutation.ClearCheck()
	return su
}

// ClearRound clears the "round" edge to the Round entity.
func (su *StatusUpdate) ClearRound() *StatusUpdate {
	su.mutation.ClearRound()
	return su
}

// ClearUser clears the "user" edge to the User entity.
func (su *StatusUpdate) ClearUser() *StatusUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StatusUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatusUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatusUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatusUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StatusUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := status.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StatusUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := status.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Status.status": %w`, err)}
		}
	}
	if _, ok := su.mutation.CheckID(); su.mutation.CheckCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.check"`)
	}
	if _, ok := su.mutation.RoundID(); su.mutation.RoundCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.round"`)
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.user"`)
	}
	return nil
}

func (su *StatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(status.Table, status.Columns, sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(status.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Error(); ok {
		_spec.SetField(status.FieldError, field.TypeString, value)
	}
	if su.mutation.ErrorCleared() {
		_spec.ClearField(status.FieldError, field.TypeString)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(status.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := su.mutation.Weight(); ok {
		_spec.SetField(status.FieldWeight, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedWeight(); ok {
		_spec.AddField(status.FieldWeight, field.TypeInt, value)
	}
	if su.mutation.CheckCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.CheckTable,
			Columns: []string{status.CheckColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CheckIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.CheckTable,
			Columns: []string{status.CheckColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.RoundTable,
			Columns: []string{status.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.RoundTable,
			Columns: []string{status.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.UserTable,
			Columns: []string{status.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.UserTable,
			Columns: []string{status.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StatusUpdateOne is the builder for updating a single Status entity.
type StatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatusMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *StatusUpdateOne) SetUpdateTime(t time.Time) *StatusUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetError sets the "error" field.
func (suo *StatusUpdateOne) SetError(s string) *StatusUpdateOne {
	suo.mutation.SetError(s)
	return suo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (suo *StatusUpdateOne) SetNillableError(s *string) *StatusUpdateOne {
	if s != nil {
		suo.SetError(*s)
	}
	return suo
}

// ClearError clears the value of the "error" field.
func (suo *StatusUpdateOne) ClearError() *StatusUpdateOne {
	suo.mutation.ClearError()
	return suo
}

// SetStatus sets the "status" field.
func (suo *StatusUpdateOne) SetStatus(s status.Status) *StatusUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *StatusUpdateOne) SetNillableStatus(s *status.Status) *StatusUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// SetWeight sets the "weight" field.
func (suo *StatusUpdateOne) SetWeight(i int) *StatusUpdateOne {
	suo.mutation.ResetWeight()
	suo.mutation.SetWeight(i)
	return suo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (suo *StatusUpdateOne) SetNillableWeight(i *int) *StatusUpdateOne {
	if i != nil {
		suo.SetWeight(*i)
	}
	return suo
}

// AddWeight adds i to the "weight" field.
func (suo *StatusUpdateOne) AddWeight(i int) *StatusUpdateOne {
	suo.mutation.AddWeight(i)
	return suo
}

// SetCheckID sets the "check" edge to the Check entity by ID.
func (suo *StatusUpdateOne) SetCheckID(id uuid.UUID) *StatusUpdateOne {
	suo.mutation.SetCheckID(id)
	return suo
}

// SetCheck sets the "check" edge to the Check entity.
func (suo *StatusUpdateOne) SetCheck(c *Check) *StatusUpdateOne {
	return suo.SetCheckID(c.ID)
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (suo *StatusUpdateOne) SetRoundID(id uuid.UUID) *StatusUpdateOne {
	suo.mutation.SetRoundID(id)
	return suo
}

// SetRound sets the "round" edge to the Round entity.
func (suo *StatusUpdateOne) SetRound(r *Round) *StatusUpdateOne {
	return suo.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *StatusUpdateOne) SetUserID(id uuid.UUID) *StatusUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *StatusUpdateOne) SetUser(u *User) *StatusUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the StatusMutation object of the builder.
func (suo *StatusUpdateOne) Mutation() *StatusMutation {
	return suo.mutation
}

// ClearCheck clears the "check" edge to the Check entity.
func (suo *StatusUpdateOne) ClearCheck() *StatusUpdateOne {
	suo.mutation.ClearCheck()
	return suo
}

// ClearRound clears the "round" edge to the Round entity.
func (suo *StatusUpdateOne) ClearRound() *StatusUpdateOne {
	suo.mutation.ClearRound()
	return suo
}

// ClearUser clears the "user" edge to the User entity.
func (suo *StatusUpdateOne) ClearUser() *StatusUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the StatusUpdate builder.
func (suo *StatusUpdateOne) Where(ps ...predicate.Status) *StatusUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StatusUpdateOne) Select(field string, fields ...string) *StatusUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Status entity.
func (suo *StatusUpdateOne) Save(ctx context.Context) (*Status, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatusUpdateOne) SaveX(ctx context.Context) *Status {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StatusUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatusUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StatusUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := status.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StatusUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := status.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Status.status": %w`, err)}
		}
	}
	if _, ok := suo.mutation.CheckID(); suo.mutation.CheckCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.check"`)
	}
	if _, ok := suo.mutation.RoundID(); suo.mutation.RoundCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.round"`)
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Status.user"`)
	}
	return nil
}

func (suo *StatusUpdateOne) sqlSave(ctx context.Context) (_node *Status, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(status.Table, status.Columns, sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Status.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, status.FieldID)
		for _, f := range fields {
			if !status.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != status.FieldID {
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
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(status.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Error(); ok {
		_spec.SetField(status.FieldError, field.TypeString, value)
	}
	if suo.mutation.ErrorCleared() {
		_spec.ClearField(status.FieldError, field.TypeString)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(status.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.Weight(); ok {
		_spec.SetField(status.FieldWeight, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedWeight(); ok {
		_spec.AddField(status.FieldWeight, field.TypeInt, value)
	}
	if suo.mutation.CheckCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.CheckTable,
			Columns: []string{status.CheckColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CheckIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.CheckTable,
			Columns: []string{status.CheckColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(check.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.RoundTable,
			Columns: []string{status.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.RoundTable,
			Columns: []string{status.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.UserTable,
			Columns: []string{status.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   status.UserTable,
			Columns: []string{status.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Status{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
