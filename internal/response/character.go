package response

import (
	"fmt"
	"github.com/boliev/wrdle/internal/domain"
)

// Character response struct
type Character struct {
	Character string `json:"char"`
	State     string `json:"state"`
}

// CreateCharacterFromDomain creates character response from domain character
func createCharacterFromDomain(char *domain.Character) *Character {
	return &Character{
		Character: fmt.Sprintf("%c", char.Character),
		State:     char.State.String(),
	}
}
