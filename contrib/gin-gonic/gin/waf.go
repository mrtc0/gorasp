package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtc0/gowaf/emitter"
)

func waf(c *gin.Context) {
	var params map[string]string

	for _, v := range c.Params {
		params[v.Key] = v.Value
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Request = r
		c.Next()
	})

	emitter.WrapHandler(handler, params).ServeHTTP(c.Writer, c.Request)
}

func WafMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		waf(c)
		c.Next()
	}
}
