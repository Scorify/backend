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
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/scorecache"
	"github.com/scorify/backend/pkg/ent/user"
)

// ScoreCacheCreate is the builder for creating a ScoreCache entity.
type ScoreCacheCreate struct {
	config
	mutation *ScoreCacheMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (scc *ScoreCacheCreate) SetCreateTime(t time.Time) *ScoreCacheCreate {
	scc.mutation.SetCreateTime(t)
	return scc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (scc *ScoreCacheCreate) SetNillableCreateTime(t *time.Time) *ScoreCacheCreate {
	if t != nil {
		scc.SetCreateTime(*t)
	}
	return scc
}

// SetUpdateTime sets the "update_time" field.
func (scc *ScoreCacheCreate) SetUpdateTime(t time.Time) *ScoreCacheCreate {
	scc.mutation.SetUpdateTime(t)
	return scc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (scc *ScoreCacheCreate) SetNillableUpdateTime(t *time.Time) *ScoreCacheCreate {
	if t != nil {
		scc.SetUpdateTime(*t)
	}
	return scc
}

// SetPoints sets the "points" field.
func (scc *ScoreCacheCreate) SetPoints(i int) *ScoreCacheCreate {
	scc.mutation.SetPoints(i)
	return scc
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (scc *ScoreCacheCreate) SetRoundID(id uuid.UUID) *ScoreCacheCreate {
	scc.mutation.SetRoundID(id)
	return scc
}

// SetRound sets the "round" edge to the Round entity.
func (scc *ScoreCacheCreate) SetRound(r *Round) *ScoreCacheCreate {
	return scc.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (scc *ScoreCacheCreate) SetUserID(id uuid.UUID) *ScoreCacheCreate {
	scc.mutation.SetUserID(id)
	return scc
}

// SetUser sets the "user" edge to the User entity.
func (scc *ScoreCacheCreate) SetUser(u *User) *ScoreCacheCreate {
	return scc.SetUserID(u.ID)
}

// Mutation returns the ScoreCacheMutation object of the builder.
func (scc *ScoreCacheCreate) Mutation() *ScoreCacheMutation {
	return scc.mutation
}

// Save creates the ScoreCache in the database.
func (scc *ScoreCacheCreate) Save(ctx context.Context) (*ScoreCache, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *ScoreCacheCreate) SaveX(ctx context.Context) *ScoreCache {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *ScoreCacheCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *ScoreCacheCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *ScoreCacheCreate) defaults() {
	if _, ok := scc.mutation.CreateTime(); !ok {
		v := scorecache.DefaultCreateTime()
		scc.mutation.SetCreateTime(v)
	}
	if _, ok := scc.mutation.UpdateTime(); !ok {
		v := scorecache.DefaultUpdateTime()
		scc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *ScoreCacheCreate) check() error {
	if _, ok := scc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "ScoreCache.create_time"`)}
	}
	if _, ok := scc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "ScoreCache.update_time"`)}
	}
	if _, ok := scc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "ScoreCache.points"`)}
	}
	if v, ok := scc.mutation.Points(); ok {
		if err := scorecache.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`ent: validator failed for field "ScoreCache.points": %w`, err)}
		}
	}
	if _, ok := scc.mutation.RoundID(); !ok {
		return &ValidationError{Name: "round", err: errors.New(`ent: missing required edge "ScoreCache.round"`)}
	}
	if _, ok := scc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ScoreCache.user"`)}
	}
	return nil
}

func (scc *ScoreCacheCreate) sqlSave(ctx context.Context) (*ScoreCache, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *ScoreCacheCreate) createSpec() (*ScoreCache, *sqlgraph.CreateSpec) {
	var (
		_node = &ScoreCache{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(scorecache.Table, sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeInt))
	)
	if value, ok := scc.mutation.CreateTime(); ok {
		_spec.SetField(scorecache.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := scc.mutation.UpdateTime(); ok {
		_spec.SetField(scorecache.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := scc.mutation.Points(); ok {
		_spec.SetField(scorecache.FieldPoints, field.TypeInt, value)
		_node.Points = value
	}
	if nodes := scc.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   scorecache.RoundTable,
			Columns: []string{scorecache.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.score_cache_round = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   scorecache.UserTable,
			Columns: []string{scorecache.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.score_cache_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScoreCacheCreateBulk is the builder for creating many ScoreCache entities in bulk.
type ScoreCacheCreateBulk struct {
	config
	err      error
	builders []*ScoreCacheCreate
}

// Save creates the ScoreCache entities in the database.
func (sccb *ScoreCacheCreateBulk) Save(ctx context.Context) ([]*ScoreCache, error) {
	if sccb.err != nil {
		return nil, sccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*ScoreCache, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScoreCacheMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *ScoreCacheCreateBulk) SaveX(ctx context.Context) []*ScoreCache {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *ScoreCacheCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *ScoreCacheCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
