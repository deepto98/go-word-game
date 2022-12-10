package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	Router *gin.Engine
}

//Factory method to initialize handler with injected services and http routea
func NewHandler(config *Config) {

	handler := &Handler{}

	accountGroup := config.Router.Group("/api/account")

	accountGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "Hello"})
	})
}
