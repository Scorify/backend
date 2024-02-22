// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
)

// StatusCreate is the builder for creating a Status entity.
type StatusCreate struct {
	config
	mutation *StatusMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *StatusCreate) SetCreateTime(t time.Time) *StatusCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *StatusCreate) SetNillableCreateTime(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *StatusCreate) SetUpdateTime(t time.Time) *StatusCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *StatusCreate) SetNillableUpdateTime(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetError sets the "error" field.
func (sc *StatusCreate) SetError(s string) *StatusCreate {
	sc.mutation.SetError(s)
	return sc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (sc *StatusCreate) SetNillableError(s *string) *StatusCreate {
	if s != nil {
		sc.SetError(*s)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *StatusCreate) SetStatus(s status.Status) *StatusCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StatusCreate) SetNillableStatus(s *status.Status) *StatusCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetPoints sets the "points" field.
func (sc *StatusCreate) SetPoints(i int) *StatusCreate {
	sc.mutation.SetPoints(i)
	return sc
}

// SetCheckID sets the "check_id" field.
func (sc *StatusCreate) SetCheckID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetCheckID(u)
	return sc
}

// SetRoundID sets the "round_id" field.
func (sc *StatusCreate) SetRoundID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetRoundID(u)
	return sc
}

// SetUserID sets the "user_id" field.
func (sc *StatusCreate) SetUserID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetUserID(u)
	return sc
}

// SetID sets the "id" field.
func (sc *StatusCreate) SetID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StatusCreate) SetNillableID(u *uuid.UUID) *StatusCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetCheck sets the "check" edge to the Check entity.
func (sc *StatusCreate) SetCheck(c *Check) *StatusCreate {
	return sc.SetCheckID(c.ID)
}

// SetRound sets the "round" edge to the Round entity.
func (sc *StatusCreate) SetRound(r *Round) *StatusCreate {
	return sc.SetRoundID(r.ID)
}

// SetUser sets the "user" edge to the User entity.
func (sc *StatusCreate) SetUser(u *User) *StatusCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the StatusMutation object of the builder.
func (sc *StatusCreate) Mutation() *StatusMutation {
	return sc.mutation
}

// Save creates the Status in the database.
func (sc *StatusCreate) Save(ctx context.Context) (*Status, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StatusCreate) SaveX(ctx context.Context) *Status {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StatusCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StatusCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StatusCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := status.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := status.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := status.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := status.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StatusCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Status.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Status.update_time"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Status.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := status.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Status.status": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "Status.points"`)}
	}
	if v, ok := sc.mutation.Points(); ok {
		if err := status.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`ent: validator failed for field "Status.points": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CheckID(); !ok {
		return &ValidationError{Name: "check_id", err: errors.New(`ent: missing required field "Status.check_id"`)}
	}
	if _, ok := sc.mutation.RoundID(); !ok {
		return &ValidationError{Name: "round_id", err: errors.New(`ent: missing required field "Status.round_id"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Status.user_id"`)}
	}
	if _, ok := sc.mutation.CheckID(); !ok {
		return &ValidationError{Name: "check", err: errors.New(`ent: missing required edge "Status.check"`)}
	}
	if _, ok := sc.mutation.RoundID(); !ok {
		return &ValidationError{Name: "round", err: errors.New(`ent: missing required edge "Status.round"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Status.user"`)}
	}
	return nil
}

func (sc *StatusCreate) sqlSave(ctx context.Context) (*Status, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StatusCreate) createSpec() (*Status, *sqlgraph.CreateSpec) {
	var (
		_node = &Status{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(status.Table, sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(status.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(status.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Error(); ok {
		_spec.SetField(status.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(status.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Points(); ok {
		_spec.SetField(status.FieldPoints, field.TypeInt, value)
		_node.Points = value
	}
	if nodes := sc.mutation.CheckIDs(); len(nodes) > 0 {
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
		_node.CheckID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.RoundIDs(); len(nodes) > 0 {
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
		_node.RoundID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StatusCreateBulk is the builder for creating many Status entities in bulk.
type StatusCreateBulk struct {
	config
	err      error
	builders []*StatusCreate
}

// Save creates the Status entities in the database.
func (scb *StatusCreateBulk) Save(ctx context.Context) ([]*Status, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Status, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StatusMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StatusCreateBulk) SaveX(ctx context.Context) []*Status {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StatusCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StatusCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
