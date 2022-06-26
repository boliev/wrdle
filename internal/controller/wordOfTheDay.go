package controller

import (
	"github.com/boliev/wrdle/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WordOfTheDay controller struct
type WordOfTheDay struct {
	WordOfTheDaySetter *service.WordOfTheDaySetter
}

// CreateWordOfTheDayController constructor
func CreateWordOfTheDayController(WordOfTheDaySetter *service.WordOfTheDaySetter) *WordOfTheDay {
	return &WordOfTheDay{
		WordOfTheDaySetter: WordOfTheDaySetter,
	}
}

// Set action for setting word of the day
func (c WordOfTheDay) Set(g *gin.Context) {
	_, err := c.WordOfTheDaySetter.SetNewWord()
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}

	g.Status(http.StatusOK)
}
