// Package mock is implemented to represent mock objects for unit-testing.
package mock

import (
	"errors"

	"github.com/igorlopushko/framey.homework/api/model"
)

// A MockProvider represents mock provider implementation.
type MockProvider struct {
	Name        string
	ReturnError bool
}

// Runs speed test simulation over the mock provider.
func (m MockProvider) Run() ([]model.SpeedTestResult, error) {
	if m.ReturnError {
		return nil, errors.New("Mock provider error")
	}

	return []model.SpeedTestResult{
		{
			Down:     95,
			Up:       90,
			DownUnit: "Mbps",
			UpUnit:   "Kbps",
		}}, nil
}
