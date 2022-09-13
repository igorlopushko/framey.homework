// Package service is implemented to represent speed test service behavior.
package service

import (
	"fmt"
	"time"

	"github.com/igorlopushko/framey.homework/api/model"
	"github.com/sirupsen/logrus"
)

// A IProvider interface which defines speed test provider behavior.
type IProvider interface {
	Run() ([]model.SpeedTestResult, error)
}

// A Service is a representation of the services which performs logic to execute speed tests over specified providers.
type Service struct {
	Providers map[string]IProvider
}

// Executes speed tests over the specified providers.
func (s Service) Exec() (map[string][]model.SpeedTestResult, error) {
	if s.Providers == nil || len(s.Providers) == 0 {
		return nil, nil
	}

	result := make(map[string][]model.SpeedTestResult)

	for k, v := range s.Providers {
		logrus.Debug(fmt.Sprintf("Start speed test for provider: '%s'", k))

		start := time.Now()

		r, err := v.Run()
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		logrus.Debug(fmt.Sprintf("Completed. Time elapsed: %.3f secs", time.Since(start).Seconds()))

		result[k] = r
	}

	return result, nil
}
