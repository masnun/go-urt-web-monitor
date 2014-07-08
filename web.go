package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gout "github.com/masnun/gout/library"
	"net/http"
)

func StartWebServer(port string, data *gout.Server) {
	app := gin.Default()

	app.LoadHTMLTemplates("static/html/*")

	// Home page
	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// Parse the response as JSON
	app.GET("/api/json", func(c *gin.Context) {
		c.JSON(200, *data)
	})

	// Serve static files
	app.ServeFiles("/static/*filepath", http.Dir("static"))

	// Listen and server on 0.0.0.0:8080
	fmt.Println("Starting web server on: " + port)
	app.Run(":" + port)
}
