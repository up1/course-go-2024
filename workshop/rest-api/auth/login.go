package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user User // Assume User is a struct that includes username and password
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Dummy check for username and password
	// Replace this with your actual database check
	if user.Username != "testuser" || user.Password != "password" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
