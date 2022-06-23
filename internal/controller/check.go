package controller

import (
	"github.com/boliev/wrdle/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Check controller struct
type Check struct {
	nounRepository repository.Noun
}

// CreateCheckController controller for /check/
func CreateCheckController(nounRepository repository.Noun) *Check {
	return &Check{
		nounRepository: nounRepository,
	}
}

// Check action for [get] /check/:word
func (c Check) Check(g *gin.Context) {
	wordForTest := g.Param("word")
	wordOfTheDay, err := c.nounRepository.FindForPuzzle([]string{})
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}

	g.JSON(http.StatusOK, []string{wordForTest, wordOfTheDay.Word})
}
