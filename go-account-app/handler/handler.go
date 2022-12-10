package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	Router *gin.Engine
}

//Factory method to initialize handler with injected services and http routea
func NewHandler(config *Config) {

	handler := &Handler{}

	accountGroup := config.Router.Group(os.Getenv("ACCOUNT_API_URL"))

	accountGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "Hello"})
	})
	accountGroup.GET("/me", handler.Me)
	accountGroup.POST("/signup", handler.Signup)
	accountGroup.POST("/signin", handler.Signin)
	accountGroup.POST("/signout", handler.Signout)
	accountGroup.POST("/tokens", handler.Tokens)
	accountGroup.POST("/image", handler.Image)
	accountGroup.DELETE("/image", handler.DeleteImage)
	accountGroup.PUT("/details", handler.Details)
}

func (handler *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Its me"})

}

func (handler *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}

func (handler *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}

func (handler *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed up"})
}

func (handler *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}

func (handler *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}

func (handler *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}

func (handler *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hi": "Signed in"})
}
