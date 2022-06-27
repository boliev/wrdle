package service

import (
	"github.com/boliev/wrdle/internal/domain"
)

// NounChecker checks a user word against the word of the day
type NounChecker struct {
}

// CreateNounChecker NounChecker constructor
func CreateNounChecker() *NounChecker {
	return &NounChecker{}
}

// Check checks a user word against the word of the day
func (c NounChecker) Check(userWord *domain.Noun, WordOfTheDay *domain.WordOfTheDay) *domain.Word {
	wod := []rune(WordOfTheDay.Word)
	word := []rune(userWord.Word)
	hash := map[int32][]int{}
	resultWord := map[int]*domain.Character{}
	for key := 0; key < 5; key++ {
		hash[wod[key]] = append(hash[wod[key]], key)
	}

	for key := 0; key < 5; key++ {
		state := domain.NotPresent
		keys, present := hash[word[key]]
		if present {
			state = domain.Present
			for _, v := range keys {
				if v == key {
					state = domain.InPlace
				}
			}
		}
		resultWord[key] = &domain.Character{Character: word[key], State: state}
	}

	return &domain.Word{
		Characters: resultWord,
	}
}
