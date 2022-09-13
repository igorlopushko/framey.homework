// Package cmd is implemented to represent the command line tool to execute the program logic
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

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

	rootCmd.PersistentFlags().StringVarP(
		&providerName,
		"provider",
		"p",
		"",
		`Speed test provider name. Skip to run all providers. Available values:
		[o] or [ookla] - for Ookla's provider
		[n] or [netflix] - for Netflix's provider`)
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

	switch strings.ToLower(providerName) {
	case "ookla", "o":
		p["ookla"] = &provider.SpeedTestProvider{}
	case "netflix", "n":
		p["netflix"] = &provider.FastProvider{}
	case "":
		// run speed test for 2 providers
		p["ookla"] = &provider.SpeedTestProvider{}
		p["netflix"] = &provider.FastProvider{}
	default:
		return errors.New("'provider' parameter value is not recognized")
	}

	fmt.Printf("Start download and upload speeds test\n")

	// execute internet speed testing
	s := &service.Service{Providers: p}
	r, err := s.Exec()
	if err != nil {
		return err
	}

	// print out the results
	for k, v := range r {
		fmt.Printf("\nSpeed test result for '%s' provider\n", k)
		for _, t := range v {
			fmt.Printf(">>Download: %.2f %s, Upload: %.2f %s\n", t.Down, t.DownUnit, t.Up, t.UpUnit)
		}
	}

	fmt.Printf("\nFinished download and upload speeds test\n")

	return nil
}
