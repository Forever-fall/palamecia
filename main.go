package main

import (
	"palamecia/config"
	"palamecia/router"
)

func main() {
	config.InitConfig()
	r := router.InitRouter()
	port := ":" + config.AppConfig.App.Port

	r.Run(port)
}
