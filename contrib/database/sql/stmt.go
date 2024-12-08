package sql

import (
	"context"
	"database/sql/driver"
	"errors"
)

// For type assertion
var _ driver.Stmt = (*raspStmt)(nil)

type raspStmt struct {
	driver.Stmt
	query string
}

func (s *raspStmt) Close() error {
	return s.Stmt.Close()
}

func (s *raspStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if stmtExecContext, ok := s.Stmt.(driver.StmtExecContext); ok {
		return stmtExecContext.ExecContext(ctx, args)
	}

	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	return s.Stmt.Exec(dargs)
}

func (s *raspStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if stmtQueryContext, ok := s.Stmt.(driver.StmtQueryContext); ok {
		return stmtQueryContext.QueryContext(ctx, args)
	}

	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	return s.Stmt.Query(dargs)
}

// copied from database/sql package: https://go.dev/src/database/sql/ctxutil.go
func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}
