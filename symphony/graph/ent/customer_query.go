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
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/symphony/graph/ent/customer"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/service"
)

// CustomerQuery is the builder for querying Customer entities.
type CustomerQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Customer
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (cq *CustomerQuery) Where(ps ...predicate.Customer) *CustomerQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CustomerQuery) Limit(limit int) *CustomerQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CustomerQuery) Offset(offset int) *CustomerQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *CustomerQuery) Order(o ...Order) *CustomerQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryServices chains the current query on the services edge.
func (cq *CustomerQuery) QueryServices() *ServiceQuery {
	query := &ServiceQuery{config: cq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(customer.Table, customer.FieldID, cq.sqlQuery()),
		sqlgraph.To(service.Table, service.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, customer.ServicesTable, customer.ServicesPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
	return query
}

// First returns the first Customer entity in the query. Returns *ErrNotFound when no customer was found.
func (cq *CustomerQuery) First(ctx context.Context) (*Customer, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &ErrNotFound{customer.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CustomerQuery) FirstX(ctx context.Context) *Customer {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first Customer id in the query. Returns *ErrNotFound when no id was found.
func (cq *CustomerQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{customer.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *CustomerQuery) FirstXID(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Customer entity in the query, returns an error if not exactly one entity was returned.
func (cq *CustomerQuery) Only(ctx context.Context) (*Customer, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &ErrNotFound{customer.Label}
	default:
		return nil, &ErrNotSingular{customer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CustomerQuery) OnlyX(ctx context.Context) *Customer {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only Customer id in the query, returns an error if not exactly one id was returned.
func (cq *CustomerQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{customer.Label}
	default:
		err = &ErrNotSingular{customer.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (cq *CustomerQuery) OnlyXID(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Customers.
func (cq *CustomerQuery) All(ctx context.Context) ([]*Customer, error) {
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CustomerQuery) AllX(ctx context.Context) []*Customer {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of Customer ids.
func (cq *CustomerQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := cq.Select(customer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CustomerQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CustomerQuery) Count(ctx context.Context) (int, error) {
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CustomerQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CustomerQuery) Exist(ctx context.Context) (bool, error) {
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CustomerQuery) Clone() *CustomerQuery {
	return &CustomerQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]Order{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.Customer{}, cq.predicates...),
		// clone intermediate query.
		sql: cq.sql.Clone(),
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
//	client.Customer.Query().
//		GroupBy(customer.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CustomerQuery) GroupBy(field string, fields ...string) *CustomerGroupBy {
	group := &CustomerGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = cq.sqlQuery()
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
//	client.Customer.Query().
//		Select(customer.FieldCreateTime).
//		Scan(ctx, &v)
//
func (cq *CustomerQuery) Select(field string, fields ...string) *CustomerSelect {
	selector := &CustomerSelect{config: cq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = cq.sqlQuery()
	return selector
}

func (cq *CustomerQuery) sqlAll(ctx context.Context) ([]*Customer, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	if unique := cq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var cs Customers
	if err := cs.FromRows(rows); err != nil {
		return nil, err
	}
	cs.config(cq.config)
	return cs, nil
}

func (cq *CustomerQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	unique := []string{customer.FieldID}
	if len(cq.unique) > 0 {
		unique = cq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
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

func (cq *CustomerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CustomerQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(customer.Table)
	selector := builder.Select(t1.Columns(customer.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(customer.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CustomerGroupBy is the builder for group-by Customer entities.
type CustomerGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CustomerGroupBy) Aggregate(fns ...Aggregate) *CustomerGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *CustomerGroupBy) Scan(ctx context.Context, v interface{}) error {
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CustomerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *CustomerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CustomerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CustomerGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *CustomerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CustomerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CustomerGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *CustomerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CustomerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CustomerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *CustomerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CustomerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CustomerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CustomerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CustomerGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// CustomerSelect is the builder for select fields of Customer entities.
type CustomerSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (cs *CustomerSelect) Scan(ctx context.Context, v interface{}) error {
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *CustomerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cs *CustomerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CustomerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *CustomerSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cs *CustomerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CustomerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *CustomerSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cs *CustomerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CustomerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *CustomerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cs *CustomerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CustomerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *CustomerSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *CustomerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *CustomerSelect) sqlQuery() sql.Querier {
	view := "customer_view"
	return sql.Dialect(cs.driver.Dialect()).
		Select(cs.fields...).From(cs.sql.As(view))
}
