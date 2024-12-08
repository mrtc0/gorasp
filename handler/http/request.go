package http

import (
	"net/http"

	"github.com/mrtc0/gorasp/emitter"
	"github.com/mrtc0/gorasp/event"
	httpListener "github.com/mrtc0/gorasp/listener/http"
)

func StartOperation(op emitter.Operation, args httpListener.HttpRequestHandlerOperationArgs) *httpListener.HttpRequestHandlerOperation {
	o := &httpListener.HttpRequestHandlerOperation{
		Operation: op,
	}

	return emitter.StartOperation(event.HTTP_REQUEST_EVENT, o, args)
}

func WrapHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rootOperation := emitter.NewOperation()
		httpListener.RegisterHTTPRequestSecurity(rootOperation)

		args := httpListener.NewHttpRequestHandlerOperationArgsFromRequest(r)
		StartOperation(rootOperation, args)

		handler.ServeHTTP(w, r)
	})
}
