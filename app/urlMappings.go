package app

import (
	"github.com/pravandkatyare/skillpanda/handler"
	"github.com/pravandkatyare/skillpanda/logging"
)

// MapURLs maps all the handlers with the functions
func mapURLs() {
	// middleware for header authentication of static keyword
	logging.Infof("Attaching middleware for authentication")
	server.Use(handler.AuthMiddleware())
	logging.Infof("Middleware attached for authentication")

	//mapping handlers with endpoint
	logging.Infof("Mapping endpoints with handlers")

	// server.POST("url-path", handlerfunction())
}
