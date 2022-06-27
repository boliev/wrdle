package main

import (
	"fmt"
	"github.com/boliev/wrdle"
	"github.com/boliev/wrdle/internal/controller"
	"github.com/boliev/wrdle/internal/domain"
	"github.com/boliev/wrdle/internal/mysql"
	"github.com/boliev/wrdle/internal/service"
	"github.com/boliev/wrdle/pkg/config"
	"github.com/go-co-op/gocron"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	cfg := DiCreateConfig()
	nounRepository := mysql.CreateNounRepository(DiCreateDB(cfg))
	wordOfTheDayRepository := mysql.CreateWordOfTheDayRepository(DiCreateDB(cfg))
	wordOfTheDaySetter := service.CreateWordOfTheDaySetter(wordOfTheDayRepository, nounRepository)
	nounChecker := service.CreateNounChecker()
	gameController := controller.CreateGameController(wordOfTheDayRepository, nounRepository, nounChecker)
	wordOfTheDayController := controller.CreateWordOfTheDayController(wordOfTheDaySetter)

	// crons
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(1).Day().At(cfg.GetString("new_word_time_utc")).Do(func() {
		newWord, err := wordOfTheDaySetter.SetNewWord()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Word setted %s\n", newWord.Word)
	})
	cron.StartAsync()

	// API
	app := wrdle.App{
		GameController:         gameController,
		WordOfTheDayController: wordOfTheDayController,
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
	err = db.AutoMigrate(&domain.WordOfTheDay{})
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
