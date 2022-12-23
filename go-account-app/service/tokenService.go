package service

import (
	"context"
	"crypto/rsa"

	"github.com/deepto98/go-word-game/go-account-app/model"
)

//Used to inject an implementation of TokenRepository for use in service methods
type TokenService struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	RefreshSecret string
}

//TokenServiceConfig will have repositories that will be injected into the service layer
type TokenServiceConfig struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	RefreshSecret string
}

func (ts *TokenService) NewTokenPairFromUser(ctx context.Context, u *model.User, previousTokenID string) (*model.TokenPair, error)

//Factory function to initialize TokenService with repository layer dependencies
func NewTokenService(config *TokenServiceConfig) model.TokenService {
	return &TokenService{
		PrivateKey:    config.PrivateKey,
		PublicKey:     config.PublicKey,
		RefreshSecret: config.RefreshSecret,
	}
}
