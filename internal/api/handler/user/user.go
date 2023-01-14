package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (c UserHandler) InitRoutes(router *gin.RouterGroup) {
	router.GET("/user", c.getAllUsers)
}

func (c UserHandler) getAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Nice"})
}
