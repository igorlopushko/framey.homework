package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/igorlopushko/framey.homework/api/config"
	"github.com/igorlopushko/framey.homework/api/provider"
	"github.com/igorlopushko/framey.homework/api/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var providerName string

var rootCmd = &cobra.Command{
	Use:   "go run main.go --provider",
	Short: "SpeedTest CLI tool",
	Long:  "SpeedTest CLI is a tool that tests your internet connection",
	RunE:  run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&providerName, "provider", "p", "", "Speed test provider name. Available values: 'speedtest.net' and 'fast.com'")
}

func initConfig() {
	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})

	logLvl, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		logrus.Warn("could not parse log level, using debug default")
		logLvl = logrus.DebugLevel
	}
	logrus.SetLevel(logLvl)
}

func run(cmd *cobra.Command, _ []string) error {
	p := make(map[string]service.IProvider)

	switch providerName {
	case "speedtest.net":
		p[providerName] = &provider.SpeedTestProvider{}
		break
	case "fast.com":
		p[providerName] = &provider.FastProvider{}
		break
	case "":
		// run speed test for 2 providers
		p["speedtest.net"] = &provider.SpeedTestProvider{}
		p["fast.com"] = &provider.FastProvider{}
		break
	default:
		return errors.New("'provider' parameter value is not recognized")
	}

	s := &service.Service{Providers: p}
	r, err := s.Exec()
	if err != nil {
		return err
	}

	for k, v := range r {
		fmt.Printf("\nSpeed test result for '%s' provider\n", k)
		for _, t := range v {
			fmt.Printf(">>Download: %.2f %s, Upload: %.2f %s\n", t.Down, t.DownUnit, t.Up, t.UpUnit)
		}
	}

	return nil
}
