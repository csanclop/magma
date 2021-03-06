// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/floorplanscale"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// FloorPlanScaleQuery is the builder for querying FloorPlanScale entities.
type FloorPlanScaleQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.FloorPlanScale
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (fpsq *FloorPlanScaleQuery) Where(ps ...predicate.FloorPlanScale) *FloorPlanScaleQuery {
	fpsq.predicates = append(fpsq.predicates, ps...)
	return fpsq
}

// Limit adds a limit step to the query.
func (fpsq *FloorPlanScaleQuery) Limit(limit int) *FloorPlanScaleQuery {
	fpsq.limit = &limit
	return fpsq
}

// Offset adds an offset step to the query.
func (fpsq *FloorPlanScaleQuery) Offset(offset int) *FloorPlanScaleQuery {
	fpsq.offset = &offset
	return fpsq
}

// Order adds an order step to the query.
func (fpsq *FloorPlanScaleQuery) Order(o ...Order) *FloorPlanScaleQuery {
	fpsq.order = append(fpsq.order, o...)
	return fpsq
}

// First returns the first FloorPlanScale entity in the query. Returns *ErrNotFound when no floorplanscale was found.
func (fpsq *FloorPlanScaleQuery) First(ctx context.Context) (*FloorPlanScale, error) {
	fpsSlice, err := fpsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fpsSlice) == 0 {
		return nil, &ErrNotFound{floorplanscale.Label}
	}
	return fpsSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) FirstX(ctx context.Context) *FloorPlanScale {
	fps, err := fpsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return fps
}

