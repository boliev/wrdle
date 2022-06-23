package domain

// Noun domain model
type Noun struct {
	Word     string `gorm:"primarykey"`
	isPuzzle bool   `gorm:"column:is_for_puzzle"`
}
