package repository

import "github.com/boliev/wrdle/internal/domain"

// Noun repository interface
type Noun interface {
	FindForPuzzle(exclude []string) (*domain.Noun, error)
	Find(word string) (*domain.Noun, error)
}
