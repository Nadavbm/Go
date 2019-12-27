package main

import (
	"net/http"
	"go.uber.org/zap"
)

var logger *zap.Logger 

func main() {
	setLogger()
	defer logger.Sync()
	getURL("www.google.com")
	getURL("https://www.google.de")
}

func setLogger() {
	logger, _ = zap.NewProduction()
}

func getURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("ERROR: ", zap.Error(err))
	} else {
		logger.Info("INFO: ", zap.String("Status code", resp.Status))
		resp.Body.Close()
	}
}
