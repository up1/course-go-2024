package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := initHanlder()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initHanlder() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", xxx)
	return r
}

func xxx(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
