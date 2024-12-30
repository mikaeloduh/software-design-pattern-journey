package e2e

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/framework"
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

func Register(w http.ResponseWriter, r *http.Request) error {
	var reqData RegisterRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	respData := RegisterResponse{
		Username: reqData.Username,
		Email:    reqData.Email,
	}

	accept := r.Header.Get("Accept")
	if accept == "application/xml" {
		if err := framework.WriteObjectAsXML(w, respData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	} else {
		if err := framework.WriteObjectAsJSON(w, respData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}

	return nil
}

func TestRegisterHandlerJSON(t *testing.T) {
	router := framework.NewRouter()
	router.Handle("/register", http.MethodPost, framework.HandlerFunc(Register))

	t.Run("test register user successfully", func(t *testing.T) {
		jsonBody := `{"username": "12345", "email": "John Doe", "password": "abc"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		expectedResponse := `{"username": "12345", "email": "John Doe"}`

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
		assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
	})
}
