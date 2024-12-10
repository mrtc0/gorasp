package event

import "errors"

type Event string

const (
	HTTP_REQUEST_EVENT = "http_request"
	SQL_QUERY_EVENT    = "sql_query"
)

var _ error = (*BlockEvent)(nil)

type BlockEvent struct{}

func (*BlockEvent) Error() string {
	return "request blocked by gorasp"
}

// IsSecurityError returns true if the error is a security event.
func IsBlockError(err error) bool {
	var blockErr *BlockEvent
	return errors.As(err, &blockErr)
}
