package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
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

func Register(ctx *framework.Context) {
	var req RegisterRequest
	if err := ctx.ReadBodyAsObject(&req); err != nil {
		ctx.AbortWithError(errors.ErrorTypeBadRequest)
		return
	}

	matchString, err := regexp.MatchString("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?", req.Email)
	if err != nil {
		ctx.AbortWithError(err)
		return
	}
	if !matchString {
		ctx.AbortWithError(errors.NewError(http.StatusBadRequest, fmt.Errorf(`Registration's format incorrect.`)))
		return
	}

	if req.Email == "exist@email.com" {
		ctx.AbortWithError(errors.NewError(http.StatusBadRequest, fmt.Errorf("Duplicate email")))
		return
	}

	resp := RegisterResponse{
		Id:    "test-id",
		Email: "test@email.com",
		Name:  "TestName",
	}

	ctx.Status(http.StatusCreated)
	ctx.JSON(resp)
}

func TestUserRegistration(t *testing.T) {
	router := framework.NewRouter()
	router.HandleError(&framework.DefaultErrorHandler{})
	router.Handle(http.MethodPost, "/api/users", Register)

	t.Run("register user successfully", func(t *testing.T) {
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
	})

	t.Run("register fail: email exists", func(t *testing.T) {
		body, _ := json.Marshal(RegisterRequest{
			Email:    "exist@email.com",
			Name:     "TestName",
			Password: "test-password",
		})
		req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code 400")
		assert.Equal(t, "Duplicate email", rr.Body.String(), "Unexpected response message")
		assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"), "Unexpected response Content-Type")
	})

	t.Run("register fail: invalid format", func(t *testing.T) {
		body, _ := json.Marshal(RegisterRequest{
			Email:    "invalid-format-email",
			Name:     "TestName",
			Password: "test-password",
		})
		req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code 400")
		assert.Equal(t, "Registration's format incorrect.", rr.Body.String(), "Unexpected response message")
		assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"), "Unexpected response Content-Type")
	})
}
