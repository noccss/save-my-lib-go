package api

import (
	"github.com/Noccss/save-my-lib/internal/drivers/http"
	"github.com/sirupsen/logrus"
)

func InitRoutes() {
	logrus.Info("Init API")

	server := http.New()

	go server.Run()
}
