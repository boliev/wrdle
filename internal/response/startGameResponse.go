package response

import "github.com/boliev/wrdle/internal/domain"

// StartGameResponse http response for /game/start endpoint
type StartGameResponse struct {
	ID int `json:"id"`
}

// CreateStartGameResponse StartGameResponse constructor
func CreateStartGameResponse(wordOfTheDay *domain.WordOfTheDay) *StartGameResponse {
	return &StartGameResponse{
		ID: int(wordOfTheDay.ID),
	}
}
