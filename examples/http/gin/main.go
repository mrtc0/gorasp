package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrtc0/gowaf"
	ginwaf "github.com/mrtc0/gowaf/contrib/gin-gonic/gin"
)

func main() {
	gowaf.Start()
	r := gin.Default()
	r.Use(ginwaf.WafMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
