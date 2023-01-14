package api

import (
	routes "github.com/Noccss/save-my-lib/internal/api/handler"
	"github.com/Noccss/save-my-lib/internal/drivers/http"
	"github.com/sirupsen/logrus"
)

func InitRoutes() {
	logrus.Info("Init API")

	server := http.New()

	privateGroup := server.Router.Group("api")
	routes.InitRoutes(privateGroup)

	go server.Run()
}
