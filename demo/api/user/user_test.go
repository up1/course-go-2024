package user_test

import (
	"api/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func initHanlder() *gin.Engine {
	r := gin.New()
	// Add hello route
	user.Routes(r)
	return r
}

func TestSuccess_with_GET_user_by_id(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := initHanlder()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
