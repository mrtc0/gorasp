package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpHandler "github.com/mrtc0/gorasp/handler/http"
)

func handle(c *gin.Context) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Request = r
		c.Next()
	})

	httpHandler.WrapHandler(handler).ServeHTTP(c.Writer, c.Request)
}

func WafMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(c)
		c.Next()
	}
}
