package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtc0/gorasp"
	ginWaf "github.com/mrtc0/gorasp/contrib/gin-gonic/gin"
)

func main() {
	gorasp.Start()
	r := gin.Default()
	r.Use(ginWaf.WafMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
