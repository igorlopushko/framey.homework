package provider

import (
	fast "github.com/igorlopushko/framey.homework/api/internal/fast"
	"github.com/igorlopushko/framey.homework/api/model"
)

// A FastProvider represents fast.com provider behavior.
type FastProvider struct {
	Name string
}

// Runs speed test for the fast.com provider.
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
