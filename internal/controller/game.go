package controller

import (
	"errors"
	"github.com/boliev/wrdle/internal/repository"
	"github.com/boliev/wrdle/internal/request"
	"github.com/boliev/wrdle/internal/response"
	"github.com/boliev/wrdle/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Game controller
type Game struct {
	WordOfTheDayRepository repository.WordOfTheDay
	NounRepository         repository.Noun
	NounChecker            *service.NounChecker
	NextWordTime           time.Time
}

// CreateGameController Game controller constructor
func CreateGameController(
	wordOfTheDayRepository repository.WordOfTheDay,
	nounRepository repository.Noun,
	nounChecker *service.NounChecker,
	nextWordTime time.Time,
) *Game {
	return &Game{
		WordOfTheDayRepository: wordOfTheDayRepository,
		NounRepository:         nounRepository,
		NounChecker:            nounChecker,
		NextWordTime:           nextWordTime,
	}
}

// Start action for starting a game
func (c Game) Start(g *gin.Context) {
	word, err := c.WordOfTheDayRepository.GetCurrent()
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusBadRequest, response.CreateStartGameResponse(word, c.NextWordTime))
}

// Check action for checking user word against word of the day
func (c Game) Check(g *gin.Context) {
	var checkRequest request.CheckRequest
	if err := g.ShouldBindUri(&checkRequest); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	word, err := c.WordOfTheDayRepository.GetCurrent()
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if checkRequest.ID != int(word.ID) {
		g.JSON(http.StatusConflict, errors.New("word was changed"))
		return
	}

	userWord, err := c.NounRepository.Find(checkRequest.Word)
	if err != nil {
		g.JSON(http.StatusConflict, err.Error())
		return
	}

	checkResult := c.NounChecker.Check(userWord, word)
	g.JSON(http.StatusBadRequest, response.CreateWordFromDomain(checkResult))
}
