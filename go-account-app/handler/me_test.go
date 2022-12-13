package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/deepto98/go-word-game/go-account-app/model/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "gotest.tools/v3/assert"
)

//Unit tests for me handler

func TestMe(t *testing.T) {

	//Setup
	gin.SetMode(gin.TestMode)

	//1. Test for success - user exists case
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		//mock response
		mockUserResponse := &model.User{
			UID:   uid,
			Email: "deepto@abcd.com",
			Name:  "Deepto",
		}

		mockUserService := new(mocks.MockUserService)

		//Define test
		mockUserService.On(
			"Get", mock.AnythingOfType("*gin.Context"), uid).
			Return(mockUserResponse, nil)

		//a response recorder
		responseRecorder := httptest.NewRecorder()

		//Middleware to set context for test

		//Serve http
		router := gin.Default()

		//Add user to request
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)

		assert.NoError(t, err)

		router.ServeHTTP(responseRecorder, request)

		//Define expected response
		respBody, err := json.Marshal(gin.H{
			"user": mockUserResponse,
		})

		assert.NoError(t, err)
		assert.Equal(t, 200, responseRecorder.Code)

		//responseRecorder.Body - actual response
		assert.Equal(t, respBody, responseRecorder.Body.Bytes())

		mockUserService.AssertExpectations(t)
	})

	//2. Test for failure - no user in context
	t.Run("Success", func(t *testing.T) {

		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Get", mock.Anything, mock.Anything).
			Return(nil, nil)

		responseRecorder := httptest.NewRecorder()

		//Middleware to set context for test
		router := gin.Default()

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)

		assert.NoError(t, err)

		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, 500, responseRecorder.Code)

		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	//3. Test for failure - user not found
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Get", mock.Anything, uid).
			Return(nil, fmt.Errorf("Some Error in call chain"))

		responseRecorder := httptest.NewRecorder()

		//Middleware to set context for test
		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(responseRecorder, request)

		respErr := apperrors.NewNotFoundError("user", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)
		assert.Equal(t, respErr.Status(), responseRecorder.Code)
		assert.Equal(t, respBody, responseRecorder.Body.Bytes())

		mockUserService.AssertExpectations(t)
	})
}
