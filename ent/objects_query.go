// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glu/video-real-time-ranking/ent/objects"
	"github.com/glu/video-real-time-ranking/ent/predicate"
	"github.com/glu/video-real-time-ranking/ent/videos"
)

// ObjectsQuery is the builder for querying Objects entities.
type ObjectsQuery struct {
	config
	ctx           *QueryContext
	order         []objects.OrderOption
	inters        []Interceptor
	predicates    []predicate.Objects
	withTblVideos *VideosQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ObjectsQuery builder.
func (oq *ObjectsQuery) Where(ps ...predicate.Objects) *ObjectsQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *ObjectsQuery) Limit(limit int) *ObjectsQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *ObjectsQuery) Offset(offset int) *ObjectsQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *ObjectsQuery) Unique(unique bool) *ObjectsQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *ObjectsQuery) Order(o ...objects.OrderOption) *ObjectsQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryTblVideos chains the current query on the "tbl_videos" edge.
func (oq *ObjectsQuery) QueryTblVideos() *VideosQuery {
	query := (&VideosClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(objects.Table, objects.FieldID, selector),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, objects.TblVideosTable, objects.TblVideosColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Objects entity from the query.
// Returns a *NotFoundError when no Objects was found.
func (oq *ObjectsQuery) First(ctx context.Context) (*Objects, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{objects.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *ObjectsQuery) FirstX(ctx context.Context) *Objects {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Objects ID from the query.
// Returns a *NotFoundError when no Objects ID was found.
func (oq *ObjectsQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{objects.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *ObjectsQuery) FirstIDX(ctx context.Context) uint {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Objects entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Objects entity is found.
// Returns a *NotFoundError when no Objects entities are found.
func (oq *ObjectsQuery) Only(ctx context.Context) (*Objects, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{objects.Label}
	default:
		return nil, &NotSingularError{objects.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *ObjectsQuery) OnlyX(ctx context.Context) *Objects {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Objects ID in the query.
// Returns a *NotSingularError when more than one Objects ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *ObjectsQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{objects.Label}
	default:
		err = &NotSingularError{objects.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *ObjectsQuery) OnlyIDX(ctx context.Context) uint {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ObjectsSlice.
func (oq *ObjectsQuery) All(ctx context.Context) ([]*Objects, error) {
	ctx = setContextOp(ctx, oq.ctx, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Objects, *ObjectsQuery]()
	return withInterceptors[[]*Objects](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *ObjectsQuery) AllX(ctx context.Context) []*Objects {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Objects IDs.
func (oq *ObjectsQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, "IDs")
	if err = oq.Select(objects.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *ObjectsQuery) IDsX(ctx context.Context) []uint {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *ObjectsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*ObjectsQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *ObjectsQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *ObjectsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *ObjectsQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ObjectsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *ObjectsQuery) Clone() *ObjectsQuery {
	if oq == nil {
		return nil
	}
	return &ObjectsQuery{
		config:        oq.config,
		ctx:           oq.ctx.Clone(),
		order:         append([]objects.OrderOption{}, oq.order...),
		inters:        append([]Interceptor{}, oq.inters...),
		predicates:    append([]predicate.Objects{}, oq.predicates...),
		withTblVideos: oq.withTblVideos.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithTblVideos tells the query-builder to eager-load the nodes that are connected to
// the "tbl_videos" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *ObjectsQuery) WithTblVideos(opts ...func(*VideosQuery)) *ObjectsQuery {
	query := (&VideosClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withTblVideos = query
	return oq
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
//	client.Objects.Query().
//		GroupBy(objects.FieldVideoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *ObjectsQuery) GroupBy(field string, fields ...string) *ObjectsGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ObjectsGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = objects.Label
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
//	client.Objects.Query().
//		Select(objects.FieldVideoID).
//		Scan(ctx, &v)
func (oq *ObjectsQuery) Select(fields ...string) *ObjectsSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &ObjectsSelect{ObjectsQuery: oq}
	sbuild.label = objects.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ObjectsSelect configured with the given aggregations.
func (oq *ObjectsQuery) Aggregate(fns ...AggregateFunc) *ObjectsSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *ObjectsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !objects.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *ObjectsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Objects, error) {
	var (
		nodes       = []*Objects{}
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withTblVideos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Objects).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Objects{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withTblVideos; query != nil {
		if err := oq.loadTblVideos(ctx, query, nodes, nil,
			func(n *Objects, e *Videos) { n.Edges.TblVideos = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *ObjectsQuery) loadTblVideos(ctx context.Context, query *VideosQuery, nodes []*Objects, init func(*Objects), assign func(*Objects, *Videos)) error {
	ids := make([]uint, 0, len(nodes))
	nodeids := make(map[uint][]*Objects)
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

func (oq *ObjectsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *ObjectsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(objects.Table, objects.Columns, sqlgraph.NewFieldSpec(objects.FieldID, field.TypeUint))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, objects.FieldID)
		for i := range fields {
			if fields[i] != objects.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oq.withTblVideos != nil {
			_spec.Node.AddColumnOnce(objects.FieldVideoID)
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *ObjectsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(objects.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = objects.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ObjectsGroupBy is the group-by builder for Objects entities.
type ObjectsGroupBy struct {
	selector
	build *ObjectsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *ObjectsGroupBy) Aggregate(fns ...AggregateFunc) *ObjectsGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *ObjectsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ObjectsQuery, *ObjectsGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *ObjectsGroupBy) sqlScan(ctx context.Context, root *ObjectsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ObjectsSelect is the builder for selecting fields of Objects entities.
type ObjectsSelect struct {
	*ObjectsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *ObjectsSelect) Aggregate(fns ...AggregateFunc) *ObjectsSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *ObjectsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ObjectsQuery, *ObjectsSelect](ctx, os.ObjectsQuery, os, os.inters, v)
}

func (os *ObjectsSelect) sqlScan(ctx context.Context, root *ObjectsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
