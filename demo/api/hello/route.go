package hello

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine, msg string) {
	resource := resource{msg: msg}
	r.GET("/ping", resource.sayHi)
}
