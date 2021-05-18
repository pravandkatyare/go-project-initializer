package init

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pravandkatyare/skillpanda/app"
	"github.com/pravandkatyare/skillpanda/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Initialize is for loading all api's before starting server
func Initialize() {
	setupViper()
	setupLogging()
	app.StartApplication()
}

func setupViper() {
	viper.AutomaticEnv()

	defaults := map[string]interface{}{
		"PORT":           8080,
		"header.api-key": "x-default",
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	// Read from application.yml file
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found: %s", err))
		} else {
			panic(fmt.Errorf("error reading config file: %s", err))
		}
	}
}

func setupLogging() {
	// --------logrus default setup-------
	host, err := os.Hostname()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Service Startup")

	// Show the version and build info
	logrus.Infof("Golang OS             : %s", runtime.GOOS)
	logrus.Infof("Golang Arch           : %s", runtime.GOARCH)
	logrus.Infof("Service Host          : %s", host)

	logging.Initialize()

	// -------logrus default setup complete------
}
