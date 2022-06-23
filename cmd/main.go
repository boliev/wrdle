package main

import (
	"github.com/boliev/wrdle"
	"github.com/boliev/wrdle/internal/controller"
	"github.com/boliev/wrdle/internal/mysql"
	"github.com/boliev/wrdle/pkg/config"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := DiCreateConfig()
	nounRepository := mysql.CreateNounRepository(DiCreateDB(cfg))
	checkController := controller.CreateCheckController(nounRepository)
	app := wrdle.App{
		CheckController: checkController,
	}
	app.Start()
}

// DiCreateDB Creates DB object
func DiCreateDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysqlDriver.Open(cfg.GetString("database_dsn")), &gorm.Config{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	return db
}

// DiCreateConfig Creates Config object
func DiCreateConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
