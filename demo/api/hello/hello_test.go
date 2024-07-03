package hello_test

import (
	"api/hello"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func initHanlder(msg string) *gin.Engine {
	r := gin.New()
	// Add hello route
	hello.Routes(r, msg)
	return r
}

func TestSuccess_with_GET_ping(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := initHanlder("pong")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
