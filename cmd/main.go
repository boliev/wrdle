package main

import (
	"github.com/boliev/wrdle"
	"github.com/boliev/wrdle/internal/controller"
)

func main() {
	checkController := controller.CreateCheckController()
	app := wrdle.App{
		CheckController: checkController,
	}
	app.Start()
}
