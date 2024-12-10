package sqli

import (
	"github.com/mrtc0/gorasp/emitter"
	"github.com/mrtc0/gorasp/event"
	"github.com/mrtc0/gorasp/inspector/sqli"
)

type SQLQueryHandlerOperation struct {
	emitter.Operation
}

type SQLQueryHandlerOperationArg struct {
	Query string
}

func (SQLQueryHandlerOperationArg) IsArgOf(*SQLQueryHandlerOperation) {}

func RegisterSQLQuerySecurity(op emitter.Operation) {
	emitter.On(event.SQL_QUERY_EVENT, op, InspectSQLQuery)
}

func InspectSQLQuery(op *SQLQueryHandlerOperation, args SQLQueryHandlerOperationArg) {
	err := sqli.IsSQLiQuery(args.Query)
	if err != nil {
		op.SetIsBlock(true)
	}
}
