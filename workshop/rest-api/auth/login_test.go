package auth_test

import (
	"api/auth"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/login", auth.Login)

	// Mock request data
	userData := `{"username":"testuser","password":"password"}`
	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(userData))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check to ensure the response status code is as expected
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code 200, got %d", w.Code)
}
