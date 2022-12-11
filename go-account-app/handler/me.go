package handler

import (
	"log"
	"net/http"

	"github.com/deepto98/go-word-game/model"
	"github.com/deepto98/go-word-game/model/apperrors"
	"github.com/gin-gonic/gin"
)

//Calls services for getting an user's details
func (handler *Handler) Me(context *gin.Context) {
	user, exists := context.Get("user")

	// c.JSON(http.StatusOK, gin.H{"hi": "Its me"})
	if !exists {
		log.Printf("Unable to extract user from context", context)
		err := apperrors.NewInternalError()
		context.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	uid := user.(*model.User).UID

	u, err := handler.UserService.Get(context, uid)

	if err != nil {
		log.Printf("Unable to find user : %v\n%v", uid, err)
		err := apperrors.NewNotFoundError("user", uid.String())
		context.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": u,
	})

}
