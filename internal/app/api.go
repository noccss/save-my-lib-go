package app

import (
	"github.com/Noccss/save-my-lib/internal/api"
	"github.com/Noccss/save-my-lib/internal/drivers/logs"
)

func StartApi() {

	logs.InitLogrus()

	api.InitRoutes()
}
