// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glu/video-real-time-ranking/ent/comments"
	"github.com/glu/video-real-time-ranking/ent/predicate"
	"github.com/glu/video-real-time-ranking/ent/videos"
)

// CommentsQuery is the builder for querying Comments entities.
type CommentsQuery struct {
	config
	ctx           *QueryContext
	order         []comments.OrderOption
	inters        []Interceptor
	predicates    []predicate.Comments
	withTblVideos *VideosQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommentsQuery builder.
func (cq *CommentsQuery) Where(ps ...predicate.Comments) *CommentsQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommentsQuery) Limit(limit int) *CommentsQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommentsQuery) Offset(offset int) *CommentsQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommentsQuery) Unique(unique bool) *CommentsQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommentsQuery) Order(o ...comments.OrderOption) *CommentsQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryTblVideos chains the current query on the "tbl_videos" edge.
func (cq *CommentsQuery) QueryTblVideos() *VideosQuery {
	query := (&VideosClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comments.Table, comments.FieldID, selector),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comments.TblVideosTable, comments.TblVideosColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Comments entity from the query.
// Returns a *NotFoundError when no Comments was found.
func (cq *CommentsQuery) First(ctx context.Context) (*Comments, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comments.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommentsQuery) FirstX(ctx context.Context) *Comments {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Comments ID from the query.
// Returns a *NotFoundError when no Comments ID was found.
func (cq *CommentsQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comments.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommentsQuery) FirstIDX(ctx context.Context) uint {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Comments entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comments entity is found.
// Returns a *NotFoundError when no Comments entities are found.
func (cq *CommentsQuery) Only(ctx context.Context) (*Comments, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comments.Label}
	default:
		return nil, &NotSingularError{comments.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommentsQuery) OnlyX(ctx context.Context) *Comments {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Comments ID in the query.
// Returns a *NotSingularError when more than one Comments ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommentsQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comments.Label}
	default:
		err = &NotSingularError{comments.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommentsQuery) OnlyIDX(ctx context.Context) uint {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommentsSlice.
func (cq *CommentsQuery) All(ctx context.Context) ([]*Comments, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Comments, *CommentsQuery]()
	return withInterceptors[[]*Comments](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommentsQuery) AllX(ctx context.Context) []*Comments {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Comments IDs.
func (cq *CommentsQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(comments.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommentsQuery) IDsX(ctx context.Context) []uint {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommentsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommentsQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommentsQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommentsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommentsQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommentsQuery) Clone() *CommentsQuery {
	if cq == nil {
		return nil
	}
	return &CommentsQuery{
		config:        cq.config,
		ctx:           cq.ctx.Clone(),
		order:         append([]comments.OrderOption{}, cq.order...),
		inters:        append([]Interceptor{}, cq.inters...),
		predicates:    append([]predicate.Comments{}, cq.predicates...),
		withTblVideos: cq.withTblVideos.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithTblVideos tells the query-builder to eager-load the nodes that are connected to
// the "tbl_videos" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentsQuery) WithTblVideos(opts ...func(*VideosQuery)) *CommentsQuery {
	query := (&VideosClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTblVideos = query
	return cq
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
//	client.Comments.Query().
//		GroupBy(comments.FieldVideoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CommentsQuery) GroupBy(field string, fields ...string) *CommentsGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentsGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = comments.Label
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
//	client.Comments.Query().
//		Select(comments.FieldVideoID).
//		Scan(ctx, &v)
func (cq *CommentsQuery) Select(fields ...string) *CommentsSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CommentsSelect{CommentsQuery: cq}
	sbuild.label = comments.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentsSelect configured with the given aggregations.
func (cq *CommentsQuery) Aggregate(fns ...AggregateFunc) *CommentsSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommentsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !comments.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CommentsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Comments, error) {
	var (
		nodes       = []*Comments{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withTblVideos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Comments).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Comments{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withTblVideos; query != nil {
		if err := cq.loadTblVideos(ctx, query, nodes, nil,
			func(n *Comments, e *Videos) { n.Edges.TblVideos = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CommentsQuery) loadTblVideos(ctx context.Context, query *VideosQuery, nodes []*Comments, init func(*Comments), assign func(*Comments, *Videos)) error {
	ids := make([]uint, 0, len(nodes))
	nodeids := make(map[uint][]*Comments)
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

func (cq *CommentsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommentsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(comments.Table, comments.Columns, sqlgraph.NewFieldSpec(comments.FieldID, field.TypeUint))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comments.FieldID)
		for i := range fields {
			if fields[i] != comments.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withTblVideos != nil {
			_spec.Node.AddColumnOnce(comments.FieldVideoID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CommentsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(comments.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = comments.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommentsGroupBy is the group-by builder for Comments entities.
type CommentsGroupBy struct {
	selector
	build *CommentsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentsGroupBy) Aggregate(fns ...AggregateFunc) *CommentsGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommentsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentsQuery, *CommentsGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommentsGroupBy) sqlScan(ctx context.Context, root *CommentsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommentsSelect is the builder for selecting fields of Comments entities.
type CommentsSelect struct {
	*CommentsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommentsSelect) Aggregate(fns ...AggregateFunc) *CommentsSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommentsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentsQuery, *CommentsSelect](ctx, cs.CommentsQuery, cs, cs.inters, v)
}

func (cs *CommentsSelect) sqlScan(ctx context.Context, root *CommentsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
