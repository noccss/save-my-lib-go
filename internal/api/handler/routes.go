package routes

import (
	"github.com/Noccss/save-my-lib/internal/api/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	userHandler := user.NewUserHandler()
	userHandler.InitRoutes(router)
}
