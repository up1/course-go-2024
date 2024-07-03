package main

import (
	"api/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	msg := "Hello, World!"
	// Setup handlers
	r := initHanlder(msg)
	// Run server
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initHanlder(msg string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Add hello route
	hello.Routes(r, msg)

	return r
}
