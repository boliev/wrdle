package response

import "github.com/boliev/wrdle/internal/domain"

// Word response struct
type Word struct {
	Characters map[int]*Character
}

// CreateWordFromDomain creates users list response from domain users list
func CreateWordFromDomain(word *domain.Word) *Word {
	characters := map[int]*Character{}
	for position, character := range word.Characters {
		characters[position] = createCharacterFromDomain(character)
	}

	return &Word{
		Characters: characters,
	}
}
