package model

import (
	"context"

	"github.com/google/uuid"
)

//Defines methods the handler layer expects
//from any service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, user *User) error
}

//Defines methods the handler layer expects to interact with
//while producing JWTs as strings
type TokenService interface {
	NewTokenPairFromUser(ctx context.Context, u *User, previousTokenID string) (*TokenPair, error)
}

//Defines methods the service layer expects
//from any repository  it interacts with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	Create(ctx context.Context, user *User) error
}
