package provider

import (
	"github.com/igorlopushko/framey.homework/api/model"
	"github.com/showwin/speedtest-go/speedtest"
)

// A SpeedTestProvider represents Ookla's provider behavior.
type SpeedTestProvider struct {
	Name string
}

// Runs speed test for the Ookla's provider.
func (f *SpeedTestProvider) Run() ([]model.SpeedTestResult, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	serverList, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, err
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return nil, err
	}

	r := make([]model.SpeedTestResult, 0)

	for _, s := range targets {
		err = s.PingTest()
		if err != nil {
			return nil, err
		}

		err = s.DownloadTest(false)
		if err != nil {
			return nil, err
		}

		err = s.UploadTest(false)
		if err != nil {
			return nil, err
		}

		r = append(r, model.SpeedTestResult{
			Down:     s.DLSpeed,
			Up:       s.ULSpeed,
			DownUnit: "Mbps",
			UpUnit:   "Mbps",
		})
	}

	return r, err
}
