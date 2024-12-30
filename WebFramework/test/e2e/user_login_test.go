package e2e

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/framework"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Login(w http.ResponseWriter, r *http.Request) error {
	var reqData LoginRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		return err
	}

	if reqData.Username == "correctName" && reqData.Password == "correctPassword" {
		respData := LoginResponse{
			Username: reqData.Username,
			Email:    "q4o5D@example.com",
		}
		if err := framework.WriteObjectAsJSON(w, respData); err != nil {
			return err
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	return nil
}

func TestUserLogin(t *testing.T) {

	router := framework.NewRouter()
	router.Handle("/login", http.MethodPost, framework.HandlerFunc(Login))

	t.Run("test login successfully", func(t *testing.T) {
		jsonBody := `{"username": "correctName", "password": "correctPassword"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		expectedResponse := `{"username": "correctName", "email": "q4o5D@example.com"}`

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
		assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
	})
}
