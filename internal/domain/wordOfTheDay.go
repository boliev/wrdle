package domain

// WordOfTheDay domain model
type WordOfTheDay struct {
	ID   uint `gorm:"primarykey"`
	Word string
}
