package service

import (
	"github.com/boliev/wrdle/internal/domain"
	"github.com/boliev/wrdle/internal/repository"
)

// WordOfTheDaySetter struct
type WordOfTheDaySetter struct {
	WordOdTheDayRepository repository.WordOfTheDay
	NounRepository         repository.Noun
}

// CreateWordOfTheDaySetter WordOfTheDaySetter constructor
func CreateWordOfTheDaySetter(
	wordOfTheDayRepository repository.WordOfTheDay,
	nounRepository repository.Noun,
) *WordOfTheDaySetter {
	return &WordOfTheDaySetter{
		WordOdTheDayRepository: wordOfTheDayRepository,
		NounRepository:         nounRepository,
	}
}

// SetNewWord sets new word of the day
func (s WordOfTheDaySetter) SetNewWord() (*domain.WordOfTheDay, error) {
	last, err := s.WordOdTheDayRepository.GetLast(100)
	if err != nil {
		return nil, err
	}
	var lastWords []string
	for _, w := range last {
		lastWords = append(lastWords, w.Word)
	}

	word, err := s.NounRepository.FindForPuzzle(lastWords)
	if err != nil {
		return nil, err
	}
	newWord := s.WordOdTheDayRepository.SetNew(word.Word)

	return newWord, nil
}
