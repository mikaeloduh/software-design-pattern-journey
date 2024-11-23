package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	jsonBody := `{"username": "12345", "email": "John Doe", "password": "abc"}`

	req := httptest.NewRequest("GET", "/register", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Register(rr, req)

	expectedResponse := jsonBody

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
}
