package sql

import (
	"context"
	"database/sql/driver"
)

// For type assertion
var _ driver.Conn = (*raspConn)(nil)

type raspConn struct {
	driver.Conn
}

func (c raspConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if connBeginTx, ok := c.Conn.(driver.ConnBeginTx); ok {
		return connBeginTx.BeginTx(ctx, opts)
	}

	tx, err := c.Conn.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c raspConn) Ping(ctx context.Context) (err error) {
	if pinger, ok := c.Conn.(driver.Pinger); ok {
		err = pinger.Ping(ctx)
	}
	return
}

func (c raspConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	if queryerContext, ok := c.Conn.(driver.QueryerContext); ok {
		rows, err = queryerContext.QueryContext(ctx, query, args)
		return rows, err
	}

	if queryer, ok := c.Conn.(driver.Queryer); ok {
		dargs, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		rows, err = queryer.Query(query, dargs)
		return rows, err
	}

	return nil, driver.ErrSkip
}

func (c raspConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (result driver.Result, err error) {
	if execContext, ok := c.Conn.(driver.ExecerContext); ok {
		result, err = execContext.ExecContext(ctx, query, args)
		return result, err
	}

	if execer, ok := c.Conn.(driver.Execer); ok {
		dargs, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		result, err = execer.Exec(query, dargs)
		return result, err
	}

	return nil, driver.ErrSkip
}

func (c raspConn) PrepareContext(ctx context.Context, query string) (stmt driver.Stmt, err error) {
	if preparerCtx, ok := c.Conn.(driver.ConnPrepareContext); ok {
		stmt, err = preparerCtx.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}

		return &raspStmt{Stmt: stmt, query: query}, nil
	}

	stmt, err = c.Prepare(query)
	if err != nil {
		return nil, err
	}

	return &raspStmt{Stmt: stmt, query: query}, nil
}
