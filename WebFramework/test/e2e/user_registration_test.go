package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"webframework/framework"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func Register(c *framework.Context) {
	var req RegisterRequest
	if err := c.ReadBodyAsObject(&req); err != nil {
		c.AbortWithError(framework.NewError(framework.ErrorTypeBadRequest, "invalid token", err))
		return
	}

	resp := RegisterResponse{
		Id:    "test-id",
		Email: "test@email.com",
		Name:  "TestName",
	}

	c.Status(http.StatusCreated)
	if err := c.WriteJSON(resp); err != nil {
		c.AbortWithError(framework.NewError(framework.ErrorTypeInternalServerError, "write response error", err))
		return
	}
}

func TestUserRegistration(t *testing.T) {
	router := framework.NewRouter()
	router.Handle(http.MethodPost, "/api/users", Register)

	body, _ := json.Marshal(RegisterRequest{
		Email:    "test@email.com",
		Name:     "TestName",
		Password: "test-password",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201 Created")

	// Assert the response body
	var bw bytes.Buffer
	_ = json.NewEncoder(&bw).Encode(RegisterResponse{
		Id:    "test-id",
		Email: "test@email.com",
		Name:  "TestName",
	})
	assert.Equal(t, bw.String(), rr.Body.String(), "Unexpected response message")

	// Assert the content-type
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Unexpected response Content-Type")
}
