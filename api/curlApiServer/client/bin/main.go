package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var webUrl string = GetEnv("POST_URL")
var body []byte

func main() {
	err := curlUrl()
	if err != "" {
		log.Fatal(err)
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return "URL is not configured"
}

func curlUrl() string {
	url := webUrl
	var jsonStr = []byte(`{"text":"I just simply curl it from crontab every minute and this is the text"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("respose status:", resp.Status)
	fmt.Println("response header:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body:", string(body))
	return string(body)
}
