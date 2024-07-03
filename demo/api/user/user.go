package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userApi struct {
	repo UserRepo
}

func (u *userApi) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}

	// GET data from mongodb
	u.repo.GetById(strconv.Itoa(id))

	c.JSON(http.StatusOK, UserResponse{})
}
