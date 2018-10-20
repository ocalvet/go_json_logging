package main

import (
	"net/http"

	"github.com/ocalvet/go_json_logging/logger"
	"github.com/ocalvet/go_json_logging/router"
)

func main() {
	log := logger.NewLogger()
	router := router.NewRouter(log)
	// logrusLogger := logrus.New()
	// logrusLogger.Formatter = &logrus.JSONFormatter{
	// 	// disable, as we set our own
	// 	DisableTimestamp: true,
	// }

	http.ListenAndServe(":2018", router)
}
