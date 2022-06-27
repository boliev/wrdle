package mysql

import (
	"errors"
	"github.com/boliev/wrdle/internal/domain"
	"github.com/boliev/wrdle/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// WordOfTheDayRepository struct
type WordOfTheDayRepository struct {
	db *gorm.DB
}

// CreateWordOfTheDayRepository Mysql WordOfTheDayRepository constructor
func CreateWordOfTheDayRepository(db *gorm.DB) repository.WordOfTheDay {
	return &WordOfTheDayRepository{
		db: db,
	}
}

// SetNew sets new word of the day
func (r WordOfTheDayRepository) SetNew(word string) *domain.WordOfTheDay {
	newWordOfTheDay := domain.WordOfTheDay{
		Word: word,
	}
	r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&newWordOfTheDay)

	return &newWordOfTheDay
}

// GetLast returns last count words of the day
func (r WordOfTheDayRepository) GetLast(count int) ([]domain.WordOfTheDay, error) {
	var words []domain.WordOfTheDay
	result := r.db.Table("word_of_the_days").Order("id DESC").Limit(count).Find(&words)

	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

// GetCurrent returns current word of the day
func (r WordOfTheDayRepository) GetCurrent() (*domain.WordOfTheDay, error) {
	words, err := r.GetLast(1)
	if err != nil {
		return nil, err
	}
	if len(words) < 1 {
		return nil, errors.New("there is no word of the day")
	}

	return &words[0], nil
}
