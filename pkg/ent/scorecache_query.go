// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/scorecache"
	"github.com/scorify/backend/pkg/ent/user"
)

// ScoreCacheQuery is the builder for querying ScoreCache entities.
type ScoreCacheQuery struct {
	config
	ctx        *QueryContext
	order      []scorecache.OrderOption
	inters     []Interceptor
	predicates []predicate.ScoreCache
	withRound  *RoundQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScoreCacheQuery builder.
func (scq *ScoreCacheQuery) Where(ps ...predicate.ScoreCache) *ScoreCacheQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ScoreCacheQuery) Limit(limit int) *ScoreCacheQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ScoreCacheQuery) Offset(offset int) *ScoreCacheQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ScoreCacheQuery) Unique(unique bool) *ScoreCacheQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ScoreCacheQuery) Order(o ...scorecache.OrderOption) *ScoreCacheQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryRound chains the current query on the "round" edge.
func (scq *ScoreCacheQuery) QueryRound() *RoundQuery {
	query := (&RoundClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scorecache.Table, scorecache.FieldID, selector),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, scorecache.RoundTable, scorecache.RoundColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (scq *ScoreCacheQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scorecache.Table, scorecache.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, scorecache.UserTable, scorecache.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScoreCache entity from the query.
// Returns a *NotFoundError when no ScoreCache was found.
func (scq *ScoreCacheQuery) First(ctx context.Context) (*ScoreCache, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scorecache.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ScoreCacheQuery) FirstX(ctx context.Context) *ScoreCache {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScoreCache ID from the query.
// Returns a *NotFoundError when no ScoreCache ID was found.
func (scq *ScoreCacheQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scorecache.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ScoreCacheQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScoreCache entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScoreCache entity is found.
// Returns a *NotFoundError when no ScoreCache entities are found.
func (scq *ScoreCacheQuery) Only(ctx context.Context) (*ScoreCache, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scorecache.Label}
	default:
		return nil, &NotSingularError{scorecache.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ScoreCacheQuery) OnlyX(ctx context.Context) *ScoreCache {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScoreCache ID in the query.
// Returns a *NotSingularError when more than one ScoreCache ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ScoreCacheQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scorecache.Label}
	default:
		err = &NotSingularError{scorecache.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ScoreCacheQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScoreCaches.
func (scq *ScoreCacheQuery) All(ctx context.Context) ([]*ScoreCache, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ScoreCache, *ScoreCacheQuery]()
	return withInterceptors[[]*ScoreCache](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ScoreCacheQuery) AllX(ctx context.Context) []*ScoreCache {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScoreCache IDs.
func (scq *ScoreCacheQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(scorecache.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ScoreCacheQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ScoreCacheQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ScoreCacheQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ScoreCacheQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ScoreCacheQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ScoreCacheQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScoreCacheQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ScoreCacheQuery) Clone() *ScoreCacheQuery {
	if scq == nil {
		return nil
	}
	return &ScoreCacheQuery{
		config:     scq.config,
		ctx:        scq.ctx.Clone(),
		order:      append([]scorecache.OrderOption{}, scq.order...),
		inters:     append([]Interceptor{}, scq.inters...),
		predicates: append([]predicate.ScoreCache{}, scq.predicates...),
		withRound:  scq.withRound.Clone(),
		withUser:   scq.withUser.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithRound tells the query-builder to eager-load the nodes that are connected to
// the "round" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ScoreCacheQuery) WithRound(opts ...func(*RoundQuery)) *ScoreCacheQuery {
	query := (&RoundClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withRound = query
	return scq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ScoreCacheQuery) WithUser(opts ...func(*UserQuery)) *ScoreCacheQuery {
	query := (&UserClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withUser = query
	return scq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ScoreCache.Query().
//		GroupBy(scorecache.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ScoreCacheQuery) GroupBy(field string, fields ...string) *ScoreCacheGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScoreCacheGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = scorecache.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ScoreCache.Query().
//		Select(scorecache.FieldCreateTime).
//		Scan(ctx, &v)
func (scq *ScoreCacheQuery) Select(fields ...string) *ScoreCacheSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ScoreCacheSelect{ScoreCacheQuery: scq}
	sbuild.label = scorecache.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScoreCacheSelect configured with the given aggregations.
func (scq *ScoreCacheQuery) Aggregate(fns ...AggregateFunc) *ScoreCacheSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ScoreCacheQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !scorecache.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ScoreCacheQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScoreCache, error) {
	var (
		nodes       = []*ScoreCache{}
		_spec       = scq.querySpec()
		loadedTypes = [2]bool{
			scq.withRound != nil,
			scq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ScoreCache).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ScoreCache{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withRound; query != nil {
		if err := scq.loadRound(ctx, query, nodes, nil,
			func(n *ScoreCache, e *Round) { n.Edges.Round = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withUser; query != nil {
		if err := scq.loadUser(ctx, query, nodes, nil,
			func(n *ScoreCache, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ScoreCacheQuery) loadRound(ctx context.Context, query *RoundQuery, nodes []*ScoreCache, init func(*ScoreCache), assign func(*ScoreCache, *Round)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ScoreCache)
	for i := range nodes {
		fk := nodes[i].RoundID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(round.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "round_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (scq *ScoreCacheQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ScoreCache, init func(*ScoreCache), assign func(*ScoreCache, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ScoreCache)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *ScoreCacheQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ScoreCacheQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(scorecache.Table, scorecache.Columns, sqlgraph.NewFieldSpec(scorecache.FieldID, field.TypeUUID))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scorecache.FieldID)
		for i := range fields {
			if fields[i] != scorecache.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if scq.withRound != nil {
			_spec.Node.AddColumnOnce(scorecache.FieldRoundID)
		}
		if scq.withUser != nil {
			_spec.Node.AddColumnOnce(scorecache.FieldUserID)
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ScoreCacheQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(scorecache.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = scorecache.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScoreCacheGroupBy is the group-by builder for ScoreCache entities.
type ScoreCacheGroupBy struct {
	selector
	build *ScoreCacheQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ScoreCacheGroupBy) Aggregate(fns ...AggregateFunc) *ScoreCacheGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ScoreCacheGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScoreCacheQuery, *ScoreCacheGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ScoreCacheGroupBy) sqlScan(ctx context.Context, root *ScoreCacheQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScoreCacheSelect is the builder for selecting fields of ScoreCache entities.
type ScoreCacheSelect struct {
	*ScoreCacheQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ScoreCacheSelect) Aggregate(fns ...AggregateFunc) *ScoreCacheSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ScoreCacheSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScoreCacheQuery, *ScoreCacheSelect](ctx, scs.ScoreCacheQuery, scs, scs.inters, v)
}

func (scs *ScoreCacheSelect) sqlScan(ctx context.Context, root *ScoreCacheQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
