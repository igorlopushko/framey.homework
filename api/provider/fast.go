package provider

import (
	fast "github.com/igorlopushko/framey.homework/api/internal/fast"
	"github.com/igorlopushko/framey.homework/api/model"
)

// A FastProvider represents Netflix's provider behavior.
type FastProvider struct {
	Name string
}

// Runs speed test for the Netflix's provider.
func (f *FastProvider) Run() ([]model.SpeedTestResult, error) {
	r, err := fast.Run()
	if err != nil {
		return nil, err
	}
	return []model.SpeedTestResult{
		{
			Down:     r.Down,
			Up:       r.Up,
			DownUnit: r.DownUnit,
			UpUnit:   r.UpUnit,
		}}, err
}
