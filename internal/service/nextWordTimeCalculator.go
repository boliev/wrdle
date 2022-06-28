package service

import "time"

// NextWordTimeCalculator service struct
type NextWordTimeCalculator struct {
	hours   int
	minutes int
}

// CreateNextWordTimeCalculator NextWordTimeCalculator constructor
func CreateNextWordTimeCalculator(hours int, minutes int) *NextWordTimeCalculator {
	return &NextWordTimeCalculator{
		hours:   hours,
		minutes: minutes,
	}
}

// GetTimeForNextWord returns time for next word of the day
func (s NextWordTimeCalculator) GetTimeForNextWord() time.Time {
	now := time.Now().UTC()
	newWordTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		s.hours,
		s.minutes,
		0,
		0,
		time.UTC,
	)

	if newWordTime.Before(now) {
		newWordTime = newWordTime.AddDate(0, 0, 1)
	}

	return newWordTime
}
