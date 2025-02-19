// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glu/video-real-time-ranking/ent/predicate"
	"github.com/glu/video-real-time-ranking/ent/videos"
	"github.com/glu/video-real-time-ranking/ent/viewers"
)

// ViewersQuery is the builder for querying Viewers entities.
type ViewersQuery struct {
	config
	ctx           *QueryContext
	order         []viewers.OrderOption
	inters        []Interceptor
	predicates    []predicate.Viewers
	withTblVideos *VideosQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ViewersQuery builder.
func (vq *ViewersQuery) Where(ps ...predicate.Viewers) *ViewersQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *ViewersQuery) Limit(limit int) *ViewersQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *ViewersQuery) Offset(offset int) *ViewersQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *ViewersQuery) Unique(unique bool) *ViewersQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *ViewersQuery) Order(o ...viewers.OrderOption) *ViewersQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryTblVideos chains the current query on the "tbl_videos" edge.
func (vq *ViewersQuery) QueryTblVideos() *VideosQuery {
	query := (&VideosClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(viewers.Table, viewers.FieldID, selector),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, viewers.TblVideosTable, viewers.TblVideosColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Viewers entity from the query.
// Returns a *NotFoundError when no Viewers was found.
func (vq *ViewersQuery) First(ctx context.Context) (*Viewers, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{viewers.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *ViewersQuery) FirstX(ctx context.Context) *Viewers {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Viewers ID from the query.
// Returns a *NotFoundError when no Viewers ID was found.
func (vq *ViewersQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{viewers.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *ViewersQuery) FirstIDX(ctx context.Context) uint {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Viewers entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Viewers entity is found.
// Returns a *NotFoundError when no Viewers entities are found.
func (vq *ViewersQuery) Only(ctx context.Context) (*Viewers, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{viewers.Label}
	default:
		return nil, &NotSingularError{viewers.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *ViewersQuery) OnlyX(ctx context.Context) *Viewers {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Viewers ID in the query.
// Returns a *NotSingularError when more than one Viewers ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *ViewersQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{viewers.Label}
	default:
		err = &NotSingularError{viewers.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *ViewersQuery) OnlyIDX(ctx context.Context) uint {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ViewersSlice.
func (vq *ViewersQuery) All(ctx context.Context) ([]*Viewers, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Viewers, *ViewersQuery]()
	return withInterceptors[[]*Viewers](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *ViewersQuery) AllX(ctx context.Context) []*Viewers {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Viewers IDs.
func (vq *ViewersQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(viewers.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *ViewersQuery) IDsX(ctx context.Context) []uint {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *ViewersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*ViewersQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *ViewersQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *ViewersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *ViewersQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ViewersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *ViewersQuery) Clone() *ViewersQuery {
	if vq == nil {
		return nil
	}
	return &ViewersQuery{
		config:        vq.config,
		ctx:           vq.ctx.Clone(),
		order:         append([]viewers.OrderOption{}, vq.order...),
		inters:        append([]Interceptor{}, vq.inters...),
		predicates:    append([]predicate.Viewers{}, vq.predicates...),
		withTblVideos: vq.withTblVideos.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithTblVideos tells the query-builder to eager-load the nodes that are connected to
// the "tbl_videos" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *ViewersQuery) WithTblVideos(opts ...func(*VideosQuery)) *ViewersQuery {
	query := (&VideosClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withTblVideos = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VideoID uint `json:"video_id" bson:"video_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Viewers.Query().
//		GroupBy(viewers.FieldVideoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *ViewersQuery) GroupBy(field string, fields ...string) *ViewersGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ViewersGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = viewers.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VideoID uint `json:"video_id" bson:"video_id"`
//	}
//
//	client.Viewers.Query().
//		Select(viewers.FieldVideoID).
//		Scan(ctx, &v)
func (vq *ViewersQuery) Select(fields ...string) *ViewersSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &ViewersSelect{ViewersQuery: vq}
	sbuild.label = viewers.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ViewersSelect configured with the given aggregations.
func (vq *ViewersQuery) Aggregate(fns ...AggregateFunc) *ViewersSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *ViewersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !viewers.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *ViewersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Viewers, error) {
	var (
		nodes       = []*Viewers{}
		_spec       = vq.querySpec()
		loadedTypes = [1]bool{
			vq.withTblVideos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Viewers).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Viewers{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withTblVideos; query != nil {
		if err := vq.loadTblVideos(ctx, query, nodes, nil,
			func(n *Viewers, e *Videos) { n.Edges.TblVideos = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *ViewersQuery) loadTblVideos(ctx context.Context, query *VideosQuery, nodes []*Viewers, init func(*Viewers), assign func(*Viewers, *Videos)) error {
	ids := make([]uint, 0, len(nodes))
	nodeids := make(map[uint][]*Viewers)
	for i := range nodes {
		fk := nodes[i].VideoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(videos.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vq *ViewersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *ViewersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(viewers.Table, viewers.Columns, sqlgraph.NewFieldSpec(viewers.FieldID, field.TypeUint))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, viewers.FieldID)
		for i := range fields {
			if fields[i] != viewers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if vq.withTblVideos != nil {
			_spec.Node.AddColumnOnce(viewers.FieldVideoID)
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *ViewersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(viewers.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = viewers.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ViewersGroupBy is the group-by builder for Viewers entities.
type ViewersGroupBy struct {
	selector
	build *ViewersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *ViewersGroupBy) Aggregate(fns ...AggregateFunc) *ViewersGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *ViewersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ViewersQuery, *ViewersGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *ViewersGroupBy) sqlScan(ctx context.Context, root *ViewersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ViewersSelect is the builder for selecting fields of Viewers entities.
type ViewersSelect struct {
	*ViewersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *ViewersSelect) Aggregate(fns ...AggregateFunc) *ViewersSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *ViewersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ViewersQuery, *ViewersSelect](ctx, vs.ViewersQuery, vs, vs.inters, v)
}

func (vs *ViewersSelect) sqlScan(ctx context.Context, root *ViewersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
