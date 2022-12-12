package service

import (
	"context"

	"github.com/deepto98/go-word-game/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

type UserConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(config *UserConfig) model.UserService {
	return &UserService{
		UserRepository: config.UserRepository,
	}
}

//
func (userService *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := userService.UserRepository.FindByID(ctx, uid)
	return u, err
}
