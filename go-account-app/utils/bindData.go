package utils

import (
	"log"

	"github.com/deepto98/go-word-game/go-account-app/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InvalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// Binds an incoming request to a struct
func BindData(ctx *gin.Context, req interface{}) bool {

	err := ctx.ShouldBind(req)

	if err != nil {
		log.Printf("Error binding data: %+v\n", err)

		//Check if error is a validator error
		errors, ok := err.(validator.ValidationErrors)

		if ok {
			var invalidArgs []InvalidArgument

			for _, err := range errors {
				invalidArgs = append(invalidArgs, InvalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			err := apperrors.NewBadRequestError("Invalid Request Parameters. See invalidArgs")

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}

		//Fallback error if the error isn't a validation error
		fallbackError := apperrors.NewInternalError()
		ctx.JSON(fallbackError.Status(), gin.H{
			"error": fallbackError,
		})
		return false
	}

	return true
}
