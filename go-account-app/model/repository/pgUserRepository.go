package repository

import (
	"log"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/net/context"
)

//Repository implementation of service layer UserRepository
type pgUserRepository struct {
	DB *sqlx.DB
}

//Factory method for initializing User Repositories
func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &pgUserRepository{
		DB: db,
	}
}

//Reaches out to database SQLX api
func (pgUserRepository *pgUserRepository) Create(ctx context.Context, user *model.User) error {

	createQuery := "INSERT into users (email,password) 	VALUES ($1,$2) RETURNING *"

	if err := pgUserRepository.DB.Get(user, createQuery, user.Email, user.Password); err != nil {

		//check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v\n", user.Email, err.Code.Name())
			return apperrors.NewConflictError("email", user.Email)
		}
		log.Printf("Could not create a user with email: %v. Reason: %v\n", user.Email, err)
		return apperrors.NewInternalError()
	}
	return nil
}

// FindByID fetches user by id
func (pgUserRepository *pgUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid=$1"

	// we need to actually check errors as it could be something other than not found
	if err := pgUserRepository.DB.Get(user, query, uid); err != nil {
		return user, apperrors.NewNotFoundError("uid", uid.String())
	}

	return user, nil
}
