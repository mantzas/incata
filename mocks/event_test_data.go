package mocks

import (
	"time"
)

// TestData for testing purposes
type TestData struct {
	Version   int       `json:"version"`
	Name      string    `json:"name"`
	Balance   float32   `json:"balance"`
	BirthDate time.Time `json:"birth_date"`
}

// GetTestData returns test data
func GetTestData() *TestData {
	location, _ := time.LoadLocation("Europe/Athens")

	return &TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}
}
