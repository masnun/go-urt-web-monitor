package main

import (
	"github.com/gin-gonic/gin"
	gout "github.com/masnun/gout/library"
	"fmt"
)

func StartWebServer(data *gout.Server) {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(200, "Nothing here yet!\n Look at: /api/json")
	})

	// Parse the response as JSON
	app.GET("/api/json", func(c *gin.Context) {
		c.JSON(200, *data)
	})

	// Listen and server on 0.0.0.0:8080
	fmt.Println("Starting web server")
	app.Run(":8080")
}
