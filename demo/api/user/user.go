package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userApi struct {
}

func (u *userApi) getUserById(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{})
}
