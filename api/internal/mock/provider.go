package mock

import (
	"errors"

	"github.com/igorlopushko/framey.homework/api/model"
)

type MockProvider struct {
	Name        string
	ReturnError bool
}

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
