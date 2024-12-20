package test

import (
	"fmt"
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

func Register(c *framework.Context) {
	var reqData RegisterRequest
	if err := c.ReadBodyAsObject(&reqData); err != nil {
		c.AbortWithError(framework.NewError(framework.ErrorTypeBadRequest, "invalid token", err))
		return
	}

	respData := RegisterResponse{
		Username: reqData.Username,
		Email:    reqData.Email,
	}

	switch c.Request.Header.Get("Accept") {
	case "application/xml":
		c.Xml(respData)
	case "application/json":
		c.JSON(respData)
	default:
		c.AbortWithError(framework.NewError(framework.ErrorTypeBadRequest, "unsupported accept", fmt.Errorf("accept header must be application/json")))
	}
}

func TestRegisterHandlerJSON(t *testing.T) {
	jsonBody := `{"username": "12345", "email": "John Doe", "password": "abc"}`

	req := httptest.NewRequest("POST", "/register", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()

	Register(&framework.Context{
		ResponseWriter: rr,
		Request:        req,
	})

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

	Register(&framework.Context{
		ResponseWriter: rr,
		Request:        req,
	})

	expectedResponse := `<RegisterResponse><username>12345</username><email>John Doe</email></RegisterResponse>`

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.Equal(t, "application/xml", rr.Header().Get("Content-Type"), "Expected Content-Type application/xml")
	assert.Equal(t, strings.TrimSpace(expectedResponse), strings.TrimSpace(rr.Body.String()), "Response body mismatch")
}
