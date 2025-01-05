package e2e

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
	"webframework/framework"
)

type User struct {
	Id       uint64
	Username string
	Email    string
	Password string
}

type UserService struct {
	store  map[uint64]*User
	nextId uint64
}

func NewUserService() *UserService {
	return &UserService{
		store:  make(map[uint64]*User),
		nextId: 1,
	}
}

func (u *UserService) CreateUser(username string, email string, password string) (*User, error) {
	user := User{
		Id:       u.nextId,
		Username: username,
		Email:    email,
		Password: password,
	}

	u.store[user.Id] = &user
	u.nextId++

	return &user, nil
}

func (u *UserService) FindUserByEmail(email string) *User {
	for _, user := range u.store {
		if user.Email == email {
			return user
		}
	}

	return nil
}

type UserController struct {
	UserService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

type RegisterResponse struct {
	Id       uint64 `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) error {
	var reqData RegisterRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if reqData.Username == "" || reqData.Email == "" || reqData.Password == "" {
		return errors.NewError(http.StatusBadRequest, fmt.Errorf("Registration's format incorrect."))
	}

	if c.UserService.FindUserByEmail(reqData.Email) != nil {
		return errors.NewError(http.StatusBadRequest, fmt.Errorf("Duplicate email"))
	}

	user, _ := c.UserService.CreateUser(reqData.Username, reqData.Email, reqData.Password)

	respData := RegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
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
	userController := NewUserController(userService)
	router := framework.NewRouter()
	router.Handle("/register", http.MethodPost, framework.HandlerFunc(userController.Register))

	t.Run("test register user successfully", func(t *testing.T) {
		jsonBody := `{"username": "John Doe", "email": "jdoe@example.com", "password": "abc"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		expectedResponse := `{"username": "John Doe", "email": "jdoe@example.com", "id": 2}`

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Expected Content-Type application/json")
		assert.JSONEq(t, expectedResponse, rr.Body.String(), "Response body mismatch")
	})

	t.Run("register fail: email exists", func(t *testing.T) {
		jsonBody := `{"username": "John Doe", "email": "jdoe@example.com", "password": "abc"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status BadRequest")
		assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"), "Expected Content-Type text/plain")
		assert.Equal(t, "Duplicate email", rr.Body.String(), "Response body mismatch")
	})

	t.Run("register fail: invalid format", func(t *testing.T) {
		jsonBody := `{"incorrectrequest": "this is a incorrect test request"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status BadRequest")
		assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"), "Expected Content-Type text/plain")
		assert.Equal(t, "Registration's format incorrect.", rr.Body.String(), "Response body mismatch")
	})
}
