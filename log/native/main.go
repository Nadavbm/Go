package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	setLogger()
	getURL("www.google.com")
	getURL("https://www.google.de")
}

func setLogger() {
	logFile, err := os.OpenFile("/var/log/example.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Log file was not created")
		panic(err)
	}
	log.SetOutput(logFile)
}

func getURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	} else {
		log.Printf("INFO: Status code %s - All OK", resp.Status)
		resp.Body.Close()
	}
}
