package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env struct {
	Client *mongo.Client
}

func (e Env) Routes(r *gin.Engine) {
	userApi := userApi{
		repo: &MyRepo{client: e.Client},
	}
	r.GET("/users/:id", userApi.getUserById)
}
