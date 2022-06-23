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
	// TODO: implement exclude
	result := r.db.Table("nouns").Where("is_for_puzzle = ? ", true).Order("rand()").First(&noun)
	if result.Error != nil {
		return nil, result.Error
	}

	return &noun, nil
}
