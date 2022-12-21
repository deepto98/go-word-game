package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/deepto98/go-word-game/go-account-app/model/mocks"

	// "github.com/deepto98/go-word-game/go-account-app/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	//Success case
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		mockUserResponse := &model.User{
			UID:   uid,
			Email: "deepto@abc.com",
			Name:  "Deepto",
		}
		mockUserRepository := new(mocks.MockUserRepository)

		userService := NewUserService(&UserConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).
			Return(mockUserResponse, nil)

		ctx := context.TODO()
		user, err := userService.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, user, mockUserResponse)
		mockUserRepository.AssertExpectations(t)
	})

	//Error case
	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserRepository := new(mocks.MockUserRepository)

		userService := NewUserService(&UserConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).
			Return(nil, fmt.Errorf("Some error down call chain"))

		ctx := context.TODO()
		user, err := userService.Get(ctx, uid)

		assert.Nil(t, user)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestSignup(t *testing.T) {
	//Case for successful signup
	t.Run("Success", func(t *testing.T) {
		userId, _ := uuid.NewRandom()

		mockUser := &model.User{
			Email:    "test@test.test",
			Password: "abcd1234",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		userService := NewUserService(&UserConfig{
			UserRepository: mockUserRepository,
		})

		mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
			Run(func(args mock.Arguments) {
				userArg := args.Get(1).(*model.User)
				userArg.UID = userId
			}).Return(nil)

		ctx := context.TODO()
		err := userService.Signup(ctx, mockUser)

		assert.NoError(t, err)

		assert.Equal(t, userId, mockUser.UID)

		mockUserRepository.AssertExpectations(t)

	})

	//Test case for unsuccessful signup
	t.Run("Error", func(t *testing.T) {
		mockUser := &model.User{
			Email:    "test@test.test",
			Password: "abcd1234",
		}

		mockUserRepository := new(mocks.MockUserRepository)

		userService := NewUserService(&UserConfig{
			UserRepository: mockUserRepository,
		})

		mockError := apperrors.NewConflictError("email", mockUser.Email)

		mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
			Return(mockError)

		ctx := context.TODO()
		err := userService.Signup(ctx, mockUser)

		assert.EqualError(t, err, mockError.Error())

		mockUserRepository.AssertExpectations(t)
	})
}
