package service

import (
	"fmt"
	"time"

	"github.com/igorlopushko/framey.homework/api/model"
	"github.com/sirupsen/logrus"
)

type IProvider interface {
	Run() ([]model.SpeedTestResult, error)
}

type Service struct {
	Providers map[string]IProvider
}

func (s *Service) Test() (map[string][]model.SpeedTestResult, error) {
	if len(s.Providers) == 0 {
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
