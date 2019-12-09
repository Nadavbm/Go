package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
This function implementing os get environment variables and print it. To use it, set your local VM with env vars as follow:
export URL1="http://localhost.local"
export PORT=8081
*/

var defaultNum = 8080

func main() {
	fmt.Println(GetEnv("URL1"))
	num := GetEnvInt("PORT")
	fmt.Println(num)
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return "No environment variable configured"
}

func GetEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return defaultNum
}
