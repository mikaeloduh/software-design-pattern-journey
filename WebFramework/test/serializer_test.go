package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webframework/framework"

	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

type RegisterResponse struct {
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var reqData RegisterRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respData := RegisterResponse{
		Username: reqData.Username,
		Email:    reqData.Email,
	}

	accept := r.Header.Get("Accept")
	if accept == "application/xml" {
		if err := framework.WriteObjectAsXML(w, respData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := framework.WriteObjectAsJSON(w, respData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TestRegisterHandlerJSON(t *testing.T) {
	jsonBody := `{"username": "12345", "email": "John Doe", "password": "abc"}`

	req := httptest.NewRequest("POST", "/register", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Register(rr, req)

	expectedResponse := `{"username": "12345", "email": "John Doe"}`

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

	expectedResponse := `<RegisterResponse><username>12345</username><email>John Doe</email></RegisterResponse>`

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.Equal(t, "application/xml", rr.Header().Get("Content-Type"), "Expected Content-Type application/xml")
	assert.Equal(t, strings.TrimSpace(expectedResponse), strings.TrimSpace(rr.Body.String()), "Response body mismatch")
}
