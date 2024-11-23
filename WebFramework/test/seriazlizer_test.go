package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandlerJSON(t *testing.T) {
	jsonBody := `{"username": "12345", "email": "John Doe", "password": "abc"}`

	req := httptest.NewRequest("POST", "/register", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Register(rr, req)

	expectedResponse := jsonBody

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
}

func TestRegisterHandlerXML(t *testing.T) {
	xmlBody := `<RegisterRequest><username>12345</username><email>John Doe</email><password>abc</password></RegisterRequest>`

	req := httptest.NewRequest("POST", "/register", strings.NewReader(xmlBody))
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")

	rr := httptest.NewRecorder()

	Register(rr, req)

	expectedResponse := xmlBody

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.Equal(t, "application/xml", rr.Header().Get("Content-Type"), "Expected Content-Type application/xml")
	assert.Equal(t, strings.TrimSpace(expectedResponse), strings.TrimSpace(rr.Body.String()), "Response body mismatch")
}
