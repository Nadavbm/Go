package main

import (
	"fmt"
	"log"
	"os"
	//"io"
	"io/ioutil"
)

var path = "./README.txt"

func main() {
	fmt.Println("The log file located at: ", PrintPath())
	fmt.Println("===========================================")
	fmt.Println("The file content in bytes slice:")
	fmt.Println("================================")
	ReadAsBuffer()
	fmt.Println("\nAnd now convert bytes to text:")
	fmt.Println("==============================")
	ReadAsText()
}

func ReadAsText() {
	text, err := ioutil.ReadFile("README.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot open file", err)
	}
	fmt.Println("", string(text))
}

func ReadAsBuffer() {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	// Retrun error is cannot open the file
	if err != nil {
		log.Fatal(err)
	}
	// Close file at the end of this function
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	fmt.Print(text)
}

func PrintPath() string {
	return path
}
