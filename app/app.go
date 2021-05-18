package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pravandkatyare/skillpanda/logging"
	"github.com/spf13/viper"
)

var (
	server = gin.Default()
)

// StartApplication initializes the server and start application
func startApplication() {
	logging.Infof("Starting application on port: %s", viper.GetString("port"))
	mapURLs()
	logging.Fatalf("Error while starting application: %s", server.Run(":"+viper.GetString("port")))
}
