package handler

import (
	"log"
	"net/http"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/deepto98/go-word-game/go-account-app/utils"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (handler *Handler) Signup(ctx *gin.Context) {
	var request SignupRequest

	//Bind incoming request to SignupRequest struct
	if ok := utils.BindData(ctx, &request); !ok {
		return
	}

	user := &model.User{
		Email:    request.Email,
		Password: request.Password,
	}

	err := handler.UserService.Signup(ctx, user)
	if err != nil {
		log.Printf("Failed to signup user: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	//Create token pair for user
	tokenPair, err := handler.TokenService.NewTokenPairFromUser(ctx, user, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	//For successful token creation
	ctx.JSON(http.StatusCreated, gin.H{
		"tokens": tokenPair,
	})
}
