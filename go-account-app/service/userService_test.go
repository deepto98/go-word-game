package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepto98/go-word-game/go-account-app/model"
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
