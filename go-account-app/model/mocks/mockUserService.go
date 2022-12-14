package mocks

import (
	"context"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

//mock of userService.Get
func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {

	ret := m.Called(ctx, uid)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return r0, r1
}

//mock of userService.Signup
func (m *MockUserService) Signup(ctx context.Context, user *model.User) error {

	ret := m.Called(ctx, user)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
