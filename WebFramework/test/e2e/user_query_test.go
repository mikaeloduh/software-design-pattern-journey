package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
	"webframework/framework"
)

type UserQueryResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UserQuery(w http.ResponseWriter, r *http.Request) error {
	res := UserQueryResponse{
		Username: "correctName",
		Email:    "q4o5D@example.com",
	}
	if err := framework.WriteObjectAsJSON(w, res); err != nil {
		return err
	}

	return nil
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next func()) error {
	token := r.Header.Get("Authorization")
	if token != "secret" {
		return errors.ErrorTypeUnauthorized
	}

	next()
	return nil
}

func TestUserQuery(t *testing.T) {

	router := framework.NewRouter()
	router.Use(AuthMiddleware)
	router.Handle("/query", http.MethodGet, framework.HandlerFunc(UserQuery))

	t.Run("test query user successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/query", nil)
		req.Header.Set("Authorization", "secret")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		expectedResponse := `{"username":"correctName","email":"q4o5D@example.com"}`

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
		assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
	})

	t.Run("test query user with invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/query", nil)
		req.Header.Set("Authorization", "invalid")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected status Unauthorized")
	})
}
