package repository

import "github.com/boliev/wrdle/internal/domain"

// WordOfTheDay repository interface
type WordOfTheDay interface {
	SetNew(word string) *domain.WordOfTheDay
	GetLast(count int) ([]domain.WordOfTheDay, error)
	GetCurrent() (*domain.WordOfTheDay, error)
}
