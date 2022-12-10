package model

import (
	"github.com/google/uuid"
)

//Defines methods the handler layer expects
//from any service it interacts with to implement
type UserService interface {
	Get(uid uuid.UUID) (*User, error)
}

//Defines methods the service layer expects
//from any repository  it interacts with to implement
type UserRepository interface {
	FindById(uid uuid.UUID) (*User, error)
}
