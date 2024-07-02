package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	msg := "Hello, World!"
	r := initHanlder(msg)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type resource struct {
	msg string
}

func (r resource) xxx(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": r.msg,
	})
}

func initHanlder(msg string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	resource := resource{msg: msg}
	r.GET("/ping", resource.xxx)
	return r
}
