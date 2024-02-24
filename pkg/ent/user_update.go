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
	"github.com/scorify/backend/pkg/ent/scorecache"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetNumber sets the "number" field.
func (uu *UserUpdate) SetNumber(i int) *UserUpdate {
	uu.mutation.ResetNumber()
	uu.mutation.SetNumber(i)
	return uu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNumber(i *int) *UserUpdate {
	if i != nil {
		uu.SetNumber(*i)
	}
	return uu
}

// AddNumber adds i to the "number" field.
func (uu *UserUpdate) AddNumber(i int) *UserUpdate {
	uu.mutation.AddNumber(i)
	return uu
}

// ClearNumber clears the value of the "number" field.
func (uu *UserUpdate) ClearNumber() *UserUpdate {
	uu.mutation.ClearNumber()
	return uu
}

// AddConfigIDs adds the "configs" edge to the CheckConfig entity by IDs.
func (uu *UserUpdate) AddConfigIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddConfigIDs(ids...)
	return uu
}

// AddConfigs adds the "configs" edges to the CheckConfig entity.
func (uu *UserUpdate) AddConfigs(c ...*CheckConfig) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddConfigIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (uu *UserUpdate) AddStatusIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddStatusIDs(ids...)
	return uu
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (uu *UserUpdate) AddStatuses(s ...*Status) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddStatusIDs(ids...)
}

// AddScoreCachIDs adds the "scoreCaches" edge to the ScoreCache entity by IDs.
func (uu *UserUpdate) AddScoreCachIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddScoreCachIDs(ids...)
	return uu
}

// AddScoreCaches adds the "scoreCaches" edges to the ScoreCache entity.
func (uu *UserUpdate) AddScoreCaches(s ...*ScoreCache) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddScoreCachIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearConfigs clears all "configs" edges to the CheckConfig entity.
func (uu *UserUpdate) ClearConfigs() *UserUpdate {
	uu.mutation.ClearConfigs()
	return uu
}

// RemoveConfigIDs removes the "configs" edge to CheckConfig entities by IDs.
func (uu *UserUpdate) RemoveConfigIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveConfigIDs(ids...)
	return uu
}

// RemoveConfigs removes "configs" edges to CheckConfig entities.
func (uu *UserUpdate) RemoveConfigs(c ...*CheckConfig) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveConfigIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (uu *UserUpdate) ClearStatuses() *UserUpdate {
	uu.mutation.ClearStatuses()
	return uu
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (uu *UserUpdate) RemoveStatusIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveStatusIDs(ids...)
	return uu
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (uu *UserUpdate) RemoveStatuses(s ...*Status) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveStatusIDs(ids...)
}

// ClearScoreCaches clears all "scoreCaches" edges to the ScoreCache entity.
func (uu *UserUpdate) ClearScoreCaches() *UserUpdate {
	uu.mutation.ClearScoreCaches()
	return uu
}

// RemoveScoreCachIDs removes the "scoreCaches" edge to ScoreCache entities by IDs.
func (uu *UserUpdate) RemoveScoreCachIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveScoreCachIDs(ids...)
	return uu
}

