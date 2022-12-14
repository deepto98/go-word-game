package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/deepto98/go-word-game/go-account-app/model/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {

	//Setup
	gin.SetMode(gin.TestMode)

	//1. Test for email and password required
	t.Run("Email and Password Required", func(t *testing.T) {

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		//Create req body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, 400, responseRecorder.Code)
		mockUserService.AssertNotCalled(t, "Signup")

	})

	//2. Test for invalid email
	t.Run("Invalid email", func(t *testing.T) {

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		//Create req body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "",
			"password": "pwdsssss",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, 400, responseRecorder.Code)
		mockUserService.AssertNotCalled(t, "Signup")

	})

	//3. Test for short password
	t.Run("Short Password", func(t *testing.T) {

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		//Create req body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "abc@abc.com",
			"password": "pwd",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, 400, responseRecorder.Code)
		mockUserService.AssertNotCalled(t, "Signup")

	})

	//4. Test for errors thrown by UserService eg : User exists
	t.Run("Error returned from UserService", func(t *testing.T) {

		user := &model.User{
			Email:    "abc@abc.com",
			Password: "abcdef",
		}
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).
			Return(apperrors.NewConflictError("User already exists", user.Email))

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		//Create req body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    user.Email,
			"password": user.Password,
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, 409, responseRecorder.Code)
		mockUserService.AssertExpectations(t)

	})
}
