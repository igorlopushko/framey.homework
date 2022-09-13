// Package model is implemented to represent domain models.
package model

// A SpeedTestResult represents speed test data.
type SpeedTestResult struct {
	Down     float64
	DownUnit string
	Up       float64
	UpUnit   string
}
