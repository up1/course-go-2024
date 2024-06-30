package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup the router, similar to how it's done in the main function
	router := setupHandler()

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Create a new request to the "/health" endpoint
	req, _ := http.NewRequest("GET", "/health", nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, w.Code)

	// Unmarshal the response body into a map
	var responseBody map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Could not unmarshal response body: %v", err)
	}

	// Check the "message" field in the response body
	expectedMessage := "healthy"
	if assert.Contains(t, responseBody, "message", "Response body does not contain 'message' key") {
		assert.Equal(t, expectedMessage, responseBody["message"], "The 'message' field does not match")
	}
}
