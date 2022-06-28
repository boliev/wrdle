package response

import (
	"github.com/boliev/wrdle/internal/domain"
	"time"
)

// StartGameResponse http response for /game/start endpoint
type StartGameResponse struct {
	ID          int       `json:"id"`
	NewWordTime time.Time `json:"newWordTime"`
}

// CreateStartGameResponse StartGameResponse constructor
func CreateStartGameResponse(wordOfTheDay *domain.WordOfTheDay, nextWordTime time.Time) *StartGameResponse {
	return &StartGameResponse{
		ID:          int(wordOfTheDay.ID),
		NewWordTime: nextWordTime,
	}
}
