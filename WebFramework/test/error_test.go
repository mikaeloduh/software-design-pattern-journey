package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"webframework/framework"
)

func TestRouter_MethodNotAllowed(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/", mockHandler)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, resp["error"], http.StatusText(http.StatusMethodNotAllowed))
}

func TestRouter_NotFound(t *testing.T) {
	r := framework.NewRouter()
	// no routes added

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, resp["error"], http.StatusText(http.StatusNotFound))
}
