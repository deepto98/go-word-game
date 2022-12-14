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

	//5. Test for Successful Token Creation
	t.Run("Successful Token Creation", func(t *testing.T) {

		user := &model.User{
			Email:    "abc@abc.com",
			Password: "abcdef",
		}
		mockTokenResponse := &model.TokenPair{
			IDToken:      "idToken",
			RefreshToken: "refreshToken",
		}

		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), user).
			Return(nil)
		mockTokenService.On("NewTokenPairFromUser", mock.AnythingOfType("*gin.Context"), user, "").
			Return(mockTokenResponse, nil)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:       router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
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

		//Define Expected Response Body
		responseBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResponse,
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, responseRecorder.Code)
		assert.Equal(t, responseBody, responseRecorder.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	//6. Test for Failed Token Creation
	t.Run("Successful Token Creation", func(t *testing.T) {

		user := &model.User{
			Email:    "abc@abc.com",
			Password: "abcdef",
		}

		mockErrorResponse := apperrors.NewInternalError()

		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), user).
			Return(nil)
		mockTokenService.On("NewTokenPairFromUser", mock.AnythingOfType("*gin.Context"), user, "").
			Return(nil, mockErrorResponse)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			Router:       router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
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

		//Define Expected Response Body
		responseBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})

		assert.NoError(t, err)
		assert.Equal(t, mockErrorResponse.Status(), responseRecorder.Code)
		assert.Equal(t, responseBody, responseRecorder.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
