// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/server/pkg/store/models/predicate"
	"github.com/gitbundle/server/pkg/store/models/repotransfer"
)

// RepoTransferQuery is the builder for querying RepoTransfer entities.
type RepoTransferQuery struct {
	config
	ctx        *QueryContext
	order      []repotransfer.OrderOption
	inters     []Interceptor
	predicates []predicate.RepoTransfer
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RepoTransferQuery builder.
func (rtq *RepoTransferQuery) Where(ps ...predicate.RepoTransfer) *RepoTransferQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit the number of records to be returned by this query.
func (rtq *RepoTransferQuery) Limit(limit int) *RepoTransferQuery {
	rtq.ctx.Limit = &limit
	return rtq
}

// Offset to start from.
func (rtq *RepoTransferQuery) Offset(offset int) *RepoTransferQuery {
	rtq.ctx.Offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *RepoTransferQuery) Unique(unique bool) *RepoTransferQuery {
	rtq.ctx.Unique = &unique
	return rtq
}

// Order specifies how the records should be ordered.
func (rtq *RepoTransferQuery) Order(o ...repotransfer.OrderOption) *RepoTransferQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// First returns the first RepoTransfer entity from the query.
// Returns a *NotFoundError when no RepoTransfer was found.
func (rtq *RepoTransferQuery) First(ctx context.Context) (*RepoTransfer, error) {
	nodes, err := rtq.Limit(1).All(setContextOp(ctx, rtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{repotransfer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *RepoTransferQuery) FirstX(ctx context.Context) *RepoTransfer {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RepoTransfer ID from the query.
// Returns a *NotFoundError when no RepoTransfer ID was found.
func (rtq *RepoTransferQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(setContextOp(ctx, rtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{repotransfer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *RepoTransferQuery) FirstIDX(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RepoTransfer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RepoTransfer entity is found.
// Returns a *NotFoundError when no RepoTransfer entities are found.
func (rtq *RepoTransferQuery) Only(ctx context.Context) (*RepoTransfer, error) {
	nodes, err := rtq.Limit(2).All(setContextOp(ctx, rtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{repotransfer.Label}
	default:
		return nil, &NotSingularError{repotransfer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *RepoTransferQuery) OnlyX(ctx context.Context) *RepoTransfer {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RepoTransfer ID in the query.
// Returns a *NotSingularError when more than one RepoTransfer ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *RepoTransferQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(setContextOp(ctx, rtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{repotransfer.Label}
	default:
		err = &NotSingularError{repotransfer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *RepoTransferQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RepoTransfers.
func (rtq *RepoTransferQuery) All(ctx context.Context) ([]*RepoTransfer, error) {
	ctx = setContextOp(ctx, rtq.ctx, "All")
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RepoTransfer, *RepoTransferQuery]()
	return withInterceptors[[]*RepoTransfer](ctx, rtq, qr, rtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rtq *RepoTransferQuery) AllX(ctx context.Context) []*RepoTransfer {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RepoTransfer IDs.
func (rtq *RepoTransferQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rtq.ctx.Unique == nil && rtq.path != nil {
		rtq.Unique(true)
	}
	ctx = setContextOp(ctx, rtq.ctx, "IDs")
	if err = rtq.Select(repotransfer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *RepoTransferQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *RepoTransferQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Count")
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rtq, querierCount[*RepoTransferQuery](), rtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *RepoTransferQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *RepoTransferQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Exist")
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *RepoTransferQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RepoTransferQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *RepoTransferQuery) Clone() *RepoTransferQuery {
	if rtq == nil {
		return nil
	}
	return &RepoTransferQuery{
		config:     rtq.config,
		ctx:        rtq.ctx.Clone(),
		order:      append([]repotransfer.OrderOption{}, rtq.order...),
		inters:     append([]Interceptor{}, rtq.inters...),
		predicates: append([]predicate.RepoTransfer{}, rtq.predicates...),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DoerID int64 `json:"doer_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RepoTransfer.Query().
//		GroupBy(repotransfer.FieldDoerID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
func (rtq *RepoTransferQuery) GroupBy(field string, fields ...string) *RepoTransferGroupBy {
	rtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RepoTransferGroupBy{build: rtq}
	grbuild.flds = &rtq.ctx.Fields
	grbuild.label = repotransfer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DoerID int64 `json:"doer_id,omitempty"`
//	}
//
//	client.RepoTransfer.Query().
//		Select(repotransfer.FieldDoerID).
//		Scan(ctx, &v)
func (rtq *RepoTransferQuery) Select(fields ...string) *RepoTransferSelect {
	rtq.ctx.Fields = append(rtq.ctx.Fields, fields...)
	sbuild := &RepoTransferSelect{RepoTransferQuery: rtq}
	sbuild.label = repotransfer.Label
	sbuild.flds, sbuild.scan = &rtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RepoTransferSelect configured with the given aggregations.
func (rtq *RepoTransferQuery) Aggregate(fns ...AggregateFunc) *RepoTransferSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *RepoTransferQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rtq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rtq.ctx.Fields {
		if !repotransfer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *RepoTransferQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RepoTransfer, error) {
	var (
		nodes = []*RepoTransfer{}
		_spec = rtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RepoTransfer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RepoTransfer{config: rtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rtq *RepoTransferQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	_spec.Node.Columns = rtq.ctx.Fields
	if len(rtq.ctx.Fields) > 0 {
		_spec.Unique = rtq.ctx.Unique != nil && *rtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *RepoTransferQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(repotransfer.Table, repotransfer.Columns, sqlgraph.NewFieldSpec(repotransfer.FieldID, field.TypeInt))
	_spec.From = rtq.sql
	if unique := rtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rtq.path != nil {
		_spec.Unique = true
	}
	if fields := rtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repotransfer.FieldID)
		for i := range fields {
			if fields[i] != repotransfer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *RepoTransferQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(repotransfer.Table)
	columns := rtq.ctx.Fields
	if len(columns) == 0 {
		columns = repotransfer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.ctx.Unique != nil && *rtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RepoTransferGroupBy is the group-by builder for RepoTransfer entities.
type RepoTransferGroupBy struct {
	selector
	build *RepoTransferQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *RepoTransferGroupBy) Aggregate(fns ...AggregateFunc) *RepoTransferGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rtgb *RepoTransferGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rtgb.build.ctx, "GroupBy")
	if err := rtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepoTransferQuery, *RepoTransferGroupBy](ctx, rtgb.build, rtgb, rtgb.build.inters, v)
}

func (rtgb *RepoTransferGroupBy) sqlScan(ctx context.Context, root *RepoTransferQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rtgb.flds)+len(rtgb.fns))
		for _, f := range *rtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RepoTransferSelect is the builder for selecting fields of RepoTransfer entities.
type RepoTransferSelect struct {
	*RepoTransferQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *RepoTransferSelect) Aggregate(fns ...AggregateFunc) *RepoTransferSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *RepoTransferSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rts.ctx, "Select")
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepoTransferQuery, *RepoTransferSelect](ctx, rts.RepoTransferQuery, rts, rts.inters, v)
}

func (rts *RepoTransferSelect) sqlScan(ctx context.Context, root *RepoTransferQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}