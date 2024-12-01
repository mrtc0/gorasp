package emitter

import (
	"context"
	"net/http"
)

type Params struct {
	Method     string
	RequestURI string
	Host       string
	RemoteAddr string
	PathParams map[string]string
}

func emit(ctx context.Context, params Params) {
	// emit to event listener
	emitter := GetEventEmitter()
	emitter.Emit("http", params)
}

func wafHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) (http.ResponseWriter, *http.Request) {
	emit(r.Context(), Params{
		Method:     r.Method,
		RequestURI: r.RequestURI,
		Host:       r.Host,
		RemoteAddr: r.RemoteAddr,
		PathParams: pathParams,
	})

	return w, r
}

func WrapHandler(handler http.Handler, pathParams map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww, wr := wafHandler(w, r, pathParams)

		/* hook
		defer func() {
			// e.g. log request
		}()
		*/
		handler.ServeHTTP(ww, wr)
	})
}
