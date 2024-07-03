package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	userApi := userApi{}
	r.GET("/users/:id", userApi.getUserById)
}