// RemoveScoreCaches removes "scoreCaches" edges to ScoreCache entities.
func (uu *UserUpdate) RemoveScoreCaches(s ...*ScoreCache) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveScoreCachIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Number(); ok {
		if err := user.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "User.number": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Number(); ok {
		_spec.SetField(user.FieldNumber, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedNumber(); ok {
		_spec.AddField(user.FieldNumber, field.TypeInt, value)
	}
	if uu.mutation.NumberCleared() {
		_spec.ClearField(user.FieldNumber, field.TypeInt)
	}
	if uu.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !uu.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
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
	if nodes := uu.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
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
	if uu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !uu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
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
	if nodes := uu.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
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
	if uu.mutation.ScoreCachesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedScoreCachesIDs(); len(nodes) > 0 && !uu.mutation.ScoreCachesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ScoreCachesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetNumber sets the "number" field.
func (uuo *UserUpdateOne) SetNumber(i int) *UserUpdateOne {
	uuo.mutation.ResetNumber()
	uuo.mutation.SetNumber(i)
	return uuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNumber(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetNumber(*i)
	}
	return uuo
}

// AddNumber adds i to the "number" field.
func (uuo *UserUpdateOne) AddNumber(i int) *UserUpdateOne {
	uuo.mutation.AddNumber(i)
	return uuo
}

// ClearNumber clears the value of the "number" field.
func (uuo *UserUpdateOne) ClearNumber() *UserUpdateOne {
	uuo.mutation.ClearNumber()
	return uuo
}

// AddConfigIDs adds the "configs" edge to the CheckConfig entity by IDs.
func (uuo *UserUpdateOne) AddConfigIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddConfigIDs(ids...)
	return uuo
}

// AddConfigs adds the "configs" edges to the CheckConfig entity.
func (uuo *UserUpdateOne) AddConfigs(c ...*CheckConfig) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddConfigIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (uuo *UserUpdateOne) AddStatusIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddStatusIDs(ids...)
	return uuo
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (uuo *UserUpdateOne) AddStatuses(s ...*Status) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddStatusIDs(ids...)
}

// AddScoreCachIDs adds the "scoreCaches" edge to the ScoreCache entity by IDs.
func (uuo *UserUpdateOne) AddScoreCachIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddScoreCachIDs(ids...)
	return uuo
}

// AddScoreCaches adds the "scoreCaches" edges to the ScoreCache entity.
func (uuo *UserUpdateOne) AddScoreCaches(s ...*ScoreCache) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddScoreCachIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearConfigs clears all "configs" edges to the CheckConfig entity.
func (uuo *UserUpdateOne) ClearConfigs() *UserUpdateOne {
	uuo.mutation.ClearConfigs()
	return uuo
}

// RemoveConfigIDs removes the "configs" edge to CheckConfig entities by IDs.
func (uuo *UserUpdateOne) RemoveConfigIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveConfigIDs(ids...)
	return uuo
}

// RemoveConfigs removes "configs" edges to CheckConfig entities.
func (uuo *UserUpdateOne) RemoveConfigs(c ...*CheckConfig) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveConfigIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (uuo *UserUpdateOne) ClearStatuses() *UserUpdateOne {
	uuo.mutation.ClearStatuses()
	return uuo
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (uuo *UserUpdateOne) RemoveStatusIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveStatusIDs(ids...)
	return uuo
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (uuo *UserUpdateOne) RemoveStatuses(s ...*Status) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveStatusIDs(ids...)
}

// ClearScoreCaches clears all "scoreCaches" edges to the ScoreCache entity.
func (uuo *UserUpdateOne) ClearScoreCaches() *UserUpdateOne {
	uuo.mutation.ClearScoreCaches()
	return uuo
}

// RemoveScoreCachIDs removes the "scoreCaches" edge to ScoreCache entities by IDs.
func (uuo *UserUpdateOne) RemoveScoreCachIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveScoreCachIDs(ids...)
	return uuo
}

// RemoveScoreCaches removes "scoreCaches" edges to ScoreCache entities.
func (uuo *UserUpdateOne) RemoveScoreCaches(s ...*ScoreCache) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveScoreCachIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Number(); ok {
		if err := user.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "User.number": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Number(); ok {
		_spec.SetField(user.FieldNumber, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedNumber(); ok {
		_spec.AddField(user.FieldNumber, field.TypeInt, value)
	}
	if uuo.mutation.NumberCleared() {
		_spec.ClearField(user.FieldNumber, field.TypeInt)
	}
	if uuo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(checkconfig.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedConfigsIDs(); len(nodes) > 0 && !uuo.mutation.ConfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
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
	if nodes := uuo.mutation.ConfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ConfigsTable,
			Columns: []string{user.ConfigsColumn},
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
	if uuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !uuo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
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
	if nodes := uuo.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.StatusesTable,
			Columns: []string{user.StatusesColumn},
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
	if uuo.mutation.ScoreCachesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedScoreCachesIDs(); len(nodes) > 0 && !uuo.mutation.ScoreCachesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ScoreCachesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ScoreCachesTable,
			Columns: []string{user.ScoreCachesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