// FirstID returns the first FloorPlanScale id in the query. Returns *ErrNotFound when no id was found.
func (fpsq *FloorPlanScaleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fpsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{floorplanscale.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) FirstXID(ctx context.Context) string {
	id, err := fpsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FloorPlanScale entity in the query, returns an error if not exactly one entity was returned.
func (fpsq *FloorPlanScaleQuery) Only(ctx context.Context) (*FloorPlanScale, error) {
	fpsSlice, err := fpsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fpsSlice) {
	case 1:
		return fpsSlice[0], nil
	case 0:
		return nil, &ErrNotFound{floorplanscale.Label}
	default:
		return nil, &ErrNotSingular{floorplanscale.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) OnlyX(ctx context.Context) *FloorPlanScale {
	fps, err := fpsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return fps
}

// OnlyID returns the only FloorPlanScale id in the query, returns an error if not exactly one id was returned.
func (fpsq *FloorPlanScaleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fpsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{floorplanscale.Label}
	default:
		err = &ErrNotSingular{floorplanscale.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) OnlyXID(ctx context.Context) string {
	id, err := fpsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FloorPlanScales.
func (fpsq *FloorPlanScaleQuery) All(ctx context.Context) ([]*FloorPlanScale, error) {
	return fpsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) AllX(ctx context.Context) []*FloorPlanScale {
	fpsSlice, err := fpsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fpsSlice
}

// IDs executes the query and returns a list of FloorPlanScale ids.
func (fpsq *FloorPlanScaleQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fpsq.Select(floorplanscale.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) IDsX(ctx context.Context) []string {
	ids, err := fpsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fpsq *FloorPlanScaleQuery) Count(ctx context.Context) (int, error) {
	return fpsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) CountX(ctx context.Context) int {
	count, err := fpsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fpsq *FloorPlanScaleQuery) Exist(ctx context.Context) (bool, error) {
	return fpsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fpsq *FloorPlanScaleQuery) ExistX(ctx context.Context) bool {
	exist, err := fpsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fpsq *FloorPlanScaleQuery) Clone() *FloorPlanScaleQuery {
	return &FloorPlanScaleQuery{
		config:     fpsq.config,
		limit:      fpsq.limit,
		offset:     fpsq.offset,
		order:      append([]Order{}, fpsq.order...),
		unique:     append([]string{}, fpsq.unique...),
		predicates: append([]predicate.FloorPlanScale{}, fpsq.predicates...),
		// clone intermediate query.
		sql: fpsq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FloorPlanScale.Query().
//		GroupBy(floorplanscale.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fpsq *FloorPlanScaleQuery) GroupBy(field string, fields ...string) *FloorPlanScaleGroupBy {
	group := &FloorPlanScaleGroupBy{config: fpsq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = fpsq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.FloorPlanScale.Query().
//		Select(floorplanscale.FieldCreateTime).
//		Scan(ctx, &v)
//
func (fpsq *FloorPlanScaleQuery) Select(field string, fields ...string) *FloorPlanScaleSelect {
	selector := &FloorPlanScaleSelect{config: fpsq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = fpsq.sqlQuery()
	return selector
}

func (fpsq *FloorPlanScaleQuery) sqlAll(ctx context.Context) ([]*FloorPlanScale, error) {
	rows := &sql.Rows{}
	selector := fpsq.sqlQuery()
	if unique := fpsq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := fpsq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var fpsSlice FloorPlanScales
	if err := fpsSlice.FromRows(rows); err != nil {
		return nil, err
	}
	fpsSlice.config(fpsq.config)
	return fpsSlice, nil
}

func (fpsq *FloorPlanScaleQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := fpsq.sqlQuery()
	unique := []string{floorplanscale.FieldID}
	if len(fpsq.unique) > 0 {
		unique = fpsq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := fpsq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (fpsq *FloorPlanScaleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fpsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fpsq *FloorPlanScaleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(fpsq.driver.Dialect())
	t1 := builder.Table(floorplanscale.Table)
	selector := builder.Select(t1.Columns(floorplanscale.Columns...)...).From(t1)
	if fpsq.sql != nil {
		selector = fpsq.sql
		selector.Select(selector.Columns(floorplanscale.Columns...)...)
	}
	for _, p := range fpsq.predicates {
		p(selector)
	}
	for _, p := range fpsq.order {
		p(selector)
	}
	if offset := fpsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fpsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FloorPlanScaleGroupBy is the builder for group-by FloorPlanScale entities.
type FloorPlanScaleGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fpsgb *FloorPlanScaleGroupBy) Aggregate(fns ...Aggregate) *FloorPlanScaleGroupBy {
	fpsgb.fns = append(fpsgb.fns, fns...)
	return fpsgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fpsgb *FloorPlanScaleGroupBy) Scan(ctx context.Context, v interface{}) error {
	return fpsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fpsgb *FloorPlanScaleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fpsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fpsgb *FloorPlanScaleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fpsgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fpsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fpsgb *FloorPlanScaleGroupBy) StringsX(ctx context.Context) []string {
	v, err := fpsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fpsgb *FloorPlanScaleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fpsgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fpsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fpsgb *FloorPlanScaleGroupBy) IntsX(ctx context.Context) []int {
	v, err := fpsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fpsgb *FloorPlanScaleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fpsgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fpsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fpsgb *FloorPlanScaleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fpsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fpsgb *FloorPlanScaleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fpsgb.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fpsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fpsgb *FloorPlanScaleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fpsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fpsgb *FloorPlanScaleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fpsgb.sqlQuery().Query()
	if err := fpsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fpsgb *FloorPlanScaleGroupBy) sqlQuery() *sql.Selector {
	selector := fpsgb.sql
	columns := make([]string, 0, len(fpsgb.fields)+len(fpsgb.fns))
	columns = append(columns, fpsgb.fields...)
	for _, fn := range fpsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(fpsgb.fields...)
}

// FloorPlanScaleSelect is the builder for select fields of FloorPlanScale entities.
type FloorPlanScaleSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (fpss *FloorPlanScaleSelect) Scan(ctx context.Context, v interface{}) error {
	return fpss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fpss *FloorPlanScaleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fpss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fpss *FloorPlanScaleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fpss.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fpss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fpss *FloorPlanScaleSelect) StringsX(ctx context.Context) []string {
	v, err := fpss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fpss *FloorPlanScaleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fpss.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fpss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fpss *FloorPlanScaleSelect) IntsX(ctx context.Context) []int {
	v, err := fpss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fpss *FloorPlanScaleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fpss.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fpss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fpss *FloorPlanScaleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fpss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fpss *FloorPlanScaleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fpss.fields) > 1 {
		return nil, errors.New("ent: FloorPlanScaleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fpss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fpss *FloorPlanScaleSelect) BoolsX(ctx context.Context) []bool {
	v, err := fpss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fpss *FloorPlanScaleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fpss.sqlQuery().Query()
	if err := fpss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fpss *FloorPlanScaleSelect) sqlQuery() sql.Querier {
	view := "floorplanscale_view"
	return sql.Dialect(fpss.driver.Dialect()).
		Select(fpss.fields...).From(fpss.sql.As(view))
}
