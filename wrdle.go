package wrdle

import (
	"github.com/boliev/wrdle/internal/controller"
	"github.com/gin-gonic/gin"
)

// App the app
type App struct {
	CheckController        *controller.Check
	WordOfTheDayController *controller.WordOfTheDay
}

// Start the app
func (app App) Start() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		check := v1.Group("/check")
		{
			check.GET("/:word", app.CheckController.Check)
		}
		wordOfTheDay := v1.Group("/word-of-the-day")
		{
			wordOfTheDay.POST("/", app.WordOfTheDayController.Set)
		}
	}
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
