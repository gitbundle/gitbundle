// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitbundle/gitbundle/storage/database/models/commitstatusindex"
	"github.com/gitbundle/gitbundle/storage/database/models/predicate"
)

// CommitStatusIndexQuery is the builder for querying CommitStatusIndex entities.
type CommitStatusIndexQuery struct {
	config
	ctx        *QueryContext
	order      []commitstatusindex.OrderOption
	inters     []Interceptor
	predicates []predicate.CommitStatusIndex
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommitStatusIndexQuery builder.
func (csiq *CommitStatusIndexQuery) Where(ps ...predicate.CommitStatusIndex) *CommitStatusIndexQuery {
	csiq.predicates = append(csiq.predicates, ps...)
	return csiq
}

// Limit the number of records to be returned by this query.
func (csiq *CommitStatusIndexQuery) Limit(limit int) *CommitStatusIndexQuery {
	csiq.ctx.Limit = &limit
	return csiq
}

// Offset to start from.
func (csiq *CommitStatusIndexQuery) Offset(offset int) *CommitStatusIndexQuery {
	csiq.ctx.Offset = &offset
	return csiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csiq *CommitStatusIndexQuery) Unique(unique bool) *CommitStatusIndexQuery {
	csiq.ctx.Unique = &unique
	return csiq
}

// Order specifies how the records should be ordered.
func (csiq *CommitStatusIndexQuery) Order(o ...commitstatusindex.OrderOption) *CommitStatusIndexQuery {
	csiq.order = append(csiq.order, o...)
	return csiq
}

// First returns the first CommitStatusIndex entity from the query.
// Returns a *NotFoundError when no CommitStatusIndex was found.
func (csiq *CommitStatusIndexQuery) First(ctx context.Context) (*CommitStatusIndex, error) {
	nodes, err := csiq.Limit(1).All(setContextOp(ctx, csiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commitstatusindex.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) FirstX(ctx context.Context) *CommitStatusIndex {
	node, err := csiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommitStatusIndex ID from the query.
// Returns a *NotFoundError when no CommitStatusIndex ID was found.
func (csiq *CommitStatusIndexQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csiq.Limit(1).IDs(setContextOp(ctx, csiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commitstatusindex.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) FirstIDX(ctx context.Context) int {
	id, err := csiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommitStatusIndex entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommitStatusIndex entity is found.
// Returns a *NotFoundError when no CommitStatusIndex entities are found.
func (csiq *CommitStatusIndexQuery) Only(ctx context.Context) (*CommitStatusIndex, error) {
	nodes, err := csiq.Limit(2).All(setContextOp(ctx, csiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commitstatusindex.Label}
	default:
		return nil, &NotSingularError{commitstatusindex.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) OnlyX(ctx context.Context) *CommitStatusIndex {
	node, err := csiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommitStatusIndex ID in the query.
// Returns a *NotSingularError when more than one CommitStatusIndex ID is found.
// Returns a *NotFoundError when no entities are found.
func (csiq *CommitStatusIndexQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csiq.Limit(2).IDs(setContextOp(ctx, csiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commitstatusindex.Label}
	default:
		err = &NotSingularError{commitstatusindex.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) OnlyIDX(ctx context.Context) int {
	id, err := csiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommitStatusIndexes.
func (csiq *CommitStatusIndexQuery) All(ctx context.Context) ([]*CommitStatusIndex, error) {
	ctx = setContextOp(ctx, csiq.ctx, "All")
	if err := csiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommitStatusIndex, *CommitStatusIndexQuery]()
	return withInterceptors[[]*CommitStatusIndex](ctx, csiq, qr, csiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) AllX(ctx context.Context) []*CommitStatusIndex {
	nodes, err := csiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommitStatusIndex IDs.
func (csiq *CommitStatusIndexQuery) IDs(ctx context.Context) (ids []int, err error) {
	if csiq.ctx.Unique == nil && csiq.path != nil {
		csiq.Unique(true)
	}
	ctx = setContextOp(ctx, csiq.ctx, "IDs")
	if err = csiq.Select(commitstatusindex.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) IDsX(ctx context.Context) []int {
	ids, err := csiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csiq *CommitStatusIndexQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, csiq.ctx, "Count")
	if err := csiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csiq, querierCount[*CommitStatusIndexQuery](), csiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) CountX(ctx context.Context) int {
	count, err := csiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csiq *CommitStatusIndexQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, csiq.ctx, "Exist")
	switch _, err := csiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (csiq *CommitStatusIndexQuery) ExistX(ctx context.Context) bool {
	exist, err := csiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommitStatusIndexQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csiq *CommitStatusIndexQuery) Clone() *CommitStatusIndexQuery {
	if csiq == nil {
		return nil
	}
	return &CommitStatusIndexQuery{
		config:     csiq.config,
		ctx:        csiq.ctx.Clone(),
		order:      append([]commitstatusindex.OrderOption{}, csiq.order...),
		inters:     append([]Interceptor{}, csiq.inters...),
		predicates: append([]predicate.CommitStatusIndex{}, csiq.predicates...),
		// clone intermediate query.
		sql:  csiq.sql.Clone(),
		path: csiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RepoID int64 `json:"repo_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CommitStatusIndex.Query().
//		GroupBy(commitstatusindex.FieldRepoID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
func (csiq *CommitStatusIndexQuery) GroupBy(field string, fields ...string) *CommitStatusIndexGroupBy {
	csiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommitStatusIndexGroupBy{build: csiq}
	grbuild.flds = &csiq.ctx.Fields
	grbuild.label = commitstatusindex.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RepoID int64 `json:"repo_id,omitempty"`
//	}
//
//	client.CommitStatusIndex.Query().
//		Select(commitstatusindex.FieldRepoID).
//		Scan(ctx, &v)
func (csiq *CommitStatusIndexQuery) Select(fields ...string) *CommitStatusIndexSelect {
	csiq.ctx.Fields = append(csiq.ctx.Fields, fields...)
	sbuild := &CommitStatusIndexSelect{CommitStatusIndexQuery: csiq}
	sbuild.label = commitstatusindex.Label
	sbuild.flds, sbuild.scan = &csiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommitStatusIndexSelect configured with the given aggregations.
func (csiq *CommitStatusIndexQuery) Aggregate(fns ...AggregateFunc) *CommitStatusIndexSelect {
	return csiq.Select().Aggregate(fns...)
}

func (csiq *CommitStatusIndexQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csiq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csiq); err != nil {
				return err
			}
		}
	}
	for _, f := range csiq.ctx.Fields {
		if !commitstatusindex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if csiq.path != nil {
		prev, err := csiq.path(ctx)
		if err != nil {
			return err
		}
		csiq.sql = prev
	}
	return nil
}

func (csiq *CommitStatusIndexQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommitStatusIndex, error) {
	var (
		nodes = []*CommitStatusIndex{}
		_spec = csiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommitStatusIndex).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommitStatusIndex{config: csiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (csiq *CommitStatusIndexQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csiq.querySpec()
	_spec.Node.Columns = csiq.ctx.Fields
	if len(csiq.ctx.Fields) > 0 {
		_spec.Unique = csiq.ctx.Unique != nil && *csiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csiq.driver, _spec)
}

func (csiq *CommitStatusIndexQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commitstatusindex.Table, commitstatusindex.Columns, sqlgraph.NewFieldSpec(commitstatusindex.FieldID, field.TypeInt))
	_spec.From = csiq.sql
	if unique := csiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csiq.path != nil {
		_spec.Unique = true
	}
	if fields := csiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commitstatusindex.FieldID)
		for i := range fields {
			if fields[i] != commitstatusindex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csiq *CommitStatusIndexQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csiq.driver.Dialect())
	t1 := builder.Table(commitstatusindex.Table)
	columns := csiq.ctx.Fields
	if len(columns) == 0 {
		columns = commitstatusindex.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csiq.sql != nil {
		selector = csiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csiq.ctx.Unique != nil && *csiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range csiq.predicates {
		p(selector)
	}
	for _, p := range csiq.order {
		p(selector)
	}
	if offset := csiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommitStatusIndexGroupBy is the group-by builder for CommitStatusIndex entities.
type CommitStatusIndexGroupBy struct {
	selector
	build *CommitStatusIndexQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csigb *CommitStatusIndexGroupBy) Aggregate(fns ...AggregateFunc) *CommitStatusIndexGroupBy {
	csigb.fns = append(csigb.fns, fns...)
	return csigb
}

// Scan applies the selector query and scans the result into the given value.
func (csigb *CommitStatusIndexGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csigb.build.ctx, "GroupBy")
	if err := csigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommitStatusIndexQuery, *CommitStatusIndexGroupBy](ctx, csigb.build, csigb, csigb.build.inters, v)
}

func (csigb *CommitStatusIndexGroupBy) sqlScan(ctx context.Context, root *CommitStatusIndexQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csigb.fns))
	for _, fn := range csigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csigb.flds)+len(csigb.fns))
		for _, f := range *csigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommitStatusIndexSelect is the builder for selecting fields of CommitStatusIndex entities.
type CommitStatusIndexSelect struct {
	*CommitStatusIndexQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (csis *CommitStatusIndexSelect) Aggregate(fns ...AggregateFunc) *CommitStatusIndexSelect {
	csis.fns = append(csis.fns, fns...)
	return csis
}

// Scan applies the selector query and scans the result into the given value.
func (csis *CommitStatusIndexSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csis.ctx, "Select")
	if err := csis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommitStatusIndexQuery, *CommitStatusIndexSelect](ctx, csis.CommitStatusIndexQuery, csis, csis.inters, v)
}

func (csis *CommitStatusIndexSelect) sqlScan(ctx context.Context, root *CommitStatusIndexQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(csis.fns))
	for _, fn := range csis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*csis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}