package mysql

import (
	"github.com/boliev/wrdle/internal/domain"
	"github.com/boliev/wrdle/internal/repository"
	"gorm.io/gorm"
)

// NounRepository struct
type NounRepository struct {
	db *gorm.DB
}

// CreateNounRepository Mysql UserRepository constructor
func CreateNounRepository(db *gorm.DB) repository.Noun {
	return &NounRepository{
		db: db,
	}
}

// FindForPuzzle find a random word for puzzle
func (r NounRepository) FindForPuzzle(exclude []string) (*domain.Noun, error) {
	var noun domain.Noun
	result := r.db.
		Table("nouns").
		Where("is_for_puzzle = ? and word not in ?", true, exclude).
		Order("rand()").
		First(&noun)
	if result.Error != nil {
		return nil, result.Error
	}

	return &noun, nil
}

// Find find a word by primary key
func (r NounRepository) Find(word string) (*domain.Noun, error) {
	var noun domain.Noun
	result := r.db.
		Table("nouns").
		First(&noun, "word = ?", word)
	if result.Error != nil {
		return nil, result.Error
	}

	return &noun, nil
}
