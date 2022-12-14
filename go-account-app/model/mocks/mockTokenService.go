package mocks

import (
	"context"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

//Mocks NewTokenPairFromUser
func (mockTokenService *MockTokenService) NewTokenPairFromUser(ctx context.Context, user *model.User, previousTokenId string) (*model.TokenPair, error) {
	ret := mockTokenService.Called(ctx, user, previousTokenId)

	var r0 *model.TokenPair
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.TokenPair)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
