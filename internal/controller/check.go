package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Check controller struct
type Check struct {
}

// CreateCheckController controller for /check/
func CreateCheckController() *Check {
	return &Check{}
}

// Check action for [get] /check/:word
func (c Check) Check(g *gin.Context) {
	wordForTest := g.Param("word")
	g.JSON(http.StatusOK, wordForTest)
}
