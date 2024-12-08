package http

import (
	"net/http"

	"github.com/mrtc0/gorasp/emitter"
	"github.com/mrtc0/gorasp/event"
)

type HttpRequestHandlerOperation struct {
	emitter.Operation
}

type HttpRequestHandlerOperationArgs struct {
	Method  string
	URI     string
	Queries map[string]string
	Headers map[string]string
	Body    string
}

func NewHttpRequestHandlerOperationArgsFromRequest(r *http.Request) HttpRequestHandlerOperationArgs {
	headers := make(map[string]string)
	for k, v := range r.Header {
		value := ""
		for _, vv := range v {
			value += vv
		}

		headers[k] = value
	}

	query := r.URL.Query()

	queries := make(map[string]string)

	for k, v := range query {
		value := ""
		for _, vv := range v {
			value += vv
		}

		queries[k] = value
	}

	return HttpRequestHandlerOperationArgs{
		Method:  r.Method,
		URI:     r.RequestURI,
		Queries: queries,
		Headers: headers,
		Body:    "",
	}
}

func (HttpRequestHandlerOperationArgs) IsArgOf(*HttpRequestHandlerOperation) {}

func RegisterHTTPRequestSecurity(op emitter.Operation) {
	emitter.On(event.HTTP_REQUEST_EVENT, op, OnHttpRequest)
}

func OnHttpRequest(op *HttpRequestHandlerOperation, args HttpRequestHandlerOperationArgs) {
	/*
		for k, v := range args.Headers {
			err := sqli.Inspect(sqli.SQLiInspectArgs{Value: v})
			fmt.Println("Inspecting header", k, ":", v, "=>", err)
		}
	*/
}
