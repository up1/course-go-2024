package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	msg string
}

func (r resource) sayHi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": r.msg,
	})
}
