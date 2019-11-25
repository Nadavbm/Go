package main

import (
	"fmt"
	"time"
)

func cronJob() {
	tic := time.NewTicker(10 * time.Second)
	for _ = range tic.C {
		fmt.Println("This check will be executed every 10 seconds")
	}
}

func main() {
	go cronJob()
	// This select will keep the main function on
	select {}
}
