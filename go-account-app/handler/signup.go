package handler

import (
	"log"

	"github.com/deepto98/go-word-game/go-account-app/model"
	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/deepto98/go-word-game/go-account-app/utils"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password,gte=66,lte=30"`
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
}
