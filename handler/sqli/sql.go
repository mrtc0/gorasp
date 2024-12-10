package sqli

import (
	"context"

	"github.com/mrtc0/gorasp/emitter"
	"github.com/mrtc0/gorasp/event"
	"github.com/mrtc0/gorasp/listener/sqli"
)

func StartOperation(op emitter.Operation, args sqli.SQLQueryHandlerOperationArg) *sqli.SQLQueryHandlerOperation {
	o := &sqli.SQLQueryHandlerOperation{
		Operation: op,
	}

	return emitter.StartOperation(event.SQL_QUERY_EVENT, o, args)
}

func ProtectSQLOperation(ctx context.Context, query string) error {
	rootOperation := emitter.NewOperation()
	sqli.RegisterSQLQuerySecurity(rootOperation)

	args := sqli.SQLQueryHandlerOperationArg{
		Query: query,
	}

	StartOperation(rootOperation, args)

	if rootOperation.IsBlocked() {
		return &event.BlockEvent{}
	}

	return nil
}
