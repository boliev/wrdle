package controller

import (
	"github.com/boliev/wrdle/internal/repository"
	"github.com/boliev/wrdle/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Game controller
type Game struct {
	WordOfTheDayRepository repository.WordOfTheDay
}

// CreateGameController Game controller constructor
func CreateGameController(WordOfTheDayRepository repository.WordOfTheDay) *Game {
	return &Game{
		WordOfTheDayRepository: WordOfTheDayRepository,
	}
}

// Start action for starting a game
func (c Game) Start(g *gin.Context) {
	word, err := c.WordOfTheDayRepository.GetCurrent()
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusBadRequest, response.CreateStartGameResponse(word))
}
